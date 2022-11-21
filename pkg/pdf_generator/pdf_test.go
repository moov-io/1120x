// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pdf_generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jbowtie/ratago/xslt"
	"github.com/moov-io/1120x/pkg/irs_990"
	"github.com/stretchr/testify/assert"
)

func TestPdfGenerator(t *testing.T) {
	InputXML, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_return.xml"))
	assert.Equal(t, nil, err)

	// 1. create tax return
	rInstance, err := CreateReturn(InputXML)
	assert.Equal(t, nil, err)
	assert.NotNil(t, rInstance)

	// 2. create return form
	form, err := CreateReturnForm(rInstance)
	assert.Equal(t, nil, err)
	assert.NotNil(t, form)

	// 3. create html and pdf using xsltproc
	_, err = GetHtmlGenerator(form, nil, "")
	assert.Equal(t, nil, err)

	params := XMLParameters{}
	generator, err := GetHtmlGenerator(form, &params, GeneratorApplicationMode)
	assert.Equal(t, nil, err)
	assert.NotNil(t, generator)

	documentCnt := len(generator.GetDocuments())
	htmls, err := generator.GenerateHtml()
	assert.Equal(t, nil, err)
	assert.Equal(t, documentCnt, len(htmls))

	pdfs, err := generator.GeneratePDF(htmls)
	assert.Equal(t, nil, err)
	assert.Equal(t, documentCnt, len(pdfs))

	// 4. create html and pdf using ratago/xslt
	generator, err = GetHtmlGenerator(form, &params, GeneratorPackageMode)
	assert.Equal(t, nil, err)
	assert.NotNil(t, generator)

	htmls, err = generator.GenerateHtml()
	assert.Equal(t, nil, err)
	assert.Equal(t, documentCnt, len(htmls))

	pdfs, err = generator.GeneratePDF(htmls)
	assert.Equal(t, nil, err)
	assert.Equal(t, documentCnt, len(pdfs))
}

func TestUnusedStructs(t *testing.T) {
	InputXML, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_invalid_return.xml"))
	assert.Equal(t, nil, err)

	_, err = CreateReturn([]byte("test"))
	assert.NotNil(t, err)

	_, err = CreateReturn(InputXML)
	assert.NotNil(t, err)

	r := &irs_990.Return{}
	_, err = CreateReturnForm(r)
	assert.NotNil(t, err)

	doc := XMLDocument{}
	_, err = runHtmlConvert(doc)
	assert.NotNil(t, err)

	options := xslt.StylesheetOptions{IndentOutput: true}
	_, err = runHtmlConvertWithOptions("test", doc, options)
	assert.NotNil(t, err)

}
