// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIrs990TestSuite(t *testing.T) {
	InputXML, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_public.xml"))
	assert.Equal(t, nil, err)

	// 1. parse from xml data
	returnData := &ReturnDataReturn{}

	err = returnData.Validate()
	assert.NotNil(t, err)

	err = xml.Unmarshal(InputXML, returnData)
	assert.Equal(t, nil, err)

	// 2. struct to xml buf
	xmlOrgBuf, err := xml.MarshalIndent(returnData, "", "\t")
	assert.Equal(t, nil, err)

	// 3. struct to json buf
	jsonBuf, err := json.MarshalIndent(returnData, "", "\t")
	assert.Equal(t, nil, err)

	// 4. json buf to struct
	newReturnData := &ReturnDataReturn{}

	err = json.Unmarshal(jsonBuf, newReturnData)
	assert.Equal(t, nil, err)

	// 5. struct to xml buf
	xmlBuf, err := xml.MarshalIndent(newReturnData, "", "\t")
	assert.Equal(t, nil, err)

	// 6. check xml buffers
	assert.Equal(t, xmlOrgBuf, xmlBuf)
}
