// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import "github.com/moov-io/1120x/pkg/utils"

type IRS990ScheduleA struct {
	ChurchInd                      CheckboxType                     `xml:"ChurchInd"`
	SchoolInd                      CheckboxType                     `xml:"SchoolInd"`
	HospitalInd                    CheckboxType                     `xml:"HospitalInd"`
	MedicalResearchOrganizationInd CheckboxType                     `xml:"MedicalResearchOrganizationInd"`
	HospitalNameAndAddressGrp      []HospitalNameAndAddressGrpType  `xml:"HospitalNameAndAddressGrp"`
	CollegeOrganizationInd         CheckboxType                     `xml:"CollegeOrganizationInd"`
	GovernmentalUnitInd            CheckboxType                     `xml:"GovernmentalUnitInd"`
	PublicOrganization170Ind       CheckboxType                     `xml:"PublicOrganization170Ind"`
	CommunityTrustInd              CheckboxType                     `xml:"CommunityTrustInd"`
	PubliclySupportedOrg509a2Ind   CheckboxType                     `xml:"PubliclySupportedOrg509a2Ind"`
	TestPublicSafetyInd            CheckboxType                     `xml:"TestPublicSafetyInd"`
	SupportingOrganization509a3Ind CheckboxType                     `xml:"SupportingOrganization509a3Ind"`
	SupportingOrgType1Ind          CheckboxType                     `xml:"SupportingOrgType1Ind"`
	SupportingOrgType2Ind          CheckboxType                     `xml:"SupportingOrgType2Ind"`
	SupportingOrgType3FuncIntInd   CheckboxType                     `xml:"SupportingOrgType3FuncIntInd"`
	SupportingOrgType3NonFuncInd   CheckboxType                     `xml:"SupportingOrgType3NonFuncInd"`
	IRSWrittenDeterminationInd     CheckboxType                     `xml:"IRSWrittenDeterminationInd,omitempty" json:",omitempty"`
	SupportedOrganizationsCnt      int                              `xml:"SupportedOrganizationsCnt,omitempty" json:",omitempty"`
	SupportedOrgInformationGrp     []SupportedOrgInformationGrpType `xml:"SupportedOrgInformationGrp,omitempty" json:",omitempty"`
	SupportedOrganizationsTotalCnt int                              `xml:"SupportedOrganizationsTotalCnt,omitempty" json:",omitempty"`
	SupportSumAmt                  int                              `xml:"SupportSumAmt,omitempty" json:",omitempty"`
	OtherSupportSumAmt             int                              `xml:"OtherSupportSumAmt,omitempty" json:",omitempty"`
	GiftsGrantsContriRcvd170Grp    *Form990SchAPartIIGroup1Type     `xml:"GiftsGrantsContriRcvd170Grp,omitempty" json:",omitempty"`
	TaxRevLeviedOrgnztnlBnft170Grp *Form990SchAPartIIGroup1Type     `xml:"TaxRevLeviedOrgnztnlBnft170Grp,omitempty" json:",omitempty"`
	GovtFurnSrvcFcltsVl170Grp      *Form990SchAPartIIGroup1Type     `xml:"GovtFurnSrvcFcltsVl170Grp,omitempty" json:",omitempty"`
	TotalCalendarYear170Grp        *Form990SchAPartIIGroup1Type     `xml:"TotalCalendarYear170Grp,omitempty" json:",omitempty"`
	SubstantialContributorsTotAmt  int                              `xml:"SubstantialContributorsTotAmt,omitempty" json:",omitempty"`
	PublicSupportTotal170Amt       int                              `xml:"PublicSupportTotal170Amt,omitempty" json:",omitempty"`
	GrossInvestmentIncome170Grp    *Form990SchAPartIIGroup1Type     `xml:"GrossInvestmentIncome170Grp,omitempty" json:",omitempty"`
	UnrelatedBusinessNetIncm170Grp *Form990SchAPartIIGroup1Type     `xml:"UnrelatedBusinessNetIncm170Grp,omitempty" json:",omitempty"`
	OtherIncome170Grp              *Form990SchAPartIIGroup1Type     `xml:"OtherIncome170Grp,omitempty" json:",omitempty"`
	TotalSupportAmt                int                              `xml:"TotalSupportAmt,omitempty" json:",omitempty"`
	GrossReceiptsRltdActivitiesAmt int                              `xml:"GrossReceiptsRltdActivitiesAmt,omitempty" json:",omitempty"`
	First5Years170Ind              CheckboxType                     `xml:"First5Years170Ind,omitempty" json:",omitempty"`
	PublicSupportCY170Pct          float64                          `xml:"PublicSupportCY170Pct,omitempty" json:",omitempty"`
	PublicSupportPY170Pct          float64                          `xml:"PublicSupportPY170Pct,omitempty" json:",omitempty"`
	ThirtyThrPctSuprtTestsCY170Ind CheckboxType                     `xml:"ThirtyThrPctSuprtTestsCY170Ind,omitempty" json:",omitempty"`
	ThirtyThrPctSuprtTestsPY170Ind CheckboxType                     `xml:"ThirtyThrPctSuprtTestsPY170Ind,omitempty" json:",omitempty"`
	TenPctFactsCrcmstncsTestCYInd  CheckboxType                     `xml:"TenPctFactsCrcmstncsTestCYInd,omitempty" json:",omitempty"`
	TenPctFactsCrcmstncsTestPYInd  CheckboxType                     `xml:"TenPctFactsCrcmstncsTestPYInd,omitempty" json:",omitempty"`
	PrivateFoundation170Ind        CheckboxType                     `xml:"PrivateFoundation170Ind,omitempty" json:",omitempty"`
	GiftsGrantsContrisRcvd509Grp   *Form990SchAPartIIGroup1Type     `xml:"GiftsGrantsContrisRcvd509Grp,omitempty" json:",omitempty"`
	GrossReceiptsAdmissionsGrp     *Form990SchAPartIIGroup1Type     `xml:"GrossReceiptsAdmissionsGrp,omitempty" json:",omitempty"`
	GrossReceiptsNonUnrltBusGrp    *Form990SchAPartIIGroup1Type     `xml:"GrossReceiptsNonUnrltBusGrp,omitempty" json:",omitempty"`
	TaxRevLeviedOrgnztnlBnft509Grp *Form990SchAPartIIGroup1Type     `xml:"TaxRevLeviedOrgnztnlBnft509Grp,omitempty" json:",omitempty"`
	GovtFurnSrvcFcltsVl509Grp      *Form990SchAPartIIGroup1Type     `xml:"GovtFurnSrvcFcltsVl509Grp,omitempty" json:",omitempty"`
	Total509Grp                    *Form990SchAPartIIGroup1Type     `xml:"Total509Grp,omitempty" json:",omitempty"`
	AmountsRcvdDsqlfyPersonGrp     *Form990SchAPartIIGroup1Type     `xml:"AmountsRcvdDsqlfyPersonGrp,omitempty" json:",omitempty"`
	SubstantialContributorsAmtGrp  *Form990SchAPartIIGroup1Type     `xml:"SubstantialContributorsAmtGrp,omitempty" json:",omitempty"`
	SubstAndDsqlfyPrsnsTotGrp      *Form990SchAPartIIGroup1Type     `xml:"SubstAndDsqlfyPrsnsTotGrp,omitempty" json:",omitempty"`
	PublicSupportTotal509Amt       int                              `xml:"PublicSupportTotal509Amt,omitempty" json:",omitempty"`
	GrossInvestmentIncome509Grp    *Form990SchAPartIIGroup1Type     `xml:"GrossInvestmentIncome509Grp,omitempty" json:",omitempty"`
	Post1975UBTIGrp                *Form990SchAPartIIGroup1Type     `xml:"Post1975UBTIGrp,omitempty" json:",omitempty"`
	InvestmentIncomeAndUBTIGrp     *Form990SchAPartIIGroup1Type     `xml:"InvestmentIncomeAndUBTIGrp,omitempty" json:",omitempty"`
	NetIncomeFromOtherUBIGrp       *Form990SchAPartIIGroup1Type     `xml:"NetIncomeFromOtherUBIGrp,omitempty" json:",omitempty"`
	OtherIncome509Grp              *Form990SchAPartIIGroup1Type     `xml:"OtherIncome509Grp,omitempty" json:",omitempty"`
	TotalSupportCalendarYearGrp    *Form990SchAPartIIGroup1Type     `xml:"TotalSupportCalendarYearGrp,omitempty" json:",omitempty"`
	First5Years509Ind              CheckboxType                     `xml:"First5Years509Ind,omitempty" json:",omitempty"`
	PublicSupportCY509Pct          float64                          `xml:"PublicSupportCY509Pct,omitempty" json:",omitempty"`
	PublicSupportPY509Pct          float64                          `xml:"PublicSupportPY509Pct,omitempty" json:",omitempty"`
	InvestmentIncomeCYPct          float64                          `xml:"InvestmentIncomeCYPct,omitempty" json:",omitempty"`
	InvestmentIncomePYPct          float64                          `xml:"InvestmentIncomePYPct,omitempty" json:",omitempty"`
	ThirtyThrPctSuprtTestsCY509Ind CheckboxType                     `xml:"ThirtyThrPctSuprtTestsCY509Ind,omitempty" json:",omitempty"`
	ThirtyThrPctSuprtTestsPY509Ind CheckboxType                     `xml:"ThirtyThrPctSuprtTestsPY509Ind,omitempty" json:",omitempty"`
	PrivateFoundation509Ind        CheckboxType                     `xml:"PrivateFoundation509Ind,omitempty" json:",omitempty"`
	Form990SchASupportingOrgGrp    *Form990SchASupportingOrgGrp     `xml:"Form990SchASupportingOrgGrp,omitempty" json:",omitempty"`
	Form990SchAType1SuprtOrgGrp    *Form990SchAType1SuprtOrgGrp     `xml:"Form990SchAType1SuprtOrgGrp,omitempty" json:",omitempty"`
	MajorityDirTrstSupportedOrgInd bool                             `xml:"MajorityDirTrstSupportedOrgInd,omitempty" json:",omitempty"`
	Form990SchAType3SprtOrgAllGrp  *Form990SchAType3SprtOrgAllGrp   `xml:"Form990SchAType3SprtOrgAllGrp,omitempty" json:",omitempty"`
	Form990SchAType3FuncIntGrp     *Form990SchAType3FuncIntGrp      `xml:"Form990SchAType3FuncIntGrp,omitempty" json:",omitempty"`
	TrustIntegralPartTestInd       CheckboxType                     `xml:"TrustIntegralPartTestInd,omitempty" json:",omitempty"`
	AdjustedNetIncomeGrp           *AdjustedNetIncomeGrp            `xml:"AdjustedNetIncomeGrp,omitempty" json:",omitempty"`
	MinimumAssetAmountGrp          *MinimumAssetAmountGrp           `xml:"MinimumAssetAmountGrp,omitempty" json:",omitempty"`
	DistributableAmountGrp         *DistributableAmountGrp          `xml:"DistributableAmountGrp,omitempty" json:",omitempty"`
	DistributionsGrp               *DistributionsGrp                `xml:"DistributionsGrp,omitempty" json:",omitempty"`
	DistributionAllocationsGrp     *DistributionAllocationsGrp      `xml:"DistributionAllocationsGrp,omitempty" json:",omitempty"`
	FactsAndCircumstancesTestTxt   string                           `xml:"FactsAndCircumstancesTestTxt,omitempty" json:",omitempty"`
	Form990ScheduleAPartVIGrp      []Form990ScheduleAPartVIGrp      `xml:"Form990ScheduleAPartVIGrp,omitempty" json:",omitempty"`
	DocumentId                     IdType                           `xml:"documentId,attr"`
	SoftwareId                     SoftwareIdType                   `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                           `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                           `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleA) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule A
type IRS990ScheduleAType struct {
	ChurchInd                      CheckboxType                     `xml:"ChurchInd"`
	SchoolInd                      CheckboxType                     `xml:"SchoolInd"`
	HospitalInd                    CheckboxType                     `xml:"HospitalInd"`
	MedicalResearchOrganizationInd CheckboxType                     `xml:"MedicalResearchOrganizationInd"`
	HospitalNameAndAddressGrp      []HospitalNameAndAddressGrpType  `xml:"HospitalNameAndAddressGrp"`
	CollegeOrganizationInd         CheckboxType                     `xml:"CollegeOrganizationInd"`
	GovernmentalUnitInd            CheckboxType                     `xml:"GovernmentalUnitInd"`
	PublicOrganization170Ind       CheckboxType                     `xml:"PublicOrganization170Ind"`
	CommunityTrustInd              CheckboxType                     `xml:"CommunityTrustInd"`
	PubliclySupportedOrg509a2Ind   CheckboxType                     `xml:"PubliclySupportedOrg509a2Ind"`
	TestPublicSafetyInd            CheckboxType                     `xml:"TestPublicSafetyInd"`
	SupportingOrganization509a3Ind CheckboxType                     `xml:"SupportingOrganization509a3Ind"`
	SupportingOrgType1Ind          CheckboxType                     `xml:"SupportingOrgType1Ind"`
	SupportingOrgType2Ind          CheckboxType                     `xml:"SupportingOrgType2Ind"`
	SupportingOrgType3FuncIntInd   CheckboxType                     `xml:"SupportingOrgType3FuncIntInd"`
	SupportingOrgType3NonFuncInd   CheckboxType                     `xml:"SupportingOrgType3NonFuncInd"`
	IRSWrittenDeterminationInd     CheckboxType                     `xml:"IRSWrittenDeterminationInd,omitempty" json:",omitempty"`
	SupportedOrganizationsCnt      int                              `xml:"SupportedOrganizationsCnt,omitempty" json:",omitempty"`
	SupportedOrgInformationGrp     []SupportedOrgInformationGrpType `xml:"SupportedOrgInformationGrp,omitempty" json:",omitempty"`
	SupportedOrganizationsTotalCnt int                              `xml:"SupportedOrganizationsTotalCnt,omitempty" json:",omitempty"`
	SupportSumAmt                  int                              `xml:"SupportSumAmt,omitempty" json:",omitempty"`
	OtherSupportSumAmt             int                              `xml:"OtherSupportSumAmt,omitempty" json:",omitempty"`
	GiftsGrantsContriRcvd170Grp    *Form990SchAPartIIGroup1Type     `xml:"GiftsGrantsContriRcvd170Grp,omitempty" json:",omitempty"`
	TaxRevLeviedOrgnztnlBnft170Grp *Form990SchAPartIIGroup1Type     `xml:"TaxRevLeviedOrgnztnlBnft170Grp,omitempty" json:",omitempty"`
	GovtFurnSrvcFcltsVl170Grp      *Form990SchAPartIIGroup1Type     `xml:"GovtFurnSrvcFcltsVl170Grp,omitempty" json:",omitempty"`
	TotalCalendarYear170Grp        *Form990SchAPartIIGroup1Type     `xml:"TotalCalendarYear170Grp,omitempty" json:",omitempty"`
	SubstantialContributorsTotAmt  int                              `xml:"SubstantialContributorsTotAmt,omitempty" json:",omitempty"`
	PublicSupportTotal170Amt       int                              `xml:"PublicSupportTotal170Amt,omitempty" json:",omitempty"`
	GrossInvestmentIncome170Grp    *Form990SchAPartIIGroup1Type     `xml:"GrossInvestmentIncome170Grp,omitempty" json:",omitempty"`
	UnrelatedBusinessNetIncm170Grp *Form990SchAPartIIGroup1Type     `xml:"UnrelatedBusinessNetIncm170Grp,omitempty" json:",omitempty"`
	OtherIncome170Grp              *Form990SchAPartIIGroup1Type     `xml:"OtherIncome170Grp,omitempty" json:",omitempty"`
	TotalSupportAmt                int                              `xml:"TotalSupportAmt,omitempty" json:",omitempty"`
	GrossReceiptsRltdActivitiesAmt int                              `xml:"GrossReceiptsRltdActivitiesAmt,omitempty" json:",omitempty"`
	First5Years170Ind              CheckboxType                     `xml:"First5Years170Ind,omitempty" json:",omitempty"`
	PublicSupportCY170Pct          float64                          `xml:"PublicSupportCY170Pct,omitempty" json:",omitempty"`
	PublicSupportPY170Pct          float64                          `xml:"PublicSupportPY170Pct,omitempty" json:",omitempty"`
	ThirtyThrPctSuprtTestsCY170Ind CheckboxType                     `xml:"ThirtyThrPctSuprtTestsCY170Ind,omitempty" json:",omitempty"`
	ThirtyThrPctSuprtTestsPY170Ind CheckboxType                     `xml:"ThirtyThrPctSuprtTestsPY170Ind,omitempty" json:",omitempty"`
	TenPctFactsCrcmstncsTestCYInd  CheckboxType                     `xml:"TenPctFactsCrcmstncsTestCYInd,omitempty" json:",omitempty"`
	TenPctFactsCrcmstncsTestPYInd  CheckboxType                     `xml:"TenPctFactsCrcmstncsTestPYInd,omitempty" json:",omitempty"`
	PrivateFoundation170Ind        CheckboxType                     `xml:"PrivateFoundation170Ind,omitempty" json:",omitempty"`
	GiftsGrantsContrisRcvd509Grp   *Form990SchAPartIIGroup1Type     `xml:"GiftsGrantsContrisRcvd509Grp,omitempty" json:",omitempty"`
	GrossReceiptsAdmissionsGrp     *Form990SchAPartIIGroup1Type     `xml:"GrossReceiptsAdmissionsGrp,omitempty" json:",omitempty"`
	GrossReceiptsNonUnrltBusGrp    *Form990SchAPartIIGroup1Type     `xml:"GrossReceiptsNonUnrltBusGrp,omitempty" json:",omitempty"`
	TaxRevLeviedOrgnztnlBnft509Grp *Form990SchAPartIIGroup1Type     `xml:"TaxRevLeviedOrgnztnlBnft509Grp,omitempty" json:",omitempty"`
	GovtFurnSrvcFcltsVl509Grp      *Form990SchAPartIIGroup1Type     `xml:"GovtFurnSrvcFcltsVl509Grp,omitempty" json:",omitempty"`
	Total509Grp                    *Form990SchAPartIIGroup1Type     `xml:"Total509Grp,omitempty" json:",omitempty"`
	AmountsRcvdDsqlfyPersonGrp     *Form990SchAPartIIGroup1Type     `xml:"AmountsRcvdDsqlfyPersonGrp,omitempty" json:",omitempty"`
	SubstantialContributorsAmtGrp  *Form990SchAPartIIGroup1Type     `xml:"SubstantialContributorsAmtGrp,omitempty" json:",omitempty"`
	SubstAndDsqlfyPrsnsTotGrp      *Form990SchAPartIIGroup1Type     `xml:"SubstAndDsqlfyPrsnsTotGrp,omitempty" json:",omitempty"`
	PublicSupportTotal509Amt       int                              `xml:"PublicSupportTotal509Amt,omitempty" json:",omitempty"`
	GrossInvestmentIncome509Grp    *Form990SchAPartIIGroup1Type     `xml:"GrossInvestmentIncome509Grp,omitempty" json:",omitempty"`
	Post1975UBTIGrp                *Form990SchAPartIIGroup1Type     `xml:"Post1975UBTIGrp,omitempty" json:",omitempty"`
	InvestmentIncomeAndUBTIGrp     *Form990SchAPartIIGroup1Type     `xml:"InvestmentIncomeAndUBTIGrp,omitempty" json:",omitempty"`
	NetIncomeFromOtherUBIGrp       *Form990SchAPartIIGroup1Type     `xml:"NetIncomeFromOtherUBIGrp,omitempty" json:",omitempty"`
	OtherIncome509Grp              *Form990SchAPartIIGroup1Type     `xml:"OtherIncome509Grp,omitempty" json:",omitempty"`
	TotalSupportCalendarYearGrp    *Form990SchAPartIIGroup1Type     `xml:"TotalSupportCalendarYearGrp,omitempty" json:",omitempty"`
	First5Years509Ind              CheckboxType                     `xml:"First5Years509Ind,omitempty" json:",omitempty"`
	PublicSupportCY509Pct          float64                          `xml:"PublicSupportCY509Pct,omitempty" json:",omitempty"`
	PublicSupportPY509Pct          float64                          `xml:"PublicSupportPY509Pct,omitempty" json:",omitempty"`
	InvestmentIncomeCYPct          float64                          `xml:"InvestmentIncomeCYPct,omitempty" json:",omitempty"`
	InvestmentIncomePYPct          float64                          `xml:"InvestmentIncomePYPct,omitempty" json:",omitempty"`
	ThirtyThrPctSuprtTestsCY509Ind CheckboxType                     `xml:"ThirtyThrPctSuprtTestsCY509Ind,omitempty" json:",omitempty"`
	ThirtyThrPctSuprtTestsPY509Ind CheckboxType                     `xml:"ThirtyThrPctSuprtTestsPY509Ind,omitempty" json:",omitempty"`
	PrivateFoundation509Ind        CheckboxType                     `xml:"PrivateFoundation509Ind,omitempty" json:",omitempty"`
	Form990SchASupportingOrgGrp    *Form990SchASupportingOrgGrp     `xml:"Form990SchASupportingOrgGrp,omitempty" json:",omitempty"`
	Form990SchAType1SuprtOrgGrp    *Form990SchAType1SuprtOrgGrp     `xml:"Form990SchAType1SuprtOrgGrp,omitempty" json:",omitempty"`
	MajorityDirTrstSupportedOrgInd bool                             `xml:"MajorityDirTrstSupportedOrgInd,omitempty" json:",omitempty"`
	Form990SchAType3SprtOrgAllGrp  *Form990SchAType3SprtOrgAllGrp   `xml:"Form990SchAType3SprtOrgAllGrp,omitempty" json:",omitempty"`
	Form990SchAType3FuncIntGrp     *Form990SchAType3FuncIntGrp      `xml:"Form990SchAType3FuncIntGrp,omitempty" json:",omitempty"`
	TrustIntegralPartTestInd       CheckboxType                     `xml:"TrustIntegralPartTestInd,omitempty" json:",omitempty"`
	AdjustedNetIncomeGrp           *AdjustedNetIncomeGrp            `xml:"AdjustedNetIncomeGrp,omitempty" json:",omitempty"`
	MinimumAssetAmountGrp          *MinimumAssetAmountGrp           `xml:"MinimumAssetAmountGrp,omitempty" json:",omitempty"`
	DistributableAmountGrp         *DistributableAmountGrp          `xml:"DistributableAmountGrp,omitempty" json:",omitempty"`
	DistributionsGrp               *DistributionsGrp                `xml:"DistributionsGrp,omitempty" json:",omitempty"`
	DistributionAllocationsGrp     *DistributionAllocationsGrp      `xml:"DistributionAllocationsGrp,omitempty" json:",omitempty"`
	FactsAndCircumstancesTestTxt   string                           `xml:"FactsAndCircumstancesTestTxt,omitempty" json:",omitempty"`
	Form990ScheduleAPartVIGrp      []Form990ScheduleAPartVIGrp      `xml:"Form990ScheduleAPartVIGrp,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleAType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleB struct {
	Organization501cInd            Organization501cInd                  `xml:"Organization501cInd"`
	Organization4947a1NotPFInd     CheckboxType                         `xml:"Organization4947a1NotPFInd"`
	Organization527Ind             CheckboxType                         `xml:"Organization527Ind"`
	Organization501c3ExemptPFInd   CheckboxType                         `xml:"Organization501c3ExemptPFInd"`
	Organization4947a1TrtdPFInd    CheckboxType                         `xml:"Organization4947a1TrtdPFInd"`
	Organization501c3TaxablePFInd  CheckboxType                         `xml:"Organization501c3TaxablePFInd"`
	GeneralRuleInd                 CheckboxType                         `xml:"GeneralRuleInd,omitempty" json:",omitempty"`
	SpclRuleMetOne3rdSuprtTestInd  CheckboxType                         `xml:"SpclRuleMetOne3rdSuprtTestInd,omitempty" json:",omitempty"`
	TotContriRcvdMore1000Ind       CheckboxType                         `xml:"TotContriRcvdMore1000Ind,omitempty" json:",omitempty"`
	TotContriRcvdUpTo1000Ind       *TotContriRcvdUpTo1000Ind            `xml:"TotContriRcvdUpTo1000Ind,omitempty" json:",omitempty"`
	ContributorInformationGrp      []ContributorInformationGrpType      `xml:"ContributorInformationGrp,omitempty" json:",omitempty"`
	NonCashPropertyContributionGrp []NonCashPropertyContributionGrpType `xml:"NonCashPropertyContributionGrp,omitempty" json:",omitempty"`
	TotalUnder1000ContributionsAmt int                                  `xml:"TotalUnder1000ContributionsAmt,omitempty" json:",omitempty"`
	CharitableContributionsDetail  []CharitableContributionsDetailType  `xml:"CharitableContributionsDetail,omitempty" json:",omitempty"`
	DocumentId                     IdType                               `xml:"documentId,attr"`
	SoftwareId                     SoftwareIdType                       `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                               `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                               `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleB) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule B
type IRS990ScheduleBType struct {
	Organization501cInd            Organization501cInd                  `xml:"Organization501cInd"`
	Organization4947a1NotPFInd     CheckboxType                         `xml:"Organization4947a1NotPFInd"`
	Organization527Ind             CheckboxType                         `xml:"Organization527Ind"`
	Organization501c3ExemptPFInd   CheckboxType                         `xml:"Organization501c3ExemptPFInd"`
	Organization4947a1TrtdPFInd    CheckboxType                         `xml:"Organization4947a1TrtdPFInd"`
	Organization501c3TaxablePFInd  CheckboxType                         `xml:"Organization501c3TaxablePFInd"`
	GeneralRuleInd                 CheckboxType                         `xml:"GeneralRuleInd,omitempty" json:",omitempty"`
	SpclRuleMetOne3rdSuprtTestInd  CheckboxType                         `xml:"SpclRuleMetOne3rdSuprtTestInd,omitempty" json:",omitempty"`
	TotContriRcvdMore1000Ind       CheckboxType                         `xml:"TotContriRcvdMore1000Ind,omitempty" json:",omitempty"`
	TotContriRcvdUpTo1000Ind       *TotContriRcvdUpTo1000Ind            `xml:"TotContriRcvdUpTo1000Ind,omitempty" json:",omitempty"`
	ContributorInformationGrp      []ContributorInformationGrpType      `xml:"ContributorInformationGrp,omitempty" json:",omitempty"`
	NonCashPropertyContributionGrp []NonCashPropertyContributionGrpType `xml:"NonCashPropertyContributionGrp,omitempty" json:",omitempty"`
	TotalUnder1000ContributionsAmt int                                  `xml:"TotalUnder1000ContributionsAmt,omitempty" json:",omitempty"`
	CharitableContributionsDetail  []CharitableContributionsDetailType  `xml:"CharitableContributionsDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleBType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleC struct {
	PoliticalExpendituresAmt       int                             `xml:"PoliticalExpendituresAmt,omitempty" json:",omitempty"`
	VolunteerHoursCnt              int                             `xml:"VolunteerHoursCnt,omitempty" json:",omitempty"`
	Section4955OrganizationTaxAmt  int                             `xml:"Section4955OrganizationTaxAmt,omitempty" json:",omitempty"`
	Section4955ManagersTaxAmt      int                             `xml:"Section4955ManagersTaxAmt,omitempty" json:",omitempty"`
	Form4720FiledSection4955TaxInd bool                            `xml:"Form4720FiledSection4955TaxInd,omitempty" json:",omitempty"`
	CorrectionMadeInd              bool                            `xml:"CorrectionMadeInd,omitempty" json:",omitempty"`
	Expended527ActivitiesAmt       int                             `xml:"Expended527ActivitiesAmt,omitempty" json:",omitempty"`
	InternalFundsContributedAmt    int                             `xml:"InternalFundsContributedAmt,omitempty" json:",omitempty"`
	TotalExemptFunctionExpendAmt   int                             `xml:"TotalExemptFunctionExpendAmt,omitempty" json:",omitempty"`
	Form1120POLFiledInd            bool                            `xml:"Form1120POLFiledInd,omitempty" json:",omitempty"`
	Section527PoliticalOrgGrp      []Section527PoliticalOrgGrp     `xml:"Section527PoliticalOrgGrp,omitempty" json:",omitempty"`
	OrganizationBelongsAffltGrpInd *OrganizationBelongsAffltGrpInd `xml:"OrganizationBelongsAffltGrpInd,omitempty" json:",omitempty"`
	LimitedControlProvisionsAppInd CheckboxType                    `xml:"LimitedControlProvisionsAppInd,omitempty" json:",omitempty"`
	TotalGrassrootsLobbyingGrp     *Form990SchCPartIIAGroup1Type   `xml:"TotalGrassrootsLobbyingGrp,omitempty" json:",omitempty"`
	TotalDirectLobbyingGrp         *Form990SchCPartIIAGroup1Type   `xml:"TotalDirectLobbyingGrp,omitempty" json:",omitempty"`
	TotalLobbyingExpendGrp         *Form990SchCPartIIAGroup1Type   `xml:"TotalLobbyingExpendGrp,omitempty" json:",omitempty"`
	OtherExemptPurposeExpendGrp    *Form990SchCPartIIAGroup1Type   `xml:"OtherExemptPurposeExpendGrp,omitempty" json:",omitempty"`
	TotalExemptPurposeExpendGrp    *Form990SchCPartIIAGroup1Type   `xml:"TotalExemptPurposeExpendGrp,omitempty" json:",omitempty"`
	LobbyingNontaxableAmountGrp    *Form990SchCPartIIAGroup1Type   `xml:"LobbyingNontaxableAmountGrp,omitempty" json:",omitempty"`
	GrassrootsNontaxableGrp        *Form990SchCPartIIAGroup1Type   `xml:"GrassrootsNontaxableGrp,omitempty" json:",omitempty"`
	TotLbbyngGrassrootMnsNonTxGrp  *Form990SchCPartIIAGroup1Type   `xml:"TotLbbyngGrassrootMnsNonTxGrp,omitempty" json:",omitempty"`
	TotLbbyExpendMnsLbbyngNonTxGrp *Form990SchCPartIIAGroup1Type   `xml:"TotLbbyExpendMnsLbbyngNonTxGrp,omitempty" json:",omitempty"`
	Form4720FiledInd               bool                            `xml:"Form4720FiledInd,omitempty" json:",omitempty"`
	AvgLobbyingNontaxableAmountGrp *Form990SchCPartIIAGroup2Type   `xml:"AvgLobbyingNontaxableAmountGrp,omitempty" json:",omitempty"`
	LobbyingCeilingAmt             int                             `xml:"LobbyingCeilingAmt,omitempty" json:",omitempty"`
	AvgTotalLobbyingExpendGrp      *Form990SchCPartIIAGroup2Type   `xml:"AvgTotalLobbyingExpendGrp,omitempty" json:",omitempty"`
	AvgGrassrootsNontaxableGrp     *Form990SchCPartIIAGroup2Type   `xml:"AvgGrassrootsNontaxableGrp,omitempty" json:",omitempty"`
	GrassrootsCeilingAmt           int                             `xml:"GrassrootsCeilingAmt,omitempty" json:",omitempty"`
	AvgGrassrootsLobbyingExpendGrp *Form990SchCPartIIAGroup2Type   `xml:"AvgGrassrootsLobbyingExpendGrp,omitempty" json:",omitempty"`
	VolunteersInd                  bool                            `xml:"VolunteersInd,omitempty" json:",omitempty"`
	PaidStaffOrManagementInd       bool                            `xml:"PaidStaffOrManagementInd,omitempty" json:",omitempty"`
	MediaAdvertisementsInd         bool                            `xml:"MediaAdvertisementsInd,omitempty" json:",omitempty"`
	MediaAdvertisementsAmt         int                             `xml:"MediaAdvertisementsAmt,omitempty" json:",omitempty"`
	MailingsMembersInd             bool                            `xml:"MailingsMembersInd,omitempty" json:",omitempty"`
	MailingsMembersAmt             int                             `xml:"MailingsMembersAmt,omitempty" json:",omitempty"`
	PublicationsOrBroadcastInd     bool                            `xml:"PublicationsOrBroadcastInd,omitempty" json:",omitempty"`
	PublicationsOrBroadcastAmt     int                             `xml:"PublicationsOrBroadcastAmt,omitempty" json:",omitempty"`
	GrantsOtherOrganizationsInd    bool                            `xml:"GrantsOtherOrganizationsInd,omitempty" json:",omitempty"`
	GrantsOtherOrganizationsAmt    int                             `xml:"GrantsOtherOrganizationsAmt,omitempty" json:",omitempty"`
	DirectContactLegislatorsInd    bool                            `xml:"DirectContactLegislatorsInd,omitempty" json:",omitempty"`
	DirectContactLegislatorsAmt    int                             `xml:"DirectContactLegislatorsAmt,omitempty" json:",omitempty"`
	RalliesDemonstrationsInd       bool                            `xml:"RalliesDemonstrationsInd,omitempty" json:",omitempty"`
	RalliesDemonstrationsAmt       int                             `xml:"RalliesDemonstrationsAmt,omitempty" json:",omitempty"`
	OtherActivitiesInd             bool                            `xml:"OtherActivitiesInd,omitempty" json:",omitempty"`
	OtherActivitiesAmt             int                             `xml:"OtherActivitiesAmt,omitempty" json:",omitempty"`
	TotalLobbyingExpendituresAmt   int                             `xml:"TotalLobbyingExpendituresAmt,omitempty" json:",omitempty"`
	NotDescribedSection501c3Ind    bool                            `xml:"NotDescribedSection501c3Ind,omitempty" json:",omitempty"`
	Tax4912Amt                     int                             `xml:"Tax4912Amt,omitempty" json:",omitempty"`
	Managers4912TaxAmt             int                             `xml:"Managers4912TaxAmt,omitempty" json:",omitempty"`
	Form4720Filed4912TaxInd        bool                            `xml:"Form4720Filed4912TaxInd,omitempty" json:",omitempty"`
	SubstantiallyAllDuesNondedInd  bool                            `xml:"SubstantiallyAllDuesNondedInd,omitempty" json:",omitempty"`
	OnlyInHouseLobbyingInd         bool                            `xml:"OnlyInHouseLobbyingInd,omitempty" json:",omitempty"`
	AgreeCarryoverPriorYearInd     bool                            `xml:"AgreeCarryoverPriorYearInd,omitempty" json:",omitempty"`
	DuesAssessmentsAmt             int                             `xml:"DuesAssessmentsAmt,omitempty" json:",omitempty"`
	NonDeductibleLbbyngPltclCYAmt  int                             `xml:"NonDeductibleLbbyngPltclCYAmt,omitempty" json:",omitempty"`
	NonDedLbbyngPltclCyovAmt       int                             `xml:"NonDedLbbyngPltclCyovAmt,omitempty" json:",omitempty"`
	NonDeductibleLbbyngPltclTotAmt int                             `xml:"NonDeductibleLbbyngPltclTotAmt,omitempty" json:",omitempty"`
	AggregateReportedDuesNtcAmt    int                             `xml:"AggregateReportedDuesNtcAmt,omitempty" json:",omitempty"`
	CarriedOverAmt                 int                             `xml:"CarriedOverAmt,omitempty" json:",omitempty"`
	TaxableAmt                     int                             `xml:"TaxableAmt,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                     IdType                          `xml:"documentId,attr"`
	SoftwareId                     SoftwareIdType                  `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                          `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                          `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleC) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule C
