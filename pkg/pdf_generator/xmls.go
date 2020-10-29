// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pdf_generator

import (
	"bytes"
	"github.com/antchfx/xmlquery"
	"github.com/moov-io/1120x/pkg/irs_990"
	"github.com/moov-io/1120x/pkg/utils"
)

var (
	rootReturn   = "//Return"
	returnTypeCd = "//ReturnHeader/ReturnTypeCd"
)

type XMLDocument struct {
	XML  []byte
	Type string
	Year int
}

// General return data interface
type ReturnFormData interface {
	GenerateXMLDocument(params *XMLParameters) []XMLDocument
}

// Create return instance
func CreateReturn(buf []byte) (utils.Return, error) {
	returnType, err := getReturnTypeFromRawXML(buf)
	if err != nil {
		return nil, err
	}

	switch *returnType {
	case utils.IRS990ReturnTypeCode:
		var r irs_990.Return
		err = r.Parse(buf)
		if err != nil {
			return nil, err
		}
		return &r, err
	}
	return nil, utils.ErrFailedCreateTaxReturn
}

func getReturnTypeFromRawXML(buf []byte) (*string, error) {
	doc, err := xmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	root := xmlquery.FindOne(doc, rootReturn)
	if root == nil {
		return nil, utils.ErrUnknownReturnType
	}

	if n := root.SelectElement(returnTypeCd); n != nil {
		returnType := n.InnerText()
		return &returnType, nil
	}

	return nil, utils.ErrUnknownReturnType
}

// Create return form
func CreateReturnForm(instance utils.Return) (ReturnFormData, error) {
	info := instance.InspectData()
	if info == nil {
		return nil, utils.ErrEmptyXML
	}
	returnData := &FormDataCreator{
		Header:            info.Header,
		Data:              info.Data,
		SubmissionYear:    instance.ReturnYear(),
		SubmissionVersion: instance.ReturnVersion(),
		SubmissionType:    instance.ReturnType(),
	}
	return returnData, nil
}
