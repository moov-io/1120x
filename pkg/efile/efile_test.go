// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package efile

import (
	"encoding/json"
	"encoding/xml"
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

	err = xml.Unmarshal(InputXML, transmission)
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
}
