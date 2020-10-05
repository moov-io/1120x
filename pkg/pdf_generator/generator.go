package pdf_generator

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/jbowtie/ratago/xslt"
	"github.com/moov-io/1120x/pkg/utils"
	"github.com/moov-io/1120x/pkg/xmls"
)

// Html generator form xml
type PdfGenerator interface {
	GenerateHtml() ([]HtmlData, error)
	GeneratePDF([]HtmlData) ([]PdfData, error)
	GetDocuments() []xmls.XMLDocument
}

type HtmlData []byte
type PdfData []byte

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)

	pathStylesheets = filepath.Join(basePath, "..", "..", "data", "mef", "Stylesheets")
	pathMefAsset    = filepath.Join(basePath, "..", "..", "data")
	pdfNamePattern  = "generated_pdf_"
	xmlNamePattern  = "generated_xml_"
	tempDir         = ""
	xsltProc        = "xsltproc"

	GeneratorPackageMode     = "package"
	GeneratorApplicationMode = "application"
	GeneratorDefaultMode     = GeneratorApplicationMode
)

type htmlGenerator struct {
	Documents []xmls.XMLDocument
	Mode      string
}

// GetHtmlGenerator returns instance of html generator
func GetHtmlGenerator(r xmls.ReturnFormData, params *xmls.XMLParameters, mode string) (PdfGenerator, error) {
	xmlArray := r.GenerateXMLDocument(params)
	if len(xmlArray) == 0 {
		return nil, utils.ErrEmptyXML
	}
	if mode != GeneratorPackageMode && mode != GeneratorApplicationMode {
		mode = GeneratorDefaultMode
	}
	return &htmlGenerator{Documents: xmlArray, Mode: mode}, nil
}

// GetDocuments returns documents
func (g *htmlGenerator) GetDocuments() []xmls.XMLDocument {
	return g.Documents
}

// GeneratePDF returns pdf data
func (g *htmlGenerator) GeneratePDF(htmls []HtmlData) ([]PdfData, error) {
	tmltopdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	var pdfs []PdfData
	for _, htmlData := range htmls {
		r := bytes.NewReader(htmlData)
		page := wkhtmltopdf.NewPageReader(r)
		tmltopdf.AddPage(page)
		err = tmltopdf.Create()
		if err != nil {
			return nil, err
		}
		pdfs = append(pdfs, tmltopdf.Bytes())
		tmltopdf.ResetPages()
	}
	return pdfs, nil
}

// GenerateHtml returns html data
func (g *htmlGenerator) GenerateHtml() ([]HtmlData, error) {
	var data []HtmlData
	for _, document := range g.Documents {
		if g.Mode == GeneratorPackageMode {
			html, err := runHtmlConvert(document)
			if err != nil {
				return nil, err
			}
			data = append(data, html)
		} else if g.Mode == GeneratorApplicationMode {
			xslFile, err := getStylesheetFile(document)
			if err != nil {
				return nil, err
			}
			xmlFile, err := generateXmlFile(document)
			if err != nil {
				return nil, err
			}
			html, err := processXslt(*xslFile, xmlFile.Name())
			defer os.Remove(xmlFile.Name())
			if err != nil {
				return nil, err
			}
			data = append(data, html)
		}
	}
	return data, nil
}

func getXSLFileName(formCode string) string {
	return "IRS" + formCode + ".xsl"
}

func getStylesheetFile(document xmls.XMLDocument) (*string, error) {
	filePath := filepath.Join(pathStylesheets, strconv.Itoa(document.Year), getXSLFileName(document.Type))
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, err
	}
	return &filePath, nil
}

func generateXmlFile(document xmls.XMLDocument) (*os.File, error) {
	tmpFile, err := ioutil.TempFile(tempDir, xmlNamePattern)
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(tmpFile.Name(), document.XML, 0644)
	if err != nil {
		defer os.Remove(tmpFile.Name())
		return nil, err
	}
	return tmpFile, nil
}

// using http://xmlsoft.org/XSLT/xsltproc.html
func processXslt(xslFile string, xmlFile string) ([]byte, error) {
	execFile, err := exec.LookPath(xsltProc)
	if err != nil {
		return nil, err
	}

	cmd := exec.Cmd{
		Args: []string{xsltProc, xslFile, xmlFile},
		Env:  os.Environ(),
		Path: execFile,
	}

	buf, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return buf, err
}

// using github.com/jbowtie/ratago/xslt
func runHtmlConvert(document xmls.XMLDocument) (HtmlData, error) {
	path, err := getStylesheetFile(document)
	if err != nil {
		return nil, err
	}
	options := xslt.StylesheetOptions{IndentOutput: true}
	return runHtmlConvertWithOptions(*path, document, options)
}

func runHtmlConvertWithOptions(xslFile string, inputBuf xmls.XMLDocument, options xslt.StylesheetOptions) (HtmlData, error) {
	style, err := xml.ReadFile(xslFile, xml.StrictParseOption|xml.XML_PARSE_XINCLUDE)
	if err != nil {
		return nil, err
	}

	input, err := xml.Parse(inputBuf.XML, xml.DefaultEncodingBytes, nil, xml.DefaultParseOption, xml.DefaultEncodingBytes)
	if err != nil {
		return nil, err
	}
	stylesheet, err := xslt.ParseStylesheet(style, xslFile)
	if err != nil {
		return nil, err
	}
	output, err := stylesheet.Process(input, options)
	if err != nil {
		return nil, err
	}

	return HtmlData(output), nil
}
