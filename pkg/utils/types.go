// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"regexp"
)

//////////////////////////////////////////////
//		Transmission File                   //
//  MIME Multi-part Content Header          //
//  MIME Part Boundary and Content Header   //
//    SOAP Envelope (SOAPEnvelope)          //
//       SOAP Header                        //
//       SOAP Body                          //
//  MIME Part Boundary                      //
//    SOAP Attachment (SOAPAttachment)      //
//  MIME Multi-part End Boundary            //
//////////////////////////////////////////////

// interface for irs e-file
type IrsTransmissionFile interface {
	Validate() error
	SOAPEnvelope() ([]byte, error)
	SOAPAttachment() ([]byte, error)
	Version() string
}

// interface for return
type IrsReturnFile interface {
	Validate() error
	ZipData() ([]byte, error)
	Version() string
}

// interface for manifest
type ManifestXml interface {
	Validate() error
	XmlData() ([]byte, error)
	SubmissionIdentifier() SubmissionIdType
	SetSubmissionIdentifier(id SubmissionIdType)
}

type ReturnInspectInfo struct {
	Header interface{}
	Data   []ReturnInspectData
}

type ReturnInspectData struct {
	Data     interface{}
	DataType string
}

// General return data interface
type Return interface {
	Parse([]byte) error
	InspectData() *ReturnInspectInfo
	ReturnVersion() string
	ReturnYear() int
	ReturnType() string
}

func FormatXML(data []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return b.Bytes(), nil
		}
		if err != nil {
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
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
