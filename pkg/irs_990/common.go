// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import (
	"github.com/moov-io/1120x/pkg/utils"
)

type AccountActivitiesOutsideUSGrpType struct {
	RegionTxt                    string `xml:"RegionTxt,omitempty" json:",omitempty"`
	OfficesCnt                   int    `xml:"OfficesCnt,omitempty" json:",omitempty"`
	EmployeeCnt                  int    `xml:"EmployeeCnt,omitempty" json:",omitempty"`
	TypeOfActivitiesConductedTxt string `xml:"TypeOfActivitiesConductedTxt,omitempty" json:",omitempty"`
	SpecificServicesProvidedTxt  string `xml:"SpecificServicesProvidedTxt,omitempty" json:",omitempty"`
	RegionTotalExpendituresAmt   int    `xml:"RegionTotalExpendituresAmt,omitempty" json:",omitempty"`
}

func (r AccountActivitiesOutsideUSGrpType) Validate() error {
	return utils.Validate(&r)
}

type ActivitiesConductedPrtshpInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ActivitiesConductedPrtshpInd) Validate() error {
	return utils.Validate(&r)
}

type AdjustedNetIncomeGrp struct {
	NetSTCapitalGainAdjNetIncmGrp Form990SchAPartVGroup1Type `xml:"NetSTCapitalGainAdjNetIncmGrp"`
	RecoveriesPYDistributionsGrp  Form990SchAPartVGroup1Type `xml:"RecoveriesPYDistributionsGrp"`
	OtherGrossIncomeGrp           Form990SchAPartVGroup1Type `xml:"OtherGrossIncomeGrp"`
	AdjustedGrossIncomeGrp        Form990SchAPartVGroup1Type `xml:"AdjustedGrossIncomeGrp"`
	DepreciationDepletionGrp      Form990SchAPartVGroup1Type `xml:"DepreciationDepletionGrp"`
	ProductionIncomeGrp           Form990SchAPartVGroup1Type `xml:"ProductionIncomeGrp"`
	OtherExpensesGrp              Form990SchAPartVGroup1Type `xml:"OtherExpensesGrp"`
	TotalAdjustedNetIncomeGrp     Form990SchAPartVGroup2Type `xml:"TotalAdjustedNetIncomeGrp"`
}

func (r AdjustedNetIncomeGrp) Validate() error {
	return utils.Validate(&r)
}

// Content model for affiliate
type AffiliateListingGrpType struct {
	BusinessName           *BusinessNameType       `xml:"BusinessName,omitempty" json:",omitempty"`
	USAddress              USAddressType           `xml:"USAddress"`
	ForeignAddress         ForeignAddressType      `xml:"ForeignAddress"`
	EIN                    EINType                 `xml:"EIN"`
	BusinessNameControlTxt BusinessNameControlType `xml:"BusinessNameControlTxt"`
}

func (r AffiliateListingGrpType) Validate() error {
	return utils.Validate(&r)
}

type AffiliatedScheduleGrp struct {
	BusinessName                  *BusinessNameType   `xml:"BusinessName,omitempty" json:",omitempty"`
	USAddress                     *USAddressType      `xml:"USAddress,omitempty" json:",omitempty"`
	ForeignAddress                *ForeignAddressType `xml:"ForeignAddress,omitempty" json:",omitempty"`
	EIN                           EINType             `xml:"EIN"`
	ElectingOrganizationInd       CheckboxType        `xml:"ElectingOrganizationInd,omitempty" json:",omitempty"`
	TotalGrassrootsLobbyAmt       int                 `xml:"TotalGrassrootsLobbyAmt"`
	TotalDirectLobbyAmt           int                 `xml:"TotalDirectLobbyAmt"`
	TotalLobbyExpenditureAmt      int                 `xml:"TotalLobbyExpenditureAmt"`
	OtherExemptPurposeExpendAmt   int                 `xml:"OtherExemptPurposeExpendAmt"`
	TotalExemptPurposeExpendAmt   int                 `xml:"TotalExemptPurposeExpendAmt"`
	LobbyNontxAmt                 int                 `xml:"LobbyNontxAmt"`
	GrassrootsNontxAmt            int                 `xml:"GrassrootsNontxAmt"`
	TotalLobbyGrssrootMnsNontxAmt int                 `xml:"TotalLobbyGrssrootMnsNontxAmt"`
	TotLbbyExpendMnsLobbyNontxAmt int                 `xml:"TotLbbyExpendMnsLobbyNontxAmt"`
	ShareExcessLobbyExpendAmt     int                 `xml:"ShareExcessLobbyExpendAmt"`
}

func (r AffiliatedScheduleGrp) Validate() error {
	return utils.Validate(&r)
}

type AllAffiliatesIncludedInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r AllAffiliatesIncludedInd) Validate() error {
	return utils.Validate(&r)
}

type AuditedFinancialStmtAttInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r AuditedFinancialStmtAttInd) Validate() error {
	return utils.Validate(&r)
}

type BooksInCareOfDetail struct {
	PersonNm       PersonNameType     `xml:"PersonNm"`
	BusinessName   BusinessNameType   `xml:"BusinessName"`
	PhoneNum       PhoneNumberType    `xml:"PhoneNum"`
	USAddress      USAddressType      `xml:"USAddress"`
	ForeignAddress ForeignAddressType `xml:"ForeignAddress"`
}

func (r BooksInCareOfDetail) Validate() error {
	return utils.Validate(&r)
}

type BusTrInvolveInterestedPrsnGrp struct {
	NameOfInterested           *NameOfInterested `xml:"NameOfInterested,omitempty" json:",omitempty"`
	RelationshipDescriptionTxt string            `xml:"RelationshipDescriptionTxt,omitempty" json:",omitempty"`
	TransactionAmt             int               `xml:"TransactionAmt,omitempty" json:",omitempty"`
	TransactionDesc            string            `xml:"TransactionDesc,omitempty" json:",omitempty"`
	SharingOfRevenuesInd       bool              `xml:"SharingOfRevenuesInd,omitempty" json:",omitempty"`
}

func (r BusTrInvolveInterestedPrsnGrp) Validate() error {
	return utils.Validate(&r)
}

type BusinessNameType struct {
	BusinessNameLine1Txt BusinessNameLine1Type  `xml:"BusinessNameLine1Txt"`
	BusinessNameLine2Txt *BusinessNameLine2Type `xml:"BusinessNameLine2Txt,omitempty" json:",omitempty"`
}

func (r BusinessNameType) Validate() error {
	return utils.Validate(&r)
}

type BusinessOfficerGrp struct {
	PersonNm                   PersonNameType  `xml:"PersonNm"`
	PersonTitleTxt             PersonTitleType `xml:"PersonTitleTxt"`
	PhoneNum                   PhoneNumberType `xml:"PhoneNum,omitempty" json:",omitempty"`
	EmailAddressTxt            string          `xml:"EmailAddressTxt,omitempty" json:",omitempty"`
	SignatureDt                DateType        `xml:"SignatureDt"`
	TaxpayerPIN                PINType         `xml:"TaxpayerPIN,omitempty" json:",omitempty"`
	DiscussWithPaidPreparerInd bool            `xml:"DiscussWithPaidPreparerInd,omitempty" json:",omitempty"`
}

func (r BusinessOfficerGrp) Validate() error {
	return utils.Validate(&r)
}

type BusinessRlnWithFamMemInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r BusinessRlnWithFamMemInd) Validate() error {
	return utils.Validate(&r)
}

type BusinessRlnWithOfficerEntInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r BusinessRlnWithOfficerEntInd) Validate() error {
	return utils.Validate(&r)
}

type BusinessRlnWithOrgMemInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r BusinessRlnWithOrgMemInd) Validate() error {
	return utils.Validate(&r)
}

type CharitableContributionsDetailType struct {
	ContributorNum                 int                `xml:"ContributorNum"`
	GiftPurposeTxt                 string             `xml:"GiftPurposeTxt"`
	GiftUseTxt                     string             `xml:"GiftUseTxt"`
	HowGiftIsHeldDesc              string             `xml:"HowGiftIsHeldDesc,omitempty" json:",omitempty"`
	TransfereeNameBusiness         BusinessNameType   `xml:"TransfereeNameBusiness"`
	TransfereeNameIndividual       PersonNameType     `xml:"TransfereeNameIndividual"`
	TransfereeUSAddress            USAddressType      `xml:"TransfereeUSAddress"`
	TransfereeForeignAddress       ForeignAddressType `xml:"TransfereeForeignAddress"`
	RlnOfTransferorToTransfereeTxt string             `xml:"RlnOfTransferorToTransfereeTxt,omitempty" json:",omitempty"`
}

func (r CharitableContributionsDetailType) Validate() error {
	return utils.Validate(&r)
}

type CollectionUsedOtherPurposesGrp struct {
	CollectionUsedOtherPurposesInd CheckboxType `xml:"CollectionUsedOtherPurposesInd"`
	OtherPurposesDesc              string       `xml:"OtherPurposesDesc"`
}

func (r CollectionUsedOtherPurposesGrp) Validate() error {
	return utils.Validate(&r)
}

type CollectionsOfArtInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r CollectionsOfArtInd) Validate() error {
	return utils.Validate(&r)
}

type ConservationEasementsInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ConservationEasementsInd) Validate() error {
	return utils.Validate(&r)
}

type ConsolidatedAuditFinclStmtInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ConsolidatedAuditFinclStmtInd) Validate() error {
	return utils.Validate(&r)
}

type ContractorAddress struct {
	USAddress      USAddressType      `xml:"USAddress"`
	ForeignAddress ForeignAddressType `xml:"ForeignAddress"`
}

func (r ContractorAddress) Validate() error {
	return utils.Validate(&r)
}

type ContractorName struct {
	PersonNm     PersonNameType   `xml:"PersonNm"`
	BusinessName BusinessNameType `xml:"BusinessName"`
}

func (r ContractorName) Validate() error {
	return utils.Validate(&r)
}

type ContributorInformationGrpType struct {
	ContributorNum            int                `xml:"ContributorNum"`
	ContributorBusinessName   BusinessNameType   `xml:"ContributorBusinessName"`
	ContributorPersonNm       PersonNameType     `xml:"ContributorPersonNm"`
	Paid527j1Ind              CheckboxType       `xml:"Paid527j1Ind"`
	ContributorUSAddress      USAddressType      `xml:"ContributorUSAddress"`
	ContributorForeignAddress ForeignAddressType `xml:"ContributorForeignAddress"`
	TotalContributionsAmt     int                `xml:"TotalContributionsAmt"`
	PersonContributionInd     CheckboxType       `xml:"PersonContributionInd,omitempty" json:",omitempty"`
	PayrollContributionInd    CheckboxType       `xml:"PayrollContributionInd,omitempty" json:",omitempty"`
	NoncashContributionInd    CheckboxType       `xml:"NoncashContributionInd,omitempty" json:",omitempty"`
}

func (r ContributorInformationGrpType) Validate() error {
	return utils.Validate(&r)
}

type CostingMethodologyUsedGrp struct {
	CostAccountingSystemInd CheckboxType `xml:"CostAccountingSystemInd"`
	CostToChargeRatioInd    CheckboxType `xml:"CostToChargeRatioInd"`
	OtherInd                CheckboxType `xml:"OtherInd"`
}

func (r CostingMethodologyUsedGrp) Validate() error {
	return utils.Validate(&r)
}

type CreditCounselingInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r CreditCounselingInd) Validate() error {
	return utils.Validate(&r)
}

type DeductibleArtContributionInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r DeductibleArtContributionInd) Validate() error {
	return utils.Validate(&r)
}

type DeductibleNonCashContriInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r DeductibleNonCashContriInd) Validate() error {
	return utils.Validate(&r)
}

type DescribedInSection501c3Ind struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r DescribedInSection501c3Ind) Validate() error {
	return utils.Validate(&r)
}

type DiscountedCareOthPercentageGrp struct {
	OtherInd               CheckboxType `xml:"OtherInd"`
	DiscountedCareOtherPct float64      `xml:"DiscountedCareOtherPct"`
}

func (r DiscountedCareOthPercentageGrp) Validate() error {
	return utils.Validate(&r)
}

type DisqualifiedPersonExBnftTrGrpType struct {
	PersonNm                    PersonNameType   `xml:"PersonNm"`
	BusinessName                BusinessNameType `xml:"BusinessName"`
	RlnDisqualifiedPersonOrgTxt string           `xml:"RlnDisqualifiedPersonOrgTxt,omitempty" json:",omitempty"`
	TransactionDesc             string           `xml:"TransactionDesc,omitempty" json:",omitempty"`
	TransactionCorrectedInd     bool             `xml:"TransactionCorrectedInd,omitempty" json:",omitempty"`
}

func (r DisqualifiedPersonExBnftTrGrpType) Validate() error {
	return utils.Validate(&r)
}

type DisregardedEntityInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r DisregardedEntityInd) Validate() error {
	return utils.Validate(&r)
}

type DistributableAmountGrp struct {
	CYAdjNetIncomeDistributableAmt int          `xml:"CYAdjNetIncomeDistributableAmt"`
	CYPct85AdjustedNetIncomeAmt    int          `xml:"CYPct85AdjustedNetIncomeAmt"`
	CYTotalMinAstDistributableAmt  int          `xml:"CYTotalMinAstDistributableAmt"`
	CYGreaterAdjustedMinimumAmt    int          `xml:"CYGreaterAdjustedMinimumAmt"`
	CYIncomeTaxImposedPYAmt        int          `xml:"CYIncomeTaxImposedPYAmt"`
	CYDistributableAsAdjustedAmt   int          `xml:"CYDistributableAsAdjustedAmt"`
	FirstYearType3NonFuncInd       CheckboxType `xml:"FirstYearType3NonFuncInd,omitempty" json:",omitempty"`
}

func (r DistributableAmountGrp) Validate() error {
	return utils.Validate(&r)
}

type DistributionAllocationsGrp struct {
	CYDistributableAsAdjustedAmt   int `xml:"CYDistributableAsAdjustedAmt"`
	UnderdistributionsAmt          int `xml:"UnderdistributionsAmt"`
	ExcessDistributionCyovYr2Amt   int `xml:"ExcessDistributionCyovYr2Amt"`
	ExcessDistributionCyovYr1Amt   int `xml:"ExcessDistributionCyovYr1Amt"`
	TotalExcessDistributionCyovAmt int `xml:"TotalExcessDistributionCyovAmt"`
	CyovAppliedUnderdistriPYAmt    int `xml:"CyovAppliedUnderdistriPYAmt"`
	CyovAppliedUnderdistrCPYAmt    int `xml:"CyovAppliedUnderdistrCPYAmt"`
	ExcessDistributionCyovAmt      int `xml:"ExcessDistributionCyovAmt"`
	CYTotalAnnualDistributionsAmt  int `xml:"CYTotalAnnualDistributionsAmt"`
	CYDistribAppUnderdistriPYAmt   int `xml:"CYDistribAppUnderdistriPYAmt"`
	CYDistriAppDistributableAmt    int `xml:"CYDistriAppDistributableAmt,omitempty" json:",omitempty"`
	ExcessDistributionAmt          int `xml:"ExcessDistributionAmt"`
	RemainingUnderdistriPYAmt      int `xml:"RemainingUnderdistriPYAmt"`
	RemainingUnderdistriCYAmt      int `xml:"RemainingUnderdistriCYAmt"`
	ExcessDistriCyovToNextYrAmt    int `xml:"ExcessDistriCyovToNextYrAmt"`
	ExcessFromYear3Amt             int `xml:"ExcessFromYear3Amt"`
	ExcessFromYear2Amt             int `xml:"ExcessFromYear2Amt"`
	ExcessFromYear1Amt             int `xml:"ExcessFromYear1Amt"`
}

