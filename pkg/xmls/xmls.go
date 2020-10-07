// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package xmls

import (
	"bytes"
	"github.com/antchfx/xmlquery"
	"github.com/moov-io/1120x/pkg/utils"
)

// General return data interface
type Return interface {
	Parse([]byte) error
	InspectData() *ReturnInspectInfo
	ReturnVersion() string
	ReturnYear() int
	ReturnType() string
}

// General return data interface
type ReturnFormData interface {
	GenerateXMLDocument(params *XMLParameters) []XMLDocument
}

type XMLDocument struct {
	XML  []byte
	Type string
	Year int
}

type ReturnInspectInfo struct {
	Header interface{}
	Data   []ReturnInspectData
}

type ReturnInspectData struct {
	Data     interface{}
	DataType string
}

var (
	rootReturn   = "//Return"
	returnTypeCd = "//ReturnHeader/ReturnTypeCd"
)

// Create return instance
func CreateReturn(buf []byte) (Return, error) {
	returnType, err := getReturnTypeFromRawXML(buf)
	if err != nil {
		return nil, err
	}

	switch *returnType {
	case utils.IRS990ReturnTypeCode:
		var r Return1099
		err = r.Parse(buf)
		if err != nil {
			return nil, err
		}
		return &r, err
	}
	return nil, utils.ErrFailedCreateTaxReturn
}

// Create return form
func CreateReturnForm(instance Return) (ReturnFormData, error) {
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

func getReturnTypeFromRawXML(buf []byte) (*string, error) {
	doc, err := xmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	root := xmlquery.FindOne(doc, rootReturn)
	if n := root.SelectElement(returnTypeCd); n != nil {
		returnType := n.InnerText()
		return &returnType, nil
	}

	return nil, utils.ErrUnknownReturnType
}
