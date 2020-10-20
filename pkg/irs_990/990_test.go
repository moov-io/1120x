// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import (
	"encoding/json"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type Irs990TestSuite struct {
	suite.Suite
	InputXML []byte
}

func (suite *Irs990TestSuite) SetupTest() {
	buf, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_public.xml"))
	assert.Equal(suite.T(), nil, err)
	suite.InputXML = buf
}

func (suite *Irs990TestSuite) TestAllSteps() {
	returnData := &Return{}

	// 1. parse from xml data
	err := xml.Unmarshal(suite.InputXML, returnData)
	assert.Equal(suite.T(), nil, err)

	// 2. struct to xml buf
	xmlOrgBuf, err := xml.MarshalIndent(returnData, "", "\t")
	assert.Equal(suite.T(), nil, err)

	// 3. struct to json buf
	jsonBuf, err := json.MarshalIndent(returnData, "", "\t")
	assert.Equal(suite.T(), nil, err)

	// 4. json buf to struct
	newReturnData := &Return{}
	err = json.Unmarshal(jsonBuf, newReturnData)
	assert.Equal(suite.T(), nil, err)

	// 5. struct to xml buf
	xmlBuf, err := xml.MarshalIndent(newReturnData, "", "\t")
	assert.Equal(suite.T(), nil, err)

	// 6. check xml buffers
	assert.Equal(suite.T(), xmlOrgBuf, xmlBuf)
}

func TestIrs990TestSuite(t *testing.T) {
	suite.Run(t, new(Irs990TestSuite))
}
