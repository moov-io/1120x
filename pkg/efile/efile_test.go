// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package efile

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/1120x/pkg/irs_990"
	"github.com/stretchr/testify/assert"
)

func TestIrs990TransmissionFile(t *testing.T) {
	InputXML, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_transmission_file.xml"))
	assert.Equal(t, nil, err)

	// 1. parse from xml data
	transmission := &Irs990TransmissionFile{}

	err = xml.Unmarshal(InputXML, &transmission.Soap)
	assert.Equal(t, nil, err)

	err = transmission.Soap.Init()
	assert.Equal(t, nil, err)

	// 2. struct to xml buf
	xmlOrgBuf, err := xml.MarshalIndent(transmission, "", "\t")
	assert.Equal(t, nil, err)

	// 3. struct to json buf
	jsonBuf, err := json.MarshalIndent(transmission, "", "\t")
	assert.Equal(t, nil, err)

	// 4. json buf to struct
	newTransmission := &Irs990TransmissionFile{}

	err = json.Unmarshal(jsonBuf, newTransmission)
	assert.Equal(t, nil, err)

	// 5. struct to xml buf
	xmlBuf, err := xml.MarshalIndent(newTransmission, "", "\t")
	assert.Equal(t, nil, err)

	err = newTransmission.Validate()
	assert.Equal(t, nil, err)

	// 6. check xml buffers
	assert.Equal(t, xmlOrgBuf, xmlBuf)

	// 7. check functions
	_, err = newTransmission.SOAPEnvelope()
	assert.Equal(t, nil, err)

	returnBuf, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_return.xml"))
	assert.Equal(t, nil, err)

	manifestBuf, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_submission_manifest.xml"))
	assert.Equal(t, nil, err)

	file := &irs_990.Irs990File{}

	err = xml.Unmarshal(returnBuf, &file.XmlData)
	assert.Equal(t, nil, err)

	file.Manifest = &irs_990.IRSSubmissionManifest{}
	err = xml.Unmarshal(manifestBuf, file.Manifest)
	assert.Equal(t, nil, err)

	newTransmission.Attachments = append(newTransmission.Attachments, file)

	_, err = newTransmission.SOAPAttachment()
	assert.Equal(t, nil, err)

	version := newTransmission.Version()
	assert.Equal(t, "2015v2.0", version)
}

func TestUnusedStructs(t *testing.T) {
	InputXML, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_transmission_file.xml"))
	assert.Equal(t, nil, err)

	fault := &FaultDetail{}
	err = xml.Unmarshal(InputXML, fault)
	assert.Equal(t, nil, err)
	_, err = xml.Marshal(fault)
	assert.Equal(t, nil, err)

	style := &EncodingStyle{"test1", "test2"}
	err = xml.Unmarshal(InputXML, fault)
	assert.Equal(t, nil, err)
	_, err = xml.Marshal(style)
	assert.Equal(t, nil, err)
	_, err = style.MarshalText()
	assert.Equal(t, nil, err)
	err = style.UnmarshalText([]byte("example"))
	assert.Equal(t, nil, err)

	var idType MessageIdType
	err = idType.Validate()
	assert.NotNil(t, err)
	idType = "012345678912abcdefgh"
	err = idType.Validate()
	assert.Equal(t, nil, err)

	var detail TransmitterDetail
	err = detail.Validate()
	assert.NotNil(t, err)

	var header IFATransmissionHeaderType
	err = header.Validate()
	assert.NotNil(t, err)

	mani := TransmissionManifest{
		SubmissionDataList: SubmissionDataListType{
			Cnt: 1,
			SubmissionData: []SubmissionDataType{
				{SubmissionId: irs_990.SubmissionIdType("111")},
			},
		},
	}
	err = mani.Validate()
	assert.NotNil(t, err)
}
