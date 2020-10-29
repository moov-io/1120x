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

func TestReturnXmlTest(t *testing.T) {
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

func TestSubmissionManifestTest(t *testing.T) {
	InputXML, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_submission_manifest.xml"))
	assert.Equal(t, nil, err)

	// 1. parse from xml data
	manifest := &IRSSubmissionManifest{}

	err = manifest.Validate()
	assert.NotNil(t, err)

	err = xml.Unmarshal(InputXML, manifest)
	assert.Equal(t, nil, err)

	// 2. struct to xml buf
	xmlOrgBuf, err := xml.MarshalIndent(manifest, "", "\t")
	assert.Equal(t, nil, err)

	// 3. struct to json buf
	jsonBuf, err := json.MarshalIndent(manifest, "", "\t")
	assert.Equal(t, nil, err)

	// 4. json buf to struct
	newManifest := &IRSSubmissionManifest{}

	err = json.Unmarshal(jsonBuf, newManifest)
	assert.Equal(t, nil, err)

	// 5. struct to xml buf
	xmlBuf, err := xml.MarshalIndent(newManifest, "", "\t")
	assert.Equal(t, nil, err)

	err = newManifest.Validate()
	assert.Equal(t, nil, err)

	// 6. check xml buffers
	assert.Equal(t, xmlOrgBuf, xmlBuf)
}

func TestStateSubmissionManifestTest(t *testing.T) {
	InputXML, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_state_submission_manifest.xml"))
	assert.Equal(t, nil, err)

	// 1. parse from xml data
	manifest := &StateSubmissionManifest{}

	err = manifest.Validate()
	assert.NotNil(t, err)

	err = xml.Unmarshal(InputXML, manifest)
	assert.Equal(t, nil, err)

	// 2. struct to xml buf
	xmlOrgBuf, err := xml.MarshalIndent(manifest, "", "\t")
	assert.Equal(t, nil, err)

	// 3. struct to json buf
	jsonBuf, err := json.MarshalIndent(manifest, "", "\t")
	assert.Equal(t, nil, err)

	// 4. json buf to struct
	newManifest := &StateSubmissionManifest{}

	err = json.Unmarshal(jsonBuf, newManifest)
	assert.Equal(t, nil, err)

	// 5. struct to xml buf
	xmlBuf, err := xml.MarshalIndent(newManifest, "", "\t")
	assert.Equal(t, nil, err)

	err = newManifest.Validate()
	assert.Equal(t, nil, err)

	// 6. check xml buffers
	assert.Equal(t, xmlOrgBuf, xmlBuf)
}

func Test990FileTest(t *testing.T) {
	returnBuf, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_return.xml"))
	assert.Equal(t, nil, err)

	manifestBuf, err := ioutil.ReadFile(filepath.Join("..", "..", "test", "testdata", "irs990_submission_manifest.xml"))
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

func TestUnmeaningTest(t *testing.T) {
	ret := &Return{}
	_ = ret.InspectData()

	ret = &Return{ReturnData: ReturnData{
		IRS990:          &IRS990{},
		IRS990ScheduleA: &IRS990ScheduleA{},
		IRS990ScheduleB: &IRS990ScheduleB{},
		IRS990ScheduleC: &IRS990ScheduleC{},
		IRS990ScheduleD: &IRS990ScheduleD{},
		IRS990ScheduleE: &IRS990ScheduleE{},
		IRS990ScheduleF: &IRS990ScheduleF{},
		IRS990ScheduleG: &IRS990ScheduleG{},
		IRS990ScheduleH: &IRS990ScheduleH{},
		IRS990ScheduleI: &IRS990ScheduleI{},
		IRS990ScheduleJ: &IRS990ScheduleJ{},
		IRS990ScheduleK: []IRS990ScheduleK{},
		IRS990ScheduleL: &IRS990ScheduleL{},
		IRS990ScheduleM: &IRS990ScheduleM{},
		IRS990ScheduleN: &IRS990ScheduleN{},
		IRS990ScheduleO: &IRS990ScheduleO{},
		IRS990ScheduleR: &IRS990ScheduleR{},
	}}
	err := ret.Parse([]byte("test"))
	assert.NotNil(t, err)
	_ = ret.Init()
	_ = ret.InspectData()
	_ = ret.ReturnYear()
	_ = ret.Validate()
	_ = ret.String()
	_ = ret.ReturnVersion()
	_ = ret.ReturnType()

	manifest := &StateSubmissionManifest{}
	_ = manifest.Init()
	_, err = manifest.XmlData()
	assert.Equal(t, nil, err)
	_ = manifest.Validate()
	_ = manifest.Init()
	_ = manifest.SubmissionIdentifier()
	manifest.SetSubmissionIdentifier("")

	submission := &IRSSubmissionManifest{}
	_ = submission.Init()
	_, err = submission.XmlData()
	assert.Equal(t, nil, err)
	_ = submission.Validate()
	_ = submission.Init()
	_ = submission.SubmissionIdentifier()
	submission.SetSubmissionIdentifier("")
}

func TestUnusedStructs(t *testing.T) {
	var scheduleAT IRS990ScheduleAType
	err := scheduleAT.Validate()
	assert.Nil(t, err)

	var scheduleBT IRS990ScheduleBType
	err = scheduleBT.Validate()
	assert.Nil(t, err)

	var scheduleC IRS990ScheduleC
	err = scheduleC.Validate()
	assert.NotNil(t, err)

	var scheduleCT IRS990ScheduleCType
	err = scheduleCT.Validate()
	assert.Nil(t, err)

	var scheduleDT IRS990ScheduleDType
	err = scheduleDT.Validate()
	assert.Nil(t, err)

	var scheduleE IRS990ScheduleE
	err = scheduleE.Validate()
	assert.NotNil(t, err)

	var scheduleET IRS990ScheduleEType
	err = scheduleET.Validate()
	assert.Nil(t, err)

	var scheduleF IRS990ScheduleF
	err = scheduleF.Validate()
	assert.NotNil(t, err)

	var scheduleFT IRS990ScheduleFType
	err = scheduleFT.Validate()
	assert.Nil(t, err)

	var scheduleG IRS990ScheduleG
	err = scheduleG.Validate()
	assert.NotNil(t, err)

	var scheduleGT IRS990ScheduleGType
	err = scheduleGT.Validate()
	assert.NotNil(t, err)

	var scheduleH IRS990ScheduleH
	err = scheduleH.Validate()
	assert.NotNil(t, err)

	var scheduleHT IRS990ScheduleHType
	err = scheduleHT.Validate()
	assert.Nil(t, err)

	var scheduleI IRS990ScheduleI
	err = scheduleI.Validate()
	assert.NotNil(t, err)

	var scheduleIT IRS990ScheduleIType
	err = scheduleIT.Validate()
	assert.Nil(t, err)

	var scheduleJ IRS990ScheduleJ
	err = scheduleJ.Validate()
	assert.NotNil(t, err)

	var scheduleJT IRS990ScheduleJType
	err = scheduleJT.Validate()
	assert.Nil(t, err)

	var scheduleK IRS990ScheduleK
	err = scheduleK.Validate()
	assert.NotNil(t, err)

	var scheduleKT IRS990ScheduleKType
	err = scheduleKT.Validate()
	assert.Nil(t, err)

	var scheduleL IRS990ScheduleL
	err = scheduleL.Validate()
	assert.NotNil(t, err)

	var scheduleLT IRS990ScheduleLType
	err = scheduleLT.Validate()
	assert.Nil(t, err)

	var scheduleM IRS990ScheduleM
	err = scheduleM.Validate()
	assert.NotNil(t, err)

	var scheduleMT IRS990ScheduleMType
	err = scheduleMT.Validate()
	assert.Nil(t, err)

	var scheduleN IRS990ScheduleN
	err = scheduleN.Validate()
	assert.NotNil(t, err)

	var scheduleNT IRS990ScheduleNType
	err = scheduleNT.Validate()
	assert.Nil(t, err)

	var scheduleO IRS990ScheduleO
	err = scheduleO.Validate()
	assert.NotNil(t, err)

	var scheduleOT IRS990ScheduleOType
	err = scheduleOT.Validate()
	assert.Nil(t, err)

	var scheduleR IRS990ScheduleR
	err = scheduleR.Validate()
	assert.NotNil(t, err)

	var scheduleRT IRS990ScheduleRType
	err = scheduleRT.Validate()
	assert.Nil(t, err)

	var affiliate AffiliateListing
	err = affiliate.Validate()
	assert.NotNil(t, err)

	var affiliateT AffiliateListingType
	err = affiliateT.Validate()
	assert.Nil(t, err)

	var acknolist AckNotificationList
	err = acknolist.Validate()
	assert.Nil(t, err)

	var ackno AckNotification
	err = ackno.Validate()
	assert.NotNil(t, err)

	var usGrp AccountActivitiesOutsideUSGrpType
	err = usGrp.Validate()
	assert.Nil(t, err)

	var incomeGrp AdjustedNetIncomeGrp
	err = incomeGrp.Validate()
	assert.Nil(t, err)
}