func (r DistributionAllocationsGrp) Validate() error {
	return utils.Validate(&r)
}

type DistributionsGrp struct {
	CYPaidAccomplishExemptPrpsAmt  int     `xml:"CYPaidAccomplishExemptPrpsAmt"`
	CYPdInExcessIncomeActivityAmt  int     `xml:"CYPdInExcessIncomeActivityAmt"`
	CYAdministrativeExpensePaidAmt int     `xml:"CYAdministrativeExpensePaidAmt"`
	ExemptUseAssetsAcquisPaidAmt   int     `xml:"ExemptUseAssetsAcquisPaidAmt"`
	QualifiedSetAsideAmt           int     `xml:"QualifiedSetAsideAmt"`
	CYOtherDistributionsAmt        int     `xml:"CYOtherDistributionsAmt"`
	CYTotalAnnualDistributionsAmt  int     `xml:"CYTotalAnnualDistributionsAmt"`
	CYDistriAttentiveSuprtOrgAmt   int     `xml:"CYDistriAttentiveSuprtOrgAmt"`
	CYDistributableAsAdjustedAmt   int     `xml:"CYDistributableAsAdjustedAmt"`
	CYDistributionYrRt             float64 `xml:"CYDistributionYrRt"`
}

func (r DistributionsGrp) Validate() error {
	return utils.Validate(&r)
}

type DonorAdvisedFundInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r DonorAdvisedFundInd) Validate() error {
	return utils.Validate(&r)
}

type EngagedInExcessBenefitTransInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r EngagedInExcessBenefitTransInd) Validate() error {
	return utils.Validate(&r)
}

type FeesForServicesProfFundraising struct {
	TotalAmt       int `xml:"TotalAmt,omitempty" json:",omitempty"`
	FundraisingAmt int `xml:"FundraisingAmt,omitempty" json:",omitempty"`
}

func (r FeesForServicesProfFundraising) Validate() error {
	return utils.Validate(&r)
}

type Filer struct {
	EIN                    EINType                 `xml:"EIN"`
	BusinessName           BusinessNameType        `xml:"BusinessName"`
	InCareOfNm             *InCareOfNameType       `xml:"InCareOfNm,omitempty" json:",omitempty"`
	BusinessNameControlTxt BusinessNameControlType `xml:"BusinessNameControlTxt"`
	PhoneNum               *PhoneNumberType        `xml:"PhoneNum,omitempty" json:",omitempty"`
	USAddress              USAddressType           `xml:"USAddress"`
	ForeignAddress         ForeignAddressType      `xml:"ForeignAddress"`
}

func (r Filer) Validate() error {
	return utils.Validate(&r)
}

type FinancialStatementType struct {
	SeparateBasisFinclStmtInd     CheckboxType `xml:"SeparateBasisFinclStmtInd"`
	ConsolidatedBasisFinclStmtInd CheckboxType `xml:"ConsolidatedBasisFinclStmtInd"`
	ConsolAndSepBasisFinclStmtInd CheckboxType `xml:"ConsolAndSepBasisFinclStmtInd"`
}

func (r FinancialStatementType) Validate() error {
	return utils.Validate(&r)
}

type ForeignActivitiesInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ForeignActivitiesInd) Validate() error {
	return utils.Validate(&r)
}

type ForeignAddressType struct {
	AddressLine1Txt   StreetAddressType  `xml:"AddressLine1Txt"`
	AddressLine2Txt   *StreetAddressType `xml:"AddressLine2Txt,omitempty" json:",omitempty"`
	CityNm            string             `xml:"CityNm,omitempty" json:",omitempty"`
	ProvinceOrStateNm string             `xml:"ProvinceOrStateNm,omitempty" json:",omitempty"`
	CountryCd         CountryType        `xml:"CountryCd"`
	ForeignPostalCd   string             `xml:"ForeignPostalCd,omitempty" json:",omitempty"`
}

func (r ForeignAddressType) Validate() error {
	return utils.Validate(&r)
}

type ForeignEntityIdentificationGrpType struct {
	ForeignEntityReferenceIdNum ForeignEntityReferenceIdNum `xml:"ForeignEntityReferenceIdNum"`
}

func (r ForeignEntityIdentificationGrpType) Validate() error {
	return utils.Validate(&r)
}

type ForeignIndividualsGrantsGrpType struct {
	TypeOfAssistanceTxt         string `xml:"TypeOfAssistanceTxt,omitempty" json:",omitempty"`
	RegionTxt                   string `xml:"RegionTxt,omitempty" json:",omitempty"`
	RecipientCnt                int    `xml:"RecipientCnt,omitempty" json:",omitempty"`
	CashGrantAmt                int    `xml:"CashGrantAmt,omitempty" json:",omitempty"`
	MannerOfCashDisbursementTxt string `xml:"MannerOfCashDisbursementTxt,omitempty" json:",omitempty"`
	NonCashAssistanceAmt        int    `xml:"NonCashAssistanceAmt,omitempty" json:",omitempty"`
	DescriptionOfNonCashAsstTxt string `xml:"DescriptionOfNonCashAsstTxt,omitempty" json:",omitempty"`
	ValuationMethodUsedDesc     string `xml:"ValuationMethodUsedDesc,omitempty" json:",omitempty"`
}

func (r ForeignIndividualsGrantsGrpType) Validate() error {
	return utils.Validate(&r)
}

type ForeignItemizedEntryType struct {
	Desc       string `xml:"Desc"`
	ForeignAmt int    `xml:"ForeignAmt"`
}

func (r ForeignItemizedEntryType) Validate() error {
	return utils.Validate(&r)
}

type Form990PartIXGroup1Type struct {
	TotalAmt           int `xml:"TotalAmt,omitempty" json:",omitempty"`
	ProgramServicesAmt int `xml:"ProgramServicesAmt,omitempty" json:",omitempty"`
}

