// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"bytes"
	"encoding/xml"
	"io"
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