type IRS990ScheduleCType struct {
	PoliticalExpendituresAmt       int                             `xml:"PoliticalExpendituresAmt,omitempty" json:",omitempty"`
	VolunteerHoursCnt              int                             `xml:"VolunteerHoursCnt,omitempty" json:",omitempty"`
	Section4955OrganizationTaxAmt  int                             `xml:"Section4955OrganizationTaxAmt,omitempty" json:",omitempty"`
	Section4955ManagersTaxAmt      int                             `xml:"Section4955ManagersTaxAmt,omitempty" json:",omitempty"`
	Form4720FiledSection4955TaxInd bool                            `xml:"Form4720FiledSection4955TaxInd,omitempty" json:",omitempty"`
	CorrectionMadeInd              bool                            `xml:"CorrectionMadeInd,omitempty" json:",omitempty"`
	Expended527ActivitiesAmt       int                             `xml:"Expended527ActivitiesAmt,omitempty" json:",omitempty"`
	InternalFundsContributedAmt    int                             `xml:"InternalFundsContributedAmt,omitempty" json:",omitempty"`
	TotalExemptFunctionExpendAmt   int                             `xml:"TotalExemptFunctionExpendAmt,omitempty" json:",omitempty"`
	Form1120POLFiledInd            bool                            `xml:"Form1120POLFiledInd,omitempty" json:",omitempty"`
	Section527PoliticalOrgGrp      []Section527PoliticalOrgGrp     `xml:"Section527PoliticalOrgGrp,omitempty" json:",omitempty"`
	OrganizationBelongsAffltGrpInd *OrganizationBelongsAffltGrpInd `xml:"OrganizationBelongsAffltGrpInd,omitempty" json:",omitempty"`
	LimitedControlProvisionsAppInd CheckboxType                    `xml:"LimitedControlProvisionsAppInd,omitempty" json:",omitempty"`
	TotalGrassrootsLobbyingGrp     *Form990SchCPartIIAGroup1Type   `xml:"TotalGrassrootsLobbyingGrp,omitempty" json:",omitempty"`
	TotalDirectLobbyingGrp         *Form990SchCPartIIAGroup1Type   `xml:"TotalDirectLobbyingGrp,omitempty" json:",omitempty"`
	TotalLobbyingExpendGrp         *Form990SchCPartIIAGroup1Type   `xml:"TotalLobbyingExpendGrp,omitempty" json:",omitempty"`
	OtherExemptPurposeExpendGrp    *Form990SchCPartIIAGroup1Type   `xml:"OtherExemptPurposeExpendGrp,omitempty" json:",omitempty"`
	TotalExemptPurposeExpendGrp    *Form990SchCPartIIAGroup1Type   `xml:"TotalExemptPurposeExpendGrp,omitempty" json:",omitempty"`
	LobbyingNontaxableAmountGrp    *Form990SchCPartIIAGroup1Type   `xml:"LobbyingNontaxableAmountGrp,omitempty" json:",omitempty"`
	GrassrootsNontaxableGrp        *Form990SchCPartIIAGroup1Type   `xml:"GrassrootsNontaxableGrp,omitempty" json:",omitempty"`
	TotLbbyngGrassrootMnsNonTxGrp  *Form990SchCPartIIAGroup1Type   `xml:"TotLbbyngGrassrootMnsNonTxGrp,omitempty" json:",omitempty"`
	TotLbbyExpendMnsLbbyngNonTxGrp *Form990SchCPartIIAGroup1Type   `xml:"TotLbbyExpendMnsLbbyngNonTxGrp,omitempty" json:",omitempty"`
	Form4720FiledInd               bool                            `xml:"Form4720FiledInd,omitempty" json:",omitempty"`
	AvgLobbyingNontaxableAmountGrp *Form990SchCPartIIAGroup2Type   `xml:"AvgLobbyingNontaxableAmountGrp,omitempty" json:",omitempty"`
	LobbyingCeilingAmt             int                             `xml:"LobbyingCeilingAmt,omitempty" json:",omitempty"`
	AvgTotalLobbyingExpendGrp      *Form990SchCPartIIAGroup2Type   `xml:"AvgTotalLobbyingExpendGrp,omitempty" json:",omitempty"`
	AvgGrassrootsNontaxableGrp     *Form990SchCPartIIAGroup2Type   `xml:"AvgGrassrootsNontaxableGrp,omitempty" json:",omitempty"`
	GrassrootsCeilingAmt           int                             `xml:"GrassrootsCeilingAmt,omitempty" json:",omitempty"`
	AvgGrassrootsLobbyingExpendGrp *Form990SchCPartIIAGroup2Type   `xml:"AvgGrassrootsLobbyingExpendGrp,omitempty" json:",omitempty"`
	VolunteersInd                  bool                            `xml:"VolunteersInd,omitempty" json:",omitempty"`
	PaidStaffOrManagementInd       bool                            `xml:"PaidStaffOrManagementInd,omitempty" json:",omitempty"`
	MediaAdvertisementsInd         bool                            `xml:"MediaAdvertisementsInd,omitempty" json:",omitempty"`
	MediaAdvertisementsAmt         int                             `xml:"MediaAdvertisementsAmt,omitempty" json:",omitempty"`
	MailingsMembersInd             bool                            `xml:"MailingsMembersInd,omitempty" json:",omitempty"`
	MailingsMembersAmt             int                             `xml:"MailingsMembersAmt,omitempty" json:",omitempty"`
	PublicationsOrBroadcastInd     bool                            `xml:"PublicationsOrBroadcastInd,omitempty" json:",omitempty"`
	PublicationsOrBroadcastAmt     int                             `xml:"PublicationsOrBroadcastAmt,omitempty" json:",omitempty"`
	GrantsOtherOrganizationsInd    bool                            `xml:"GrantsOtherOrganizationsInd,omitempty" json:",omitempty"`
	GrantsOtherOrganizationsAmt    int                             `xml:"GrantsOtherOrganizationsAmt,omitempty" json:",omitempty"`
	DirectContactLegislatorsInd    bool                            `xml:"DirectContactLegislatorsInd,omitempty" json:",omitempty"`
	DirectContactLegislatorsAmt    int                             `xml:"DirectContactLegislatorsAmt,omitempty" json:",omitempty"`
	RalliesDemonstrationsInd       bool                            `xml:"RalliesDemonstrationsInd,omitempty" json:",omitempty"`
	RalliesDemonstrationsAmt       int                             `xml:"RalliesDemonstrationsAmt,omitempty" json:",omitempty"`
	OtherActivitiesInd             bool                            `xml:"OtherActivitiesInd,omitempty" json:",omitempty"`
	OtherActivitiesAmt             int                             `xml:"OtherActivitiesAmt,omitempty" json:",omitempty"`
	TotalLobbyingExpendituresAmt   int                             `xml:"TotalLobbyingExpendituresAmt,omitempty" json:",omitempty"`
	NotDescribedSection501c3Ind    bool                            `xml:"NotDescribedSection501c3Ind,omitempty" json:",omitempty"`
	Tax4912Amt                     int                             `xml:"Tax4912Amt,omitempty" json:",omitempty"`
	Managers4912TaxAmt             int                             `xml:"Managers4912TaxAmt,omitempty" json:",omitempty"`
	Form4720Filed4912TaxInd        bool                            `xml:"Form4720Filed4912TaxInd,omitempty" json:",omitempty"`
	SubstantiallyAllDuesNondedInd  bool                            `xml:"SubstantiallyAllDuesNondedInd,omitempty" json:",omitempty"`
	OnlyInHouseLobbyingInd         bool                            `xml:"OnlyInHouseLobbyingInd,omitempty" json:",omitempty"`
	AgreeCarryoverPriorYearInd     bool                            `xml:"AgreeCarryoverPriorYearInd,omitempty" json:",omitempty"`
	DuesAssessmentsAmt             int                             `xml:"DuesAssessmentsAmt,omitempty" json:",omitempty"`
	NonDeductibleLbbyngPltclCYAmt  int                             `xml:"NonDeductibleLbbyngPltclCYAmt,omitempty" json:",omitempty"`
	NonDedLbbyngPltclCyovAmt       int                             `xml:"NonDedLbbyngPltclCyovAmt,omitempty" json:",omitempty"`
	NonDeductibleLbbyngPltclTotAmt int                             `xml:"NonDeductibleLbbyngPltclTotAmt,omitempty" json:",omitempty"`
	AggregateReportedDuesNtcAmt    int                             `xml:"AggregateReportedDuesNtcAmt,omitempty" json:",omitempty"`
	CarriedOverAmt                 int                             `xml:"CarriedOverAmt,omitempty" json:",omitempty"`
	TaxableAmt                     int                             `xml:"TaxableAmt,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleCType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleD struct {
	DonorAdvisedFundsHeldCnt       int                             `xml:"DonorAdvisedFundsHeldCnt,omitempty" json:",omitempty"`
	FundsAndOtherAccountsHeldCnt   int                             `xml:"FundsAndOtherAccountsHeldCnt,omitempty" json:",omitempty"`
	DonorAdvisedFundsContriAmt     int                             `xml:"DonorAdvisedFundsContriAmt,omitempty" json:",omitempty"`
	FundsAndOtherAccountsContriAmt int                             `xml:"FundsAndOtherAccountsContriAmt,omitempty" json:",omitempty"`
	DonorAdvisedFundsGrantsAmt     int                             `xml:"DonorAdvisedFundsGrantsAmt,omitempty" json:",omitempty"`
	FundsAndOtherAccountsGrantsAmt int                             `xml:"FundsAndOtherAccountsGrantsAmt,omitempty" json:",omitempty"`
	DonorAdvisedFundsVlEOYAmt      int                             `xml:"DonorAdvisedFundsVlEOYAmt,omitempty" json:",omitempty"`
	FundsAndOtherAccountsVlEOYAmt  int                             `xml:"FundsAndOtherAccountsVlEOYAmt,omitempty" json:",omitempty"`
	DisclosedOrgLegCtrlInd         bool                            `xml:"DisclosedOrgLegCtrlInd,omitempty" json:",omitempty"`
	DisclosedForCharitablePrpsInd  bool                            `xml:"DisclosedForCharitablePrpsInd,omitempty" json:",omitempty"`
	PreservationPublicUseInd       CheckboxType                    `xml:"PreservationPublicUseInd,omitempty" json:",omitempty"`
	ProtectionNaturalHabitatInd    CheckboxType                    `xml:"ProtectionNaturalHabitatInd,omitempty" json:",omitempty"`
	PreservationOpenSpaceInd       CheckboxType                    `xml:"PreservationOpenSpaceInd,omitempty" json:",omitempty"`
	HistoricLandAreaInd            CheckboxType                    `xml:"HistoricLandAreaInd,omitempty" json:",omitempty"`
	HistoricStructureInd           CheckboxType                    `xml:"HistoricStructureInd,omitempty" json:",omitempty"`
	TotalEasementsCnt              int                             `xml:"TotalEasementsCnt,omitempty" json:",omitempty"`
	TotalAcreageCnt                float64                         `xml:"TotalAcreageCnt,omitempty" json:",omitempty"`
	HistoricStructureEasementsCnt  int                             `xml:"HistoricStructureEasementsCnt,omitempty" json:",omitempty"`
	HistoricStrctrEasementsAftrCnt int                             `xml:"HistoricStrctrEasementsAftrCnt,omitempty" json:",omitempty"`
	ModifiedEasementsCnt           int                             `xml:"ModifiedEasementsCnt,omitempty" json:",omitempty"`
	StatesEasementsHeldCnt         int                             `xml:"StatesEasementsHeldCnt,omitempty" json:",omitempty"`
	WrittenPolicyMonitoringInd     bool                            `xml:"WrittenPolicyMonitoringInd,omitempty" json:",omitempty"`
	StaffHoursSpentEnforcementCnt  float64                         `xml:"StaffHoursSpentEnforcementCnt,omitempty" json:",omitempty"`
	ExpensesIncurredEnforcementAmt int                             `xml:"ExpensesIncurredEnforcementAmt,omitempty" json:",omitempty"`
	Section170hRqrStsfdInd         bool                            `xml:"Section170hRqrStsfdInd,omitempty" json:",omitempty"`
	ArtPublicExhibitionAmountsGrp  *Form990SchDPartIIIGroup1Type   `xml:"ArtPublicExhibitionAmountsGrp,omitempty" json:",omitempty"`
	HeldWorksArtGrp                *Form990SchDPartIIIGroup1Type   `xml:"HeldWorksArtGrp,omitempty" json:",omitempty"`
	CollectionUsedPubExhibitionInd CheckboxType                    `xml:"CollectionUsedPubExhibitionInd,omitempty" json:",omitempty"`
	CollUsedScholarlyRsrchInd      CheckboxType                    `xml:"CollUsedScholarlyRsrchInd,omitempty" json:",omitempty"`
	CollectionUsedPreservationInd  CheckboxType                    `xml:"CollectionUsedPreservationInd,omitempty" json:",omitempty"`
	CollUsedLoanOrExchProgInd      CheckboxType                    `xml:"CollUsedLoanOrExchProgInd,omitempty" json:",omitempty"`
	CollectionUsedOtherPurposesGrp *CollectionUsedOtherPurposesGrp `xml:"CollectionUsedOtherPurposesGrp,omitempty" json:",omitempty"`
	SolicitedAssetsSaleInd         bool                            `xml:"SolicitedAssetsSaleInd,omitempty" json:",omitempty"`
	AgentTrusteeEtcInd             bool                            `xml:"AgentTrusteeEtcInd,omitempty" json:",omitempty"`
	BeginningBalanceAmt            int                             `xml:"BeginningBalanceAmt,omitempty" json:",omitempty"`
	AdditionsDuringYearAmt         int                             `xml:"AdditionsDuringYearAmt,omitempty" json:",omitempty"`
	DistributionsDuringYearAmt     int                             `xml:"DistributionsDuringYearAmt,omitempty" json:",omitempty"`
	EndingBalanceAmt               int                             `xml:"EndingBalanceAmt,omitempty" json:",omitempty"`
	InclEscrowCustodialAcctLiabInd bool                            `xml:"InclEscrowCustodialAcctLiabInd,omitempty" json:",omitempty"`
	ExplanationProvidedInd         CheckboxType                    `xml:"ExplanationProvidedInd,omitempty" json:",omitempty"`
	CYEndwmtFundGrp                *Form990SchDPartVGroup1Type     `xml:"CYEndwmtFundGrp,omitempty" json:",omitempty"`
	CYMinus1YrEndwmtFundGrp        *Form990SchDPartVGroup1Type     `xml:"CYMinus1YrEndwmtFundGrp,omitempty" json:",omitempty"`
	CYMinus2YrEndwmtFundGrp        *Form990SchDPartVGroup1Type     `xml:"CYMinus2YrEndwmtFundGrp,omitempty" json:",omitempty"`
	CYMinus3YrEndwmtFundGrp        *Form990SchDPartVGroup1Type     `xml:"CYMinus3YrEndwmtFundGrp,omitempty" json:",omitempty"`
	CYMinus4YrEndwmtFundGrp        *Form990SchDPartVGroup1Type     `xml:"CYMinus4YrEndwmtFundGrp,omitempty" json:",omitempty"`
	BoardDesignatedBalanceEOYPct   float64                         `xml:"BoardDesignatedBalanceEOYPct,omitempty" json:",omitempty"`
	PrmnntEndowmentBalanceEOYPct   float64                         `xml:"PrmnntEndowmentBalanceEOYPct,omitempty" json:",omitempty"`
	TermEndowmentBalanceEOYPct     float64                         `xml:"TermEndowmentBalanceEOYPct,omitempty" json:",omitempty"`
	EndowmentsHeldUnrelatedOrgInd  bool                            `xml:"EndowmentsHeldUnrelatedOrgInd,omitempty" json:",omitempty"`
	EndowmentsHeldRelatedOrgInd    bool                            `xml:"EndowmentsHeldRelatedOrgInd,omitempty" json:",omitempty"`
	RelatedOrgListSchRInd          bool                            `xml:"RelatedOrgListSchRInd,omitempty" json:",omitempty"`
	LandGrp                        *Form990SchDPartVIGroup2Type    `xml:"LandGrp,omitempty" json:",omitempty"`
	BuildingsGrp                   *Form990SchDPartVIGroup1Type    `xml:"BuildingsGrp,omitempty" json:",omitempty"`
	LeaseholdImprovementsGrp       *Form990SchDPartVIGroup1Type    `xml:"LeaseholdImprovementsGrp,omitempty" json:",omitempty"`
	EquipmentGrp                   *Form990SchDPartVIGroup1Type    `xml:"EquipmentGrp,omitempty" json:",omitempty"`
	OtherLandBuildingsGrp          *Form990SchDPartVIGroup1Type    `xml:"OtherLandBuildingsGrp,omitempty" json:",omitempty"`
	TotalBookValueLandBuildingsAmt int                             `xml:"TotalBookValueLandBuildingsAmt,omitempty" json:",omitempty"`
	FinancialDerivativesGrp        *Form990SchDPartVIIGroup1Type   `xml:"FinancialDerivativesGrp,omitempty" json:",omitempty"`
	CloselyHeldEquityInterestsGrp  *Form990SchDPartVIIGroup1Type   `xml:"CloselyHeldEquityInterestsGrp,omitempty" json:",omitempty"`
	OtherSecuritiesGrp             []Form990SchDPartVIIGroup2Type  `xml:"OtherSecuritiesGrp,omitempty" json:",omitempty"`
	TotalBookValueSecuritiesAmt    int                             `xml:"TotalBookValueSecuritiesAmt,omitempty" json:",omitempty"`
	InvstProgramRelatedOrgGrp      []Form990SchDPartVIIIGroup1Type `xml:"InvstProgramRelatedOrgGrp,omitempty" json:",omitempty"`
	TotalBookValueProgramRltdAmt   int                             `xml:"TotalBookValueProgramRltdAmt,omitempty" json:",omitempty"`
	OtherAssetsOrgGrp              []Form990SchDPartIXGroup1Type   `xml:"OtherAssetsOrgGrp,omitempty" json:",omitempty"`
	TotalBookValueOtherAssetsAmt   int                             `xml:"TotalBookValueOtherAssetsAmt,omitempty" json:",omitempty"`
	FederalIncomeTaxLiabilityAmt   int                             `xml:"FederalIncomeTaxLiabilityAmt,omitempty" json:",omitempty"`
	OtherLiabilitiesOrgGrp         []Form990SchDPartXGroup1Type    `xml:"OtherLiabilitiesOrgGrp,omitempty" json:",omitempty"`
	TotalLiabilityAmt              int                             `xml:"TotalLiabilityAmt,omitempty" json:",omitempty"`
	FootnoteTextInd                CheckboxType                    `xml:"FootnoteTextInd,omitempty" json:",omitempty"`
	TotalRevEtcAuditedFinclStmtAmt int                             `xml:"TotalRevEtcAuditedFinclStmtAmt,omitempty" json:",omitempty"`
	NetUnrealizedGainsInvstAmt     int                             `xml:"NetUnrealizedGainsInvstAmt,omitempty" json:",omitempty"`
	DonatedServicesAndUseFcltsAmt  int                             `xml:"DonatedServicesAndUseFcltsAmt,omitempty" json:",omitempty"`
	RecoveriesPriorYearGrantsAmt   int                             `xml:"RecoveriesPriorYearGrantsAmt,omitempty" json:",omitempty"`
	OtherRevenueAmt                int                             `xml:"OtherRevenueAmt,omitempty" json:",omitempty"`
	RevenueNotReportedAmt          int                             `xml:"RevenueNotReportedAmt,omitempty" json:",omitempty"`
	RevenueSubtotalAmt             int                             `xml:"RevenueSubtotalAmt,omitempty" json:",omitempty"`
	InvestmentExpensesNotIncldAmt  int                             `xml:"InvestmentExpensesNotIncldAmt,omitempty" json:",omitempty"`
	OtherRevenuesNotIncludedAmt    int                             `xml:"OtherRevenuesNotIncludedAmt,omitempty" json:",omitempty"`
	RevenueNotReportedFinclStmtAmt int                             `xml:"RevenueNotReportedFinclStmtAmt,omitempty" json:",omitempty"`
	TotalRevenuePerForm990Amt      int                             `xml:"TotalRevenuePerForm990Amt,omitempty" json:",omitempty"`
	TotExpnsEtcAuditedFinclStmtAmt int                             `xml:"TotExpnsEtcAuditedFinclStmtAmt,omitempty" json:",omitempty"`
	DonatedServicesUseFcltsAmt     int                             `xml:"DonatedServicesUseFcltsAmt,omitempty" json:",omitempty"`
	PriorYearAdjustmentsAmt        int                             `xml:"PriorYearAdjustmentsAmt,omitempty" json:",omitempty"`
	LossesReportedAmt              int                             `xml:"LossesReportedAmt,omitempty" json:",omitempty"`
	OtherExpensesIncludedAmt       int                             `xml:"OtherExpensesIncludedAmt,omitempty" json:",omitempty"`
	ExpensesNotReportedAmt         int                             `xml:"ExpensesNotReportedAmt,omitempty" json:",omitempty"`
	ExpensesSubtotalAmt            int                             `xml:"ExpensesSubtotalAmt,omitempty" json:",omitempty"`
	InvestmentExpensesNotIncld2Amt int                             `xml:"InvestmentExpensesNotIncld2Amt,omitempty" json:",omitempty"`
	OtherExpensesNotIncludedAmt    int                             `xml:"OtherExpensesNotIncludedAmt,omitempty" json:",omitempty"`
	ExpensesNotRptFinclStmtAmt     int                             `xml:"ExpensesNotRptFinclStmtAmt,omitempty" json:",omitempty"`
	TotalExpensesPerForm990Amt     int                             `xml:"TotalExpensesPerForm990Amt,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                     IdType                          `xml:"documentId,attr"`
	SoftwareId                     SoftwareIdType                  `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                          `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                          `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleD) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule D
type IRS990ScheduleDType struct {
	DonorAdvisedFundsHeldCnt       int                             `xml:"DonorAdvisedFundsHeldCnt,omitempty" json:",omitempty"`
	FundsAndOtherAccountsHeldCnt   int                             `xml:"FundsAndOtherAccountsHeldCnt,omitempty" json:",omitempty"`
	DonorAdvisedFundsContriAmt     int                             `xml:"DonorAdvisedFundsContriAmt,omitempty" json:",omitempty"`
	FundsAndOtherAccountsContriAmt int                             `xml:"FundsAndOtherAccountsContriAmt,omitempty" json:",omitempty"`
	DonorAdvisedFundsGrantsAmt     int                             `xml:"DonorAdvisedFundsGrantsAmt,omitempty" json:",omitempty"`
	FundsAndOtherAccountsGrantsAmt int                             `xml:"FundsAndOtherAccountsGrantsAmt,omitempty" json:",omitempty"`
	DonorAdvisedFundsVlEOYAmt      int                             `xml:"DonorAdvisedFundsVlEOYAmt,omitempty" json:",omitempty"`
	FundsAndOtherAccountsVlEOYAmt  int                             `xml:"FundsAndOtherAccountsVlEOYAmt,omitempty" json:",omitempty"`
	DisclosedOrgLegCtrlInd         bool                            `xml:"DisclosedOrgLegCtrlInd,omitempty" json:",omitempty"`
	DisclosedForCharitablePrpsInd  bool                            `xml:"DisclosedForCharitablePrpsInd,omitempty" json:",omitempty"`
	PreservationPublicUseInd       CheckboxType                    `xml:"PreservationPublicUseInd,omitempty" json:",omitempty"`
	ProtectionNaturalHabitatInd    CheckboxType                    `xml:"ProtectionNaturalHabitatInd,omitempty" json:",omitempty"`
	PreservationOpenSpaceInd       CheckboxType                    `xml:"PreservationOpenSpaceInd,omitempty" json:",omitempty"`
	HistoricLandAreaInd            CheckboxType                    `xml:"HistoricLandAreaInd,omitempty" json:",omitempty"`
	HistoricStructureInd           CheckboxType                    `xml:"HistoricStructureInd,omitempty" json:",omitempty"`
	TotalEasementsCnt              int                             `xml:"TotalEasementsCnt,omitempty" json:",omitempty"`
	TotalAcreageCnt                float64                         `xml:"TotalAcreageCnt,omitempty" json:",omitempty"`
	HistoricStructureEasementsCnt  int                             `xml:"HistoricStructureEasementsCnt,omitempty" json:",omitempty"`
	HistoricStrctrEasementsAftrCnt int                             `xml:"HistoricStrctrEasementsAftrCnt,omitempty" json:",omitempty"`
	ModifiedEasementsCnt           int                             `xml:"ModifiedEasementsCnt,omitempty" json:",omitempty"`
	StatesEasementsHeldCnt         int                             `xml:"StatesEasementsHeldCnt,omitempty" json:",omitempty"`
	WrittenPolicyMonitoringInd     bool                            `xml:"WrittenPolicyMonitoringInd,omitempty" json:",omitempty"`
	StaffHoursSpentEnforcementCnt  float64                         `xml:"StaffHoursSpentEnforcementCnt,omitempty" json:",omitempty"`
	ExpensesIncurredEnforcementAmt int                             `xml:"ExpensesIncurredEnforcementAmt,omitempty" json:",omitempty"`
	Section170hRqrStsfdInd         bool                            `xml:"Section170hRqrStsfdInd,omitempty" json:",omitempty"`
	ArtPublicExhibitionAmountsGrp  *Form990SchDPartIIIGroup1Type   `xml:"ArtPublicExhibitionAmountsGrp,omitempty" json:",omitempty"`
	HeldWorksArtGrp                *Form990SchDPartIIIGroup1Type   `xml:"HeldWorksArtGrp,omitempty" json:",omitempty"`
	CollectionUsedPubExhibitionInd CheckboxType                    `xml:"CollectionUsedPubExhibitionInd,omitempty" json:",omitempty"`
	CollUsedScholarlyRsrchInd      CheckboxType                    `xml:"CollUsedScholarlyRsrchInd,omitempty" json:",omitempty"`
	CollectionUsedPreservationInd  CheckboxType                    `xml:"CollectionUsedPreservationInd,omitempty" json:",omitempty"`
	CollUsedLoanOrExchProgInd      CheckboxType                    `xml:"CollUsedLoanOrExchProgInd,omitempty" json:",omitempty"`
	CollectionUsedOtherPurposesGrp *CollectionUsedOtherPurposesGrp `xml:"CollectionUsedOtherPurposesGrp,omitempty" json:",omitempty"`
	SolicitedAssetsSaleInd         bool                            `xml:"SolicitedAssetsSaleInd,omitempty" json:",omitempty"`
	AgentTrusteeEtcInd             bool                            `xml:"AgentTrusteeEtcInd,omitempty" json:",omitempty"`
	BeginningBalanceAmt            int                             `xml:"BeginningBalanceAmt,omitempty" json:",omitempty"`
	AdditionsDuringYearAmt         int                             `xml:"AdditionsDuringYearAmt,omitempty" json:",omitempty"`
	DistributionsDuringYearAmt     int                             `xml:"DistributionsDuringYearAmt,omitempty" json:",omitempty"`
	EndingBalanceAmt               int                             `xml:"EndingBalanceAmt,omitempty" json:",omitempty"`
	InclEscrowCustodialAcctLiabInd bool                            `xml:"InclEscrowCustodialAcctLiabInd,omitempty" json:",omitempty"`
	ExplanationProvidedInd         CheckboxType                    `xml:"ExplanationProvidedInd,omitempty" json:",omitempty"`
	CYEndwmtFundGrp                *Form990SchDPartVGroup1Type     `xml:"CYEndwmtFundGrp,omitempty" json:",omitempty"`
	CYMinus1YrEndwmtFundGrp        *Form990SchDPartVGroup1Type     `xml:"CYMinus1YrEndwmtFundGrp,omitempty" json:",omitempty"`
	CYMinus2YrEndwmtFundGrp        *Form990SchDPartVGroup1Type     `xml:"CYMinus2YrEndwmtFundGrp,omitempty" json:",omitempty"`
	CYMinus3YrEndwmtFundGrp        *Form990SchDPartVGroup1Type     `xml:"CYMinus3YrEndwmtFundGrp,omitempty" json:",omitempty"`
	CYMinus4YrEndwmtFundGrp        *Form990SchDPartVGroup1Type     `xml:"CYMinus4YrEndwmtFundGrp,omitempty" json:",omitempty"`
	BoardDesignatedBalanceEOYPct   float64                         `xml:"BoardDesignatedBalanceEOYPct,omitempty" json:",omitempty"`
	PrmnntEndowmentBalanceEOYPct   float64                         `xml:"PrmnntEndowmentBalanceEOYPct,omitempty" json:",omitempty"`
	TermEndowmentBalanceEOYPct     float64                         `xml:"TermEndowmentBalanceEOYPct,omitempty" json:",omitempty"`
	EndowmentsHeldUnrelatedOrgInd  bool                            `xml:"EndowmentsHeldUnrelatedOrgInd,omitempty" json:",omitempty"`
	EndowmentsHeldRelatedOrgInd    bool                            `xml:"EndowmentsHeldRelatedOrgInd,omitempty" json:",omitempty"`
	RelatedOrgListSchRInd          bool                            `xml:"RelatedOrgListSchRInd,omitempty" json:",omitempty"`
	LandGrp                        *Form990SchDPartVIGroup2Type    `xml:"LandGrp,omitempty" json:",omitempty"`
	BuildingsGrp                   *Form990SchDPartVIGroup1Type    `xml:"BuildingsGrp,omitempty" json:",omitempty"`
	LeaseholdImprovementsGrp       *Form990SchDPartVIGroup1Type    `xml:"LeaseholdImprovementsGrp,omitempty" json:",omitempty"`
	EquipmentGrp                   *Form990SchDPartVIGroup1Type    `xml:"EquipmentGrp,omitempty" json:",omitempty"`
	OtherLandBuildingsGrp          *Form990SchDPartVIGroup1Type    `xml:"OtherLandBuildingsGrp,omitempty" json:",omitempty"`
	TotalBookValueLandBuildingsAmt int                             `xml:"TotalBookValueLandBuildingsAmt,omitempty" json:",omitempty"`
	FinancialDerivativesGrp        *Form990SchDPartVIIGroup1Type   `xml:"FinancialDerivativesGrp,omitempty" json:",omitempty"`
	CloselyHeldEquityInterestsGrp  *Form990SchDPartVIIGroup1Type   `xml:"CloselyHeldEquityInterestsGrp,omitempty" json:",omitempty"`
	OtherSecuritiesGrp             []Form990SchDPartVIIGroup2Type  `xml:"OtherSecuritiesGrp,omitempty" json:",omitempty"`
	TotalBookValueSecuritiesAmt    int                             `xml:"TotalBookValueSecuritiesAmt,omitempty" json:",omitempty"`
	InvstProgramRelatedOrgGrp      []Form990SchDPartVIIIGroup1Type `xml:"InvstProgramRelatedOrgGrp,omitempty" json:",omitempty"`
	TotalBookValueProgramRltdAmt   int                             `xml:"TotalBookValueProgramRltdAmt,omitempty" json:",omitempty"`
	OtherAssetsOrgGrp              []Form990SchDPartIXGroup1Type   `xml:"OtherAssetsOrgGrp,omitempty" json:",omitempty"`
	TotalBookValueOtherAssetsAmt   int                             `xml:"TotalBookValueOtherAssetsAmt,omitempty" json:",omitempty"`
	FederalIncomeTaxLiabilityAmt   int                             `xml:"FederalIncomeTaxLiabilityAmt,omitempty" json:",omitempty"`
	OtherLiabilitiesOrgGrp         []Form990SchDPartXGroup1Type    `xml:"OtherLiabilitiesOrgGrp,omitempty" json:",omitempty"`
	TotalLiabilityAmt              int                             `xml:"TotalLiabilityAmt,omitempty" json:",omitempty"`
	FootnoteTextInd                CheckboxType                    `xml:"FootnoteTextInd,omitempty" json:",omitempty"`
	TotalRevEtcAuditedFinclStmtAmt int                             `xml:"TotalRevEtcAuditedFinclStmtAmt,omitempty" json:",omitempty"`
	NetUnrealizedGainsInvstAmt     int                             `xml:"NetUnrealizedGainsInvstAmt,omitempty" json:",omitempty"`
	DonatedServicesAndUseFcltsAmt  int                             `xml:"DonatedServicesAndUseFcltsAmt,omitempty" json:",omitempty"`
	RecoveriesPriorYearGrantsAmt   int                             `xml:"RecoveriesPriorYearGrantsAmt,omitempty" json:",omitempty"`
	OtherRevenueAmt                int                             `xml:"OtherRevenueAmt,omitempty" json:",omitempty"`
	RevenueNotReportedAmt          int                             `xml:"RevenueNotReportedAmt,omitempty" json:",omitempty"`
	RevenueSubtotalAmt             int                             `xml:"RevenueSubtotalAmt,omitempty" json:",omitempty"`
	InvestmentExpensesNotIncldAmt  int                             `xml:"InvestmentExpensesNotIncldAmt,omitempty" json:",omitempty"`
	OtherRevenuesNotIncludedAmt    int                             `xml:"OtherRevenuesNotIncludedAmt,omitempty" json:",omitempty"`
	RevenueNotReportedFinclStmtAmt int                             `xml:"RevenueNotReportedFinclStmtAmt,omitempty" json:",omitempty"`
	TotalRevenuePerForm990Amt      int                             `xml:"TotalRevenuePerForm990Amt,omitempty" json:",omitempty"`
	TotExpnsEtcAuditedFinclStmtAmt int                             `xml:"TotExpnsEtcAuditedFinclStmtAmt,omitempty" json:",omitempty"`
	DonatedServicesUseFcltsAmt     int                             `xml:"DonatedServicesUseFcltsAmt,omitempty" json:",omitempty"`
	PriorYearAdjustmentsAmt        int                             `xml:"PriorYearAdjustmentsAmt,omitempty" json:",omitempty"`
	LossesReportedAmt              int                             `xml:"LossesReportedAmt,omitempty" json:",omitempty"`
	OtherExpensesIncludedAmt       int                             `xml:"OtherExpensesIncludedAmt,omitempty" json:",omitempty"`
	ExpensesNotReportedAmt         int                             `xml:"ExpensesNotReportedAmt,omitempty" json:",omitempty"`
	ExpensesSubtotalAmt            int                             `xml:"ExpensesSubtotalAmt,omitempty" json:",omitempty"`
	InvestmentExpensesNotIncld2Amt int                             `xml:"InvestmentExpensesNotIncld2Amt,omitempty" json:",omitempty"`
	OtherExpensesNotIncludedAmt    int                             `xml:"OtherExpensesNotIncludedAmt,omitempty" json:",omitempty"`
	ExpensesNotRptFinclStmtAmt     int                             `xml:"ExpensesNotRptFinclStmtAmt,omitempty" json:",omitempty"`
	TotalExpensesPerForm990Amt     int                             `xml:"TotalExpensesPerForm990Amt,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleDType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleE struct {
	NondiscriminatoryPolicyStmtInd bool                            `xml:"NondiscriminatoryPolicyStmtInd,omitempty" json:",omitempty"`
	PolicyStmtInBrochuresEtcInd    bool                            `xml:"PolicyStmtInBrochuresEtcInd,omitempty" json:",omitempty"`
	PlcyPblczdViaBroadcastMediaInd bool                            `xml:"PlcyPblczdViaBroadcastMediaInd,omitempty" json:",omitempty"`
	MaintainRacialCompRecsInd      bool                            `xml:"MaintainRacialCompRecsInd,omitempty" json:",omitempty"`
	MaintainScholarshipsRecsInd    bool                            `xml:"MaintainScholarshipsRecsInd,omitempty" json:",omitempty"`
	MaintainCpyOfBrochuresEtcInd   bool                            `xml:"MaintainCpyOfBrochuresEtcInd,omitempty" json:",omitempty"`
	MaintainCpyOfAllSolInd         bool                            `xml:"MaintainCpyOfAllSolInd,omitempty" json:",omitempty"`
	DiscriminateRaceStdntsRghtsInd bool                            `xml:"DiscriminateRaceStdntsRghtsInd,omitempty" json:",omitempty"`
	DiscriminateRaceAdmissPlcyInd  bool                            `xml:"DiscriminateRaceAdmissPlcyInd,omitempty" json:",omitempty"`
	DiscriminateRaceEmplmFcultyInd bool                            `xml:"DiscriminateRaceEmplmFcultyInd,omitempty" json:",omitempty"`
	DiscriminateRaceSchsInd        bool                            `xml:"DiscriminateRaceSchsInd,omitempty" json:",omitempty"`
	DiscriminateRaceEducPlcyInd    bool                            `xml:"DiscriminateRaceEducPlcyInd,omitempty" json:",omitempty"`
	DiscriminateRaceUseOfFcltsInd  bool                            `xml:"DiscriminateRaceUseOfFcltsInd,omitempty" json:",omitempty"`
	DiscriminateRaceAthltProgInd   bool                            `xml:"DiscriminateRaceAthltProgInd,omitempty" json:",omitempty"`
	DiscriminateRaceOtherActyInd   bool                            `xml:"DiscriminateRaceOtherActyInd,omitempty" json:",omitempty"`
	GovernmentFinancialAidRcvdInd  bool                            `xml:"GovernmentFinancialAidRcvdInd,omitempty" json:",omitempty"`
	GovernmentFinancialAidRvkdInd  bool                            `xml:"GovernmentFinancialAidRvkdInd,omitempty" json:",omitempty"`
	ComplianceWithRevProc7550Ind   bool                            `xml:"ComplianceWithRevProc7550Ind,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                     IdType                          `xml:"documentId,attr"`
	SoftwareId                     SoftwareIdType                  `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                          `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                          `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleE) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule E
