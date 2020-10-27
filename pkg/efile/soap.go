// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package efile

import (
	"encoding/xml"
	"github.com/moov-io/1120x/pkg/utils"
)

type Envelope struct {
	Xmlns  string      `xml:"xmlns,attr,omitempty" json:",omitempty"`
	Items  []string    `xml:",any"`
	Header *SoapHeader `xml:"Header,omitempty" json:",omitempty"`
	Body   *SoapBody   `xml:"Body" json:",omitempty"`
}

func (r *Envelope) Init() error {
	r.Xmlns = "http://www.irs.gov/efile"
	return nil
}

func (r Envelope) Validate() error {
	return utils.Validate(&r)
}

type SoapBody struct {
	Body     []string
	Manifest *TransmissionManifest
}

func (r SoapBody) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ArrayType string   `xml:"http://schemas.xmlsoap.org/wsdl/ arrayType,attr"`
		Items     []string `xml:"item"`
	}
	output.Items = r.Body
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: "xmlns:ns1"}, Value: "http://www.w3.org/2001/XMLSchema"})
	output.ArrayType = "ns1:anyType[]"
	return e.EncodeElement(&output, start)
}
func (r *SoapBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	r.Manifest = nil
	var tok xml.Token
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if tok, ok := tok.(xml.StartElement); ok {
			var item string
			if err = d.DecodeElement(&item, &tok); err == nil {
				r.Body = append(r.Body, item)
				if r.Manifest == nil {
					var manifest TransmissionManifest
					if err = xml.Unmarshal([]byte(item), &manifest); err == nil {
						r.Manifest = &manifest
					}
				}
			}
		}
		if _, ok := tok.(xml.EndElement); ok {
			break
		}
	}
	return err
}

type SoapHeader struct {
	Header       []string
	Transmission *IFATransmissionHeaderType
}

func (r SoapHeader) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ArrayType string   `xml:"http://schemas.xmlsoap.org/wsdl/ arrayType,attr"`
		Items     []string `xml:"item"`
	}
	output.Items = r.Header
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Space: "", Local: "xmlns:ns1"}, Value: "http://www.w3.org/2001/XMLSchema"})
	output.ArrayType = "ns1:anyType[]"
	return e.EncodeElement(&output, start)
}
func (r *SoapHeader) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var tok xml.Token
	r.Transmission = nil
	for tok, err = d.Token(); err == nil; tok, err = d.Token() {
		if tok, ok := tok.(xml.StartElement); ok {
			{
				var item string
				if err = d.DecodeElement(&item, &tok); err == nil {
					r.Header = append(r.Header, item)
					if r.Transmission == nil {
						var header IFATransmissionHeaderType
						if err = xml.Unmarshal([]byte(item), &header); err == nil {
							r.Transmission = &header
						}
					}
				}
			}
		}
		if _, ok := tok.(xml.EndElement); ok {
			break
		}
	}

	return err
}