func (r Form990PartIXGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartIXGroup2Type struct {
	TotalAmt                int `xml:"TotalAmt,omitempty" json:",omitempty"`
	ProgramServicesAmt      int `xml:"ProgramServicesAmt,omitempty" json:",omitempty"`
	ManagementAndGeneralAmt int `xml:"ManagementAndGeneralAmt,omitempty" json:",omitempty"`
	FundraisingAmt          int `xml:"FundraisingAmt,omitempty" json:",omitempty"`
}

func (r Form990PartIXGroup2Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartIXGroup3Type struct {
	TotalAmt                int `xml:"TotalAmt"`
	ProgramServicesAmt      int `xml:"ProgramServicesAmt,omitempty" json:",omitempty"`
	ManagementAndGeneralAmt int `xml:"ManagementAndGeneralAmt,omitempty" json:",omitempty"`
	FundraisingAmt          int `xml:"FundraisingAmt,omitempty" json:",omitempty"`
}

func (r Form990PartIXGroup3Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartIXGroup4Type struct {
	Desc                    string `xml:"Desc,omitempty" json:",omitempty"`
	TotalAmt                int    `xml:"TotalAmt,omitempty" json:",omitempty"`
	ProgramServicesAmt      int    `xml:"ProgramServicesAmt,omitempty" json:",omitempty"`
	ManagementAndGeneralAmt int    `xml:"ManagementAndGeneralAmt,omitempty" json:",omitempty"`
	FundraisingAmt          int    `xml:"FundraisingAmt,omitempty" json:",omitempty"`
}

func (r Form990PartIXGroup4Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartVIIGroup1Type struct {
	ContractorName    *ContractorName    `xml:"ContractorName,omitempty" json:",omitempty"`
	ContractorAddress *ContractorAddress `xml:"ContractorAddress,omitempty" json:",omitempty"`
	ServicesDesc      string             `xml:"ServicesDesc,omitempty" json:",omitempty"`
	CompensationAmt   int                `xml:"CompensationAmt,omitempty" json:",omitempty"`
}

func (r Form990PartVIIGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartVIIIGroup1Type struct {
	BusinessCd                   BusinessCd `xml:"BusinessCd,omitempty" json:",omitempty"`
	TotalRevenueColumnAmt        int        `xml:"TotalRevenueColumnAmt,omitempty" json:",omitempty"`
	RelatedOrExemptFuncIncomeAmt int        `xml:"RelatedOrExemptFuncIncomeAmt,omitempty" json:",omitempty"`
	UnrelatedBusinessRevenueAmt  int        `xml:"UnrelatedBusinessRevenueAmt,omitempty" json:",omitempty"`
	ExclusionAmt                 int        `xml:"ExclusionAmt,omitempty" json:",omitempty"`
}

func (r Form990PartVIIIGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartVIIIGroup2Type struct {
	Desc                         string     `xml:"Desc,omitempty" json:",omitempty"`
	BusinessCd                   BusinessCd `xml:"BusinessCd,omitempty" json:",omitempty"`
	TotalRevenueColumnAmt        int        `xml:"TotalRevenueColumnAmt,omitempty" json:",omitempty"`
	RelatedOrExemptFuncIncomeAmt int        `xml:"RelatedOrExemptFuncIncomeAmt,omitempty" json:",omitempty"`
	UnrelatedBusinessRevenueAmt  int        `xml:"UnrelatedBusinessRevenueAmt,omitempty" json:",omitempty"`
	ExclusionAmt                 int        `xml:"ExclusionAmt,omitempty" json:",omitempty"`
}

func (r Form990PartVIIIGroup2Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartVIIIGroup3Type struct {
	TotalRevenueColumnAmt        int `xml:"TotalRevenueColumnAmt,omitempty" json:",omitempty"`
	RelatedOrExemptFuncIncomeAmt int `xml:"RelatedOrExemptFuncIncomeAmt,omitempty" json:",omitempty"`
	UnrelatedBusinessRevenueAmt  int `xml:"UnrelatedBusinessRevenueAmt,omitempty" json:",omitempty"`
	ExclusionAmt                 int `xml:"ExclusionAmt,omitempty" json:",omitempty"`
}

func (r Form990PartVIIIGroup3Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartVIIIGroup4Type struct {
	RealAmt     int `xml:"RealAmt,omitempty" json:",omitempty"`
	PersonalAmt int `xml:"PersonalAmt,omitempty" json:",omitempty"`
}

func (r Form990PartVIIIGroup4Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartVIIIGroup5Type struct {
	SecuritiesAmt int `xml:"SecuritiesAmt,omitempty" json:",omitempty"`
	OtherAmt      int `xml:"OtherAmt,omitempty" json:",omitempty"`
}

func (r Form990PartVIIIGroup5Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartVIIIGroup6Type struct {
	TotalRevenueColumnAmt        int `xml:"TotalRevenueColumnAmt"`
	RelatedOrExemptFuncIncomeAmt int `xml:"RelatedOrExemptFuncIncomeAmt,omitempty" json:",omitempty"`
	UnrelatedBusinessRevenueAmt  int `xml:"UnrelatedBusinessRevenueAmt,omitempty" json:",omitempty"`
	ExclusionAmt                 int `xml:"ExclusionAmt,omitempty" json:",omitempty"`
}

func (r Form990PartVIIIGroup6Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartVIIIGroup7Type struct {
	TotalRevenueColumnAmt       int `xml:"TotalRevenueColumnAmt"`
	UnrelatedBusinessRevenueAmt int `xml:"UnrelatedBusinessRevenueAmt,omitempty" json:",omitempty"`
	ExclusionAmt                int `xml:"ExclusionAmt,omitempty" json:",omitempty"`
}

func (r Form990PartVIIIGroup7Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartVIISectionAGrp struct {
	PersonNm                       PersonNameType   `xml:"PersonNm"`
	BusinessName                   BusinessNameType `xml:"BusinessName"`
	TitleTxt                       string           `xml:"TitleTxt,omitempty" json:",omitempty"`
	AverageHoursPerWeekRt          float64          `xml:"AverageHoursPerWeekRt,omitempty" json:",omitempty"`
	AverageHoursPerWeekRltdOrgRt   float64          `xml:"AverageHoursPerWeekRltdOrgRt,omitempty" json:",omitempty"`
	IndividualTrusteeOrDirectorInd CheckboxType     `xml:"IndividualTrusteeOrDirectorInd,omitempty" json:",omitempty"`
	InstitutionalTrusteeInd        CheckboxType     `xml:"InstitutionalTrusteeInd,omitempty" json:",omitempty"`
	OfficerInd                     CheckboxType     `xml:"OfficerInd,omitempty" json:",omitempty"`
	KeyEmployeeInd                 CheckboxType     `xml:"KeyEmployeeInd,omitempty" json:",omitempty"`
	HighestCompensatedEmployeeInd  CheckboxType     `xml:"HighestCompensatedEmployeeInd,omitempty" json:",omitempty"`
	FormerOfcrDirectorTrusteeInd   CheckboxType     `xml:"FormerOfcrDirectorTrusteeInd,omitempty" json:",omitempty"`
	ReportableCompFromOrgAmt       int              `xml:"ReportableCompFromOrgAmt,omitempty" json:",omitempty"`
	ReportableCompFromRltdOrgAmt   int              `xml:"ReportableCompFromRltdOrgAmt,omitempty" json:",omitempty"`
	OtherCompensationAmt           int              `xml:"OtherCompensationAmt,omitempty" json:",omitempty"`
}

func (r Form990PartVIISectionAGrp) Validate() error {
	return utils.Validate(&r)
}

type Form990PartXGroup1Type struct {
	BOYAmt int `xml:"BOYAmt,omitempty" json:",omitempty"`
	EOYAmt int `xml:"EOYAmt,omitempty" json:",omitempty"`
}

func (r Form990PartXGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990PartXGroup2Type struct {
	BOYAmt int `xml:"BOYAmt"`
	EOYAmt int `xml:"EOYAmt"`
}

func (r Form990PartXGroup2Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchAPartIIGroup1Type struct {
	CurrentTaxYearMinus4YearsAmt int `xml:"CurrentTaxYearMinus4YearsAmt,omitempty" json:",omitempty"`
	CurrentTaxYearMinus3YearsAmt int `xml:"CurrentTaxYearMinus3YearsAmt,omitempty" json:",omitempty"`
	CurrentTaxYearMinus2YearsAmt int `xml:"CurrentTaxYearMinus2YearsAmt,omitempty" json:",omitempty"`
	CurrentTaxYearMinus1YearAmt  int `xml:"CurrentTaxYearMinus1YearAmt,omitempty" json:",omitempty"`
	CurrentTaxYearAmt            int `xml:"CurrentTaxYearAmt,omitempty" json:",omitempty"`
	TotalAmt                     int `xml:"TotalAmt,omitempty" json:",omitempty"`
}

func (r Form990SchAPartIIGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchAPartVGroup1Type struct {
	PriorYearAmt   int `xml:"PriorYearAmt"`
	CurrentYearAmt int `xml:"CurrentYearAmt,omitempty" json:",omitempty"`
}

func (r Form990SchAPartVGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchAPartVGroup2Type struct {
	PriorYearAmt   int `xml:"PriorYearAmt"`
	CurrentYearAmt int `xml:"CurrentYearAmt,omitempty" json:",omitempty"`
}

func (r Form990SchAPartVGroup2Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchASupportingOrgGrp struct {
	ListedByNameGoverningDocInd    bool `xml:"ListedByNameGoverningDocInd"`
	SuprtOrgNoIRSDeterminationInd  bool `xml:"SuprtOrgNoIRSDeterminationInd"`
	SupportedOrgSectionC456Ind     bool `xml:"SupportedOrgSectionC456Ind"`
	SupportedOrgQualifiedInd       bool `xml:"SupportedOrgQualifiedInd,omitempty" json:",omitempty"`
	SuprtExclusivelySec170c2BInd   bool `xml:"SuprtExclusivelySec170c2BInd,omitempty" json:",omitempty"`
	SupportedOrgNotOrganizedUSInd  bool `xml:"SupportedOrgNotOrganizedUSInd"`
	ControlDecidingGrntFrgnOrgInd  bool `xml:"ControlDecidingGrntFrgnOrgInd,omitempty" json:",omitempty"`
	SupportForeignOrgNoDetermInd   bool `xml:"SupportForeignOrgNoDetermInd,omitempty" json:",omitempty"`
	OrganizationChangeSuprtOrgInd  bool `xml:"OrganizationChangeSuprtOrgInd"`
	SupportedOrgClassDesignatedInd bool `xml:"SupportedOrgClassDesignatedInd,omitempty" json:",omitempty"`
	SubstitutionBeyondOrgCntlInd   bool `xml:"SubstitutionBeyondOrgCntlInd,omitempty" json:",omitempty"`
	SupportNonSupportedOrgInd      bool `xml:"SupportNonSupportedOrgInd"`
	PaymentSubstantialContribtrInd bool `xml:"PaymentSubstantialContribtrInd"`
	LoanDisqualifiedPersonInd      bool `xml:"LoanDisqualifiedPersonInd"`
	ControlledDisqualifiedPrsnInd  bool `xml:"ControlledDisqualifiedPrsnInd"`
	DisqualifiedPrsnControllIntInd bool `xml:"DisqualifiedPrsnControllIntInd"`
	DisqualifiedPrsnOwnrIntInd     bool `xml:"DisqualifiedPrsnOwnrIntInd"`
	ExcessBusinessHoldingsRulesInd bool `xml:"ExcessBusinessHoldingsRulesInd"`
	ExcessBusinessHoldingsInd      bool `xml:"ExcessBusinessHoldingsInd,omitempty" json:",omitempty"`
	ContributionControllerInd      bool `xml:"ContributionControllerInd"`
	ContributionFamilyInd          bool `xml:"ContributionFamilyInd"`
	Contribution35ControlledInd    bool `xml:"Contribution35ControlledInd"`
}

func (r Form990SchASupportingOrgGrp) Validate() error {
	return utils.Validate(&r)
}

type Form990SchAType1SuprtOrgGrp struct {
	PowerAppointMajorityDirTrstInd bool `xml:"PowerAppointMajorityDirTrstInd"`
	OperateBenefitNonSuprtOrgInd   bool `xml:"OperateBenefitNonSuprtOrgInd"`
}

func (r Form990SchAType1SuprtOrgGrp) Validate() error {
	return utils.Validate(&r)
}

type Form990SchAType3FuncIntGrp struct {
	ActivitiesTestInd              CheckboxType `xml:"ActivitiesTestInd,omitempty" json:",omitempty"`
	ParentSupportedOrgInd          CheckboxType `xml:"ParentSupportedOrgInd,omitempty" json:",omitempty"`
	GovernmentalEntityInd          CheckboxType `xml:"GovernmentalEntityInd,omitempty" json:",omitempty"`
	ActivitiesFurtherExemptPrpsInd bool         `xml:"ActivitiesFurtherExemptPrpsInd,omitempty" json:",omitempty"`
	ActivitiesEngagedOrgInvlmntInd bool         `xml:"ActivitiesEngagedOrgInvlmntInd,omitempty" json:",omitempty"`
	AppointElectMajorityOfficerInd bool         `xml:"AppointElectMajorityOfficerInd,omitempty" json:",omitempty"`
	ExerciseDirectionPoliciesInd   bool         `xml:"ExerciseDirectionPoliciesInd,omitempty" json:",omitempty"`
}

func (r Form990SchAType3FuncIntGrp) Validate() error {
	return utils.Validate(&r)
}

type Form990SchAType3SprtOrgAllGrp struct {
	TimelyProvidedDocumentsInd     bool `xml:"TimelyProvidedDocumentsInd"`
	OfficersCloseRelationshipInd   bool `xml:"OfficersCloseRelationshipInd"`
	SupportedOrgVoiceInvestmentInd bool `xml:"SupportedOrgVoiceInvestmentInd"`
}

func (r Form990SchAType3SprtOrgAllGrp) Validate() error {
	return utils.Validate(&r)
}

type Form990SchCPartIIAGroup1Type struct {
	FilingOrganizationsTotalAmt int `xml:"FilingOrganizationsTotalAmt,omitempty" json:",omitempty"`
	AffiliatedGroupTotalAmt     int `xml:"AffiliatedGroupTotalAmt,omitempty" json:",omitempty"`
}

func (r Form990SchCPartIIAGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchCPartIIAGroup2Type struct {
	CurrentYearMinus3Amt int `xml:"CurrentYearMinus3Amt,omitempty" json:",omitempty"`
	CurrentYearMinus2Amt int `xml:"CurrentYearMinus2Amt,omitempty" json:",omitempty"`
	CurrentYearMinus1Amt int `xml:"CurrentYearMinus1Amt,omitempty" json:",omitempty"`
	CurrentYearAmt       int `xml:"CurrentYearAmt,omitempty" json:",omitempty"`
	TotalAmt             int `xml:"TotalAmt,omitempty" json:",omitempty"`
}

func (r Form990SchCPartIIAGroup2Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchDPartIIIGroup1Type struct {
	RevenuesIncludedAmt int `xml:"RevenuesIncludedAmt,omitempty" json:",omitempty"`
	AssetsIncludedAmt   int `xml:"AssetsIncludedAmt,omitempty" json:",omitempty"`
}

func (r Form990SchDPartIIIGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchDPartIXGroup1Type struct {
	Desc         string `xml:"Desc,omitempty" json:",omitempty"`
	BookValueAmt int    `xml:"BookValueAmt,omitempty" json:",omitempty"`
}

func (r Form990SchDPartIXGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchDPartVGroup1Type struct {
	BeginningYearBalanceAmt       int `xml:"BeginningYearBalanceAmt,omitempty" json:",omitempty"`
	ContributionsAmt              int `xml:"ContributionsAmt,omitempty" json:",omitempty"`
	InvestmentEarningsOrLossesAmt int `xml:"InvestmentEarningsOrLossesAmt,omitempty" json:",omitempty"`
	GrantsOrScholarshipsAmt       int `xml:"GrantsOrScholarshipsAmt,omitempty" json:",omitempty"`
	OtherExpendituresAmt          int `xml:"OtherExpendituresAmt,omitempty" json:",omitempty"`
	AdministrativeExpensesAmt     int `xml:"AdministrativeExpensesAmt,omitempty" json:",omitempty"`
	EndYearBalanceAmt             int `xml:"EndYearBalanceAmt,omitempty" json:",omitempty"`
}

func (r Form990SchDPartVGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchDPartVIGroup1Type struct {
	InvestmentCostOrOtherBasisAmt int `xml:"InvestmentCostOrOtherBasisAmt,omitempty" json:",omitempty"`
	OtherCostOrOtherBasisAmt      int `xml:"OtherCostOrOtherBasisAmt,omitempty" json:",omitempty"`
	DepreciationAmt               int `xml:"DepreciationAmt,omitempty" json:",omitempty"`
	BookValueAmt                  int `xml:"BookValueAmt,omitempty" json:",omitempty"`
}

func (r Form990SchDPartVIGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchDPartVIGroup2Type struct {
	InvestmentCostOrOtherBasisAmt int `xml:"InvestmentCostOrOtherBasisAmt,omitempty" json:",omitempty"`
	OtherCostOrOtherBasisAmt      int `xml:"OtherCostOrOtherBasisAmt,omitempty" json:",omitempty"`
	BookValueAmt                  int `xml:"BookValueAmt,omitempty" json:",omitempty"`
}

func (r Form990SchDPartVIGroup2Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchDPartVIIGroup1Type struct {
	BookValueAmt      int               `xml:"BookValueAmt,omitempty" json:",omitempty"`
	MethodValuationCd MethodValuationCd `xml:"MethodValuationCd,omitempty" json:",omitempty"`
}

func (r Form990SchDPartVIIGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchDPartVIIGroup2Type struct {
	Desc              string            `xml:"Desc,omitempty" json:",omitempty"`
	BookValueAmt      int               `xml:"BookValueAmt,omitempty" json:",omitempty"`
	MethodValuationCd MethodValuationCd `xml:"MethodValuationCd,omitempty" json:",omitempty"`
}

func (r Form990SchDPartVIIGroup2Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchDPartVIIIGroup1Type struct {
	Desc              string            `xml:"Desc,omitempty" json:",omitempty"`
	BookValueAmt      int               `xml:"BookValueAmt,omitempty" json:",omitempty"`
	MethodValuationCd MethodValuationCd `xml:"MethodValuationCd,omitempty" json:",omitempty"`
}

func (r Form990SchDPartVIIIGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchDPartXGroup1Type struct {
	Desc string `xml:"Desc,omitempty" json:",omitempty"`
	Amt  int    `xml:"Amt,omitempty" json:",omitempty"`
}

func (r Form990SchDPartXGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchHPartIGroup1Type struct {
	ActivitiesOrProgramsCnt       int     `xml:"ActivitiesOrProgramsCnt,omitempty" json:",omitempty"`
	PersonsServedCnt              int     `xml:"PersonsServedCnt,omitempty" json:",omitempty"`
	TotalCommunityBenefitExpnsAmt int     `xml:"TotalCommunityBenefitExpnsAmt,omitempty" json:",omitempty"`
	DirectOffsettingRevenueAmt    int     `xml:"DirectOffsettingRevenueAmt,omitempty" json:",omitempty"`
	NetCommunityBenefitExpnsAmt   int     `xml:"NetCommunityBenefitExpnsAmt,omitempty" json:",omitempty"`
	TotalExpensePct               float64 `xml:"TotalExpensePct,omitempty" json:",omitempty"`
}

func (r Form990SchHPartIGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchHPartIIGroup1Type struct {
	ActivitiesOrProgramsCnt       int     `xml:"ActivitiesOrProgramsCnt,omitempty" json:",omitempty"`
	PersonsServedCnt              int     `xml:"PersonsServedCnt,omitempty" json:",omitempty"`
	TotalCommunityBenefitExpnsAmt int     `xml:"TotalCommunityBenefitExpnsAmt,omitempty" json:",omitempty"`
	DirectOffsettingRevenueAmt    int     `xml:"DirectOffsettingRevenueAmt,omitempty" json:",omitempty"`
	NetCommunityBenefitExpnsAmt   int     `xml:"NetCommunityBenefitExpnsAmt,omitempty" json:",omitempty"`
	TotalExpensePct               float64 `xml:"TotalExpensePct,omitempty" json:",omitempty"`
}

func (r Form990SchHPartIIGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchMGroup1Type struct {
	NonCashCheckboxInd             CheckboxType `xml:"NonCashCheckboxInd,omitempty" json:",omitempty"`
	ContributionCnt                int          `xml:"ContributionCnt,omitempty" json:",omitempty"`
	NoncashContributionsRptF990Amt int          `xml:"NoncashContributionsRptF990Amt,omitempty" json:",omitempty"`
	MethodOfDeterminingRevenuesTxt string       `xml:"MethodOfDeterminingRevenuesTxt,omitempty" json:",omitempty"`
}

func (r Form990SchMGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchMGroup2Type struct {
	NonCashCheckboxInd             CheckboxType `xml:"NonCashCheckboxInd,omitempty" json:",omitempty"`
	NoncashContributionsRptF990Amt int          `xml:"NoncashContributionsRptF990Amt,omitempty" json:",omitempty"`
	MethodOfDeterminingRevenuesTxt string       `xml:"MethodOfDeterminingRevenuesTxt,omitempty" json:",omitempty"`
}

func (r Form990SchMGroup2Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchMGroup3Type struct {
	NonCashCheckboxInd             CheckboxType `xml:"NonCashCheckboxInd,omitempty" json:",omitempty"`
	Desc                           string       `xml:"Desc,omitempty" json:",omitempty"`
	ContributionCnt                int          `xml:"ContributionCnt,omitempty" json:",omitempty"`
	NoncashContributionsRptF990Amt int          `xml:"NoncashContributionsRptF990Amt,omitempty" json:",omitempty"`
	MethodOfDeterminingRevenuesTxt string       `xml:"MethodOfDeterminingRevenuesTxt,omitempty" json:",omitempty"`
}

func (r Form990SchMGroup3Type) Validate() error {
	return utils.Validate(&r)
}

type Form990SchNGroup1Type struct {
	AssetsDistriOrExpnssPaidDesc string             `xml:"AssetsDistriOrExpnssPaidDesc,omitempty" json:",omitempty"`
	DistributionDt               DateType           `xml:"DistributionDt,omitempty" json:",omitempty"`
	FairMarketValueOfAssetAmt    int                `xml:"FairMarketValueOfAssetAmt,omitempty" json:",omitempty"`
	MethodOfFMVDeterminationTxt  string             `xml:"MethodOfFMVDeterminationTxt,omitempty" json:",omitempty"`
	EIN                          EINType            `xml:"EIN,omitempty" json:",omitempty"`
	PersonNm                     PersonNameType     `xml:"PersonNm"`
	BusinessName                 BusinessNameType   `xml:"BusinessName"`
	USAddress                    USAddressType      `xml:"USAddress"`
	ForeignAddress               ForeignAddressType `xml:"ForeignAddress"`
	IRCSectionTxt                string             `xml:"IRCSectionTxt,omitempty" json:",omitempty"`
}

func (r Form990SchNGroup1Type) Validate() error {
	return utils.Validate(&r)
}

type Form990ScheduleAPartVIGrp struct {
	FormAndLineReferenceDesc string `xml:"FormAndLineReferenceDesc,omitempty" json:",omitempty"`
	ExplanationTxt           string `xml:"ExplanationTxt,omitempty" json:",omitempty"`
}

func (r Form990ScheduleAPartVIGrp) Validate() error {
	return utils.Validate(&r)
}

type FreeCareOthPercentageGrp struct {
	OtherInd         CheckboxType `xml:"OtherInd"`
	FreeCareOtherPct float64      `xml:"FreeCareOtherPct"`
}

func (r FreeCareOthPercentageGrp) Validate() error {
	return utils.Validate(&r)
}

type FundraiserActivityInfoGrpType struct {
	PersonNm                    PersonNameType     `xml:"PersonNm"`
	OrganizationBusinessName    BusinessNameType   `xml:"OrganizationBusinessName"`
	USAddress                   USAddressType      `xml:"USAddress"`
	ForeignAddress              ForeignAddressType `xml:"ForeignAddress"`
	ActivityTxt                 string             `xml:"ActivityTxt,omitempty" json:",omitempty"`
	FundraiserControlOfFundsInd bool               `xml:"FundraiserControlOfFundsInd,omitempty" json:",omitempty"`
	GrossReceiptsAmt            int                `xml:"GrossReceiptsAmt,omitempty" json:",omitempty"`
	RetainedByContractorAmt     int                `xml:"RetainedByContractorAmt,omitempty" json:",omitempty"`
	NetToOrganizationAmt        int                `xml:"NetToOrganizationAmt,omitempty" json:",omitempty"`
}

func (r FundraiserActivityInfoGrpType) Validate() error {
	return utils.Validate(&r)
}

type FundraisingActivitiesInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r FundraisingActivitiesInd) Validate() error {
	return utils.Validate(&r)
}

type FundraisingEventInformationGrpType struct {
	Event1Nm                       string `xml:"Event1Nm,omitempty" json:",omitempty"`
	GrossReceiptsEvent1Amt         int    `xml:"GrossReceiptsEvent1Amt,omitempty" json:",omitempty"`
	CharitableContriEvent1Amt      int    `xml:"CharitableContriEvent1Amt,omitempty" json:",omitempty"`
	GrossRevenueEvent1Amt          int    `xml:"GrossRevenueEvent1Amt,omitempty" json:",omitempty"`
	CashPrizesEvent1Amt            int    `xml:"CashPrizesEvent1Amt,omitempty" json:",omitempty"`
	NonCashPrizesEvent1Amt         int    `xml:"NonCashPrizesEvent1Amt,omitempty" json:",omitempty"`
	RentFacilityCostsEvent1Amt     int    `xml:"RentFacilityCostsEvent1Amt,omitempty" json:",omitempty"`
	FoodAndBeverageEvent1Amt       int    `xml:"FoodAndBeverageEvent1Amt,omitempty" json:",omitempty"`
	EntertainmentEvent1Amt         int    `xml:"EntertainmentEvent1Amt,omitempty" json:",omitempty"`
	OtherDirectExpensesEvent1Amt   int    `xml:"OtherDirectExpensesEvent1Amt,omitempty" json:",omitempty"`
	Event2Nm                       string `xml:"Event2Nm,omitempty" json:",omitempty"`
	GrossReceiptsEvent2Amt         int    `xml:"GrossReceiptsEvent2Amt,omitempty" json:",omitempty"`
	CharitableContriEvent2Amt      int    `xml:"CharitableContriEvent2Amt,omitempty" json:",omitempty"`
	GrossRevenueEvent2Amt          int    `xml:"GrossRevenueEvent2Amt,omitempty" json:",omitempty"`
	CashPrizesEvent2Amt            int    `xml:"CashPrizesEvent2Amt,omitempty" json:",omitempty"`
	NonCashPrizesEvent2Amt         int    `xml:"NonCashPrizesEvent2Amt,omitempty" json:",omitempty"`
	RentFacilityCostsEvent2Amt     int    `xml:"RentFacilityCostsEvent2Amt,omitempty" json:",omitempty"`
	FoodAndBeverageEvent2Amt       int    `xml:"FoodAndBeverageEvent2Amt,omitempty" json:",omitempty"`
	EntertainmentEvent2Amt         int    `xml:"EntertainmentEvent2Amt,omitempty" json:",omitempty"`
	OtherDirectExpensesEvent2Amt   int    `xml:"OtherDirectExpensesEvent2Amt,omitempty" json:",omitempty"`
	OtherEventsTotalCnt            int    `xml:"OtherEventsTotalCnt,omitempty" json:",omitempty"`
	GrossReceiptsOtherEventsAmt    int    `xml:"GrossReceiptsOtherEventsAmt,omitempty" json:",omitempty"`
	CharitableContriOtherEventsAmt int    `xml:"CharitableContriOtherEventsAmt,omitempty" json:",omitempty"`
	GrossRevenueOtherEventsAmt     int    `xml:"GrossRevenueOtherEventsAmt,omitempty" json:",omitempty"`
	CashPrizesOtherEventsAmt       int    `xml:"CashPrizesOtherEventsAmt,omitempty" json:",omitempty"`
	NonCashPrizesOtherEventsAmt    int    `xml:"NonCashPrizesOtherEventsAmt,omitempty" json:",omitempty"`
	RentFcltyCostsOtherEventsAmt   int    `xml:"RentFcltyCostsOtherEventsAmt,omitempty" json:",omitempty"`
	FoodAndBeverageOtherEventsAmt  int    `xml:"FoodAndBeverageOtherEventsAmt,omitempty" json:",omitempty"`
	EntertainmentOtherEventsAmt    int    `xml:"EntertainmentOtherEventsAmt,omitempty" json:",omitempty"`
	OthDirectExpnssOtherEventsAmt  int    `xml:"OthDirectExpnssOtherEventsAmt,omitempty" json:",omitempty"`
	GrossReceiptsTotalAmt          int    `xml:"GrossReceiptsTotalAmt,omitempty" json:",omitempty"`
	CharitableContributionsTotAmt  int    `xml:"CharitableContributionsTotAmt,omitempty" json:",omitempty"`
	GrossRevenueTotalEventsAmt     int    `xml:"GrossRevenueTotalEventsAmt,omitempty" json:",omitempty"`
	CashPrizesTotalEventsAmt       int    `xml:"CashPrizesTotalEventsAmt,omitempty" json:",omitempty"`
	NonCashPrizesTotalEventsAmt    int    `xml:"NonCashPrizesTotalEventsAmt,omitempty" json:",omitempty"`
	RentFcltyCostsTotalEventsAmt   int    `xml:"RentFcltyCostsTotalEventsAmt,omitempty" json:",omitempty"`
	FoodAndBeverageTotalEventsAmt  int    `xml:"FoodAndBeverageTotalEventsAmt,omitempty" json:",omitempty"`
	EntertainmentTotalEventsAmt    int    `xml:"EntertainmentTotalEventsAmt,omitempty" json:",omitempty"`
	OthDirectExpnssTotalEventsAmt  int    `xml:"OthDirectExpnssTotalEventsAmt,omitempty" json:",omitempty"`
	DirectExpenseSummaryEventsAmt  int    `xml:"DirectExpenseSummaryEventsAmt,omitempty" json:",omitempty"`
	NetIncomeSummaryAmt            int    `xml:"NetIncomeSummaryAmt,omitempty" json:",omitempty"`
}

func (r FundraisingEventInformationGrpType) Validate() error {
	return utils.Validate(&r)
}

type GamingActivitiesInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r GamingActivitiesInd) Validate() error {
	return utils.Validate(&r)
}

type GamingInformationGrpType struct {
	GrossRevenueBingoAmt           int     `xml:"GrossRevenueBingoAmt,omitempty" json:",omitempty"`
	CashPrizesBingoAmt             int     `xml:"CashPrizesBingoAmt,omitempty" json:",omitempty"`
	NonCashPrizesBingoAmt          int     `xml:"NonCashPrizesBingoAmt,omitempty" json:",omitempty"`
	RentFacilityCostsBingoAmt      int     `xml:"RentFacilityCostsBingoAmt,omitempty" json:",omitempty"`
	OtherDirectExpensesBingoAmt    int     `xml:"OtherDirectExpensesBingoAmt,omitempty" json:",omitempty"`
	VolunteerLaborBingoInd         bool    `xml:"VolunteerLaborBingoInd,omitempty" json:",omitempty"`
	VolunteerLaborBingoPct         float64 `xml:"VolunteerLaborBingoPct,omitempty" json:",omitempty"`
	GrossRevenuePullTabsAmt        int     `xml:"GrossRevenuePullTabsAmt,omitempty" json:",omitempty"`
	CashPrizesPullTabsAmt          int     `xml:"CashPrizesPullTabsAmt,omitempty" json:",omitempty"`
	NonCashPrizesPullTabsAmt       int     `xml:"NonCashPrizesPullTabsAmt,omitempty" json:",omitempty"`
	RentFacilityCostsPullTabsAmt   int     `xml:"RentFacilityCostsPullTabsAmt,omitempty" json:",omitempty"`
	OtherDirectExpensesPullTabsAmt int     `xml:"OtherDirectExpensesPullTabsAmt,omitempty" json:",omitempty"`
	VolunteerLaborPullTabsInd      bool    `xml:"VolunteerLaborPullTabsInd,omitempty" json:",omitempty"`
	VolunteerLaborPullTabsPct      float64 `xml:"VolunteerLaborPullTabsPct,omitempty" json:",omitempty"`
	GrossRevenueOtherGamingAmt     int     `xml:"GrossRevenueOtherGamingAmt,omitempty" json:",omitempty"`
	CashPrizesOtherGamingAmt       int     `xml:"CashPrizesOtherGamingAmt,omitempty" json:",omitempty"`
	NonCashPrizesOtherGamingAmt    int     `xml:"NonCashPrizesOtherGamingAmt,omitempty" json:",omitempty"`
	RentFcltyCostsOtherGamingAmt   int     `xml:"RentFcltyCostsOtherGamingAmt,omitempty" json:",omitempty"`
	OthDirectExpnssOtherGamingAmt  int     `xml:"OthDirectExpnssOtherGamingAmt,omitempty" json:",omitempty"`
	VolunteerLaborOtherGamingInd   bool    `xml:"VolunteerLaborOtherGamingInd,omitempty" json:",omitempty"`
	VolunteerLaborOtherGamingPct   float64 `xml:"VolunteerLaborOtherGamingPct,omitempty" json:",omitempty"`
	GrossRevenueTotalGamingAmt     int     `xml:"GrossRevenueTotalGamingAmt,omitempty" json:",omitempty"`
	CashPrizesTotalGamingAmt       int     `xml:"CashPrizesTotalGamingAmt,omitempty" json:",omitempty"`
	NonCashPrizesTotalGamingAmt    int     `xml:"NonCashPrizesTotalGamingAmt,omitempty" json:",omitempty"`
	RentFcltyCostsTotalGamingAmt   int     `xml:"RentFcltyCostsTotalGamingAmt,omitempty" json:",omitempty"`
	OthDirectExpnssTotalGamingAmt  int     `xml:"OthDirectExpnssTotalGamingAmt,omitempty" json:",omitempty"`
	DirectExpenseSummaryGamingAmt  int     `xml:"DirectExpenseSummaryGamingAmt,omitempty" json:",omitempty"`
	NetGamingIncomeSummaryAmt      int     `xml:"NetGamingIncomeSummaryAmt,omitempty" json:",omitempty"`
}

func (r GamingInformationGrpType) Validate() error {
	return utils.Validate(&r)
}

type GrantToRelatedPersonInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r GrantToRelatedPersonInd) Validate() error {
	return utils.Validate(&r)
}

type GrantsOtherAsstToIndivInUSGrp struct {
	GrantTypeTxt            string `xml:"GrantTypeTxt,omitempty" json:",omitempty"`
	RecipientCnt            int    `xml:"RecipientCnt,omitempty" json:",omitempty"`
	CashGrantAmt            int    `xml:"CashGrantAmt,omitempty" json:",omitempty"`
	NonCashAssistanceAmt    int    `xml:"NonCashAssistanceAmt,omitempty" json:",omitempty"`
	ValuationMethodUsedDesc string `xml:"ValuationMethodUsedDesc,omitempty" json:",omitempty"`
	NonCashAssistanceDesc   string `xml:"NonCashAssistanceDesc,omitempty" json:",omitempty"`
}

func (r GrantsOtherAsstToIndivInUSGrp) Validate() error {
	return utils.Validate(&r)
}

type GrantsToIndividualsInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r GrantsToIndividualsInd) Validate() error {
	return utils.Validate(&r)
}

type GrantsToOrgOutsideUSGrpType struct {
	RegionTxt                   string `xml:"RegionTxt,omitempty" json:",omitempty"`
	PurposeOfGrantTxt           string `xml:"PurposeOfGrantTxt,omitempty" json:",omitempty"`
	CashGrantAmt                int    `xml:"CashGrantAmt,omitempty" json:",omitempty"`
	MannerOfCashDisbursementTxt string `xml:"MannerOfCashDisbursementTxt,omitempty" json:",omitempty"`
	NonCashAssistanceAmt        int    `xml:"NonCashAssistanceAmt,omitempty" json:",omitempty"`
	DescriptionOfNonCashAsstTxt string `xml:"DescriptionOfNonCashAsstTxt,omitempty" json:",omitempty"`
	ValuationMethodUsedDesc     string `xml:"ValuationMethodUsedDesc,omitempty" json:",omitempty"`
}

func (r GrantsToOrgOutsideUSGrpType) Validate() error {
	return utils.Validate(&r)
}

type GrantsToOrganizationsInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r GrantsToOrganizationsInd) Validate() error {
	return utils.Validate(&r)
}

type GrntAsstBnftInterestedPrsnGrp struct {
	PersonNm               PersonNameType   `xml:"PersonNm"`
	BusinessName           BusinessNameType `xml:"BusinessName"`
	RelationshipWithOrgTxt string           `xml:"RelationshipWithOrgTxt,omitempty" json:",omitempty"`
	CashGrantAmt           int              `xml:"CashGrantAmt,omitempty" json:",omitempty"`
	TypeOfAssistanceTxt    string           `xml:"TypeOfAssistanceTxt,omitempty" json:",omitempty"`
	AssistancePurposeTxt   string           `xml:"AssistancePurposeTxt,omitempty" json:",omitempty"`
}

func (r GrntAsstBnftInterestedPrsnGrp) Validate() error {
	return utils.Validate(&r)
}

type HospitalFacilitiesGrp struct {
	FacilityNum                  int               `xml:"FacilityNum"`
	BusinessName                 *BusinessNameType `xml:"BusinessName,omitempty" json:",omitempty"`
	USAddress                    *USAddressType    `xml:"USAddress,omitempty" json:",omitempty"`
	WebsiteAddressTxt            string            `xml:"WebsiteAddressTxt,omitempty" json:",omitempty"`
	StateLicenseNum              string            `xml:"StateLicenseNum,omitempty" json:",omitempty"`
	SubordinateHospitalName      *BusinessNameType `xml:"SubordinateHospitalName,omitempty" json:",omitempty"`
	SubordinateHospitalEIN       EINType           `xml:"SubordinateHospitalEIN,omitempty" json:",omitempty"`
	LicensedHospitalInd          CheckboxType      `xml:"LicensedHospitalInd,omitempty" json:",omitempty"`
	GeneralMedicalAndSurgicalInd CheckboxType      `xml:"GeneralMedicalAndSurgicalInd,omitempty" json:",omitempty"`
	ChildrensHospitalInd         CheckboxType      `xml:"ChildrensHospitalInd,omitempty" json:",omitempty"`
	TeachingHospitalInd          CheckboxType      `xml:"TeachingHospitalInd,omitempty" json:",omitempty"`
	CriticalAccessHospitalInd    CheckboxType      `xml:"CriticalAccessHospitalInd,omitempty" json:",omitempty"`
	ResearchFacilityInd          CheckboxType      `xml:"ResearchFacilityInd,omitempty" json:",omitempty"`
	EmergencyRoom24HrsInd        CheckboxType      `xml:"EmergencyRoom24HrsInd,omitempty" json:",omitempty"`
	EmergencyRoomOtherInd        CheckboxType      `xml:"EmergencyRoomOtherInd,omitempty" json:",omitempty"`
	OtherDesc                    string            `xml:"OtherDesc,omitempty" json:",omitempty"`
	FacilityReportingGroupCd     string            `xml:"FacilityReportingGroupCd,omitempty" json:",omitempty"`
}

func (r HospitalFacilitiesGrp) Validate() error {
	return utils.Validate(&r)
}

type HospitalFcltyPoliciesPrctcGrp struct {
	HospitalFacilityName           *BusinessNameType            `xml:"HospitalFacilityName,omitempty" json:",omitempty"`
	FacilityNum                    int                          `xml:"FacilityNum,omitempty" json:",omitempty"`
	FirstLicensedCYOrPYInd         bool                         `xml:"FirstLicensedCYOrPYInd"`
	TaxExemptHospitalCYOrPYInd     bool                         `xml:"TaxExemptHospitalCYOrPYInd"`
	CHNAConductedInd               bool                         `xml:"CHNAConductedInd"`
	CommunityDefinitionInd         CheckboxType                 `xml:"CommunityDefinitionInd,omitempty" json:",omitempty"`
	CommunityDemographicsInd       CheckboxType                 `xml:"CommunityDemographicsInd,omitempty" json:",omitempty"`
	ExistingResourcesInd           CheckboxType                 `xml:"ExistingResourcesInd,omitempty" json:",omitempty"`
	HowDataObtainedInd             CheckboxType                 `xml:"HowDataObtainedInd,omitempty" json:",omitempty"`
	CommunityHealthNeedsInd        CheckboxType                 `xml:"CommunityHealthNeedsInd,omitempty" json:",omitempty"`
	OtherHealthIssuesInd           CheckboxType                 `xml:"OtherHealthIssuesInd,omitempty" json:",omitempty"`
	CommunityHlthNeedsIdProcessInd CheckboxType                 `xml:"CommunityHlthNeedsIdProcessInd,omitempty" json:",omitempty"`
	ConsultingProcessInd           CheckboxType                 `xml:"ConsultingProcessInd,omitempty" json:",omitempty"`
	InformationGapsInd             CheckboxType                 `xml:"InformationGapsInd,omitempty" json:",omitempty"`
	OtherInd                       CheckboxType                 `xml:"OtherInd,omitempty" json:",omitempty"`
	CHNAConductedYr                CHNAConductedYr              `xml:"CHNAConductedYr,omitempty" json:",omitempty"`
	TakeIntoAccountOthersInputInd  bool                         `xml:"TakeIntoAccountOthersInputInd,omitempty" json:",omitempty"`
	CHNAConductedWithOtherFcltsInd bool                         `xml:"CHNAConductedWithOtherFcltsInd,omitempty" json:",omitempty"`
	CHNAConductedWithNonFcltsInd   bool                         `xml:"CHNAConductedWithNonFcltsInd,omitempty" json:",omitempty"`
	CHNAReportWidelyAvailableInd   bool                         `xml:"CHNAReportWidelyAvailableInd,omitempty" json:",omitempty"`
	RptAvailableOnOwnWebsiteInd    CheckboxType                 `xml:"RptAvailableOnOwnWebsiteInd,omitempty" json:",omitempty"`
	OwnWebsiteURLTxt               string                       `xml:"OwnWebsiteURLTxt,omitempty" json:",omitempty"`
	OtherWebsiteInd                CheckboxType                 `xml:"OtherWebsiteInd,omitempty" json:",omitempty"`
	OtherWebsiteURLTxt             string                       `xml:"OtherWebsiteURLTxt,omitempty" json:",omitempty"`
	PaperCopyPublicInspectionInd   CheckboxType                 `xml:"PaperCopyPublicInspectionInd,omitempty" json:",omitempty"`
	RptAvailableThruOtherMethodInd CheckboxType                 `xml:"RptAvailableThruOtherMethodInd,omitempty" json:",omitempty"`
	ImplementationStrategyAdoptInd bool                         `xml:"ImplementationStrategyAdoptInd,omitempty" json:",omitempty"`
	ImplementationStrategyAdptYr   ImplementationStrategyAdptYr `xml:"ImplementationStrategyAdptYr,omitempty" json:",omitempty"`
	StrategyPostedWebsiteInd       bool                         `xml:"StrategyPostedWebsiteInd,omitempty" json:",omitempty"`
	StrategyWebsiteURLTxt          string                       `xml:"StrategyWebsiteURLTxt,omitempty" json:",omitempty"`
	StrategyAttachedInd            StrategyAttachedInd          `xml:"StrategyAttachedInd,omitempty" json:",omitempty"`
	OrganizationIncurExciseTaxInd  bool                         `xml:"OrganizationIncurExciseTaxInd,omitempty" json:",omitempty"`
	Form4720FiledInd               bool                         `xml:"Form4720FiledInd,omitempty" json:",omitempty"`
	ExciseReportForm4720ForAllAmt  int                          `xml:"ExciseReportForm4720ForAllAmt,omitempty" json:",omitempty"`
	EligCriteriaExplainedInd       bool                         `xml:"EligCriteriaExplainedInd"`
	FPGFamilyIncmLmtFreeDscntInd   CheckboxType                 `xml:"FPGFamilyIncmLmtFreeDscntInd,omitempty" json:",omitempty"`
	FPGFamilyIncmLmtFreeCarePct    float64                      `xml:"FPGFamilyIncmLmtFreeCarePct,omitempty" json:",omitempty"`
	FPGFamilyIncmLmtDscntCarePct   float64                      `xml:"FPGFamilyIncmLmtDscntCarePct,omitempty" json:",omitempty"`
	IncomeLevelCriteriaInd         CheckboxType                 `xml:"IncomeLevelCriteriaInd,omitempty" json:",omitempty"`
	AssetLevelCriteriaInd          CheckboxType                 `xml:"AssetLevelCriteriaInd,omitempty" json:",omitempty"`
	MedicalIndigencyCriteriaInd    CheckboxType                 `xml:"MedicalIndigencyCriteriaInd,omitempty" json:",omitempty"`
	InsuranceStatusCriteriaInd     CheckboxType                 `xml:"InsuranceStatusCriteriaInd,omitempty" json:",omitempty"`
	UnderinsuranceStatCriteriaInd  CheckboxType                 `xml:"UnderinsuranceStatCriteriaInd,omitempty" json:",omitempty"`
	ResidencyCriteriaInd           CheckboxType                 `xml:"ResidencyCriteriaInd,omitempty" json:",omitempty"`
	OtherCriteriaInd               CheckboxType                 `xml:"OtherCriteriaInd,omitempty" json:",omitempty"`
	ExplainedBasisInd              bool                         `xml:"ExplainedBasisInd"`
	AppFinancialAsstExplnInd       bool                         `xml:"AppFinancialAsstExplnInd"`
	DescribedInfoInd               CheckboxType                 `xml:"DescribedInfoInd,omitempty" json:",omitempty"`
	DescribedSuprtDocInd           CheckboxType                 `xml:"DescribedSuprtDocInd,omitempty" json:",omitempty"`
	ProvidedHospitalContactInd     CheckboxType                 `xml:"ProvidedHospitalContactInd,omitempty" json:",omitempty"`
	ProvidedNonprofitContactInd    CheckboxType                 `xml:"ProvidedNonprofitContactInd,omitempty" json:",omitempty"`
	OtherMethodInd                 CheckboxType                 `xml:"OtherMethodInd,omitempty" json:",omitempty"`
	IncludesPublicityMeasuresInd   bool                         `xml:"IncludesPublicityMeasuresInd"`
	FAPAvailableOnWebsiteInd       CheckboxType                 `xml:"FAPAvailableOnWebsiteInd,omitempty" json:",omitempty"`
	FAPAvailableOnWebsiteURLTxt    string                       `xml:"FAPAvailableOnWebsiteURLTxt,omitempty" json:",omitempty"`
	FAPAppAvailableOnWebsiteInd    CheckboxType                 `xml:"FAPAppAvailableOnWebsiteInd,omitempty" json:",omitempty"`
	FAPAppAvailableOnWebsiteURLTxt string                       `xml:"FAPAppAvailableOnWebsiteURLTxt,omitempty" json:",omitempty"`
	FAPSummaryOnWebsiteInd         CheckboxType                 `xml:"FAPSummaryOnWebsiteInd,omitempty" json:",omitempty"`
	FAPSummaryOnWebsiteURLTxt      string                       `xml:"FAPSummaryOnWebsiteURLTxt,omitempty" json:",omitempty"`
	FAPAvlblOnRequestNoChargeInd   CheckboxType                 `xml:"FAPAvlblOnRequestNoChargeInd,omitempty" json:",omitempty"`
	FAPAppAvlblOnRequestNoChrgInd  CheckboxType                 `xml:"FAPAppAvlblOnRequestNoChrgInd,omitempty" json:",omitempty"`
	FAPSumAvlblOnRequestNoChrgInd  CheckboxType                 `xml:"FAPSumAvlblOnRequestNoChrgInd,omitempty" json:",omitempty"`
	FAPNoticeDisplayedInd          CheckboxType                 `xml:"FAPNoticeDisplayedInd,omitempty" json:",omitempty"`
	CommuntityNotifiedFAPInd       CheckboxType                 `xml:"CommuntityNotifiedFAPInd,omitempty" json:",omitempty"`
	OtherPublicityInd              CheckboxType                 `xml:"OtherPublicityInd,omitempty" json:",omitempty"`
	FAPActionsOnNonpaymentInd      bool                         `xml:"FAPActionsOnNonpaymentInd"`
	PermitReportToCreditAgencyInd  CheckboxType                 `xml:"PermitReportToCreditAgencyInd,omitempty" json:",omitempty"`
	PermitSellingDebtInd           CheckboxType                 `xml:"PermitSellingDebtInd,omitempty" json:",omitempty"`
	PermitLegalJudicialProcessInd  CheckboxType                 `xml:"PermitLegalJudicialProcessInd,omitempty" json:",omitempty"`
	PermitOtherActionsInd          CheckboxType                 `xml:"PermitOtherActionsInd,omitempty" json:",omitempty"`
	PermitNoActionsInd             CheckboxType                 `xml:"PermitNoActionsInd,omitempty" json:",omitempty"`
	CollectionActivitiesInd        bool                         `xml:"CollectionActivitiesInd"`
	ReportingToCreditAgencyInd     CheckboxType                 `xml:"ReportingToCreditAgencyInd,omitempty" json:",omitempty"`
	EngagedSellingDebtInd          CheckboxType                 `xml:"EngagedSellingDebtInd,omitempty" json:",omitempty"`
	EngagedLegalJudicialProcessInd CheckboxType                 `xml:"EngagedLegalJudicialProcessInd,omitempty" json:",omitempty"`
	OtherActionsInd                CheckboxType                 `xml:"OtherActionsInd,omitempty" json:",omitempty"`
	FAPNotifiedUponAdmissionInd    CheckboxType                 `xml:"FAPNotifiedUponAdmissionInd,omitempty" json:",omitempty"`
	FAPNotifiedBeforeDischargeInd  CheckboxType                 `xml:"FAPNotifiedBeforeDischargeInd,omitempty" json:",omitempty"`
	FAPNotifiedAllPatientsInd      CheckboxType                 `xml:"FAPNotifiedAllPatientsInd,omitempty" json:",omitempty"`
	DocumentedEligDeterminationInd CheckboxType                 `xml:"DocumentedEligDeterminationInd,omitempty" json:",omitempty"`
	OtherActionsTakenInd           CheckboxType                 `xml:"OtherActionsTakenInd,omitempty" json:",omitempty"`
	NoneMadeInd                    CheckboxType                 `xml:"NoneMadeInd,omitempty" json:",omitempty"`
	NondisEmergencyCarePolicyInd   bool                         `xml:"NondisEmergencyCarePolicyInd"`
	NoEmergencyCareInd             CheckboxType                 `xml:"NoEmergencyCareInd,omitempty" json:",omitempty"`
	NoEmergencyCarePolicyInd       CheckboxType                 `xml:"NoEmergencyCarePolicyInd,omitempty" json:",omitempty"`
	EmergencyCareLimitedInd        CheckboxType                 `xml:"EmergencyCareLimitedInd,omitempty" json:",omitempty"`
	OtherReasonInd                 CheckboxType                 `xml:"OtherReasonInd,omitempty" json:",omitempty"`
	LowestNegotiatedRatesInd       CheckboxType                 `xml:"LowestNegotiatedRatesInd,omitempty" json:",omitempty"`
	AverageNegotiatedRatesInd      CheckboxType                 `xml:"AverageNegotiatedRatesInd,omitempty" json:",omitempty"`
	MedicareRatesInd               CheckboxType                 `xml:"MedicareRatesInd,omitempty" json:",omitempty"`
	OtherMethodUsedInd             CheckboxType                 `xml:"OtherMethodUsedInd,omitempty" json:",omitempty"`
	AmountsGenerallyBilledInd      bool                         `xml:"AmountsGenerallyBilledInd"`
	GrossChargesInd                bool                         `xml:"GrossChargesInd"`
}

func (r HospitalFcltyPoliciesPrctcGrp) Validate() error {
	return utils.Validate(&r)
}

type HospitalNameAndAddressGrpType struct {
	SupportedOrganizationName *BusinessNameType `xml:"SupportedOrganizationName,omitempty" json:",omitempty"`
	CityNm                    CityType          `xml:"CityNm"`
	StateAbbreviationCd       StateType         `xml:"StateAbbreviationCd"`
	CountryCd                 CountryType       `xml:"CountryCd"`
}

func (r HospitalNameAndAddressGrpType) Validate() error {
	return utils.Validate(&r)
}

// IP address type to include either decimal or hexi decimal format
type IPAddressType struct {
	IPv4AddressTxt IPv4Type `xml:"IPv4AddressTxt"`
	IPv6AddressTxt IPv6Type `xml:"IPv6AddressTxt"`
}

func (r IPAddressType) Validate() error {
	return utils.Validate(&r)
}

type IdDisregardedEntitiesGrp struct {
	DisregardedEntityName         *BusinessNameType     `xml:"DisregardedEntityName,omitempty" json:",omitempty"`
	USAddress                     USAddressType         `xml:"USAddress"`
	ForeignAddress                ForeignAddressType    `xml:"ForeignAddress"`
	EIN                           EINType               `xml:"EIN,omitempty" json:",omitempty"`
	PrimaryActivitiesTxt          string                `xml:"PrimaryActivitiesTxt,omitempty" json:",omitempty"`
	LegalDomicileStateCd          StateType             `xml:"LegalDomicileStateCd"`
	LegalDomicileForeignCountryCd CountryType           `xml:"LegalDomicileForeignCountryCd"`
	TotalIncomeAmt                int                   `xml:"TotalIncomeAmt,omitempty" json:",omitempty"`
	EndOfYearAssetsAmt            int                   `xml:"EndOfYearAssetsAmt,omitempty" json:",omitempty"`
	DirectControllingEntityName   BusinessNameType      `xml:"DirectControllingEntityName"`
	DirectControllingNACd         DirectControllingNACd `xml:"DirectControllingNACd"`
}

func (r IdDisregardedEntitiesGrp) Validate() error {
	return utils.Validate(&r)
}

type IdRelatedOrgTxblCorpTrGrp struct {
	RelatedOrganizationName       *BusinessNameType     `xml:"RelatedOrganizationName,omitempty" json:",omitempty"`
	USAddress                     USAddressType         `xml:"USAddress"`
	ForeignAddress                ForeignAddressType    `xml:"ForeignAddress"`
	EIN                           EINType               `xml:"EIN,omitempty" json:",omitempty"`
	PrimaryActivitiesTxt          string                `xml:"PrimaryActivitiesTxt,omitempty" json:",omitempty"`
	LegalDomicileStateCd          StateType             `xml:"LegalDomicileStateCd"`
	LegalDomicileForeignCountryCd CountryType           `xml:"LegalDomicileForeignCountryCd"`
	DirectControllingEntityName   BusinessNameType      `xml:"DirectControllingEntityName"`
	DirectControllingNACd         DirectControllingNACd `xml:"DirectControllingNACd"`
	EntityTypeTxt                 string                `xml:"EntityTypeTxt,omitempty" json:",omitempty"`
	ShareOfTotalIncomeAmt         int                   `xml:"ShareOfTotalIncomeAmt,omitempty" json:",omitempty"`
	ShareOfEOYAssetsAmt           int                   `xml:"ShareOfEOYAssetsAmt,omitempty" json:",omitempty"`
	OwnershipPct                  float64               `xml:"OwnershipPct,omitempty" json:",omitempty"`
	ControlledOrganizationInd     bool                  `xml:"ControlledOrganizationInd,omitempty" json:",omitempty"`
}

func (r IdRelatedOrgTxblCorpTrGrp) Validate() error {
	return utils.Validate(&r)
}

type IdRelatedOrgTxblPartnershipGrp struct {
	RelatedOrganizationName        *BusinessNameType     `xml:"RelatedOrganizationName,omitempty" json:",omitempty"`
	USAddress                      USAddressType         `xml:"USAddress"`
	ForeignAddress                 ForeignAddressType    `xml:"ForeignAddress"`
	EIN                            EINType               `xml:"EIN,omitempty" json:",omitempty"`
	PrimaryActivitiesTxt           string                `xml:"PrimaryActivitiesTxt,omitempty" json:",omitempty"`
	LegalDomicileStateCd           StateType             `xml:"LegalDomicileStateCd"`
	LegalDomicileForeignCountryCd  CountryType           `xml:"LegalDomicileForeignCountryCd"`
	DirectControllingEntityName    BusinessNameType      `xml:"DirectControllingEntityName"`
	DirectControllingNACd          DirectControllingNACd `xml:"DirectControllingNACd"`
	PredominantIncomeTypeTxt       string                `xml:"PredominantIncomeTypeTxt,omitempty" json:",omitempty"`
	ShareOfTotalIncomeAmt          int                   `xml:"ShareOfTotalIncomeAmt,omitempty" json:",omitempty"`
	ShareOfEOYAssetsAmt            int                   `xml:"ShareOfEOYAssetsAmt,omitempty" json:",omitempty"`
	DisproportionateAllocationsInd bool                  `xml:"DisproportionateAllocationsInd,omitempty" json:",omitempty"`
	UBICodeVAmt                    int                   `xml:"UBICodeVAmt,omitempty" json:",omitempty"`
	GeneralOrManagingPartnerInd    bool                  `xml:"GeneralOrManagingPartnerInd,omitempty" json:",omitempty"`
	OwnershipPct                   float64               `xml:"OwnershipPct,omitempty" json:",omitempty"`
}

func (r IdRelatedOrgTxblPartnershipGrp) Validate() error {
	return utils.Validate(&r)
}

type IdRelatedTaxExemptOrgGrp struct {
	DisregardedEntityName         *BusinessNameType     `xml:"DisregardedEntityName,omitempty" json:",omitempty"`
	USAddress                     USAddressType         `xml:"USAddress"`
	ForeignAddress                ForeignAddressType    `xml:"ForeignAddress"`
	EIN                           EINType               `xml:"EIN,omitempty" json:",omitempty"`
	PrimaryActivitiesTxt          string                `xml:"PrimaryActivitiesTxt,omitempty" json:",omitempty"`
	LegalDomicileStateCd          StateType             `xml:"LegalDomicileStateCd"`
	LegalDomicileForeignCountryCd CountryType           `xml:"LegalDomicileForeignCountryCd"`
	ExemptCodeSectionTxt          string                `xml:"ExemptCodeSectionTxt,omitempty" json:",omitempty"`
	PublicCharityStatusTxt        string                `xml:"PublicCharityStatusTxt,omitempty" json:",omitempty"`
	DirectControllingEntityName   BusinessNameType      `xml:"DirectControllingEntityName"`
	DirectControllingNACd         DirectControllingNACd `xml:"DirectControllingNACd"`
	ControlledOrganizationInd     bool                  `xml:"ControlledOrganizationInd,omitempty" json:",omitempty"`
}

func (r IdRelatedTaxExemptOrgGrp) Validate() error {
	return utils.Validate(&r)
}

type IncludeFIN48FootnoteInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r IncludeFIN48FootnoteInd) Validate() error {
	return utils.Validate(&r)
}

type IndependentAuditFinclStmtInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r IndependentAuditFinclStmtInd) Validate() error {
	return utils.Validate(&r)
}

type LiquidationOfAssetsTableGrp struct {
	LiquidationOfAssetsDetail []Form990SchNGroup1Type `xml:"LiquidationOfAssetsDetail,omitempty" json:",omitempty"`
	ReferenceDocumentId       IdListType              `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName     string                  `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r LiquidationOfAssetsTableGrp) Validate() error {
	return utils.Validate(&r)
}

type LoanOutstandingInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r LoanOutstandingInd) Validate() error {
	return utils.Validate(&r)
}

type LoansBtwnOrgInterestedPrsnGrpType struct {
	PersonNm                    PersonNameType   `xml:"PersonNm"`
	BusinessName                BusinessNameType `xml:"BusinessName"`
	RelationshipWithOrgTxt      string           `xml:"RelationshipWithOrgTxt,omitempty" json:",omitempty"`
	LoanPurposeTxt              string           `xml:"LoanPurposeTxt,omitempty" json:",omitempty"`
	LoanToOrganizationInd       CheckboxType     `xml:"LoanToOrganizationInd"`
	LoanFromOrganizationInd     CheckboxType     `xml:"LoanFromOrganizationInd"`
	OriginalPrincipalAmt        int              `xml:"OriginalPrincipalAmt,omitempty" json:",omitempty"`
	BalanceDueAmt               int              `xml:"BalanceDueAmt,omitempty" json:",omitempty"`
	DefaultInd                  bool             `xml:"DefaultInd,omitempty" json:",omitempty"`
	BoardOrCommitteeApprovalInd bool             `xml:"BoardOrCommitteeApprovalInd,omitempty" json:",omitempty"`
	WrittenAgreementInd         bool             `xml:"WrittenAgreementInd,omitempty" json:",omitempty"`
}

func (r LoansBtwnOrgInterestedPrsnGrpType) Validate() error {
	return utils.Validate(&r)
}

type LobbyingActivitiesInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r LobbyingActivitiesInd) Validate() error {
	return utils.Validate(&r)
}

type ManagementCoAndJntVenturesGrp struct {
	EntityName                     *BusinessNameType `xml:"EntityName,omitempty" json:",omitempty"`
	PrimaryActivitiesTxt           string            `xml:"PrimaryActivitiesTxt,omitempty" json:",omitempty"`
	OrgProfitOrOwnershipPct        float64           `xml:"OrgProfitOrOwnershipPct,omitempty" json:",omitempty"`
	OfcrEtcProfitOrOwnershipPct    float64           `xml:"OfcrEtcProfitOrOwnershipPct,omitempty" json:",omitempty"`
	PhysiciansProfitOrOwnershipPct float64           `xml:"PhysiciansProfitOrOwnershipPct,omitempty" json:",omitempty"`
}

func (r ManagementCoAndJntVenturesGrp) Validate() error {
	return utils.Validate(&r)
}

type MethodOfAccountingOtherInd struct {
	CheckboxType                CheckboxType `xml:",chardata"`
	MethodOfAccountingOtherDesc string       `xml:"methodOfAccountingOtherDesc,attr,omitempty" json:",omitempty"`
}

func (r MethodOfAccountingOtherInd) Validate() error {
	return utils.Validate(&r)
}

type MinimumAssetAmountGrp struct {
	AverageMonthlyFMVOfSecGrp      Form990SchAPartVGroup2Type `xml:"AverageMonthlyFMVOfSecGrp"`
	AverageMonthlyCashBalancesGrp  Form990SchAPartVGroup2Type `xml:"AverageMonthlyCashBalancesGrp"`
	FMVOtherNonExemptUseAssetGrp   Form990SchAPartVGroup2Type `xml:"FMVOtherNonExemptUseAssetGrp"`
	TotalFMVOfNonExemptUseAssetGrp Form990SchAPartVGroup2Type `xml:"TotalFMVOfNonExemptUseAssetGrp"`
	DiscountClaimedAmt             int                        `xml:"DiscountClaimedAmt"`
	AcquisitionIndebtednessGrp     Form990SchAPartVGroup1Type `xml:"AcquisitionIndebtednessGrp"`
	AdjustedFMVLessIndebtednessGrp Form990SchAPartVGroup1Type `xml:"AdjustedFMVLessIndebtednessGrp"`
	CashDeemedCharitableGrp        Form990SchAPartVGroup1Type `xml:"CashDeemedCharitableGrp"`
	NetVlNonExemptUseAssetsGrp     Form990SchAPartVGroup1Type `xml:"NetVlNonExemptUseAssetsGrp"`
	PctOfNetVlNonExemptUseAstGrp   Form990SchAPartVGroup1Type `xml:"PctOfNetVlNonExemptUseAstGrp"`
	RecoveriesPYDistriMinAssetGrp  Form990SchAPartVGroup1Type `xml:"RecoveriesPYDistriMinAssetGrp"`
	TotalMinimumAssetGrp           Form990SchAPartVGroup1Type `xml:"TotalMinimumAssetGrp"`
}

func (r MinimumAssetAmountGrp) Validate() error {
	return utils.Validate(&r)
}

type MoreThan5000KToIndividualsInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r MoreThan5000KToIndividualsInd) Validate() error {
	return utils.Validate(&r)
}

type MoreThan5000KToOrgInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r MoreThan5000KToOrgInd) Validate() error {
	return utils.Validate(&r)
}

// Recurring Name and Address Type
type NameAndAddressType struct {
	PersonNm       PersonNameType     `xml:"PersonNm"`
	BusinessName   BusinessNameType   `xml:"BusinessName"`
	USAddress      USAddressType      `xml:"USAddress"`
	ForeignAddress ForeignAddressType `xml:"ForeignAddress"`
}

func (r NameAndAddressType) Validate() error {
	return utils.Validate(&r)
}

type NameOfInterested struct {
	PersonNm     PersonNameType   `xml:"PersonNm"`
	BusinessName BusinessNameType `xml:"BusinessName"`
}

func (r NameOfInterested) Validate() error {
	return utils.Validate(&r)
}

type NonCashPropertyContributionGrpType struct {
	ContributorNum      int      `xml:"ContributorNum"`
	NoncashPropertyDesc string   `xml:"NoncashPropertyDesc"`
	FairMarketValueAmt  int      `xml:"FairMarketValueAmt"`
	ReceivedDt          DateType `xml:"ReceivedDt,omitempty" json:",omitempty"`
}

func (r NonCashPropertyContributionGrpType) Validate() error {
	return utils.Validate(&r)
}

type OperateHospitalInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r OperateHospitalInd) Validate() error {
	return utils.Validate(&r)
}

type Organization501cInd struct {
	CheckboxType            CheckboxType            `xml:",chardata"`
	Organization501cTypeTxt Organization501cTypeTxt `xml:"organization501cTypeTxt,attr,omitempty" json:",omitempty"`
}

func (r Organization501cInd) Validate() error {
	return utils.Validate(&r)
}

type OrganizationBelongsAffltGrpInd struct {
	CheckboxType          CheckboxType `xml:",chardata"`
	ReferenceDocumentId   IdListType   `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string       `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r OrganizationBelongsAffltGrpInd) Validate() error {
	return utils.Validate(&r)
}

type OriginatorGrp struct {
	EFIN               EFINType            `xml:"EFIN"`
	OriginatorTypeCd   OriginatorType      `xml:"OriginatorTypeCd"`
	PractitionerPINGrp *PractitionerPINGrp `xml:"PractitionerPINGrp,omitempty" json:",omitempty"`
}

func (r OriginatorGrp) Validate() error {
	return utils.Validate(&r)
}

type OthHlthCareFcltsGrp struct {
	BusinessName BusinessNameType `xml:"BusinessName"`
	USAddress    USAddressType    `xml:"USAddress"`
	FacilityTxt  string           `xml:"FacilityTxt"`
}

func (r OthHlthCareFcltsGrp) Validate() error {
	return utils.Validate(&r)
}

type OthHlthCareFcltsNotHospitalGrp struct {
	OthHlthCareFcltsGrp []OthHlthCareFcltsGrp `xml:"OthHlthCareFcltsGrp"`
}

func (r OthHlthCareFcltsNotHospitalGrp) Validate() error {
	return utils.Validate(&r)
}

type OtherForeignAddressType struct {
	AddressLine1Txt   *StreetAddressType `xml:"AddressLine1Txt,omitempty" json:",omitempty"`
	AddressLine2Txt   *StreetAddressType `xml:"AddressLine2Txt,omitempty" json:",omitempty"`
	CityNm            string             `xml:"CityNm,omitempty" json:",omitempty"`
	ProvinceOrStateNm string             `xml:"ProvinceOrStateNm,omitempty" json:",omitempty"`
	CountryCd         CountryType        `xml:"CountryCd,omitempty" json:",omitempty"`
	ForeignPostalCd   string             `xml:"ForeignPostalCd,omitempty" json:",omitempty"`
}

func (r OtherForeignAddressType) Validate() error {
	return utils.Validate(&r)
}

type OtherUSAddressType struct {
	AddressLine1Txt     *StreetAddressType `xml:"AddressLine1Txt,omitempty" json:",omitempty"`
	AddressLine2Txt     *StreetAddressType `xml:"AddressLine2Txt,omitempty" json:",omitempty"`
	CityNm              CityType           `xml:"CityNm,omitempty" json:",omitempty"`
	StateAbbreviationCd *StateType         `xml:"StateAbbreviationCd,omitempty" json:",omitempty"`
	ZIPCd               ZIPCodeType        `xml:"ZIPCd,omitempty" json:",omitempty"`
}

func (r OtherUSAddressType) Validate() error {
	return utils.Validate(&r)
}

type PYExcessBenefitTransInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r PYExcessBenefitTransInd) Validate() error {
	return utils.Validate(&r)
}

type PartialLiquidationInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r PartialLiquidationInd) Validate() error {
	return utils.Validate(&r)
}

type PersonFullNameType struct {
	PersonFirstNm PersonFirstNameType `xml:"PersonFirstNm"`
	PersonLastNm  PersonLastNameType  `xml:"PersonLastNm,omitempty" json:",omitempty"`
}

func (r PersonFullNameType) Validate() error {
	return utils.Validate(&r)
}

type PoliticalCampaignActyInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r PoliticalCampaignActyInd) Validate() error {
	return utils.Validate(&r)
}

type PractitionerPINGrp struct {
	EFIN EFINType `xml:"EFIN"`
	PIN  PINType  `xml:"PIN"`
}

func (r PractitionerPINGrp) Validate() error {
	return utils.Validate(&r)
}

type PreparerFirmGrp struct {
	PreparerFirmEIN        EINType            `xml:"PreparerFirmEIN,omitempty" json:",omitempty"`
	PreparerFirmName       BusinessNameType   `xml:"PreparerFirmName"`
	PreparerUSAddress      USAddressType      `xml:"PreparerUSAddress"`
	PreparerForeignAddress ForeignAddressType `xml:"PreparerForeignAddress"`
}

func (r PreparerFirmGrp) Validate() error {
	return utils.Validate(&r)
}

type PreparerPersonGrp struct {
	PreparerPersonNm PersonNameType  `xml:"PreparerPersonNm,omitempty" json:",omitempty"`
	SSN              SSNType         `xml:"SSN"`
	PTIN             PTINType        `xml:"PTIN"`
	PhoneNum         PhoneNumberType `xml:"PhoneNum,omitempty" json:",omitempty"`
	EmailAddressTxt  string          `xml:"EmailAddressTxt,omitempty" json:",omitempty"`
	PreparationDt    DateType        `xml:"PreparationDt,omitempty" json:",omitempty"`
	SelfEmployedInd  CheckboxType    `xml:"SelfEmployedInd,omitempty" json:",omitempty"`
}

func (r PreparerPersonGrp) Validate() error {
	return utils.Validate(&r)
}

type ProceduresCorrectiveActionGrpType struct {
	BondReferenceCd               BondReferenceCd `xml:"BondReferenceCd"`
	ProceduresCorrectiveActionInd bool            `xml:"ProceduresCorrectiveActionInd,omitempty" json:",omitempty"`
}

func (r ProceduresCorrectiveActionGrpType) Validate() error {
	return utils.Validate(&r)
}

type ProfessionalFundraisingInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ProfessionalFundraisingInd) Validate() error {
	return utils.Validate(&r)
}

type ProgSrvcAccomplishmentActyGrpType struct {
	ActivityCd int    `xml:"ActivityCd,omitempty" json:",omitempty"`
	ExpenseAmt int    `xml:"ExpenseAmt,omitempty" json:",omitempty"`
	GrantAmt   int    `xml:"GrantAmt,omitempty" json:",omitempty"`
	RevenueAmt int    `xml:"RevenueAmt,omitempty" json:",omitempty"`
	Desc       string `xml:"Desc,omitempty" json:",omitempty"`
}

func (r ProgSrvcAccomplishmentActyGrpType) Validate() error {
	return utils.Validate(&r)
}

type RecipientTable struct {
	RecipientBusinessName   *BusinessNameType  `xml:"RecipientBusinessName,omitempty" json:",omitempty"`
	USAddress               USAddressType      `xml:"USAddress"`
	ForeignAddress          ForeignAddressType `xml:"ForeignAddress"`
	RecipientEIN            EINType            `xml:"RecipientEIN,omitempty" json:",omitempty"`
	IRCSectionDesc          string             `xml:"IRCSectionDesc,omitempty" json:",omitempty"`
	CashGrantAmt            int                `xml:"CashGrantAmt,omitempty" json:",omitempty"`
	NonCashAssistanceAmt    int                `xml:"NonCashAssistanceAmt,omitempty" json:",omitempty"`
	ValuationMethodUsedDesc string             `xml:"ValuationMethodUsedDesc,omitempty" json:",omitempty"`
	NonCashAssistanceDesc   string             `xml:"NonCashAssistanceDesc,omitempty" json:",omitempty"`
	PurposeOfGrantTxt       string             `xml:"PurposeOfGrantTxt,omitempty" json:",omitempty"`
}

func (r RecipientTable) Validate() error {
	return utils.Validate(&r)
}

type RelatedEntityInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r RelatedEntityInd) Validate() error {
	return utils.Validate(&r)
}

type ReportInvestmentsOtherSecInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ReportInvestmentsOtherSecInd) Validate() error {
	return utils.Validate(&r)
}

type ReportLandBuildingEquipmentInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ReportLandBuildingEquipmentInd) Validate() error {
	return utils.Validate(&r)
}

type ReportOtherAssetsInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ReportOtherAssetsInd) Validate() error {
	return utils.Validate(&r)
}

type ReportOtherLiabilitiesInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ReportOtherLiabilitiesInd) Validate() error {
	return utils.Validate(&r)
}

type ReportProgramRelatedInvstInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ReportProgramRelatedInvstInd) Validate() error {
	return utils.Validate(&r)
}

type RltdOrgOfficerTrstKeyEmplGrp struct {
	PersonNm                       PersonNameType   `xml:"PersonNm"`
	BusinessName                   BusinessNameType `xml:"BusinessName"`
	TitleTxt                       string           `xml:"TitleTxt,omitempty" json:",omitempty"`
	BaseCompensationFilingOrgAmt   int              `xml:"BaseCompensationFilingOrgAmt,omitempty" json:",omitempty"`
	CompensationBasedOnRltdOrgsAmt int              `xml:"CompensationBasedOnRltdOrgsAmt,omitempty" json:",omitempty"`
	BonusFilingOrganizationAmount  int              `xml:"BonusFilingOrganizationAmount,omitempty" json:",omitempty"`
	BonusRelatedOrganizationsAmt   int              `xml:"BonusRelatedOrganizationsAmt,omitempty" json:",omitempty"`
	OtherCompensationFilingOrgAmt  int              `xml:"OtherCompensationFilingOrgAmt,omitempty" json:",omitempty"`
	OtherCompensationRltdOrgsAmt   int              `xml:"OtherCompensationRltdOrgsAmt,omitempty" json:",omitempty"`
	DeferredCompensationFlngOrgAmt int              `xml:"DeferredCompensationFlngOrgAmt,omitempty" json:",omitempty"`
	DeferredCompRltdOrgsAmt        int              `xml:"DeferredCompRltdOrgsAmt,omitempty" json:",omitempty"`
	NontaxableBenefitsFilingOrgAmt int              `xml:"NontaxableBenefitsFilingOrgAmt,omitempty" json:",omitempty"`
	NontaxableBenefitsRltdOrgsAmt  int              `xml:"NontaxableBenefitsRltdOrgsAmt,omitempty" json:",omitempty"`
	TotalCompensationFilingOrgAmt  int              `xml:"TotalCompensationFilingOrgAmt,omitempty" json:",omitempty"`
	TotalCompensationRltdOrgsAmt   int              `xml:"TotalCompensationRltdOrgsAmt,omitempty" json:",omitempty"`
	CompReportPrior990FilingOrgAmt int              `xml:"CompReportPrior990FilingOrgAmt,omitempty" json:",omitempty"`
	CompReportPrior990RltdOrgsAmt  int              `xml:"CompReportPrior990RltdOrgsAmt,omitempty" json:",omitempty"`
}

func (r RltdOrgOfficerTrstKeyEmplGrp) Validate() error {
	return utils.Validate(&r)
}

type ScheduleBRequiredInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ScheduleBRequiredInd) Validate() error {
	return utils.Validate(&r)
}

type ScheduleJRequiredInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r ScheduleJRequiredInd) Validate() error {
	return utils.Validate(&r)
}

type SchoolOperatingInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r SchoolOperatingInd) Validate() error {
	return utils.Validate(&r)
}

type Section527PoliticalOrgGrp struct {
	OrganizationBusinessName *BusinessNameType  `xml:"OrganizationBusinessName,omitempty" json:",omitempty"`
	USAddress                USAddressType      `xml:"USAddress"`
	ForeignAddress           ForeignAddressType `xml:"ForeignAddress"`
	EIN                      EINType            `xml:"EIN,omitempty" json:",omitempty"`
	PaidInternalFundsAmt     int                `xml:"PaidInternalFundsAmt,omitempty" json:",omitempty"`
	ContributionsRcvdDlvrAmt int                `xml:"ContributionsRcvdDlvrAmt,omitempty" json:",omitempty"`
}

func (r Section527PoliticalOrgGrp) Validate() error {
	return utils.Validate(&r)
}

type StrategyAttachedInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r StrategyAttachedInd) Validate() error {
	return utils.Validate(&r)
}

type SubjectToProxyTaxInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r SubjectToProxyTaxInd) Validate() error {
	return utils.Validate(&r)
}

type TaxExemptBondsInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r TaxExemptBondsInd) Validate() error {
	return utils.Validate(&r)
}

type TempOrPermanentEndowmentsInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r TempOrPermanentEndowmentsInd) Validate() error {
	return utils.Validate(&r)
}

type TerminateOperationsInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r TerminateOperationsInd) Validate() error {
	return utils.Validate(&r)
}

type TrnsfrExmptNonChrtblRltdOrgInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r TrnsfrExmptNonChrtblRltdOrgInd) Validate() error {
	return utils.Validate(&r)
}

type USAddressType struct {
	AddressLine1Txt     StreetAddressType  `xml:"AddressLine1Txt"`
	AddressLine2Txt     *StreetAddressType `xml:"AddressLine2Txt,omitempty" json:",omitempty"`
	CityNm              CityType           `xml:"CityNm"`
	StateAbbreviationCd StateType          `xml:"StateAbbreviationCd"`
	ZIPCd               ZIPCodeType        `xml:"ZIPCd"`
}

func (r USAddressType) Validate() error {
	return utils.Validate(&r)
}

type USItemizedEntryType struct {
	Desc string `xml:"Desc"`
	Amt  int    `xml:"Amt"`
}

func (r USItemizedEntryType) Validate() error {
	return utils.Validate(&r)
}

type UnrelatedOrgTxblPartnershipGrp struct {
	BusinessName                   *BusinessNameType  `xml:"BusinessName,omitempty" json:",omitempty"`
	USAddress                      USAddressType      `xml:"USAddress"`
	ForeignAddress                 ForeignAddressType `xml:"ForeignAddress"`
	EIN                            EINType            `xml:"EIN,omitempty" json:",omitempty"`
	PrimaryActivitiesTxt           string             `xml:"PrimaryActivitiesTxt,omitempty" json:",omitempty"`
	LegalDomicileStateCd           StateType          `xml:"LegalDomicileStateCd"`
	LegalDomicileForeignCountryCd  CountryType        `xml:"LegalDomicileForeignCountryCd"`
	PredominateIncomeDesc          string             `xml:"PredominateIncomeDesc,omitempty" json:",omitempty"`
	AllPartnersC3SInd              bool               `xml:"AllPartnersC3SInd,omitempty" json:",omitempty"`
	ShareOfTotalIncomeAmt          int                `xml:"ShareOfTotalIncomeAmt,omitempty" json:",omitempty"`
	ShareOfEOYAssetsAmt            int                `xml:"ShareOfEOYAssetsAmt,omitempty" json:",omitempty"`
	DisproportionateAllocationsInd bool               `xml:"DisproportionateAllocationsInd,omitempty" json:",omitempty"`
	UBICodeVAmt                    int                `xml:"UBICodeVAmt,omitempty" json:",omitempty"`
	GeneralOrManagingPartnerInd    bool               `xml:"GeneralOrManagingPartnerInd,omitempty" json:",omitempty"`
	OwnershipPct                   float64            `xml:"OwnershipPct,omitempty" json:",omitempty"`
}

func (r UnrelatedOrgTxblPartnershipGrp) Validate() error {
	return utils.Validate(&r)
}

type VehicleDescriptionGrpType struct {
	VehicleModelYr      YearType `xml:"VehicleModelYr"`
	VehicleMakeNameTxt  string   `xml:"VehicleMakeNameTxt"`
	VehicleModelNameTxt string   `xml:"VehicleModelNameTxt"`
}

func (r VehicleDescriptionGrpType) Validate() error {
	return utils.Validate(&r)
}

type SupplementalInformationDetail struct {
	FormAndLineReferenceDesc string `xml:"FormAndLineReferenceDesc,omitempty" json:",omitempty"`
	ExplanationTxt           string `xml:"ExplanationTxt,omitempty" json:",omitempty"`
}

func (r SupplementalInformationDetail) Validate() error {
	return utils.Validate(&r)
}

// Content model for general explanation, IRS Form 990 Schedule O
type SupplementalInformationDetailType struct {
	FormAndLineReferenceDesc string `xml:"FormAndLineReferenceDesc,omitempty" json:",omitempty"`
	ExplanationTxt           string `xml:"ExplanationTxt"`
}

func (r SupplementalInformationDetailType) Validate() error {
	return utils.Validate(&r)
}

type SupplementalInformationGrp struct {
	FormAndLineReferenceDesc string `xml:"FormAndLineReferenceDesc,omitempty" json:",omitempty"`
	ExplanationTxt           string `xml:"ExplanationTxt,omitempty" json:",omitempty"`
}

func (r SupplementalInformationGrp) Validate() error {
	return utils.Validate(&r)
}

type SupportedOrgInformationGrpType struct {
	SupportedOrganizationName  BusinessNameType   `xml:"SupportedOrganizationName"`
	EIN                        EINType            `xml:"EIN"`
	OrganizationTypeCd         OrganizationTypeCd `xml:"OrganizationTypeCd"`
	GoverningDocumentListedInd bool               `xml:"GoverningDocumentListedInd"`
	SupportAmt                 int                `xml:"SupportAmt"`
	OtherSupportAmt            int                `xml:"OtherSupportAmt"`
}

func (r SupportedOrgInformationGrpType) Validate() error {
	return utils.Validate(&r)
}

type TaxExemptBondsArbitrageGrpType struct {
	BondReferenceCd               BondReferenceCd   `xml:"BondReferenceCd"`
	Form8038TFiledInd             bool              `xml:"Form8038TFiledInd,omitempty" json:",omitempty"`
	RebateNotDueYetInd            bool              `xml:"RebateNotDueYetInd,omitempty" json:",omitempty"`
	ExceptionToRebateInd          bool              `xml:"ExceptionToRebateInd,omitempty" json:",omitempty"`
	NoRebateDueInd                bool              `xml:"NoRebateDueInd,omitempty" json:",omitempty"`
	VariableRateIssueInd          bool              `xml:"VariableRateIssueInd,omitempty" json:",omitempty"`
	HedgeIdentifiedInBksAndRecInd bool              `xml:"HedgeIdentifiedInBksAndRecInd,omitempty" json:",omitempty"`
	HedgeProviderName             *BusinessNameType `xml:"HedgeProviderName,omitempty" json:",omitempty"`
	TermOfHedgePct                float64           `xml:"TermOfHedgePct,omitempty" json:",omitempty"`
	SuperintegratedHedgeInd       bool              `xml:"SuperintegratedHedgeInd,omitempty" json:",omitempty"`
	HedgeTerminatedInd            bool              `xml:"HedgeTerminatedInd,omitempty" json:",omitempty"`
	GrossProceedsInvestedInGICInd bool              `xml:"GrossProceedsInvestedInGICInd,omitempty" json:",omitempty"`
	GICProviderName               *BusinessNameType `xml:"GICProviderName,omitempty" json:",omitempty"`
	TermOfGICPct                  float64           `xml:"TermOfGICPct,omitempty" json:",omitempty"`
	RegulatorySafeHarborStsfdInd  bool              `xml:"RegulatorySafeHarborStsfdInd,omitempty" json:",omitempty"`
	GrossProceedsInvestedInd      bool              `xml:"GrossProceedsInvestedInd,omitempty" json:",omitempty"`
	WrittenProcToMonitorReqsInd   bool              `xml:"WrittenProcToMonitorReqsInd,omitempty" json:",omitempty"`
}

func (r TaxExemptBondsArbitrageGrpType) Validate() error {
	return utils.Validate(&r)
}

type TaxExemptBondsIssuesGrpType struct {
	BondReferenceCd     BondReferenceCd   `xml:"BondReferenceCd"`
	IssuerName          *BusinessNameType `xml:"IssuerName,omitempty" json:",omitempty"`
	BondIssuerEIN       EINType           `xml:"BondIssuerEIN,omitempty" json:",omitempty"`
	CUSIPNum            CUSIPNumberType   `xml:"CUSIPNum,omitempty" json:",omitempty"`
	BondIssuedDt        DateType          `xml:"BondIssuedDt,omitempty" json:",omitempty"`
	IssuePriceAmt       int               `xml:"IssuePriceAmt,omitempty" json:",omitempty"`
	PurposeDesc         string            `xml:"PurposeDesc,omitempty" json:",omitempty"`
	DefeasedInd         bool              `xml:"DefeasedInd,omitempty" json:",omitempty"`
	OnBehalfOfIssuerInd bool              `xml:"OnBehalfOfIssuerInd,omitempty" json:",omitempty"`
	PoolFinancingInd    bool              `xml:"PoolFinancingInd,omitempty" json:",omitempty"`
}

func (r TaxExemptBondsIssuesGrpType) Validate() error {
	return utils.Validate(&r)
}

type TaxExemptBondsPrivateBusUseGrpType struct {
	BondReferenceCd                BondReferenceCd `xml:"BondReferenceCd,omitempty" json:",omitempty"`
	OwningBondFinancedPropertyInd  bool            `xml:"OwningBondFinancedPropertyInd,omitempty" json:",omitempty"`
	AnyLeaseArrangementsInd        bool            `xml:"AnyLeaseArrangementsInd,omitempty" json:",omitempty"`
	MgmtContractBondFincdPropInd   bool            `xml:"MgmtContractBondFincdPropInd,omitempty" json:",omitempty"`
	EngageBondCounselContractsInd  bool            `xml:"EngageBondCounselContractsInd,omitempty" json:",omitempty"`
	AnyResearchAgreementsInd       bool            `xml:"AnyResearchAgreementsInd,omitempty" json:",omitempty"`
	EngageBondCounselResearchInd   bool            `xml:"EngageBondCounselResearchInd,omitempty" json:",omitempty"`
	PrivateBusUseByOthersPct       float64         `xml:"PrivateBusUseByOthersPct,omitempty" json:",omitempty"`
	PrivateBusConcerningUBIPct     float64         `xml:"PrivateBusConcerningUBIPct,omitempty" json:",omitempty"`
	TotalPrivateBusinessUsePct     float64         `xml:"TotalPrivateBusinessUsePct,omitempty" json:",omitempty"`
	BondIssMeetPrvtSecPymtTestInd  bool            `xml:"BondIssMeetPrvtSecPymtTestInd,omitempty" json:",omitempty"`
	ChangeInUseBondFinancedPropInd bool            `xml:"ChangeInUseBondFinancedPropInd,omitempty" json:",omitempty"`
	ChangeInUseBondFinancedPropPct float64         `xml:"ChangeInUseBondFinancedPropPct,omitempty" json:",omitempty"`
	RemedialActionTakenInd         bool            `xml:"RemedialActionTakenInd,omitempty" json:",omitempty"`
	ProcsNonqualifiedBondRemdtdInd bool            `xml:"ProcsNonqualifiedBondRemdtdInd,omitempty" json:",omitempty"`
}

func (r TaxExemptBondsPrivateBusUseGrpType) Validate() error {
	return utils.Validate(&r)
}

type TaxExemptBondsProceedsGrpType struct {
	BondReferenceCd               BondReferenceCd `xml:"BondReferenceCd"`
	RetiredAmt                    int             `xml:"RetiredAmt,omitempty" json:",omitempty"`
	BondDefeasedAmt               int             `xml:"BondDefeasedAmt,omitempty" json:",omitempty"`
	TotalProceedsAmt              int             `xml:"TotalProceedsAmt,omitempty" json:",omitempty"`
	InReserveFundAmt              int             `xml:"InReserveFundAmt,omitempty" json:",omitempty"`
	CapitalizedInterestAmt        int             `xml:"CapitalizedInterestAmt,omitempty" json:",omitempty"`
	RefundingEscrowAmt            int             `xml:"RefundingEscrowAmt,omitempty" json:",omitempty"`
	IssuanceCostsFromProceedsAmt  int             `xml:"IssuanceCostsFromProceedsAmt,omitempty" json:",omitempty"`
	CreditEnhancementAmt          int             `xml:"CreditEnhancementAmt,omitempty" json:",omitempty"`
	WorkingCapitalExpendituresAmt int             `xml:"WorkingCapitalExpendituresAmt,omitempty" json:",omitempty"`
	CapitalExpendituresAmt        int             `xml:"CapitalExpendituresAmt,omitempty" json:",omitempty"`
	OtherSpentProceedsAmt         int             `xml:"OtherSpentProceedsAmt,omitempty" json:",omitempty"`
	UnspentAmt                    int             `xml:"UnspentAmt,omitempty" json:",omitempty"`
	SubstantialCompletionYr       YearType        `xml:"SubstantialCompletionYr,omitempty" json:",omitempty"`
	CurrentRefundingInd           bool            `xml:"CurrentRefundingInd,omitempty" json:",omitempty"`
	AdvanceRefundingInd           bool            `xml:"AdvanceRefundingInd,omitempty" json:",omitempty"`
	FinalAllocationMadeInd        bool            `xml:"FinalAllocationMadeInd,omitempty" json:",omitempty"`
	AdequateBooksAndRecMaintInd   bool            `xml:"AdequateBooksAndRecMaintInd,omitempty" json:",omitempty"`
}

func (r TaxExemptBondsProceedsGrpType) Validate() error {
	return utils.Validate(&r)
}

type TotContriRcvdUpTo1000Ind struct {
	CheckboxType          CheckboxType `xml:",chardata"`
	TotalContributionsAmt int          `xml:"totalContributionsAmt,attr,omitempty" json:",omitempty"`
}

func (r TotContriRcvdUpTo1000Ind) Validate() error {
	return utils.Validate(&r)
}

type TransactionWithControlEntInd struct {
	Value                 bool       `xml:",chardata"`
	ReferenceDocumentId   IdListType `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName string     `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

func (r TransactionWithControlEntInd) Validate() error {
	return utils.Validate(&r)
}

type TransactionsRelatedOrgGrpType struct {
	OtherOrganizationName          *BusinessNameType `xml:"OtherOrganizationName,omitempty" json:",omitempty"`
	TransactionTypeTxt             string            `xml:"TransactionTypeTxt,omitempty" json:",omitempty"`
	InvolvedAmt                    int               `xml:"InvolvedAmt,omitempty" json:",omitempty"`
	MethodOfAmountDeterminationTxt string            `xml:"MethodOfAmountDeterminationTxt,omitempty" json:",omitempty"`
}

func (r TransactionsRelatedOrgGrpType) Validate() error {
	return utils.Validate(&r)
}

type ValidationErrorListType struct {
	ValidationErrorGrp []ValidationErrorGrp `xml:"ValidationErrorGrp"`
	ErrorCnt           int                  `xml:"errorCnt,attr"`
}

func (r ValidationErrorListType) Validate() error {
	return utils.Validate(&r)
}

type ValidationErrorGrp struct {
	DocumentId      IdType `xml:"DocumentId"`
	XpathContentTxt string `xml:"XpathContentTxt,omitempty" json:",omitempty"`
	ErrorCategoryCd string `xml:"ErrorCategoryCd"`
	ErrorMessageTxt string `xml:"ErrorMessageTxt"`
	RuleNum         string `xml:"RuleNum"`
	SeverityCd      string `xml:"SeverityCd"`
	FieldValueTxt   string `xml:"FieldValueTxt,omitempty" json:",omitempty"`
	ErrorId         int    `xml:"errorId,attr"`
}

func (r ValidationErrorGrp) Validate() error {
	return utils.Validate(&r)
}

type ValidationAlertListType struct {
	ValidationAlertGrp []ValidationAlertGrp `xml:"ValidationAlertGrp"`
	AlertCnt           int                  `xml:"alertCnt,attr"`
}

func (r ValidationAlertListType) Validate() error {
	return utils.Validate(&r)
}

type ValidationAlertGrp struct {
	DocumentId      IdType `xml:"DocumentId"`
	XpathContentTxt string `xml:"XpathContentTxt,omitempty" json:",omitempty"`
	AlertCategoryCd string `xml:"AlertCategoryCd"`
	AlertMessageTxt string `xml:"AlertMessageTxt"`
	RuleNum         string `xml:"RuleNum"`
	SeverityCd      string `xml:"SeverityCd"`
	FieldValueTxt   string `xml:"FieldValueTxt,omitempty" json:",omitempty"`
	AlertId         int    `xml:"alertId,attr"`
}

func (r ValidationAlertGrp) Validate() error {
	return utils.Validate(&r)
}

type SubmissionReceiptGrp struct {
	SubmissionId         SubmissionIdType `xml:"SubmissionId"`
	SubmissionReceivedTs TimestampType    `xml:"SubmissionReceivedTs"`
}

func (r SubmissionReceiptGrp) Validate() error {
	return utils.Validate(&r)
}

type StatusRecordGrp struct {
	SubmissionId                  SubmissionIdType `SubmissionId"`
	SubmissionStatusTxt           string           `SubmissionStatusTxt"`
	SubmsnStatusAcknowledgementDt DateType         `SubmsnStatusAcknowledgementDt"`
	DisclaimerTxt                 string           `DisclaimerTxt,omitempty" json:",omitempty"`
}

func (r StatusRecordGrp) Validate() error {
	return utils.Validate(&r)
}

type Acknowledgement struct {
	SubmissionId                SubmissionIdType              `xml:"SubmissionId"`
	EFIN                        EFINType                      `xml:"EFIN"`
	ExtndGovernmentCd           ExtndGovernmentCdType         `xml:"ExtndGovernmentCd"`
	SubmissionTyp               SubmissionTyp                 `xml:"SubmissionTyp"`
	ExtndSubmissionCategoryCd   ExtndSubmissionCategoryCdType `xml:"ExtndSubmissionCategoryCd"`
	AcceptanceStatusTxt         string                        `xml:"AcceptanceStatusTxt"`
	ContainedAlertsInd          bool                          `xml:"ContainedAlertsInd"`
	StatusDt                    DateType                      `xml:"StatusDt"`
	TIN                         EINType                       `xml:"TIN"`
	TempId                      TempIdType                    `xml:"TempId"`
	ExpectedRefundAmt           int                           `xml:"ExpectedRefundAmt"`
	BalanceDueAmt               int                           `xml:"BalanceDueAmt"`
	TaxYr                       *YearType                     `xml:"TaxYr,omitempty" json:",omitempty"`
	ElectronicPostmarkTs        *TimestampType                `xml:"ElectronicPostmarkTs,omitempty" json:",omitempty"`
	IRSSubmissionId             *SubmissionIdType             `xml:"IRSSubmissionId,omitempty" json:",omitempty"`
	StateSubmissionCopyCnt      int                           `xml:"StateSubmissionCopyCnt,omitempty" json:",omitempty"`
	IRSReceivedDt               *DateType                     `xml:"IRSReceivedDt,omitempty" json:",omitempty"`
	ExtensionFilingTypeDesc     string                        `xml:"ExtensionFilingTypeDesc,omitempty" json:",omitempty"`
	TaxPeriodEndDt              *DateType                     `xml:"TaxPeriodEndDt,omitempty" json:",omitempty"`
	PaymentRequestRcvdCd        *PaymentRequestRcvdCd         `xml:"PaymentRequestRcvdCd,omitempty" json:",omitempty"`
	SubmissionValidationCompInd bool                          `xml:"SubmissionValidationCompInd,omitempty" json:",omitempty"`
	EmbeddedCRC32Num            *EmbeddedCRC32Num             `xml:"EmbeddedCRC32Num,omitempty" json:",omitempty"`
	ComputedCRC32Num            *ComputedCRC32Num             `xml:"ComputedCRC32Num,omitempty" json:",omitempty"`
	TaxableIncomeAmt            int                           `xml:"TaxableIncomeAmt,omitempty" json:",omitempty"`
	TotalTaxAmt                 int                           `xml:"TotalTaxAmt,omitempty" json:",omitempty"`
	NetIncomeLossAmt            int                           `xml:"NetIncomeLossAmt,omitempty" json:",omitempty"`
	ReservedIPAddressCd         string                        `xml:"ReservedIPAddressCd,omitempty" json:",omitempty"`
	BirthDtValidityCd           string                        `xml:"BirthDtValidityCd,omitempty" json:",omitempty"`
	PINPresenceCd               string                        `xml:"PINPresenceCd,omitempty" json:",omitempty"`
	ITINMismatchCd              string                        `xml:"ITINMismatchCd,omitempty" json:",omitempty"`
	ValidationErrorList         *ValidationErrorListType      `xml:"ValidationErrorList,omitempty" json:",omitempty"`
	ValidationAlertList         *ValidationAlertListType      `xml:"ValidationAlertList,omitempty" json:",omitempty"`
	SubmissionVersionNum        string                        `xml:"submissionVersionNum,attr,omitempty" json:",omitempty"`
	ValidatingSchemaVersionNum  string                        `xml:"validatingSchemaVersionNum,attr,omitempty" json:",omitempty"`
}

func (r Acknowledgement) Validate() error {
	return utils.Validate(&r)
}

type AcknowledgementList struct {
	Cnt             int               `xml:"Cnt"`
	Acknowledgement []Acknowledgement `xml:"http://www.irs.gov/efile Acknowledgement,omitempty" json:",omitempty"`
}

func (r AcknowledgementList) Validate() error {
	return utils.Validate(&r)
}

type SubmissionReceiptList struct {
	Cnt                  int                    `xml:"Cnt"`
	SubmissionReceiptGrp []SubmissionReceiptGrp `xml:"SubmissionReceiptGrp,omitempty" json:",omitempty"`
}

func (r SubmissionReceiptList) Validate() error {
	return utils.Validate(&r)
}

type StatusRecordList struct {
	Cnt             int               `xml:"Cnt"`
	StatusRecordGrp []StatusRecordGrp `xml:"StatusRecordGrp,omitempty" json:",omitempty"`
}

func (r StatusRecordList) Validate() error {
	return utils.Validate(&r)
}

type AckNotification struct {
	SubmissionId SubmissionIdType `xml:"SubmissionId"`
	Ts           TimestampType    `xml:"Ts"`
}

func (r AckNotification) Validate() error {
	return utils.Validate(&r)
}

type AckNotificationList struct {
	Cnt             int               `xml:"Cnt"`
	AckNotification []AckNotification `xml:"AckNotification,omitempty" json:",omitempty"`
}

func (r AckNotificationList) Validate() error {
	return utils.Validate(&r)
}
