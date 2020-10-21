// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package efile

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"regexp"

	"github.com/moov-io/1120x/pkg/irs_990"
	"github.com/moov-io/1120x/pkg/utils"
)

type Irs990TransmissionFile struct {
	Header       *IFATransmissionHeaderType
	Body         *TransmissionManifest
	Attachments  []IrsReturnFile
	SubmissionId string
}

func (r Irs990TransmissionFile) Validate() error {
	return utils.Validate(&r)
}

func (r Irs990TransmissionFile) SOAPEnvelope() ([]byte, error) {
	envelope := SoapEnvelope{}
	if r.Header != nil {
		buf, err := xml.Marshal(r.Header)
		if err != nil {
			return nil, err
		}
		envelope.Header = &SoapHeader{string(buf)}
	}
	if r.Body != nil {
		buf, err := xml.Marshal(r.Body)
		if err != nil {
			return nil, err
		}
		envelope.Body = &SoapBody{string(buf)}
	}
	envelope.Init()

	return xml.Marshal(envelope)
}

func (r Irs990TransmissionFile) SOAPAttachment() ([]byte, error) {
	if len(r.Body.SubmissionDataList.SubmissionData) != len(r.Attachments) {
		return nil, errors.New("don't match submission data and attachments")
	}

	// Create a buffer to write our archive to.
	fileBuf := new(bytes.Buffer)

	// Create a new zip archive.
	writer := zip.NewWriter(fileBuf)
	defer writer.Close()

	submission := r.Body.SubmissionDataList.SubmissionData
	for index, attachment := range r.Attachments {
		name := string(submission[index].SubmissionId)
		f, err := writer.Create(name)
		if err != nil {
			return nil, err
		}
		data, err := attachment.ZipData()
		if err != nil {
			return nil, err
		}
		_, err = f.Write(data)
		if err != nil {
			return nil, err
		}
	}

	return fileBuf.Bytes(), writer.Close()
}

func (r Irs990TransmissionFile) Version() string {
	return "2015v2.0"
}

// common types for elements

// Must match the pattern [0-9]{12}[a-z0-9]{8}
type MessageIdType string

func (r MessageIdType) Validate() error {
	reg := regexp.MustCompile(`[0-9]{12}[a-z0-9]{8}`)
	if !reg.MatchString(string(r)) {
		return errors.New("ETINType is invalid")
	}
	return nil
}

type TransmitterDetail struct {
	ETIN irs_990.ETINType `xml:"ETIN"`
}

func (r TransmitterDetail) Validate() error {
	return utils.Validate(&r)
}

type SubmissionDataType struct {
	SubmissionId         irs_990.SubmissionIdType `xml:"SubmissionId"`
	ElectronicPostmarkTs irs_990.TimestampType    `xml:"ElectronicPostmarkTs"`
}

func (r SubmissionDataType) Validate() error {
	return utils.Validate(&r)
}

type SubmissionDataListType struct {
	Cnt            int                  `xml:"Cnt"`
	SubmissionData []SubmissionDataType `xml:"SubmissionData"`
}

func (r SubmissionDataListType) Validate() error {
	return utils.Validate(&r)
}

type IFATransmissionHeaderType struct {
	MessageId         MessageIdType          `xml:"MessageId"`
	TransmissionTs    *irs_990.TimestampType `xml:"TransmissionTs,omitempty" json:",omitempty"`
	TransmitterDetail TransmitterDetail      `xml:"TransmitterDetail"`
}

func (r IFATransmissionHeaderType) Validate() error {
	return utils.Validate(&r)
}

type TransmissionManifest struct {
	SubmissionDataList SubmissionDataListType `xml:"SubmissionDataList"`
}

func (r TransmissionManifest) Validate() error {
	return utils.Validate(&r)
}
