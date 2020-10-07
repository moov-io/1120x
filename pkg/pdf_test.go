// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pkg

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/moov-io/1120x/pkg/pdf_generator"
	"github.com/moov-io/1120x/pkg/xmls"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PdfTestSuite struct {
	suite.Suite
	InputXML []byte
}

func (suite *PdfTestSuite) SetupTest() {
	buf, err := ioutil.ReadFile(filepath.Join("..", "test", "testdata", "irs990_public.xml"))
	assert.Equal(suite.T(), nil, err)
	suite.InputXML = buf
}

func (suite *PdfTestSuite) TestAllSteps() {
	// 1. create tax return
	rInstance, err := xmls.CreateReturn(suite.InputXML)
	assert.Equal(suite.T(), nil, err)
	assert.NotNil(suite.T(), rInstance)

	// 2. create return form
	form, err := xmls.CreateReturnForm(rInstance)
	assert.Equal(suite.T(), nil, err)
	assert.NotNil(suite.T(), form)

	// 3. create html and pdf using xsltproc
	params := xmls.XMLParameters{}
	generator, err := pdf_generator.GetHtmlGenerator(form, &params, pdf_generator.GeneratorApplicationMode)
	assert.Equal(suite.T(), nil, err)
	assert.NotNil(suite.T(), generator)

	documentCnt := len(generator.GetDocuments())
	htmls, err := generator.GenerateHtml()
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), documentCnt, len(htmls))

	pdfs, err := generator.GeneratePDF(htmls)
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), documentCnt, len(pdfs))

	// 4. create html and pdf using ratago/xslt
	generator, err = pdf_generator.GetHtmlGenerator(form, &params, pdf_generator.GeneratorPackageMode)
	assert.Equal(suite.T(), nil, err)
	assert.NotNil(suite.T(), generator)

	htmls, err = generator.GenerateHtml()
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), documentCnt, len(htmls))

	pdfs, err = generator.GeneratePDF(htmls)
	assert.Equal(suite.T(), nil, err)
	assert.Equal(suite.T(), documentCnt, len(pdfs))
}

func TestPdfTestSuite(t *testing.T) {
	suite.Run(t, new(PdfTestSuite))
}
