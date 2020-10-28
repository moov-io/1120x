// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pdf_generator

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPdfGenerator(t *testing.T) {
	InputXML, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_return.xml"))
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
