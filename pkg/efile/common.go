// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package efile

import (
	"bytes"
	"encoding/xml"
	"errors"
	"regexp"
)

type FaultDetail []string

func (r FaultDetail) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ArrayType string   `xml:"http://schemas.xmlsoap.org/wsdl/ arrayType,attr"`
		Items     []string `xml:"item"`
	}
	output.Items = []string(r)
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: "xmlns:ns1"}, Value: "http://www.w3.org/2001/XMLSchema"})
	output.ArrayType = "ns1:anyType[]"
	return e.EncodeElement(&output, start)
}
func (r *FaultDetail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var tok xml.Token
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if tok, ok := tok.(xml.StartElement); ok {
			var item string
			if err = d.DecodeElement(&item, &tok); err == nil {
				*r = append(*r, item)
			}
		}
		if _, ok := tok.(xml.EndElement); ok {
			break
		}
	}
	return err
}

type EncodingStyle []string

func (r *EncodingStyle) MarshalText() ([]byte, error) {
	result := make([][]byte, 0, len(*r))
	for _, v := range *r {
		result = append(result, []byte(v))
	}
	return bytes.Join(result, []byte(" ")), nil
}
func (r *EncodingStyle) UnmarshalText(text []byte) error {
	for _, v := range bytes.Fields(text) {
		*r = append(*r, string(v))
	}
	return nil
}

// Must match the pattern [0-9]{13}[a-z0-9]{7}
type SubmissionIdType string

func (r SubmissionIdType) Validate() error {
	reg := regexp.MustCompile(`[0-9]{13}[a-z0-9]{7}`)
	if !reg.MatchString(string(r)) {
		return errors.New("SubmissionIdType is invalid")
	}
	return nil
}

// Fault reporting structure
type Fault struct {
	Faultcode   xml.Name     `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultcode"`
	Faultstring string       `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultstring"`
	Faultactor  string       `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultactor,omitempty" json:",omitempty"`
	Detail      *FaultDetail `xml:"http://schemas.xmlsoap.org/soap/envelope/ detail,omitempty" json:",omitempty"`
}

// interface for manifest
type ManifestXml interface {
	Validate() error
	XmlData() ([]byte, error)
	SubmissionIdentifier() SubmissionIdType
	SetSubmissionIdentifier(id SubmissionIdType)
}
