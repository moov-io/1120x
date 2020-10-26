// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import (
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnXmlTestSuite(t *testing.T) {
	InputXML, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_return.xml"))
	assert.Equal(t, nil, err)

	// 1. parse from xml data
	returnData := &Return{}

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
	newReturnData := &Return{}

	err = json.Unmarshal(jsonBuf, newReturnData)
	assert.Equal(t, nil, err)

	// 5. struct to xml buf
	xmlBuf, err := xml.MarshalIndent(newReturnData, "", "\t")
	assert.Equal(t, nil, err)

	err = newReturnData.Validate()
	assert.Equal(t, nil, err)

	// 6. check xml buffers
	assert.Equal(t, xmlOrgBuf, xmlBuf)
}

func TestManifestXmlTestSuite(t *testing.T) {
	InputXML, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_manifest.xml"))
	assert.Equal(t, nil, err)

	// 1. parse from xml data
	returnData := &IRSSubmissionManifest{}

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
	newReturnData := &IRSSubmissionManifest{}

	err = json.Unmarshal(jsonBuf, newReturnData)
	assert.Equal(t, nil, err)

	// 5. struct to xml buf
	xmlBuf, err := xml.MarshalIndent(newReturnData, "", "\t")
	assert.Equal(t, nil, err)

	err = newReturnData.Validate()
	assert.Equal(t, nil, err)

	// 6. check xml buffers
	assert.Equal(t, xmlOrgBuf, xmlBuf)
}

func Test990FileTestSuite(t *testing.T) {
	returnBuf, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_return.xml"))
	assert.Equal(t, nil, err)

	manifestBuf, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_manifest.xml"))
	assert.Equal(t, nil, err)

	file := &Irs990File{}

	err = xml.Unmarshal(returnBuf, &file.XmlData)
	assert.Equal(t, nil, err)

	file.Manifest = &IRSSubmissionManifest{}
	err = xml.Unmarshal(manifestBuf, file.Manifest)
	assert.Equal(t, nil, err)

	// 2. struct to xml buf
	xmlOrgBuf, err := xml.MarshalIndent(file, "", "\t")
	assert.Equal(t, nil, err)

	// 3. struct to json buf
	jsonBuf, err := json.MarshalIndent(file, "", "\t")
	assert.Equal(t, nil, err)

	// 4. json buf to struct
	newFile := &Irs990File{}

	err = json.Unmarshal(jsonBuf, newFile)
	assert.Equal(t, nil, err)

	// 5. struct to xml buf
	xmlBuf, err := xml.MarshalIndent(newFile, "", "\t")
	assert.Equal(t, nil, err)

	// 6. check xml buffers
	assert.Equal(t, xmlOrgBuf, xmlBuf)

	// 7. validate
	err = newFile.Validate()
	assert.Equal(t, nil, err)

	version := newFile.Version()
	assert.Equal(t, "2019v1.0", version)

	zipData, err := newFile.ZipData()
	assert.Equal(t, nil, err)

	tmpFile, err := ioutil.TempFile("", "test_zip_")
	assert.Equal(t, nil, err)
	err = ioutil.WriteFile(tmpFile.Name(), zipData, 0644)
	assert.Equal(t, nil, err)

	r, err := zip.OpenReader(tmpFile.Name())
	assert.Equal(t, nil, err)

	defer r.Close()
	names := []string{
		filepath.Join("xml", "submission.xml"),
		filepath.Join("manifest", "manifest.xml"),
	}
	for _, f := range r.File {
		assert.Contains(t, names, f.Name)
	}
}