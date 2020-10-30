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

	"github.com/moov-io/1120x/pkg/utils"
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
	instances := []interface{}{
		&IRS990ScheduleAType{}, &IRS990ScheduleA{},
		&IRS990ScheduleBType{}, &IRS990ScheduleB{},
		&IRS990ScheduleCType{}, &IRS990ScheduleC{},
		&IRS990ScheduleDType{}, &IRS990ScheduleD{},
		&IRS990ScheduleEType{}, &IRS990ScheduleE{},
		&IRS990ScheduleFType{}, &IRS990ScheduleF{},
		&IRS990ScheduleGType{}, &IRS990ScheduleG{},
		&IRS990ScheduleHType{}, &IRS990ScheduleH{},
		&IRS990ScheduleIType{}, &IRS990ScheduleI{},
		&IRS990ScheduleJType{}, &IRS990ScheduleJ{},
		&IRS990ScheduleKType{}, &IRS990ScheduleK{},
		&IRS990ScheduleLType{}, &IRS990ScheduleL{},
		&IRS990ScheduleMType{}, &IRS990ScheduleM{},
		&IRS990ScheduleNType{}, &IRS990ScheduleN{},
		&IRS990ScheduleOType{}, &IRS990ScheduleO{},
		&IRS990ScheduleRType{}, &IRS990ScheduleR{},
		&AffiliateListingType{}, &AffiliateListing{},
		&AckNotificationList{},
		&AckNotification{},
		&AccountActivitiesOutsideUSGrpType{},
		&AdjustedNetIncomeGrp{},
		&AffiliateListingGrpType{},
		&AffiliatedScheduleGrp{},
		&AllAffiliatesIncludedInd{},
		&AuditedFinancialStmtAttInd{},
		&BusTrInvolveInterestedPrsnGrp{},
		&CharitableContributionsDetailType{},
		&CollectionUsedOtherPurposesGrp{},
		&ContractorAddress{},
		&ContractorName{},
		&CostingMethodologyUsedGrp{},
		&DiscountedCareOthPercentageGrp{},
		&DisqualifiedPersonExBnftTrGrpType{},
		&DistributableAmountGrp{},
		&DistributionAllocationsGrp{},
		&DistributionsGrp{},
		&FeesForServicesProfFundraising{},
		&FinancialStatementType{},
		&ForeignEntityIdentificationGrpType{},
		&ForeignIndividualsGrantsGrpType{},
		&ForeignItemizedEntryType{},
		&Form990PartIXGroup1Type{},
		&Form990PartVIIGroup1Type{},
		&Form990PartVIIIGroup1Type{},
		&Form990PartVIIIGroup4Type{},
		&Form990PartVIIIGroup7Type{},
		&Form990SchASupportingOrgGrp{},
		&Form990SchAType1SuprtOrgGrp{},
		&Form990SchAType3FuncIntGrp{},
		&Form990SchAType3SprtOrgAllGrp{},
		&Form990SchCPartIIAGroup1Type{},
		&Form990SchCPartIIAGroup2Type{},
		&Form990SchDPartIIIGroup1Type{},
		&Form990SchDPartIXGroup1Type{},
		&Form990SchDPartVGroup1Type{},
		&Form990SchDPartVIGroup2Type{},
		&Form990SchDPartVIIGroup1Type{},
		&Form990SchDPartVIIGroup2Type{},
		&Form990SchDPartVIIIGroup1Type{},
		&Form990SchDPartXGroup1Type{},
		&Form990SchHPartIGroup1Type{},
		&Form990SchHPartIIGroup1Type{},
		&Form990SchMGroup2Type{},
		&Form990SchNGroup1Type{},
		&Form990ScheduleAPartVIGrp{},
		&FreeCareOthPercentageGrp{},
		&FundraiserActivityInfoGrpType{},
		&FundraisingEventInformationGrpType{},
		&GamingInformationGrpType{},
		&GrantsOtherAsstToIndivInUSGrp{},
		&GrantsToOrgOutsideUSGrpType{},
		&GrntAsstBnftInterestedPrsnGrp{},
		&HospitalFacilitiesGrp{},
		&HospitalFcltyPoliciesPrctcGrp{},
		&HospitalNameAndAddressGrpType{},
		&IPAddressType{},
		&IdDisregardedEntitiesGrp{},
		&IdRelatedOrgTxblCorpTrGrp{},
		&IdRelatedOrgTxblPartnershipGrp{},
		&IdRelatedTaxExemptOrgGrp{},
		&LiquidationOfAssetsTableGrp{},
		&LoansBtwnOrgInterestedPrsnGrpType{},
		&ManagementCoAndJntVenturesGrp{},
		&MinimumAssetAmountGrp{},
		&NameAndAddressType{},
		&NameOfInterested{},
		&NonCashPropertyContributionGrpType{},
		&Organization501cInd{},
		&OrganizationBelongsAffltGrpInd{},
		&OthHlthCareFcltsNotHospitalGrp{},
		&OtherForeignAddressType{},
		&OtherUSAddressType{},
		&PersonFullNameType{},
		&PractitionerPINGrp{},
		&ProceduresCorrectiveActionGrpType{},
		&ProgSrvcAccomplishmentActyGrpType{},
		&RecipientTable{},
		&RltdOrgOfficerTrstKeyEmplGrp{},
		&Section527PoliticalOrgGrp{},
		&StrategyAttachedInd{},
		&USItemizedEntryType{},
		&UnrelatedOrgTxblPartnershipGrp{},
		&VehicleDescriptionGrpType{},
		&SupplementalInformationDetail{},
		&SupplementalInformationGrp{},
		&SupportedOrgInformationGrpType{},
		&TaxExemptBondsArbitrageGrpType{},
		&TaxExemptBondsIssuesGrpType{},
		&TaxExemptBondsPrivateBusUseGrpType{},
		&TaxExemptBondsProceedsGrpType{},
		&TotContriRcvdUpTo1000Ind{},
		&TransactionWithControlEntInd{},
		&TransactionsRelatedOrgGrpType{},
		&ValidationErrorListType{},
		&ValidationErrorGrp{},
		&ValidationAlertListType{},
		&ValidationAlertGrp{},
		&SubmissionReceiptGrp{},
		&StatusRecordGrp{},
		&Acknowledgement{},
		&AcknowledgementList{},
		&SubmissionReceiptList{},
		&StatusRecordList{},
	}
	for _, instance := range instances {
		utils.Validate(instance)
	}
}