type IRS990ScheduleEType struct {
	NondiscriminatoryPolicyStmtInd bool                            `xml:"NondiscriminatoryPolicyStmtInd,omitempty" json:",omitempty"`
	PolicyStmtInBrochuresEtcInd    bool                            `xml:"PolicyStmtInBrochuresEtcInd,omitempty" json:",omitempty"`
	PlcyPblczdViaBroadcastMediaInd bool                            `xml:"PlcyPblczdViaBroadcastMediaInd,omitempty" json:",omitempty"`
	MaintainRacialCompRecsInd      bool                            `xml:"MaintainRacialCompRecsInd,omitempty" json:",omitempty"`
	MaintainScholarshipsRecsInd    bool                            `xml:"MaintainScholarshipsRecsInd,omitempty" json:",omitempty"`
	MaintainCpyOfBrochuresEtcInd   bool                            `xml:"MaintainCpyOfBrochuresEtcInd,omitempty" json:",omitempty"`
	MaintainCpyOfAllSolInd         bool                            `xml:"MaintainCpyOfAllSolInd,omitempty" json:",omitempty"`
	DiscriminateRaceStdntsRghtsInd bool                            `xml:"DiscriminateRaceStdntsRghtsInd,omitempty" json:",omitempty"`
	DiscriminateRaceAdmissPlcyInd  bool                            `xml:"DiscriminateRaceAdmissPlcyInd,omitempty" json:",omitempty"`
	DiscriminateRaceEmplmFcultyInd bool                            `xml:"DiscriminateRaceEmplmFcultyInd,omitempty" json:",omitempty"`
	DiscriminateRaceSchsInd        bool                            `xml:"DiscriminateRaceSchsInd,omitempty" json:",omitempty"`
	DiscriminateRaceEducPlcyInd    bool                            `xml:"DiscriminateRaceEducPlcyInd,omitempty" json:",omitempty"`
	DiscriminateRaceUseOfFcltsInd  bool                            `xml:"DiscriminateRaceUseOfFcltsInd,omitempty" json:",omitempty"`
	DiscriminateRaceAthltProgInd   bool                            `xml:"DiscriminateRaceAthltProgInd,omitempty" json:",omitempty"`
	DiscriminateRaceOtherActyInd   bool                            `xml:"DiscriminateRaceOtherActyInd,omitempty" json:",omitempty"`
	GovernmentFinancialAidRcvdInd  bool                            `xml:"GovernmentFinancialAidRcvdInd,omitempty" json:",omitempty"`
	GovernmentFinancialAidRvkdInd  bool                            `xml:"GovernmentFinancialAidRvkdInd,omitempty" json:",omitempty"`
	ComplianceWithRevProc7550Ind   bool                            `xml:"ComplianceWithRevProc7550Ind,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleEType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleF struct {
	GrantRecordsMaintainedInd     bool                                `xml:"GrantRecordsMaintainedInd,omitempty" json:",omitempty"`
	AccountActivitiesOutsideUSGrp []AccountActivitiesOutsideUSGrpType `xml:"AccountActivitiesOutsideUSGrp,omitempty" json:",omitempty"`
	SubtotalOfficesCnt            int                                 `xml:"SubtotalOfficesCnt,omitempty" json:",omitempty"`
	SubtotalEmployeesCnt          int                                 `xml:"SubtotalEmployeesCnt,omitempty" json:",omitempty"`
	ContinutationTotalOfficeCnt   int                                 `xml:"ContinutationTotalOfficeCnt,omitempty" json:",omitempty"`
	ContinutationTotalEmployeeCnt int                                 `xml:"ContinutationTotalEmployeeCnt,omitempty" json:",omitempty"`
	TotalOfficeCnt                int                                 `xml:"TotalOfficeCnt,omitempty" json:",omitempty"`
	TotalEmployeeCnt              int                                 `xml:"TotalEmployeeCnt,omitempty" json:",omitempty"`
	SubtotalSpentAmt              int                                 `xml:"SubtotalSpentAmt,omitempty" json:",omitempty"`
	ContinuationSpentAmt          int                                 `xml:"ContinuationSpentAmt,omitempty" json:",omitempty"`
	TotalSpentAmt                 int                                 `xml:"TotalSpentAmt,omitempty" json:",omitempty"`
	GrantsToOrgOutsideUSGrp       []GrantsToOrgOutsideUSGrpType       `xml:"GrantsToOrgOutsideUSGrp,omitempty" json:",omitempty"`
	Total501c3OrgCnt              int                                 `xml:"Total501c3OrgCnt,omitempty" json:",omitempty"`
	TotalOtherOrgCnt              int                                 `xml:"TotalOtherOrgCnt,omitempty" json:",omitempty"`
	ForeignIndividualsGrantsGrp   []ForeignIndividualsGrantsGrpType   `xml:"ForeignIndividualsGrantsGrp,omitempty" json:",omitempty"`
	TransferToForeignCorpInd      bool                                `xml:"TransferToForeignCorpInd,omitempty" json:",omitempty"`
	InterestInForeignTrustInd     bool                                `xml:"InterestInForeignTrustInd,omitempty" json:",omitempty"`
	ForeignCorpOwnershipInd       bool                                `xml:"ForeignCorpOwnershipInd,omitempty" json:",omitempty"`
	PassiveForeignInvestmestCoInd bool                                `xml:"PassiveForeignInvestmestCoInd,omitempty" json:",omitempty"`
	ForeignPartnershipInd         bool                                `xml:"ForeignPartnershipInd,omitempty" json:",omitempty"`
	BoycottCountriesInd           bool                                `xml:"BoycottCountriesInd,omitempty" json:",omitempty"`
	SupplementalInformationDetail []SupplementalInformationDetail     `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                    IdType                              `xml:"documentId,attr"`
	SoftwareId                    SoftwareIdType                      `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum            string                              `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                  string                              `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleF) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule F
type IRS990ScheduleFType struct {
	GrantRecordsMaintainedInd     bool                                `xml:"GrantRecordsMaintainedInd,omitempty" json:",omitempty"`
	AccountActivitiesOutsideUSGrp []AccountActivitiesOutsideUSGrpType `xml:"AccountActivitiesOutsideUSGrp,omitempty" json:",omitempty"`
	SubtotalOfficesCnt            int                                 `xml:"SubtotalOfficesCnt,omitempty" json:",omitempty"`
	SubtotalEmployeesCnt          int                                 `xml:"SubtotalEmployeesCnt,omitempty" json:",omitempty"`
	ContinutationTotalOfficeCnt   int                                 `xml:"ContinutationTotalOfficeCnt,omitempty" json:",omitempty"`
	ContinutationTotalEmployeeCnt int                                 `xml:"ContinutationTotalEmployeeCnt,omitempty" json:",omitempty"`
	TotalOfficeCnt                int                                 `xml:"TotalOfficeCnt,omitempty" json:",omitempty"`
	TotalEmployeeCnt              int                                 `xml:"TotalEmployeeCnt,omitempty" json:",omitempty"`
	SubtotalSpentAmt              int                                 `xml:"SubtotalSpentAmt,omitempty" json:",omitempty"`
	ContinuationSpentAmt          int                                 `xml:"ContinuationSpentAmt,omitempty" json:",omitempty"`
	TotalSpentAmt                 int                                 `xml:"TotalSpentAmt,omitempty" json:",omitempty"`
	GrantsToOrgOutsideUSGrp       []GrantsToOrgOutsideUSGrpType       `xml:"GrantsToOrgOutsideUSGrp,omitempty" json:",omitempty"`
	Total501c3OrgCnt              int                                 `xml:"Total501c3OrgCnt,omitempty" json:",omitempty"`
	TotalOtherOrgCnt              int                                 `xml:"TotalOtherOrgCnt,omitempty" json:",omitempty"`
	ForeignIndividualsGrantsGrp   []ForeignIndividualsGrantsGrpType   `xml:"ForeignIndividualsGrantsGrp,omitempty" json:",omitempty"`
	TransferToForeignCorpInd      bool                                `xml:"TransferToForeignCorpInd,omitempty" json:",omitempty"`
	InterestInForeignTrustInd     bool                                `xml:"InterestInForeignTrustInd,omitempty" json:",omitempty"`
	ForeignCorpOwnershipInd       bool                                `xml:"ForeignCorpOwnershipInd,omitempty" json:",omitempty"`
	PassiveForeignInvestmestCoInd bool                                `xml:"PassiveForeignInvestmestCoInd,omitempty" json:",omitempty"`
	ForeignPartnershipInd         bool                                `xml:"ForeignPartnershipInd,omitempty" json:",omitempty"`
	BoycottCountriesInd           bool                                `xml:"BoycottCountriesInd,omitempty" json:",omitempty"`
	SupplementalInformationDetail []SupplementalInformationDetail     `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleFType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleG struct {
	MailSolicitationsInd           CheckboxType                        `xml:"MailSolicitationsInd,omitempty" json:",omitempty"`
	EmailSolicitationsInd          CheckboxType                        `xml:"EmailSolicitationsInd,omitempty" json:",omitempty"`
	PhoneSolicitationsInd          CheckboxType                        `xml:"PhoneSolicitationsInd,omitempty" json:",omitempty"`
	InPersonSolicitationsInd       CheckboxType                        `xml:"InPersonSolicitationsInd,omitempty" json:",omitempty"`
	SolicitationOfNonGovtGrantsInd CheckboxType                        `xml:"SolicitationOfNonGovtGrantsInd,omitempty" json:",omitempty"`
	SolicitationOfGovtGrantsInd    CheckboxType                        `xml:"SolicitationOfGovtGrantsInd,omitempty" json:",omitempty"`
	SpecialFundraisingEventsInd    CheckboxType                        `xml:"SpecialFundraisingEventsInd,omitempty" json:",omitempty"`
	AgrmtProfFundraisingActyInd    bool                                `xml:"AgrmtProfFundraisingActyInd,omitempty" json:",omitempty"`
	FundraiserActivityInfoGrp      []FundraiserActivityInfoGrpType     `xml:"FundraiserActivityInfoGrp,omitempty" json:",omitempty"`
	TotalGrossReceiptsAmt          int                                 `xml:"TotalGrossReceiptsAmt,omitempty" json:",omitempty"`
	TotalRetainedByContractorsAmt  int                                 `xml:"TotalRetainedByContractorsAmt,omitempty" json:",omitempty"`
	TotalNetToOrganizationAmt      int                                 `xml:"TotalNetToOrganizationAmt,omitempty" json:",omitempty"`
	LicensedStatesCd               []StateType                         `xml:"LicensedStatesCd"`
	AllStatesCd                    AllStatesCd                         `xml:"AllStatesCd"`
	FundraisingEventInformationGrp *FundraisingEventInformationGrpType `xml:"FundraisingEventInformationGrp,omitempty" json:",omitempty"`
	GamingInformationGrp           *GamingInformationGrpType           `xml:"GamingInformationGrp,omitempty" json:",omitempty"`
	StatesWhereGamingConductedCd   []StateType                         `xml:"StatesWhereGamingConductedCd,omitempty" json:",omitempty"`
	LicensedInd                    bool                                `xml:"LicensedInd,omitempty" json:",omitempty"`
	ExplanationIfNoTxt             string                              `xml:"ExplanationIfNoTxt,omitempty" json:",omitempty"`
	LicenseSuspendedEtcInd         bool                                `xml:"LicenseSuspendedEtcInd,omitempty" json:",omitempty"`
	ExplanationIfYesTxt            string                              `xml:"ExplanationIfYesTxt,omitempty" json:",omitempty"`
	GamingWithNonmembersInd        bool                                `xml:"GamingWithNonmembersInd,omitempty" json:",omitempty"`
	MemberOfOtherEntityInd         bool                                `xml:"MemberOfOtherEntityInd,omitempty" json:",omitempty"`
	GamingOwnFacilityPct           float64                             `xml:"GamingOwnFacilityPct,omitempty" json:",omitempty"`
	GamingOtherFacilityPct         float64                             `xml:"GamingOtherFacilityPct,omitempty" json:",omitempty"`
	IndividualWithBooksNm          PersonNameType                      `xml:"IndividualWithBooksNm"`
	PersonsWithBooksName           *BusinessNameType                   `xml:"PersonsWithBooksName"`
	PersonsWithBooksUSAddress      *USAddressType                      `xml:"PersonsWithBooksUSAddress"`
	PersonsWithBooksForeignAddress *ForeignAddressType                 `xml:"PersonsWithBooksForeignAddress"`
	CntrctWith3rdPrtyForGameRevInd bool                                `xml:"CntrctWith3rdPrtyForGameRevInd,omitempty" json:",omitempty"`
	GamingRevenueReceivedByOrgAmt  int                                 `xml:"GamingRevenueReceivedByOrgAmt,omitempty" json:",omitempty"`
	GamingRevenueRtnBy3PrtyAmt     int                                 `xml:"GamingRevenueRtnBy3PrtyAmt,omitempty" json:",omitempty"`
	ThirdPartyPersonNm             PersonNameType                      `xml:"ThirdPartyPersonNm"`
	ThirdPartyBusinessName         *BusinessNameType                   `xml:"ThirdPartyBusinessName"`
	ThirdPartyUSAddress            *USAddressType                      `xml:"ThirdPartyUSAddress"`
	ThirdPartyForeignAddress       *ForeignAddressType                 `xml:"ThirdPartyForeignAddress"`
	GamingManagerPersonNm          PersonNameType                      `xml:"GamingManagerPersonNm"`
	GamingManagerBusinessName      *BusinessNameType                   `xml:"GamingManagerBusinessName"`
	GamingManagerCompensationAmt   int                                 `xml:"GamingManagerCompensationAmt,omitempty" json:",omitempty"`
	GamingManagerServicesProvTxt   string                              `xml:"GamingManagerServicesProvTxt,omitempty" json:",omitempty"`
	GamingManagerIsDirectorOfcrInd CheckboxType                        `xml:"GamingManagerIsDirectorOfcrInd,omitempty" json:",omitempty"`
	GamingManagerIsEmployeeInd     CheckboxType                        `xml:"GamingManagerIsEmployeeInd,omitempty" json:",omitempty"`
	GamingManagerIsIndCntrctInd    CheckboxType                        `xml:"GamingManagerIsIndCntrctInd,omitempty" json:",omitempty"`
	CharitableDistributionRqrInd   bool                                `xml:"CharitableDistributionRqrInd,omitempty" json:",omitempty"`
	DistributedAmt                 int                                 `xml:"DistributedAmt,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail     `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                     IdType                              `xml:"documentId,attr"`
	SoftwareId                     SoftwareIdType                      `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                              `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                              `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleG) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule G
type IRS990ScheduleGType struct {
	MailSolicitationsInd           CheckboxType                        `xml:"MailSolicitationsInd,omitempty" json:",omitempty"`
	EmailSolicitationsInd          CheckboxType                        `xml:"EmailSolicitationsInd,omitempty" json:",omitempty"`
	PhoneSolicitationsInd          CheckboxType                        `xml:"PhoneSolicitationsInd,omitempty" json:",omitempty"`
	InPersonSolicitationsInd       CheckboxType                        `xml:"InPersonSolicitationsInd,omitempty" json:",omitempty"`
	SolicitationOfNonGovtGrantsInd CheckboxType                        `xml:"SolicitationOfNonGovtGrantsInd,omitempty" json:",omitempty"`
	SolicitationOfGovtGrantsInd    CheckboxType                        `xml:"SolicitationOfGovtGrantsInd,omitempty" json:",omitempty"`
	SpecialFundraisingEventsInd    CheckboxType                        `xml:"SpecialFundraisingEventsInd,omitempty" json:",omitempty"`
	AgrmtProfFundraisingActyInd    bool                                `xml:"AgrmtProfFundraisingActyInd,omitempty" json:",omitempty"`
	FundraiserActivityInfoGrp      []FundraiserActivityInfoGrpType     `xml:"FundraiserActivityInfoGrp,omitempty" json:",omitempty"`
	TotalGrossReceiptsAmt          int                                 `xml:"TotalGrossReceiptsAmt,omitempty" json:",omitempty"`
	TotalRetainedByContractorsAmt  int                                 `xml:"TotalRetainedByContractorsAmt,omitempty" json:",omitempty"`
	TotalNetToOrganizationAmt      int                                 `xml:"TotalNetToOrganizationAmt,omitempty" json:",omitempty"`
	LicensedStatesCd               []StateType                         `xml:"LicensedStatesCd"`
	AllStatesCd                    AllStatesCd                         `xml:"AllStatesCd"`
	FundraisingEventInformationGrp *FundraisingEventInformationGrpType `xml:"FundraisingEventInformationGrp,omitempty" json:",omitempty"`
	GamingInformationGrp           *GamingInformationGrpType           `xml:"GamingInformationGrp,omitempty" json:",omitempty"`
	StatesWhereGamingConductedCd   []StateType                         `xml:"StatesWhereGamingConductedCd,omitempty" json:",omitempty"`
	LicensedInd                    bool                                `xml:"LicensedInd,omitempty" json:",omitempty"`
	ExplanationIfNoTxt             string                              `xml:"ExplanationIfNoTxt,omitempty" json:",omitempty"`
	LicenseSuspendedEtcInd         bool                                `xml:"LicenseSuspendedEtcInd,omitempty" json:",omitempty"`
	ExplanationIfYesTxt            string                              `xml:"ExplanationIfYesTxt,omitempty" json:",omitempty"`
	GamingWithNonmembersInd        bool                                `xml:"GamingWithNonmembersInd,omitempty" json:",omitempty"`
	MemberOfOtherEntityInd         bool                                `xml:"MemberOfOtherEntityInd,omitempty" json:",omitempty"`
	GamingOwnFacilityPct           float64                             `xml:"GamingOwnFacilityPct,omitempty" json:",omitempty"`
	GamingOtherFacilityPct         float64                             `xml:"GamingOtherFacilityPct,omitempty" json:",omitempty"`
	IndividualWithBooksNm          PersonNameType                      `xml:"IndividualWithBooksNm"`
	PersonsWithBooksName           *BusinessNameType                   `xml:"PersonsWithBooksName"`
	PersonsWithBooksUSAddress      *USAddressType                      `xml:"PersonsWithBooksUSAddress"`
	PersonsWithBooksForeignAddress *ForeignAddressType                 `xml:"PersonsWithBooksForeignAddress"`
	CntrctWith3rdPrtyForGameRevInd bool                                `xml:"CntrctWith3rdPrtyForGameRevInd,omitempty" json:",omitempty"`
	GamingRevenueReceivedByOrgAmt  int                                 `xml:"GamingRevenueReceivedByOrgAmt,omitempty" json:",omitempty"`
	GamingRevenueRtnBy3PrtyAmt     int                                 `xml:"GamingRevenueRtnBy3PrtyAmt,omitempty" json:",omitempty"`
	ThirdPartyPersonNm             PersonNameType                      `xml:"ThirdPartyPersonNm"`
	ThirdPartyBusinessName         *BusinessNameType                   `xml:"ThirdPartyBusinessName"`
	ThirdPartyUSAddress            *USAddressType                      `xml:"ThirdPartyUSAddress"`
	ThirdPartyForeignAddress       *ForeignAddressType                 `xml:"ThirdPartyForeignAddress"`
	GamingManagerPersonNm          PersonNameType                      `xml:"GamingManagerPersonNm"`
	GamingManagerBusinessName      *BusinessNameType                   `xml:"GamingManagerBusinessName"`
	GamingManagerCompensationAmt   int                                 `xml:"GamingManagerCompensationAmt,omitempty" json:",omitempty"`
	GamingManagerServicesProvTxt   string                              `xml:"GamingManagerServicesProvTxt,omitempty" json:",omitempty"`
	GamingManagerIsDirectorOfcrInd CheckboxType                        `xml:"GamingManagerIsDirectorOfcrInd,omitempty" json:",omitempty"`
	GamingManagerIsEmployeeInd     CheckboxType                        `xml:"GamingManagerIsEmployeeInd,omitempty" json:",omitempty"`
	GamingManagerIsIndCntrctInd    CheckboxType                        `xml:"GamingManagerIsIndCntrctInd,omitempty" json:",omitempty"`
	CharitableDistributionRqrInd   bool                                `xml:"CharitableDistributionRqrInd,omitempty" json:",omitempty"`
	DistributedAmt                 int                                 `xml:"DistributedAmt,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail     `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleGType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleH struct {
	FinancialAssistancePolicyInd   bool                            `xml:"FinancialAssistancePolicyInd,omitempty" json:",omitempty"`
	WrittenPolicyInd               bool                            `xml:"WrittenPolicyInd,omitempty" json:",omitempty"`
	AllHospitalsPolicyInd          CheckboxType                    `xml:"AllHospitalsPolicyInd"`
	MostHospitalsPolicyInd         CheckboxType                    `xml:"MostHospitalsPolicyInd"`
	IndivHospitalTailoredPolicyInd CheckboxType                    `xml:"IndivHospitalTailoredPolicyInd"`
	FPGReferenceFreeCareInd        bool                            `xml:"FPGReferenceFreeCareInd,omitempty" json:",omitempty"`
	Percent100Ind                  CheckboxType                    `xml:"Percent100Ind"`
	Percent150Ind                  CheckboxType                    `xml:"Percent150Ind"`
	Percent200Ind                  CheckboxType                    `xml:"Percent200Ind"`
	FreeCareOthPercentageGrp       *FreeCareOthPercentageGrp       `xml:"FreeCareOthPercentageGrp"`
	FPGReferenceDiscountedCareInd  bool                            `xml:"FPGReferenceDiscountedCareInd,omitempty" json:",omitempty"`
	Percent200DInd                 CheckboxType                    `xml:"Percent200DInd"`
	Percent250Ind                  CheckboxType                    `xml:"Percent250Ind"`
	Percent300Ind                  CheckboxType                    `xml:"Percent300Ind"`
	Percent350Ind                  CheckboxType                    `xml:"Percent350Ind"`
	Percent400Ind                  CheckboxType                    `xml:"Percent400Ind"`
	DiscountedCareOthPercentageGrp *DiscountedCareOthPercentageGrp `xml:"DiscountedCareOthPercentageGrp"`
	FreeCareMedicallyIndigentInd   bool                            `xml:"FreeCareMedicallyIndigentInd,omitempty" json:",omitempty"`
	FinancialAssistanceBudgetInd   bool                            `xml:"FinancialAssistanceBudgetInd,omitempty" json:",omitempty"`
	ExpensesExceedBudgetInd        bool                            `xml:"ExpensesExceedBudgetInd,omitempty" json:",omitempty"`
	UnableToProvideCareInd         bool                            `xml:"UnableToProvideCareInd,omitempty" json:",omitempty"`
	AnnualCommunityBnftReportInd   bool                            `xml:"AnnualCommunityBnftReportInd,omitempty" json:",omitempty"`
	ReportPublicallyAvailableInd   bool                            `xml:"ReportPublicallyAvailableInd,omitempty" json:",omitempty"`
	FinancialAssistanceAtCostTyp   *Form990SchHPartIGroup1Type     `xml:"FinancialAssistanceAtCostTyp,omitempty" json:",omitempty"`
	UnreimbursedMedicaidGrp        *Form990SchHPartIGroup1Type     `xml:"UnreimbursedMedicaidGrp,omitempty" json:",omitempty"`
	UnreimbursedCostsGrp           *Form990SchHPartIGroup1Type     `xml:"UnreimbursedCostsGrp,omitempty" json:",omitempty"`
	TotalFinancialAssistanceTyp    *Form990SchHPartIGroup1Type     `xml:"TotalFinancialAssistanceTyp,omitempty" json:",omitempty"`
	CommunityHealthServicesGrp     *Form990SchHPartIGroup1Type     `xml:"CommunityHealthServicesGrp,omitempty" json:",omitempty"`
	HealthProfessionsEducationGrp  *Form990SchHPartIGroup1Type     `xml:"HealthProfessionsEducationGrp,omitempty" json:",omitempty"`
	SubsidizedHealthServicesGrp    *Form990SchHPartIGroup1Type     `xml:"SubsidizedHealthServicesGrp,omitempty" json:",omitempty"`
	ResearchGrp                    *Form990SchHPartIGroup1Type     `xml:"ResearchGrp,omitempty" json:",omitempty"`
	CashAndInKindContributionsGrp  *Form990SchHPartIGroup1Type     `xml:"CashAndInKindContributionsGrp,omitempty" json:",omitempty"`
	TotalOtherBenefitsGrp          *Form990SchHPartIGroup1Type     `xml:"TotalOtherBenefitsGrp,omitempty" json:",omitempty"`
	TotalCommunityBenefitsGrp      *Form990SchHPartIGroup1Type     `xml:"TotalCommunityBenefitsGrp,omitempty" json:",omitempty"`
	PhysicalImprvAndHousingGrp     *Form990SchHPartIIGroup1Type    `xml:"PhysicalImprvAndHousingGrp,omitempty" json:",omitempty"`
	EconomicDevelopmentGrp         *Form990SchHPartIIGroup1Type    `xml:"EconomicDevelopmentGrp,omitempty" json:",omitempty"`
	CommunitySupportGrp            *Form990SchHPartIIGroup1Type    `xml:"CommunitySupportGrp,omitempty" json:",omitempty"`
	EnvironmentalImprovementsGrp   *Form990SchHPartIIGroup1Type    `xml:"EnvironmentalImprovementsGrp,omitempty" json:",omitempty"`
	LeadershipDevelopmentGrp       *Form990SchHPartIIGroup1Type    `xml:"LeadershipDevelopmentGrp,omitempty" json:",omitempty"`
	CoalitionBuildingGrp           *Form990SchHPartIIGroup1Type    `xml:"CoalitionBuildingGrp,omitempty" json:",omitempty"`
	HealthImprovementAdvocacyGrp   *Form990SchHPartIIGroup1Type    `xml:"HealthImprovementAdvocacyGrp,omitempty" json:",omitempty"`
	WorkforceDevelopmentGrp        *Form990SchHPartIIGroup1Type    `xml:"WorkforceDevelopmentGrp,omitempty" json:",omitempty"`
	OtherCommuntityBuildingActyGrp *Form990SchHPartIIGroup1Type    `xml:"OtherCommuntityBuildingActyGrp,omitempty" json:",omitempty"`
	TotalCommuntityBuildingActyGrp *Form990SchHPartIIGroup1Type    `xml:"TotalCommuntityBuildingActyGrp,omitempty" json:",omitempty"`
	BadDebtExpenseReportedInd      bool                            `xml:"BadDebtExpenseReportedInd,omitempty" json:",omitempty"`
	BadDebtExpenseAmt              int                             `xml:"BadDebtExpenseAmt,omitempty" json:",omitempty"`
	BadDebtExpenseAttributableAmt  int                             `xml:"BadDebtExpenseAttributableAmt,omitempty" json:",omitempty"`
	ReimbursedByMedicareAmt        int                             `xml:"ReimbursedByMedicareAmt,omitempty" json:",omitempty"`
	CostOfCareReimbursedByMedcrAmt int                             `xml:"CostOfCareReimbursedByMedcrAmt,omitempty" json:",omitempty"`
	MedicareSurplusOrShortfallAmt  int                             `xml:"MedicareSurplusOrShortfallAmt,omitempty" json:",omitempty"`
	CostingMethodologyUsedGrp      *CostingMethodologyUsedGrp      `xml:"CostingMethodologyUsedGrp,omitempty" json:",omitempty"`
	WrittenDebtCollectionPolicyInd bool                            `xml:"WrittenDebtCollectionPolicyInd,omitempty" json:",omitempty"`
	FinancialAssistancePrvsnInd    bool                            `xml:"FinancialAssistancePrvsnInd,omitempty" json:",omitempty"`
	ManagementCoAndJntVenturesGrp  []ManagementCoAndJntVenturesGrp `xml:"ManagementCoAndJntVenturesGrp,omitempty" json:",omitempty"`
	HospitalFacilitiesCnt          int                             `xml:"HospitalFacilitiesCnt"`
	HospitalFacilitiesGrp          []HospitalFacilitiesGrp         `xml:"HospitalFacilitiesGrp"`
	HospitalFcltyPoliciesPrctcGrp  []HospitalFcltyPoliciesPrctcGrp `xml:"HospitalFcltyPoliciesPrctcGrp,omitempty" json:",omitempty"`
	SupplementalInformationGrp     []SupplementalInformationGrp    `xml:"SupplementalInformationGrp,omitempty" json:",omitempty"`
	FacilityNum                    int                             `xml:"FacilityNum,omitempty" json:",omitempty"`
	OthHlthCareFcltsNotHospitalGrp *OthHlthCareFcltsNotHospitalGrp `xml:"OthHlthCareFcltsNotHospitalGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                     IdType                          `xml:"documentId,attr"`
	SoftwareId                     SoftwareIdType                  `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                          `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                          `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleH) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule H
type IRS990ScheduleHType struct {
	FinancialAssistancePolicyInd   bool                            `xml:"FinancialAssistancePolicyInd,omitempty" json:",omitempty"`
	WrittenPolicyInd               bool                            `xml:"WrittenPolicyInd,omitempty" json:",omitempty"`
	AllHospitalsPolicyInd          CheckboxType                    `xml:"AllHospitalsPolicyInd"`
	MostHospitalsPolicyInd         CheckboxType                    `xml:"MostHospitalsPolicyInd"`
	IndivHospitalTailoredPolicyInd CheckboxType                    `xml:"IndivHospitalTailoredPolicyInd"`
	FPGReferenceFreeCareInd        bool                            `xml:"FPGReferenceFreeCareInd,omitempty" json:",omitempty"`
	Percent100Ind                  CheckboxType                    `xml:"Percent100Ind"`
	Percent150Ind                  CheckboxType                    `xml:"Percent150Ind"`
	Percent200Ind                  CheckboxType                    `xml:"Percent200Ind"`
	FreeCareOthPercentageGrp       *FreeCareOthPercentageGrp       `xml:"FreeCareOthPercentageGrp"`
	FPGReferenceDiscountedCareInd  bool                            `xml:"FPGReferenceDiscountedCareInd,omitempty" json:",omitempty"`
	Percent200DInd                 CheckboxType                    `xml:"Percent200DInd"`
	Percent250Ind                  CheckboxType                    `xml:"Percent250Ind"`
	Percent300Ind                  CheckboxType                    `xml:"Percent300Ind"`
	Percent350Ind                  CheckboxType                    `xml:"Percent350Ind"`
	Percent400Ind                  CheckboxType                    `xml:"Percent400Ind"`
	DiscountedCareOthPercentageGrp *DiscountedCareOthPercentageGrp `xml:"DiscountedCareOthPercentageGrp"`
	FreeCareMedicallyIndigentInd   bool                            `xml:"FreeCareMedicallyIndigentInd,omitempty" json:",omitempty"`
	FinancialAssistanceBudgetInd   bool                            `xml:"FinancialAssistanceBudgetInd,omitempty" json:",omitempty"`
	ExpensesExceedBudgetInd        bool                            `xml:"ExpensesExceedBudgetInd,omitempty" json:",omitempty"`
	UnableToProvideCareInd         bool                            `xml:"UnableToProvideCareInd,omitempty" json:",omitempty"`
	AnnualCommunityBnftReportInd   bool                            `xml:"AnnualCommunityBnftReportInd,omitempty" json:",omitempty"`
	ReportPublicallyAvailableInd   bool                            `xml:"ReportPublicallyAvailableInd,omitempty" json:",omitempty"`
	FinancialAssistanceAtCostTyp   *Form990SchHPartIGroup1Type     `xml:"FinancialAssistanceAtCostTyp,omitempty" json:",omitempty"`
	UnreimbursedMedicaidGrp        *Form990SchHPartIGroup1Type     `xml:"UnreimbursedMedicaidGrp,omitempty" json:",omitempty"`
	UnreimbursedCostsGrp           *Form990SchHPartIGroup1Type     `xml:"UnreimbursedCostsGrp,omitempty" json:",omitempty"`
	TotalFinancialAssistanceTyp    *Form990SchHPartIGroup1Type     `xml:"TotalFinancialAssistanceTyp,omitempty" json:",omitempty"`
	CommunityHealthServicesGrp     *Form990SchHPartIGroup1Type     `xml:"CommunityHealthServicesGrp,omitempty" json:",omitempty"`
	HealthProfessionsEducationGrp  *Form990SchHPartIGroup1Type     `xml:"HealthProfessionsEducationGrp,omitempty" json:",omitempty"`
	SubsidizedHealthServicesGrp    *Form990SchHPartIGroup1Type     `xml:"SubsidizedHealthServicesGrp,omitempty" json:",omitempty"`
	ResearchGrp                    *Form990SchHPartIGroup1Type     `xml:"ResearchGrp,omitempty" json:",omitempty"`
	CashAndInKindContributionsGrp  *Form990SchHPartIGroup1Type     `xml:"CashAndInKindContributionsGrp,omitempty" json:",omitempty"`
	TotalOtherBenefitsGrp          *Form990SchHPartIGroup1Type     `xml:"TotalOtherBenefitsGrp,omitempty" json:",omitempty"`
	TotalCommunityBenefitsGrp      *Form990SchHPartIGroup1Type     `xml:"TotalCommunityBenefitsGrp,omitempty" json:",omitempty"`
	PhysicalImprvAndHousingGrp     *Form990SchHPartIIGroup1Type    `xml:"PhysicalImprvAndHousingGrp,omitempty" json:",omitempty"`
	EconomicDevelopmentGrp         *Form990SchHPartIIGroup1Type    `xml:"EconomicDevelopmentGrp,omitempty" json:",omitempty"`
	CommunitySupportGrp            *Form990SchHPartIIGroup1Type    `xml:"CommunitySupportGrp,omitempty" json:",omitempty"`
	EnvironmentalImprovementsGrp   *Form990SchHPartIIGroup1Type    `xml:"EnvironmentalImprovementsGrp,omitempty" json:",omitempty"`
	LeadershipDevelopmentGrp       *Form990SchHPartIIGroup1Type    `xml:"LeadershipDevelopmentGrp,omitempty" json:",omitempty"`
	CoalitionBuildingGrp           *Form990SchHPartIIGroup1Type    `xml:"CoalitionBuildingGrp,omitempty" json:",omitempty"`
	HealthImprovementAdvocacyGrp   *Form990SchHPartIIGroup1Type    `xml:"HealthImprovementAdvocacyGrp,omitempty" json:",omitempty"`
	WorkforceDevelopmentGrp        *Form990SchHPartIIGroup1Type    `xml:"WorkforceDevelopmentGrp,omitempty" json:",omitempty"`
	OtherCommuntityBuildingActyGrp *Form990SchHPartIIGroup1Type    `xml:"OtherCommuntityBuildingActyGrp,omitempty" json:",omitempty"`
	TotalCommuntityBuildingActyGrp *Form990SchHPartIIGroup1Type    `xml:"TotalCommuntityBuildingActyGrp,omitempty" json:",omitempty"`
	BadDebtExpenseReportedInd      bool                            `xml:"BadDebtExpenseReportedInd,omitempty" json:",omitempty"`
	BadDebtExpenseAmt              int                             `xml:"BadDebtExpenseAmt,omitempty" json:",omitempty"`
	BadDebtExpenseAttributableAmt  int                             `xml:"BadDebtExpenseAttributableAmt,omitempty" json:",omitempty"`
	ReimbursedByMedicareAmt        int                             `xml:"ReimbursedByMedicareAmt,omitempty" json:",omitempty"`
	CostOfCareReimbursedByMedcrAmt int                             `xml:"CostOfCareReimbursedByMedcrAmt,omitempty" json:",omitempty"`
	MedicareSurplusOrShortfallAmt  int                             `xml:"MedicareSurplusOrShortfallAmt,omitempty" json:",omitempty"`
	CostingMethodologyUsedGrp      *CostingMethodologyUsedGrp      `xml:"CostingMethodologyUsedGrp,omitempty" json:",omitempty"`
	WrittenDebtCollectionPolicyInd bool                            `xml:"WrittenDebtCollectionPolicyInd,omitempty" json:",omitempty"`
	FinancialAssistancePrvsnInd    bool                            `xml:"FinancialAssistancePrvsnInd,omitempty" json:",omitempty"`
	ManagementCoAndJntVenturesGrp  []ManagementCoAndJntVenturesGrp `xml:"ManagementCoAndJntVenturesGrp,omitempty" json:",omitempty"`
	HospitalFacilitiesCnt          int                             `xml:"HospitalFacilitiesCnt"`
	HospitalFacilitiesGrp          []HospitalFacilitiesGrp         `xml:"HospitalFacilitiesGrp"`
	HospitalFcltyPoliciesPrctcGrp  []HospitalFcltyPoliciesPrctcGrp `xml:"HospitalFcltyPoliciesPrctcGrp,omitempty" json:",omitempty"`
	SupplementalInformationGrp     []SupplementalInformationGrp    `xml:"SupplementalInformationGrp,omitempty" json:",omitempty"`
	FacilityNum                    int                             `xml:"FacilityNum,omitempty" json:",omitempty"`
	OthHlthCareFcltsNotHospitalGrp *OthHlthCareFcltsNotHospitalGrp `xml:"OthHlthCareFcltsNotHospitalGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleHType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleI struct {
	GrantRecordsMaintainedInd     bool                            `xml:"GrantRecordsMaintainedInd,omitempty" json:",omitempty"`
	RecipientTable                []RecipientTable                `xml:"RecipientTable,omitempty" json:",omitempty"`
	Total501c3OrgCnt              int                             `xml:"Total501c3OrgCnt,omitempty" json:",omitempty"`
	TotalOtherOrgCnt              int                             `xml:"TotalOtherOrgCnt,omitempty" json:",omitempty"`
	GrantsOtherAsstToIndivInUSGrp []GrantsOtherAsstToIndivInUSGrp `xml:"GrantsOtherAsstToIndivInUSGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                    IdType                          `xml:"documentId,attr"`
	SoftwareId                    SoftwareIdType                  `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum            string                          `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                  string                          `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleI) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule I
type IRS990ScheduleIType struct {
	GrantRecordsMaintainedInd     bool                            `xml:"GrantRecordsMaintainedInd,omitempty" json:",omitempty"`
	RecipientTable                []RecipientTable                `xml:"RecipientTable,omitempty" json:",omitempty"`
	Total501c3OrgCnt              int                             `xml:"Total501c3OrgCnt,omitempty" json:",omitempty"`
	TotalOtherOrgCnt              int                             `xml:"TotalOtherOrgCnt,omitempty" json:",omitempty"`
	GrantsOtherAsstToIndivInUSGrp []GrantsOtherAsstToIndivInUSGrp `xml:"GrantsOtherAsstToIndivInUSGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleIType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleJ struct {
	FirstClassOrCharterTravelInd   CheckboxType                    `xml:"FirstClassOrCharterTravelInd,omitempty" json:",omitempty"`
	TravelForCompanionsInd         CheckboxType                    `xml:"TravelForCompanionsInd,omitempty" json:",omitempty"`
	IdemnificationGrossUpPmtsInd   CheckboxType                    `xml:"IdemnificationGrossUpPmtsInd,omitempty" json:",omitempty"`
	DiscretionarySpendingAcctInd   CheckboxType                    `xml:"DiscretionarySpendingAcctInd,omitempty" json:",omitempty"`
	HousingAllowanceOrResidenceInd CheckboxType                    `xml:"HousingAllowanceOrResidenceInd,omitempty" json:",omitempty"`
	PaymentsForUseOfResidenceInd   CheckboxType                    `xml:"PaymentsForUseOfResidenceInd,omitempty" json:",omitempty"`
	ClubDuesOrFeesInd              CheckboxType                    `xml:"ClubDuesOrFeesInd,omitempty" json:",omitempty"`
	PersonalServicesInd            CheckboxType                    `xml:"PersonalServicesInd,omitempty" json:",omitempty"`
	WrittenPolicyRefTAndEExpnssInd bool                            `xml:"WrittenPolicyRefTAndEExpnssInd,omitempty" json:",omitempty"`
	SubstantiationRequiredInd      bool                            `xml:"SubstantiationRequiredInd,omitempty" json:",omitempty"`
	CompensationCommitteeInd       CheckboxType                    `xml:"CompensationCommitteeInd,omitempty" json:",omitempty"`
	IndependentConsultantInd       CheckboxType                    `xml:"IndependentConsultantInd,omitempty" json:",omitempty"`
	Form990OfOtherOrganizationsInd CheckboxType                    `xml:"Form990OfOtherOrganizationsInd,omitempty" json:",omitempty"`
	WrittenEmploymentContractInd   CheckboxType                    `xml:"WrittenEmploymentContractInd,omitempty" json:",omitempty"`
	CompensationSurveyInd          CheckboxType                    `xml:"CompensationSurveyInd,omitempty" json:",omitempty"`
	BoardOrCommitteeApprovalInd    CheckboxType                    `xml:"BoardOrCommitteeApprovalInd,omitempty" json:",omitempty"`
	SeverancePaymentInd            bool                            `xml:"SeverancePaymentInd,omitempty" json:",omitempty"`
	SupplementalNonqualRtrPlanInd  bool                            `xml:"SupplementalNonqualRtrPlanInd,omitempty" json:",omitempty"`
	EquityBasedCompArrngmInd       bool                            `xml:"EquityBasedCompArrngmInd,omitempty" json:",omitempty"`
	CompBasedOnRevenueOfFlngOrgInd bool                            `xml:"CompBasedOnRevenueOfFlngOrgInd,omitempty" json:",omitempty"`
	CompBsdOnRevRelatedOrgsInd     bool                            `xml:"CompBsdOnRevRelatedOrgsInd,omitempty" json:",omitempty"`
	CompBsdNetEarnsFlngOrgInd      bool                            `xml:"CompBsdNetEarnsFlngOrgInd,omitempty" json:",omitempty"`
	CompBsdNetEarnsRltdOrgsInd     bool                            `xml:"CompBsdNetEarnsRltdOrgsInd,omitempty" json:",omitempty"`
	AnyNonFixedPaymentsInd         bool                            `xml:"AnyNonFixedPaymentsInd,omitempty" json:",omitempty"`
	InitialContractExceptionInd    bool                            `xml:"InitialContractExceptionInd,omitempty" json:",omitempty"`
	RebuttablePresumptionProcInd   bool                            `xml:"RebuttablePresumptionProcInd,omitempty" json:",omitempty"`
	RltdOrgOfficerTrstKeyEmplGrp   []RltdOrgOfficerTrstKeyEmplGrp  `xml:"RltdOrgOfficerTrstKeyEmplGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                     IdType                          `xml:"documentId,attr"`
	SoftwareId                     SoftwareIdType                  `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                          `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                          `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleJ) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule J
type IRS990ScheduleJType struct {
	FirstClassOrCharterTravelInd   CheckboxType                    `xml:"FirstClassOrCharterTravelInd,omitempty" json:",omitempty"`
	TravelForCompanionsInd         CheckboxType                    `xml:"TravelForCompanionsInd,omitempty" json:",omitempty"`
	IdemnificationGrossUpPmtsInd   CheckboxType                    `xml:"IdemnificationGrossUpPmtsInd,omitempty" json:",omitempty"`
	DiscretionarySpendingAcctInd   CheckboxType                    `xml:"DiscretionarySpendingAcctInd,omitempty" json:",omitempty"`
	HousingAllowanceOrResidenceInd CheckboxType                    `xml:"HousingAllowanceOrResidenceInd,omitempty" json:",omitempty"`
	PaymentsForUseOfResidenceInd   CheckboxType                    `xml:"PaymentsForUseOfResidenceInd,omitempty" json:",omitempty"`
	ClubDuesOrFeesInd              CheckboxType                    `xml:"ClubDuesOrFeesInd,omitempty" json:",omitempty"`
	PersonalServicesInd            CheckboxType                    `xml:"PersonalServicesInd,omitempty" json:",omitempty"`
	WrittenPolicyRefTAndEExpnssInd bool                            `xml:"WrittenPolicyRefTAndEExpnssInd,omitempty" json:",omitempty"`
	SubstantiationRequiredInd      bool                            `xml:"SubstantiationRequiredInd,omitempty" json:",omitempty"`
	CompensationCommitteeInd       CheckboxType                    `xml:"CompensationCommitteeInd,omitempty" json:",omitempty"`
	IndependentConsultantInd       CheckboxType                    `xml:"IndependentConsultantInd,omitempty" json:",omitempty"`
	Form990OfOtherOrganizationsInd CheckboxType                    `xml:"Form990OfOtherOrganizationsInd,omitempty" json:",omitempty"`
	WrittenEmploymentContractInd   CheckboxType                    `xml:"WrittenEmploymentContractInd,omitempty" json:",omitempty"`
	CompensationSurveyInd          CheckboxType                    `xml:"CompensationSurveyInd,omitempty" json:",omitempty"`
	BoardOrCommitteeApprovalInd    CheckboxType                    `xml:"BoardOrCommitteeApprovalInd,omitempty" json:",omitempty"`
	SeverancePaymentInd            bool                            `xml:"SeverancePaymentInd,omitempty" json:",omitempty"`
	SupplementalNonqualRtrPlanInd  bool                            `xml:"SupplementalNonqualRtrPlanInd,omitempty" json:",omitempty"`
	EquityBasedCompArrngmInd       bool                            `xml:"EquityBasedCompArrngmInd,omitempty" json:",omitempty"`
	CompBasedOnRevenueOfFlngOrgInd bool                            `xml:"CompBasedOnRevenueOfFlngOrgInd,omitempty" json:",omitempty"`
	CompBsdOnRevRelatedOrgsInd     bool                            `xml:"CompBsdOnRevRelatedOrgsInd,omitempty" json:",omitempty"`
	CompBsdNetEarnsFlngOrgInd      bool                            `xml:"CompBsdNetEarnsFlngOrgInd,omitempty" json:",omitempty"`
	CompBsdNetEarnsRltdOrgsInd     bool                            `xml:"CompBsdNetEarnsRltdOrgsInd,omitempty" json:",omitempty"`
	AnyNonFixedPaymentsInd         bool                            `xml:"AnyNonFixedPaymentsInd,omitempty" json:",omitempty"`
	InitialContractExceptionInd    bool                            `xml:"InitialContractExceptionInd,omitempty" json:",omitempty"`
	RebuttablePresumptionProcInd   bool                            `xml:"RebuttablePresumptionProcInd,omitempty" json:",omitempty"`
	RltdOrgOfficerTrstKeyEmplGrp   []RltdOrgOfficerTrstKeyEmplGrp  `xml:"RltdOrgOfficerTrstKeyEmplGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleJType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleK struct {
	TaxExemptBondsIssuesGrp        []TaxExemptBondsIssuesGrpType        `xml:"TaxExemptBondsIssuesGrp,omitempty" json:",omitempty"`
	TaxExemptBondsProceedsGrp      []TaxExemptBondsProceedsGrpType      `xml:"TaxExemptBondsProceedsGrp,omitempty" json:",omitempty"`
	TaxExemptBondsPrivateBusUseGrp []TaxExemptBondsPrivateBusUseGrpType `xml:"TaxExemptBondsPrivateBusUseGrp,omitempty" json:",omitempty"`
	TaxExemptBondsArbitrageGrp     []TaxExemptBondsArbitrageGrpType     `xml:"TaxExemptBondsArbitrageGrp,omitempty" json:",omitempty"`
	ProceduresCorrectiveActionGrp  []ProceduresCorrectiveActionGrpType  `xml:"ProceduresCorrectiveActionGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail      `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                     IdType                               `xml:"documentId,attr"`
	SoftwareId                     SoftwareIdType                       `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                               `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                               `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleK) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule K
type IRS990ScheduleKType struct {
	TaxExemptBondsIssuesGrp        []TaxExemptBondsIssuesGrpType        `xml:"TaxExemptBondsIssuesGrp,omitempty" json:",omitempty"`
	TaxExemptBondsProceedsGrp      []TaxExemptBondsProceedsGrpType      `xml:"TaxExemptBondsProceedsGrp,omitempty" json:",omitempty"`
	TaxExemptBondsPrivateBusUseGrp []TaxExemptBondsPrivateBusUseGrpType `xml:"TaxExemptBondsPrivateBusUseGrp,omitempty" json:",omitempty"`
	TaxExemptBondsArbitrageGrp     []TaxExemptBondsArbitrageGrpType     `xml:"TaxExemptBondsArbitrageGrp,omitempty" json:",omitempty"`
	ProceduresCorrectiveActionGrp  []ProceduresCorrectiveActionGrpType  `xml:"ProceduresCorrectiveActionGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail      `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleKType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleL struct {
	DisqualifiedPersonExBnftTrGrp []DisqualifiedPersonExBnftTrGrpType `xml:"DisqualifiedPersonExBnftTrGrp,omitempty" json:",omitempty"`
	TaxImposedAmt                 int                                 `xml:"TaxImposedAmt,omitempty" json:",omitempty"`
	TaxReimbursedAmt              int                                 `xml:"TaxReimbursedAmt,omitempty" json:",omitempty"`
	LoansBtwnOrgInterestedPrsnGrp []LoansBtwnOrgInterestedPrsnGrpType `xml:"LoansBtwnOrgInterestedPrsnGrp,omitempty" json:",omitempty"`
	TotalBalanceDueAmt            int                                 `xml:"TotalBalanceDueAmt,omitempty" json:",omitempty"`
	GrntAsstBnftInterestedPrsnGrp []GrntAsstBnftInterestedPrsnGrp     `xml:"GrntAsstBnftInterestedPrsnGrp,omitempty" json:",omitempty"`
	BusTrInvolveInterestedPrsnGrp []BusTrInvolveInterestedPrsnGrp     `xml:"BusTrInvolveInterestedPrsnGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail []SupplementalInformationDetail     `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                    IdType                              `xml:"documentId,attr"`
	SoftwareId                    SoftwareIdType                      `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum            string                              `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                  string                              `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleL) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule L
type IRS990ScheduleLType struct {
	DisqualifiedPersonExBnftTrGrp []DisqualifiedPersonExBnftTrGrpType `xml:"DisqualifiedPersonExBnftTrGrp,omitempty" json:",omitempty"`
	TaxImposedAmt                 int                                 `xml:"TaxImposedAmt,omitempty" json:",omitempty"`
	TaxReimbursedAmt              int                                 `xml:"TaxReimbursedAmt,omitempty" json:",omitempty"`
	LoansBtwnOrgInterestedPrsnGrp []LoansBtwnOrgInterestedPrsnGrpType `xml:"LoansBtwnOrgInterestedPrsnGrp,omitempty" json:",omitempty"`
	TotalBalanceDueAmt            int                                 `xml:"TotalBalanceDueAmt,omitempty" json:",omitempty"`
	GrntAsstBnftInterestedPrsnGrp []GrntAsstBnftInterestedPrsnGrp     `xml:"GrntAsstBnftInterestedPrsnGrp,omitempty" json:",omitempty"`
	BusTrInvolveInterestedPrsnGrp []BusTrInvolveInterestedPrsnGrp     `xml:"BusTrInvolveInterestedPrsnGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail []SupplementalInformationDetail     `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleLType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleM struct {
	WorksOfArtGrp                  *Form990SchMGroup1Type          `xml:"WorksOfArtGrp,omitempty" json:",omitempty"`
	ArtHistoricalTreasuresGrp      *Form990SchMGroup1Type          `xml:"ArtHistoricalTreasuresGrp,omitempty" json:",omitempty"`
	ArtFractionalInterestGrp       *Form990SchMGroup1Type          `xml:"ArtFractionalInterestGrp,omitempty" json:",omitempty"`
	BooksAndPublicationsGrp        *Form990SchMGroup2Type          `xml:"BooksAndPublicationsGrp,omitempty" json:",omitempty"`
	ClothingAndHouseholdGoodsGrp   *Form990SchMGroup2Type          `xml:"ClothingAndHouseholdGoodsGrp,omitempty" json:",omitempty"`
	CarsAndOtherVehiclesGrp        *Form990SchMGroup1Type          `xml:"CarsAndOtherVehiclesGrp,omitempty" json:",omitempty"`
	BoatsAndPlanesGrp              *Form990SchMGroup1Type          `xml:"BoatsAndPlanesGrp,omitempty" json:",omitempty"`
	IntellectualPropertyGrp        *Form990SchMGroup1Type          `xml:"IntellectualPropertyGrp,omitempty" json:",omitempty"`
	SecuritiesPubliclyTradedGrp    *Form990SchMGroup1Type          `xml:"SecuritiesPubliclyTradedGrp,omitempty" json:",omitempty"`
	SecuritiesCloselyHeldStockGrp  *Form990SchMGroup1Type          `xml:"SecuritiesCloselyHeldStockGrp,omitempty" json:",omitempty"`
	SecurPrtnrshpTrustIntrstsGrp   *Form990SchMGroup1Type          `xml:"SecurPrtnrshpTrustIntrstsGrp,omitempty" json:",omitempty"`
	SecuritiesMiscellaneousGrp     *Form990SchMGroup1Type          `xml:"SecuritiesMiscellaneousGrp,omitempty" json:",omitempty"`
	QualifiedContribHistStructGrp  *Form990SchMGroup1Type          `xml:"QualifiedContribHistStructGrp,omitempty" json:",omitempty"`
	QualifiedContribOtherGrp       *Form990SchMGroup1Type          `xml:"QualifiedContribOtherGrp,omitempty" json:",omitempty"`
	RealEstateResidentialGrp       *Form990SchMGroup1Type          `xml:"RealEstateResidentialGrp,omitempty" json:",omitempty"`
	RealEstateCommercialGrp        *Form990SchMGroup1Type          `xml:"RealEstateCommercialGrp,omitempty" json:",omitempty"`
	RealEstateOtherGrp             *Form990SchMGroup1Type          `xml:"RealEstateOtherGrp,omitempty" json:",omitempty"`
	CollectiblesGrp                *Form990SchMGroup1Type          `xml:"CollectiblesGrp,omitempty" json:",omitempty"`
	FoodInventoryGrp               *Form990SchMGroup1Type          `xml:"FoodInventoryGrp,omitempty" json:",omitempty"`
	DrugsAndMedicalSuppliesGrp     *Form990SchMGroup1Type          `xml:"DrugsAndMedicalSuppliesGrp,omitempty" json:",omitempty"`
	TaxidermyGrp                   *Form990SchMGroup1Type          `xml:"TaxidermyGrp,omitempty" json:",omitempty"`
	HistoricalArtifactsGrp         *Form990SchMGroup1Type          `xml:"HistoricalArtifactsGrp,omitempty" json:",omitempty"`
	ScientificSpecimensGrp         *Form990SchMGroup1Type          `xml:"ScientificSpecimensGrp,omitempty" json:",omitempty"`
	ArchaeologicalArtifactsGrp     *Form990SchMGroup1Type          `xml:"ArchaeologicalArtifactsGrp,omitempty" json:",omitempty"`
	OtherNonCashContriTableGrp     []Form990SchMGroup3Type         `xml:"OtherNonCashContriTableGrp,omitempty" json:",omitempty"`
	Form8283ReceivedCnt            int                             `xml:"Form8283ReceivedCnt,omitempty" json:",omitempty"`
	AnyPropertyThatMustBeHeldInd   bool                            `xml:"AnyPropertyThatMustBeHeldInd,omitempty" json:",omitempty"`
	ReviewProcessUnusualNCGiftsInd bool                            `xml:"ReviewProcessUnusualNCGiftsInd,omitempty" json:",omitempty"`
	ThirdPartiesUsedInd            bool                            `xml:"ThirdPartiesUsedInd,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                     IdType                          `xml:"documentId,attr"`
	SoftwareId                     SoftwareIdType                  `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                          `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                          `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleM) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule M
type IRS990ScheduleMType struct {
	WorksOfArtGrp                  *Form990SchMGroup1Type          `xml:"WorksOfArtGrp,omitempty" json:",omitempty"`
	ArtHistoricalTreasuresGrp      *Form990SchMGroup1Type          `xml:"ArtHistoricalTreasuresGrp,omitempty" json:",omitempty"`
	ArtFractionalInterestGrp       *Form990SchMGroup1Type          `xml:"ArtFractionalInterestGrp,omitempty" json:",omitempty"`
	BooksAndPublicationsGrp        *Form990SchMGroup2Type          `xml:"BooksAndPublicationsGrp,omitempty" json:",omitempty"`
	ClothingAndHouseholdGoodsGrp   *Form990SchMGroup2Type          `xml:"ClothingAndHouseholdGoodsGrp,omitempty" json:",omitempty"`
	CarsAndOtherVehiclesGrp        *Form990SchMGroup1Type          `xml:"CarsAndOtherVehiclesGrp,omitempty" json:",omitempty"`
	BoatsAndPlanesGrp              *Form990SchMGroup1Type          `xml:"BoatsAndPlanesGrp,omitempty" json:",omitempty"`
	IntellectualPropertyGrp        *Form990SchMGroup1Type          `xml:"IntellectualPropertyGrp,omitempty" json:",omitempty"`
	SecuritiesPubliclyTradedGrp    *Form990SchMGroup1Type          `xml:"SecuritiesPubliclyTradedGrp,omitempty" json:",omitempty"`
	SecuritiesCloselyHeldStockGrp  *Form990SchMGroup1Type          `xml:"SecuritiesCloselyHeldStockGrp,omitempty" json:",omitempty"`
	SecurPrtnrshpTrustIntrstsGrp   *Form990SchMGroup1Type          `xml:"SecurPrtnrshpTrustIntrstsGrp,omitempty" json:",omitempty"`
	SecuritiesMiscellaneousGrp     *Form990SchMGroup1Type          `xml:"SecuritiesMiscellaneousGrp,omitempty" json:",omitempty"`
	QualifiedContribHistStructGrp  *Form990SchMGroup1Type          `xml:"QualifiedContribHistStructGrp,omitempty" json:",omitempty"`
	QualifiedContribOtherGrp       *Form990SchMGroup1Type          `xml:"QualifiedContribOtherGrp,omitempty" json:",omitempty"`
	RealEstateResidentialGrp       *Form990SchMGroup1Type          `xml:"RealEstateResidentialGrp,omitempty" json:",omitempty"`
	RealEstateCommercialGrp        *Form990SchMGroup1Type          `xml:"RealEstateCommercialGrp,omitempty" json:",omitempty"`
	RealEstateOtherGrp             *Form990SchMGroup1Type          `xml:"RealEstateOtherGrp,omitempty" json:",omitempty"`
	CollectiblesGrp                *Form990SchMGroup1Type          `xml:"CollectiblesGrp,omitempty" json:",omitempty"`
	FoodInventoryGrp               *Form990SchMGroup1Type          `xml:"FoodInventoryGrp,omitempty" json:",omitempty"`
	DrugsAndMedicalSuppliesGrp     *Form990SchMGroup1Type          `xml:"DrugsAndMedicalSuppliesGrp,omitempty" json:",omitempty"`
	TaxidermyGrp                   *Form990SchMGroup1Type          `xml:"TaxidermyGrp,omitempty" json:",omitempty"`
	HistoricalArtifactsGrp         *Form990SchMGroup1Type          `xml:"HistoricalArtifactsGrp,omitempty" json:",omitempty"`
	ScientificSpecimensGrp         *Form990SchMGroup1Type          `xml:"ScientificSpecimensGrp,omitempty" json:",omitempty"`
	ArchaeologicalArtifactsGrp     *Form990SchMGroup1Type          `xml:"ArchaeologicalArtifactsGrp,omitempty" json:",omitempty"`
	OtherNonCashContriTableGrp     []Form990SchMGroup3Type         `xml:"OtherNonCashContriTableGrp,omitempty" json:",omitempty"`
	Form8283ReceivedCnt            int                             `xml:"Form8283ReceivedCnt,omitempty" json:",omitempty"`
	AnyPropertyThatMustBeHeldInd   bool                            `xml:"AnyPropertyThatMustBeHeldInd,omitempty" json:",omitempty"`
	ReviewProcessUnusualNCGiftsInd bool                            `xml:"ReviewProcessUnusualNCGiftsInd,omitempty" json:",omitempty"`
	ThirdPartiesUsedInd            bool                            `xml:"ThirdPartiesUsedInd,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleMType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleN struct {
	LiquidationOfAssetsTableGrp   *LiquidationOfAssetsTableGrp    `xml:"LiquidationOfAssetsTableGrp,omitempty" json:",omitempty"`
	DirectorOfSuccessorInd        bool                            `xml:"DirectorOfSuccessorInd,omitempty" json:",omitempty"`
	EmployeeOfSuccessorInd        bool                            `xml:"EmployeeOfSuccessorInd,omitempty" json:",omitempty"`
	OwnerOfSuccessorInd           bool                            `xml:"OwnerOfSuccessorInd,omitempty" json:",omitempty"`
	ReceiveCompensationInd        bool                            `xml:"ReceiveCompensationInd,omitempty" json:",omitempty"`
	AssetsDistributedInd          bool                            `xml:"AssetsDistributedInd,omitempty" json:",omitempty"`
	RequiredToNotifyAGInd         bool                            `xml:"RequiredToNotifyAGInd,omitempty" json:",omitempty"`
	AttorneyGeneralNotifiedInd    bool                            `xml:"AttorneyGeneralNotifiedInd,omitempty" json:",omitempty"`
	LiabilitiesPaidInd            bool                            `xml:"LiabilitiesPaidInd,omitempty" json:",omitempty"`
	BondsOutstandingInd           bool                            `xml:"BondsOutstandingInd,omitempty" json:",omitempty"`
	BondLiabilitiesDischargedInd  bool                            `xml:"BondLiabilitiesDischargedInd,omitempty" json:",omitempty"`
	DispositionOfAssetsDetail     []Form990SchNGroup1Type         `xml:"DispositionOfAssetsDetail,omitempty" json:",omitempty"`
	DirectorOfSuccessor2Ind       bool                            `xml:"DirectorOfSuccessor2Ind,omitempty" json:",omitempty"`
	EmployeeOfSuccessor2Ind       bool                            `xml:"EmployeeOfSuccessor2Ind,omitempty" json:",omitempty"`
	OwnerOfSuccessor2Ind          bool                            `xml:"OwnerOfSuccessor2Ind,omitempty" json:",omitempty"`
	ReceiveCompensation2Ind       bool                            `xml:"ReceiveCompensation2Ind,omitempty" json:",omitempty"`
	SupplementalInformationDetail []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                    IdType                          `xml:"documentId,attr"`
	SoftwareId                    SoftwareIdType                  `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum            string                          `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                  string                          `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleN) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule N
type IRS990ScheduleNType struct {
	LiquidationOfAssetsTableGrp   *LiquidationOfAssetsTableGrp    `xml:"LiquidationOfAssetsTableGrp,omitempty" json:",omitempty"`
	DirectorOfSuccessorInd        bool                            `xml:"DirectorOfSuccessorInd,omitempty" json:",omitempty"`
	EmployeeOfSuccessorInd        bool                            `xml:"EmployeeOfSuccessorInd,omitempty" json:",omitempty"`
	OwnerOfSuccessorInd           bool                            `xml:"OwnerOfSuccessorInd,omitempty" json:",omitempty"`
	ReceiveCompensationInd        bool                            `xml:"ReceiveCompensationInd,omitempty" json:",omitempty"`
	AssetsDistributedInd          bool                            `xml:"AssetsDistributedInd,omitempty" json:",omitempty"`
	RequiredToNotifyAGInd         bool                            `xml:"RequiredToNotifyAGInd,omitempty" json:",omitempty"`
	AttorneyGeneralNotifiedInd    bool                            `xml:"AttorneyGeneralNotifiedInd,omitempty" json:",omitempty"`
	LiabilitiesPaidInd            bool                            `xml:"LiabilitiesPaidInd,omitempty" json:",omitempty"`
	BondsOutstandingInd           bool                            `xml:"BondsOutstandingInd,omitempty" json:",omitempty"`
	BondLiabilitiesDischargedInd  bool                            `xml:"BondLiabilitiesDischargedInd,omitempty" json:",omitempty"`
	DispositionOfAssetsDetail     []Form990SchNGroup1Type         `xml:"DispositionOfAssetsDetail,omitempty" json:",omitempty"`
	DirectorOfSuccessor2Ind       bool                            `xml:"DirectorOfSuccessor2Ind,omitempty" json:",omitempty"`
	EmployeeOfSuccessor2Ind       bool                            `xml:"EmployeeOfSuccessor2Ind,omitempty" json:",omitempty"`
	OwnerOfSuccessor2Ind          bool                            `xml:"OwnerOfSuccessor2Ind,omitempty" json:",omitempty"`
	ReceiveCompensation2Ind       bool                            `xml:"ReceiveCompensation2Ind,omitempty" json:",omitempty"`
	SupplementalInformationDetail []SupplementalInformationDetail `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleNType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleO struct {
	SupplementalInformationDetail []SupplementalInformationDetailType `xml:"SupplementalInformationDetail"`
	DocumentId                    IdType                              `xml:"documentId,attr"`
	SoftwareId                    SoftwareIdType                      `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum            string                              `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                  string                              `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleO) Validate() error {
	return utils.Validate(&r)
}

// Content model for IRS Form 990 Schedule O
type IRS990ScheduleOType struct {
	SupplementalInformationDetail []SupplementalInformationDetailType `xml:"SupplementalInformationDetail"`
}

func (r IRS990ScheduleOType) Validate() error {
	return utils.Validate(&r)
}

type IRS990ScheduleR struct {
	IdDisregardedEntitiesGrp       []IdDisregardedEntitiesGrp       `xml:"IdDisregardedEntitiesGrp,omitempty" json:",omitempty"`
	IdRelatedTaxExemptOrgGrp       []IdRelatedTaxExemptOrgGrp       `xml:"IdRelatedTaxExemptOrgGrp,omitempty" json:",omitempty"`
	IdRelatedOrgTxblPartnershipGrp []IdRelatedOrgTxblPartnershipGrp `xml:"IdRelatedOrgTxblPartnershipGrp,omitempty" json:",omitempty"`
	IdRelatedOrgTxblCorpTrGrp      []IdRelatedOrgTxblCorpTrGrp      `xml:"IdRelatedOrgTxblCorpTrGrp,omitempty" json:",omitempty"`
	ReceiptOfIntAnntsRntsRyltsInd  bool                             `xml:"ReceiptOfIntAnntsRntsRyltsInd,omitempty" json:",omitempty"`
	GiftGrntOrCapContriToOthOrgInd bool                             `xml:"GiftGrntOrCapContriToOthOrgInd,omitempty" json:",omitempty"`
	GiftGrntCapContriFromOthOrgInd bool                             `xml:"GiftGrntCapContriFromOthOrgInd,omitempty" json:",omitempty"`
	LoansOrGuaranteesToOtherOrgInd bool                             `xml:"LoansOrGuaranteesToOtherOrgInd,omitempty" json:",omitempty"`
	LoansOrGuaranteesFromOthOrgInd bool                             `xml:"LoansOrGuaranteesFromOthOrgInd,omitempty" json:",omitempty"`
	DivRelatedOrganizationInd      bool                             `xml:"DivRelatedOrganizationInd,omitempty" json:",omitempty"`
	AssetSaleToOtherOrgInd         bool                             `xml:"AssetSaleToOtherOrgInd,omitempty" json:",omitempty"`
	AssetPurchaseFromOtherOrgInd   bool                             `xml:"AssetPurchaseFromOtherOrgInd,omitempty" json:",omitempty"`
	AssetExchangeInd               bool                             `xml:"AssetExchangeInd,omitempty" json:",omitempty"`
	RentalOfFacilitiesToOthOrgInd  bool                             `xml:"RentalOfFacilitiesToOthOrgInd,omitempty" json:",omitempty"`
	RentalOfFcltsFromOthOrgInd     bool                             `xml:"RentalOfFcltsFromOthOrgInd,omitempty" json:",omitempty"`
	PerformOfServicesForOthOrgInd  bool                             `xml:"PerformOfServicesForOthOrgInd,omitempty" json:",omitempty"`
	PerformOfServicesByOtherOrgInd bool                             `xml:"PerformOfServicesByOtherOrgInd,omitempty" json:",omitempty"`
	SharingOfFacilitiesInd         bool                             `xml:"SharingOfFacilitiesInd,omitempty" json:",omitempty"`
	PaidEmployeesSharingInd        bool                             `xml:"PaidEmployeesSharingInd,omitempty" json:",omitempty"`
	ReimbursementPaidToOtherOrgInd bool                             `xml:"ReimbursementPaidToOtherOrgInd,omitempty" json:",omitempty"`
	ReimbursementPaidByOtherOrgInd bool                             `xml:"ReimbursementPaidByOtherOrgInd,omitempty" json:",omitempty"`
	TransferToOtherOrgInd          bool                             `xml:"TransferToOtherOrgInd,omitempty" json:",omitempty"`
	TransferFromOtherOrgInd        bool                             `xml:"TransferFromOtherOrgInd,omitempty" json:",omitempty"`
	TransactionsRelatedOrgGrp      []TransactionsRelatedOrgGrpType  `xml:"TransactionsRelatedOrgGrp,omitempty" json:",omitempty"`
	UnrelatedOrgTxblPartnershipGrp []UnrelatedOrgTxblPartnershipGrp `xml:"UnrelatedOrgTxblPartnershipGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail  `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
	DocumentId                     IdType                           `xml:"documentId,attr"`
	SoftwareId                     SoftwareIdType                   `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                           `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                           `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleR) Validate() error {
	return utils.Validate(&r)
}

// Content model for Form 990 Schedule R
type IRS990ScheduleRType struct {
	IdDisregardedEntitiesGrp       []IdDisregardedEntitiesGrp       `xml:"IdDisregardedEntitiesGrp,omitempty" json:",omitempty"`
	IdRelatedTaxExemptOrgGrp       []IdRelatedTaxExemptOrgGrp       `xml:"IdRelatedTaxExemptOrgGrp,omitempty" json:",omitempty"`
	IdRelatedOrgTxblPartnershipGrp []IdRelatedOrgTxblPartnershipGrp `xml:"IdRelatedOrgTxblPartnershipGrp,omitempty" json:",omitempty"`
	IdRelatedOrgTxblCorpTrGrp      []IdRelatedOrgTxblCorpTrGrp      `xml:"IdRelatedOrgTxblCorpTrGrp,omitempty" json:",omitempty"`
	ReceiptOfIntAnntsRntsRyltsInd  bool                             `xml:"ReceiptOfIntAnntsRntsRyltsInd,omitempty" json:",omitempty"`
	GiftGrntOrCapContriToOthOrgInd bool                             `xml:"GiftGrntOrCapContriToOthOrgInd,omitempty" json:",omitempty"`
	GiftGrntCapContriFromOthOrgInd bool                             `xml:"GiftGrntCapContriFromOthOrgInd,omitempty" json:",omitempty"`
	LoansOrGuaranteesToOtherOrgInd bool                             `xml:"LoansOrGuaranteesToOtherOrgInd,omitempty" json:",omitempty"`
	LoansOrGuaranteesFromOthOrgInd bool                             `xml:"LoansOrGuaranteesFromOthOrgInd,omitempty" json:",omitempty"`
	DivRelatedOrganizationInd      bool                             `xml:"DivRelatedOrganizationInd,omitempty" json:",omitempty"`
	AssetSaleToOtherOrgInd         bool                             `xml:"AssetSaleToOtherOrgInd,omitempty" json:",omitempty"`
	AssetPurchaseFromOtherOrgInd   bool                             `xml:"AssetPurchaseFromOtherOrgInd,omitempty" json:",omitempty"`
	AssetExchangeInd               bool                             `xml:"AssetExchangeInd,omitempty" json:",omitempty"`
	RentalOfFacilitiesToOthOrgInd  bool                             `xml:"RentalOfFacilitiesToOthOrgInd,omitempty" json:",omitempty"`
	RentalOfFcltsFromOthOrgInd     bool                             `xml:"RentalOfFcltsFromOthOrgInd,omitempty" json:",omitempty"`
	PerformOfServicesForOthOrgInd  bool                             `xml:"PerformOfServicesForOthOrgInd,omitempty" json:",omitempty"`
	PerformOfServicesByOtherOrgInd bool                             `xml:"PerformOfServicesByOtherOrgInd,omitempty" json:",omitempty"`
	SharingOfFacilitiesInd         bool                             `xml:"SharingOfFacilitiesInd,omitempty" json:",omitempty"`
	PaidEmployeesSharingInd        bool                             `xml:"PaidEmployeesSharingInd,omitempty" json:",omitempty"`
	ReimbursementPaidToOtherOrgInd bool                             `xml:"ReimbursementPaidToOtherOrgInd,omitempty" json:",omitempty"`
	ReimbursementPaidByOtherOrgInd bool                             `xml:"ReimbursementPaidByOtherOrgInd,omitempty" json:",omitempty"`
	TransferToOtherOrgInd          bool                             `xml:"TransferToOtherOrgInd,omitempty" json:",omitempty"`
	TransferFromOtherOrgInd        bool                             `xml:"TransferFromOtherOrgInd,omitempty" json:",omitempty"`
	TransactionsRelatedOrgGrp      []TransactionsRelatedOrgGrpType  `xml:"TransactionsRelatedOrgGrp,omitempty" json:",omitempty"`
	UnrelatedOrgTxblPartnershipGrp []UnrelatedOrgTxblPartnershipGrp `xml:"UnrelatedOrgTxblPartnershipGrp,omitempty" json:",omitempty"`
	SupplementalInformationDetail  []SupplementalInformationDetail  `xml:"SupplementalInformationDetail,omitempty" json:",omitempty"`
}

func (r IRS990ScheduleRType) Validate() error {
	return utils.Validate(&r)
}

type AffiliateListing struct {
	AffiliateListingGrp []AffiliateListingGrpType `xml:"AffiliateListingGrp,omitempty" json:",omitempty"`
	DocumentId          IdType                    `xml:"documentId,attr"`
	SoftwareId          SoftwareIdType            `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum  string                    `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName        string                    `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r AffiliateListing) Validate() error {
	return utils.Validate(&r)
}

// Content model for affiliate listing
type AffiliateListingType struct {
	AffiliateListingGrp []AffiliateListingGrpType `xml:"AffiliateListingGrp,omitempty" json:",omitempty"`
}

func (r AffiliateListingType) Validate() error {
	return utils.Validate(&r)
}

type ReasonableCauseExplanation struct {
	ExplanationTxt     string         `xml:"ExplanationTxt,omitempty" json:",omitempty"`
	DocumentId         IdType         `xml:"documentId,attr"`
	SoftwareId         SoftwareIdType `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum string         `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName       string         `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r ReasonableCauseExplanation) Validate() error {
	return utils.Validate(&r)
}

// Content model for reasonable cause explanation
type ReasonableCauseExplanationType struct {
	ExplanationTxt string `xml:"ExplanationTxt,omitempty" json:",omitempty"`
}

func (r ReasonableCauseExplanationType) Validate() error {
	return utils.Validate(&r)
}

type AffiliatedGroupAttachment struct {
	MeduimExplanationTxt string         `xml:"MeduimExplanationTxt,omitempty" json:",omitempty"`
	DocumentId           IdType         `xml:"documentId,attr"`
	SoftwareId           SoftwareIdType `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum   string         `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName         string         `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r AffiliatedGroupAttachment) Validate() error {
	return utils.Validate(&r)
}

// Content model for Affiliated Group Attachment
type AffiliatedGroupAttachmentType struct {
	MeduimExplanationTxt string `xml:"MeduimExplanationTxt,omitempty" json:",omitempty"`
}

func (r AffiliatedGroupAttachmentType) Validate() error {
	return utils.Validate(&r)
}

type AffiliatedGroupSchedule struct {
	AffiliatedScheduleGrp []AffiliatedScheduleGrp `xml:"AffiliatedScheduleGrp,omitempty" json:",omitempty"`
	DocumentId            IdType                  `xml:"documentId,attr"`
	SoftwareId            SoftwareIdType          `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum    string                  `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName          string                  `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r AffiliatedGroupSchedule) Validate() error {
	return utils.Validate(&r)
}

// Content model for Affiliated Group Schedule
type AffiliatedGroupScheduleType struct {
	AffiliatedScheduleGrp []AffiliatedScheduleGrp `xml:"AffiliatedScheduleGrp,omitempty" json:",omitempty"`
}

func (r AffiliatedGroupScheduleType) Validate() error {
	return utils.Validate(&r)
}

type AveragingAttachment struct {
	ExplanationTxt     string         `xml:"ExplanationTxt,omitempty" json:",omitempty"`
	DocumentId         IdType         `xml:"documentId,attr"`
	SoftwareId         SoftwareIdType `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum string         `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName       string         `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r AveragingAttachment) Validate() error {
	return utils.Validate(&r)
}

// Content model for averaging attachment
type AveragingAttachmentType struct {
	ExplanationTxt string `xml:"ExplanationTxt,omitempty" json:",omitempty"`
}

func (r AveragingAttachmentType) Validate() error {
	return utils.Validate(&r)
}

type RequestForCopyAttachment struct {
	CopyReceiverGovernmentCd []GovernmentCodeType `xml:"CopyReceiverGovernmentCd"`
	PersonNm                 PersonNameType       `xml:"PersonNm"`
	PersonTitleTxt           PersonTitleType      `xml:"PersonTitleTxt"`
	PIN                      PINType              `xml:"PIN"`
	Dt                       DateType             `xml:"Dt"`
	DocumentId               IdType               `xml:"documentId,attr"`
	SoftwareId               SoftwareIdType       `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum       string               `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName             string               `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r RequestForCopyAttachment) Validate() error {
	return utils.Validate(&r)
}

// Content model for Request For Copy Attachment
type RequestForCopyAttachmentType struct {
	CopyReceiverGovernmentCd []GovernmentCodeType `xml:"CopyReceiverGovernmentCd"`
	PersonNm                 PersonNameType       `xml:"PersonNm"`
	PersonTitleTxt           PersonTitleType      `xml:"PersonTitleTxt"`
	PIN                      PINType              `xml:"PIN"`
	Dt                       DateType             `xml:"Dt"`
}

func (r RequestForCopyAttachmentType) Validate() error {
	return utils.Validate(&r)
}

type BinaryAttachment struct {
	DocumentTypeCd        DocumentTypeCd `xml:"DocumentTypeCd"`
	Desc                  string         `xml:"Desc"`
	AttachmentLocationTxt string         `xml:"AttachmentLocationTxt"`
	DocumentId            IdType         `xml:"documentId,attr"`
	SoftwareId            SoftwareIdType `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum    string         `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName          string         `xml:"documentName,attr,omitempty" json:",omitempty"`
}

func (r BinaryAttachment) Validate() error {
	return utils.Validate(&r)
}

// Content model for Binary Attachment
type BinaryAttachmentType struct {
	DocumentTypeCd        DocumentTypeCd `xml:"DocumentTypeCd"`
	Desc                  string         `xml:"Desc"`
	AttachmentLocationTxt string         `xml:"AttachmentLocationTxt"`
}

func (r BinaryAttachmentType) Validate() error {
	return utils.Validate(&r)
}
