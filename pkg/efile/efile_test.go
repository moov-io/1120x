// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package efile

import (
	"encoding/json"
	"encoding/xml"
	"github.com/moov-io/1120x/pkg/irs_990"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestIrs990TransmissionFile(t *testing.T) {
	InputXML, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_transmission_file.xml"))
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

	_, err = newTransmission.SOAPAttachment()
	assert.NotNil(t, err)

	version := newTransmission.Version()
	assert.Equal(t, "2015v2.0", version)
}

func TestUnusedStructs(t *testing.T) {
	InputXML, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_transmission_file.xml"))
	assert.Equal(t, nil, err)

	fault := &FaultDetail{}
	err = xml.Unmarshal(InputXML, fault)
	assert.Equal(t, nil, err)
	_, err = xml.Marshal(fault)
	assert.Equal(t, nil, err)

	style := &EncodingStyle{}
	err = xml.Unmarshal(InputXML, fault)
	assert.Equal(t, nil, err)
	_, err = xml.Marshal(style)
	assert.Equal(t, nil, err)

	var idType MessageIdType
	err = idType.Validate()
	assert.NotNil(t, err)

	var detail TransmitterDetail
	err = detail.Validate()
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
