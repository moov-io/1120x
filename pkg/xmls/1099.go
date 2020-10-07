// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package xmls

import (
	"bytes"
	"encoding/xml"
	"github.com/moov-io/1120x/pkg/utils"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type IRS990ScheduleOData struct {
	Text                          string `xml:",chardata"`
	DocumentId                    string `xml:"documentId,attr"`
	SupplementalInformationDetail []struct {
		Text                     string `xml:",chardata"`
		FormAndLineReferenceDesc string `xml:"FormAndLineReferenceDesc"`
		ExplanationTxt           string `xml:"ExplanationTxt"`
	} `xml:"SupplementalInformationDetail"`
}

type IRS990ScheduleMData struct {
	Text                        string `xml:",chardata"`
	DocumentId                  string `xml:"documentId,attr"`
	SecuritiesPubliclyTradedGrp struct {
		Text                           string `xml:",chardata"`
		NonCashCheckboxInd             string `xml:"NonCashCheckboxInd"`
		ContributionCnt                string `xml:"ContributionCnt"`
		NoncashContributionsRptF990Amt string `xml:"NoncashContributionsRptF990Amt"`
		MethodOfDeterminingRevenuesTxt string `xml:"MethodOfDeterminingRevenuesTxt"`
	} `xml:"SecuritiesPubliclyTradedGrp"`
	OtherNonCashContriTableGrp []struct {
		Text                           string `xml:",chardata"`
		NonCashCheckboxInd             string `xml:"NonCashCheckboxInd"`
		Desc                           string `xml:"Desc"`
		ContributionCnt                string `xml:"ContributionCnt"`
		NoncashContributionsRptF990Amt string `xml:"NoncashContributionsRptF990Amt"`
		MethodOfDeterminingRevenuesTxt string `xml:"MethodOfDeterminingRevenuesTxt"`
	} `xml:"OtherNonCashContriTableGrp"`
	AnyPropertyThatMustBeHeldInd   string `xml:"AnyPropertyThatMustBeHeldInd"`
	ReviewProcessUnusualNCGiftsInd string `xml:"ReviewProcessUnusualNCGiftsInd"`
	ThirdPartiesUsedInd            string `xml:"ThirdPartiesUsedInd"`
}

type IRS990ScheduleDData struct {
	Text                     string `xml:",chardata"`
	DocumentId               string `xml:"documentId,attr"`
	LeaseholdImprovementsGrp struct {
		Text                     string `xml:",chardata"`
		OtherCostOrOtherBasisAmt string `xml:"OtherCostOrOtherBasisAmt"`
		DepreciationAmt          string `xml:"DepreciationAmt"`
		BookValueAmt             string `xml:"BookValueAmt"`
	} `xml:"LeaseholdImprovementsGrp"`
	EquipmentGrp struct {
		Text                     string `xml:",chardata"`
		OtherCostOrOtherBasisAmt string `xml:"OtherCostOrOtherBasisAmt"`
		DepreciationAmt          string `xml:"DepreciationAmt"`
		BookValueAmt             string `xml:"BookValueAmt"`
	} `xml:"EquipmentGrp"`
	OtherLandBuildingsGrp struct {
		Text                     string `xml:",chardata"`
		OtherCostOrOtherBasisAmt string `xml:"OtherCostOrOtherBasisAmt"`
		DepreciationAmt          string `xml:"DepreciationAmt"`
		BookValueAmt             string `xml:"BookValueAmt"`
	} `xml:"OtherLandBuildingsGrp"`
	TotalBookValueLandBuildingsAmt string `xml:"TotalBookValueLandBuildingsAmt"`
}

type IRS990ScheduleBData struct {
	Text                      string `xml:",chardata"`
	DocumentId                string `xml:"documentId,attr"`
	ContributorInformationGrp struct {
		Text                    string `xml:",chardata"`
		ContributorNum          string `xml:"ContributorNum"`
		ContributorBusinessName struct {
			Text              string `xml:",chardata"`
			BusinessNameLine1 string `xml:"BusinessNameLine1"`
		} `xml:"ContributorBusinessName"`
		ContributorUSAddress struct {
			Text         string `xml:",chardata"`
			AddressLine1 string `xml:"AddressLine1"`
			AddressLine2 string `xml:"AddressLine2"`
			City         string `xml:"City"`
			State        string `xml:"State"`
			ZIPCode      string `xml:"ZIPCode"`
		} `xml:"ContributorUSAddress"`
		TotalContributionsAmt string `xml:"TotalContributionsAmt"`
	} `xml:"ContributorInformationGrp"`
}

type IRS990ScheduleAData struct {
	Text                        string `xml:",chardata"`
	DocumentId                  string `xml:"documentId,attr"`
	PublicOrganization170Ind    string `xml:"PublicOrganization170Ind"`
	GiftsGrantsContriRcvd170Grp struct {
		Text                         string `xml:",chardata"`
		CurrentTaxYearMinus4YearsAmt string `xml:"CurrentTaxYearMinus4YearsAmt"`
		CurrentTaxYearMinus3YearsAmt string `xml:"CurrentTaxYearMinus3YearsAmt"`
		CurrentTaxYearMinus2YearsAmt string `xml:"CurrentTaxYearMinus2YearsAmt"`
		CurrentTaxYearMinus1YearAmt  string `xml:"CurrentTaxYearMinus1YearAmt"`
		CurrentTaxYearAmt            string `xml:"CurrentTaxYearAmt"`
		TotalAmt                     string `xml:"TotalAmt"`
	} `xml:"GiftsGrantsContriRcvd170Grp"`
	TotalCalendarYear170Grp struct {
		Text                         string `xml:",chardata"`
		CurrentTaxYearMinus4YearsAmt string `xml:"CurrentTaxYearMinus4YearsAmt"`
		CurrentTaxYearMinus3YearsAmt string `xml:"CurrentTaxYearMinus3YearsAmt"`
		CurrentTaxYearMinus2YearsAmt string `xml:"CurrentTaxYearMinus2YearsAmt"`
		CurrentTaxYearMinus1YearAmt  string `xml:"CurrentTaxYearMinus1YearAmt"`
		CurrentTaxYearAmt            string `xml:"CurrentTaxYearAmt"`
		TotalAmt                     string `xml:"TotalAmt"`
	} `xml:"TotalCalendarYear170Grp"`
	SubstantialContributorsTotAmt string `xml:"SubstantialContributorsTotAmt"`
	PublicSupportTotal170Amt      string `xml:"PublicSupportTotal170Amt"`
	OtherIncome170Grp             struct {
		Text                         string `xml:",chardata"`
		CurrentTaxYearMinus4YearsAmt string `xml:"CurrentTaxYearMinus4YearsAmt"`
		CurrentTaxYearMinus3YearsAmt string `xml:"CurrentTaxYearMinus3YearsAmt"`
		CurrentTaxYearMinus2YearsAmt string `xml:"CurrentTaxYearMinus2YearsAmt"`
		CurrentTaxYearMinus1YearAmt  string `xml:"CurrentTaxYearMinus1YearAmt"`
		CurrentTaxYearAmt            string `xml:"CurrentTaxYearAmt"`
		TotalAmt                     string `xml:"TotalAmt"`
	} `xml:"OtherIncome170Grp"`
	TotalSupportAmt                string `xml:"TotalSupportAmt"`
	PublicSupportCY170Pct          string `xml:"PublicSupportCY170Pct"`
	PublicSupportPY170Pct          string `xml:"PublicSupportPY170Pct"`
	ThirtyThrPctSuprtTestsCY170Ind string `xml:"ThirtyThrPctSuprtTestsCY170Ind"`
}

type IRS990Data struct {
	Text                string `xml:",chardata"`
	DocumentId          string `xml:"documentId,attr"`
	ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	PrincipalOfficerNm  string `xml:"PrincipalOfficerNm"`
	USAddress           struct {
		Text                string `xml:",chardata"`
		AddressLine1Txt     string `xml:"AddressLine1Txt"`
		CityNm              string `xml:"CityNm"`
		StateAbbreviationCd string `xml:"StateAbbreviationCd"`
		ZIPCd               string `xml:"ZIPCd"`
	} `xml:"USAddress"`
	GrossReceiptsAmt               string `xml:"GrossReceiptsAmt"`
	GroupReturnForAffiliatesInd    string `xml:"GroupReturnForAffiliatesInd"`
	Organization501c3Ind           string `xml:"Organization501c3Ind"`
	WebsiteAddressTxt              string `xml:"WebsiteAddressTxt"`
	TypeOfOrganizationCorpInd      string `xml:"TypeOfOrganizationCorpInd"`
	FormationYr                    string `xml:"FormationYr"`
	LegalDomicileStateCd           string `xml:"LegalDomicileStateCd"`
	ActivityOrMissionDesc          string `xml:"ActivityOrMissionDesc"`
	VotingMembersGoverningBodyCnt  string `xml:"VotingMembersGoverningBodyCnt"`
	VotingMembersIndependentCnt    string `xml:"VotingMembersIndependentCnt"`
	TotalEmployeeCnt               string `xml:"TotalEmployeeCnt"`
	TotalVolunteersCnt             string `xml:"TotalVolunteersCnt"`
	TotalGrossUBIAmt               string `xml:"TotalGrossUBIAmt"`
	NetUnrelatedBusTxblIncmAmt     string `xml:"NetUnrelatedBusTxblIncmAmt"`
	PYContributionsGrantsAmt       string `xml:"PYContributionsGrantsAmt"`
	CYContributionsGrantsAmt       string `xml:"CYContributionsGrantsAmt"`
	PYProgramServiceRevenueAmt     string `xml:"PYProgramServiceRevenueAmt"`
	CYProgramServiceRevenueAmt     string `xml:"CYProgramServiceRevenueAmt"`
	PYInvestmentIncomeAmt          string `xml:"PYInvestmentIncomeAmt"`
	CYInvestmentIncomeAmt          string `xml:"CYInvestmentIncomeAmt"`
	PYOtherRevenueAmt              string `xml:"PYOtherRevenueAmt"`
	CYOtherRevenueAmt              string `xml:"CYOtherRevenueAmt"`
	PYTotalRevenueAmt              string `xml:"PYTotalRevenueAmt"`
	CYTotalRevenueAmt              string `xml:"CYTotalRevenueAmt"`
	PYGrantsAndSimilarPaidAmt      string `xml:"PYGrantsAndSimilarPaidAmt"`
	CYGrantsAndSimilarPaidAmt      string `xml:"CYGrantsAndSimilarPaidAmt"`
	PYBenefitsPaidToMembersAmt     string `xml:"PYBenefitsPaidToMembersAmt"`
	CYBenefitsPaidToMembersAmt     string `xml:"CYBenefitsPaidToMembersAmt"`
	PYSalariesCompEmpBnftPaidAmt   string `xml:"PYSalariesCompEmpBnftPaidAmt"`
	CYSalariesCompEmpBnftPaidAmt   string `xml:"CYSalariesCompEmpBnftPaidAmt"`
	PYTotalProfFndrsngExpnsAmt     string `xml:"PYTotalProfFndrsngExpnsAmt"`
	CYTotalProfFndrsngExpnsAmt     string `xml:"CYTotalProfFndrsngExpnsAmt"`
	CYTotalFundraisingExpenseAmt   string `xml:"CYTotalFundraisingExpenseAmt"`
	PYOtherExpensesAmt             string `xml:"PYOtherExpensesAmt"`
	CYOtherExpensesAmt             string `xml:"CYOtherExpensesAmt"`
	PYTotalExpensesAmt             string `xml:"PYTotalExpensesAmt"`
	CYTotalExpensesAmt             string `xml:"CYTotalExpensesAmt"`
	PYRevenuesLessExpensesAmt      string `xml:"PYRevenuesLessExpensesAmt"`
	CYRevenuesLessExpensesAmt      string `xml:"CYRevenuesLessExpensesAmt"`
	TotalAssetsBOYAmt              string `xml:"TotalAssetsBOYAmt"`
	TotalAssetsEOYAmt              string `xml:"TotalAssetsEOYAmt"`
	TotalLiabilitiesBOYAmt         string `xml:"TotalLiabilitiesBOYAmt"`
	TotalLiabilitiesEOYAmt         string `xml:"TotalLiabilitiesEOYAmt"`
	NetAssetsOrFundBalancesBOYAmt  string `xml:"NetAssetsOrFundBalancesBOYAmt"`
	NetAssetsOrFundBalancesEOYAmt  string `xml:"NetAssetsOrFundBalancesEOYAmt"`
	MissionDesc                    string `xml:"MissionDesc"`
	SignificantNewProgramSrvcInd   string `xml:"SignificantNewProgramSrvcInd"`
	SignificantChangeInd           string `xml:"SignificantChangeInd"`
	ExpenseAmt                     string `xml:"ExpenseAmt"`
	RevenueAmt                     string `xml:"RevenueAmt"`
	Desc                           string `xml:"Desc"`
	TotalProgramServiceExpensesAmt string `xml:"TotalProgramServiceExpensesAmt"`
	DescribedInSection501c3Ind     struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"DescribedInSection501c3Ind"`
	ScheduleBRequiredInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"ScheduleBRequiredInd"`
	PoliticalCampaignActyInd string `xml:"PoliticalCampaignActyInd"`
	LobbyingActivitiesInd    string `xml:"LobbyingActivitiesInd"`
	SubjectToProxyTaxInd     string `xml:"SubjectToProxyTaxInd"`
	DonorAdvisedFundInd      struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"DonorAdvisedFundInd"`
	ConservationEasementsInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"ConservationEasementsInd"`
	CollectionsOfArtInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"CollectionsOfArtInd"`
	CreditCounselingInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"CreditCounselingInd"`
	TempOrPermanentEndowmentsInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"TempOrPermanentEndowmentsInd"`
	ReportLandBuildingEquipmentInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"ReportLandBuildingEquipmentInd"`
	ReportInvestmentsOtherSecInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"ReportInvestmentsOtherSecInd"`
	ReportProgramRelatedInvstInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"ReportProgramRelatedInvstInd"`
	ReportOtherAssetsInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"ReportOtherAssetsInd"`
	ReportOtherLiabilitiesInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"ReportOtherLiabilitiesInd"`
	IncludeFIN48FootnoteInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"IncludeFIN48FootnoteInd"`
	IndependentAuditFinclStmtInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"IndependentAuditFinclStmtInd"`
	ConsolidatedAuditFinclStmtInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"ConsolidatedAuditFinclStmtInd"`
	SchoolOperatingInd             string `xml:"SchoolOperatingInd"`
	ForeignOfficeInd               string `xml:"ForeignOfficeInd"`
	ForeignActivitiesInd           string `xml:"ForeignActivitiesInd"`
	MoreThan5000KToOrgInd          string `xml:"MoreThan5000KToOrgInd"`
	MoreThan5000KToIndividualsInd  string `xml:"MoreThan5000KToIndividualsInd"`
	ProfessionalFundraisingInd     string `xml:"ProfessionalFundraisingInd"`
	FundraisingActivitiesInd       string `xml:"FundraisingActivitiesInd"`
	GamingActivitiesInd            string `xml:"GamingActivitiesInd"`
	OperateHospitalInd             string `xml:"OperateHospitalInd"`
	GrantsToOrganizationsInd       string `xml:"GrantsToOrganizationsInd"`
	GrantsToIndividualsInd         string `xml:"GrantsToIndividualsInd"`
	ScheduleJRequiredInd           string `xml:"ScheduleJRequiredInd"`
	TaxExemptBondsInd              string `xml:"TaxExemptBondsInd"`
	EngagedInExcessBenefitTransInd string `xml:"EngagedInExcessBenefitTransInd"`
	PYExcessBenefitTransInd        string `xml:"PYExcessBenefitTransInd"`
	LoanOutstandingInd             string `xml:"LoanOutstandingInd"`
	GrantToRelatedPersonInd        string `xml:"GrantToRelatedPersonInd"`
	BusinessRlnWithOrgMemInd       string `xml:"BusinessRlnWithOrgMemInd"`
	BusinessRlnWithFamMemInd       string `xml:"BusinessRlnWithFamMemInd"`
	BusinessRlnWithOfficerEntInd   string `xml:"BusinessRlnWithOfficerEntInd"`
	DeductibleNonCashContriInd     struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"DeductibleNonCashContriInd"`
	DeductibleArtContributionInd struct {
		Text                string `xml:",chardata"`
		ReferenceDocumentId string `xml:"referenceDocumentId,attr"`
	} `xml:"DeductibleArtContributionInd"`
	TerminateOperationsInd         string `xml:"TerminateOperationsInd"`
	PartialLiquidationInd          string `xml:"PartialLiquidationInd"`
	DisregardedEntityInd           string `xml:"DisregardedEntityInd"`
	RelatedEntityInd               string `xml:"RelatedEntityInd"`
	RelatedOrganizationCtrlEntInd  string `xml:"RelatedOrganizationCtrlEntInd"`
	TrnsfrExmptNonChrtblRltdOrgInd string `xml:"TrnsfrExmptNonChrtblRltdOrgInd"`
	ActivitiesConductedPrtshpInd   string `xml:"ActivitiesConductedPrtshpInd"`
	ScheduleORequiredInd           string `xml:"ScheduleORequiredInd"`
	IRPDocumentCnt                 string `xml:"IRPDocumentCnt"`
	IRPDocumentW2GCnt              string `xml:"IRPDocumentW2GCnt"`
	BackupWthldComplianceInd       string `xml:"BackupWthldComplianceInd"`
	EmployeeCnt                    string `xml:"EmployeeCnt"`
	EmploymentTaxReturnsFiledInd   string `xml:"EmploymentTaxReturnsFiledInd"`
	UnrelatedBusIncmOverLimitInd   string `xml:"UnrelatedBusIncmOverLimitInd"`
	Form990TFiledInd               string `xml:"Form990TFiledInd"`
	ForeignFinancialAccountInd     string `xml:"ForeignFinancialAccountInd"`
	ProhibitedTaxShelterTransInd   string `xml:"ProhibitedTaxShelterTransInd"`
	TaxablePartyNotificationInd    string `xml:"TaxablePartyNotificationInd"`
	NondeductibleContributionsInd  string `xml:"NondeductibleContributionsInd"`
	QuidProQuoContributionsInd     string `xml:"QuidProQuoContributionsInd"`
	QuidProQuoContriDisclInd       string `xml:"QuidProQuoContriDisclInd"`
	Form8282PropertyDisposedOfInd  string `xml:"Form8282PropertyDisposedOfInd"`
	RcvFndsToPayPrsnlBnftCntrctInd string `xml:"RcvFndsToPayPrsnlBnftCntrctInd"`
	PayPremiumsPrsnlBnftCntrctInd  string `xml:"PayPremiumsPrsnlBnftCntrctInd"`
	IndoorTanningServicesInd       string `xml:"IndoorTanningServicesInd"`
	InfoInScheduleOPartVIInd       string `xml:"InfoInScheduleOPartVIInd"`
	GoverningBodyVotingMembersCnt  string `xml:"GoverningBodyVotingMembersCnt"`
	IndependentVotingMemberCnt     string `xml:"IndependentVotingMemberCnt"`
	FamilyOrBusinessRlnInd         string `xml:"FamilyOrBusinessRlnInd"`
	DelegationOfMgmtDutiesInd      string `xml:"DelegationOfMgmtDutiesInd"`
	ChangeToOrgDocumentsInd        string `xml:"ChangeToOrgDocumentsInd"`
	MaterialDiversionOrMisuseInd   string `xml:"MaterialDiversionOrMisuseInd"`
	MembersOrStockholdersInd       string `xml:"MembersOrStockholdersInd"`
	ElectionOfBoardMembersInd      string `xml:"ElectionOfBoardMembersInd"`
	DecisionsSubjectToApprovaInd   string `xml:"DecisionsSubjectToApprovaInd"`
	MinutesOfGoverningBodyInd      string `xml:"MinutesOfGoverningBodyInd"`
	MinutesOfCommitteesInd         string `xml:"MinutesOfCommitteesInd"`
	OfficerMailingAddressInd       string `xml:"OfficerMailingAddressInd"`
	LocalChaptersInd               string `xml:"LocalChaptersInd"`
	Form990ProvidedToGvrnBodyInd   string `xml:"Form990ProvidedToGvrnBodyInd"`
	ConflictOfInterestPolicyInd    string `xml:"ConflictOfInterestPolicyInd"`
	AnnualDisclosureCoveredPrsnInd string `xml:"AnnualDisclosureCoveredPrsnInd"`
	RegularMonitoringEnfrcInd      string `xml:"RegularMonitoringEnfrcInd"`
	WhistleblowerPolicyInd         string `xml:"WhistleblowerPolicyInd"`
	DocumentRetentionPolicyInd     string `xml:"DocumentRetentionPolicyInd"`
	CompensationProcessCEOInd      string `xml:"CompensationProcessCEOInd"`
	CompensationProcessOtherInd    string `xml:"CompensationProcessOtherInd"`
	InvestmentInJointVentureInd    string `xml:"InvestmentInJointVentureInd"`
	StatesWhereCopyOfReturnIsFldCd string `xml:"StatesWhereCopyOfReturnIsFldCd"`
	OtherWebsiteInd                string `xml:"OtherWebsiteInd"`
	UponRequestInd                 string `xml:"UponRequestInd"`
	BooksInCareOfDetail            struct {
		Text         string `xml:",chardata"`
		BusinessName struct {
			Text                 string `xml:",chardata"`
			BusinessNameLine1Txt string `xml:"BusinessNameLine1Txt"`
		} `xml:"BusinessName"`
		PhoneNum  string `xml:"PhoneNum"`
		USAddress struct {
			Text                string `xml:",chardata"`
			AddressLine1Txt     string `xml:"AddressLine1Txt"`
			CityNm              string `xml:"CityNm"`
			StateAbbreviationCd string `xml:"StateAbbreviationCd"`
			ZIPCd               string `xml:"ZIPCd"`
		} `xml:"USAddress"`
	} `xml:"BooksInCareOfDetail"`
	Form990PartVIISectionAGrp []struct {
		Text                           string `xml:",chardata"`
		PersonNm                       string `xml:"PersonNm"`
		TitleTxt                       string `xml:"TitleTxt"`
		AverageHoursPerWeekRt          string `xml:"AverageHoursPerWeekRt"`
		IndividualTrusteeOrDirectorInd string `xml:"IndividualTrusteeOrDirectorInd"`
		ReportableCompFromOrgAmt       string `xml:"ReportableCompFromOrgAmt"`
		ReportableCompFromRltdOrgAmt   string `xml:"ReportableCompFromRltdOrgAmt"`
		OtherCompensationAmt           string `xml:"OtherCompensationAmt"`
		OfficerInd                     string `xml:"OfficerInd"`
	} `xml:"Form990PartVIISectionAGrp"`
	TotalReportableCompFromOrgAmt string `xml:"TotalReportableCompFromOrgAmt"`
	TotReportableCompRltdOrgAmt   string `xml:"TotReportableCompRltdOrgAmt"`
	TotalOtherCompensationAmt     string `xml:"TotalOtherCompensationAmt"`
	IndivRcvdGreaterThan100KCnt   string `xml:"IndivRcvdGreaterThan100KCnt"`
	FormerOfcrEmployeesListedInd  string `xml:"FormerOfcrEmployeesListedInd"`
	TotalCompGreaterThan150KInd   string `xml:"TotalCompGreaterThan150KInd"`
	CompensationFromOtherSrcsInd  string `xml:"CompensationFromOtherSrcsInd"`
	CntrctRcvdGreaterThan100KCnt  string `xml:"CntrctRcvdGreaterThan100KCnt"`
	MembershipDuesAmt             string `xml:"MembershipDuesAmt"`
	AllOtherContributionsAmt      string `xml:"AllOtherContributionsAmt"`
	NoncashContributionsAmt       string `xml:"NoncashContributionsAmt"`
	TotalContributionsAmt         string `xml:"TotalContributionsAmt"`
	ProgramServiceRevenueGrp      []struct {
		Text                         string `xml:",chardata"`
		Desc                         string `xml:"Desc"`
		BusinessCd                   string `xml:"BusinessCd"`
		TotalRevenueColumnAmt        string `xml:"TotalRevenueColumnAmt"`
		RelatedOrExemptFuncIncomeAmt string `xml:"RelatedOrExemptFuncIncomeAmt"`
		UnrelatedBusinessRevenueAmt  string `xml:"UnrelatedBusinessRevenueAmt"`
	} `xml:"ProgramServiceRevenueGrp"`
	TotalProgramServiceRevenueAmt string `xml:"TotalProgramServiceRevenueAmt"`
	GrossAmountSalesAssetsGrp     struct {
		Text          string `xml:",chardata"`
		SecuritiesAmt string `xml:"SecuritiesAmt"`
	} `xml:"GrossAmountSalesAssetsGrp"`
	LessCostOthBasisSalesExpnssGrp struct {
		Text          string `xml:",chardata"`
		SecuritiesAmt string `xml:"SecuritiesAmt"`
		OtherAmt      string `xml:"OtherAmt"`
	} `xml:"LessCostOthBasisSalesExpnssGrp"`
	GainOrLossGrp struct {
		Text          string `xml:",chardata"`
		SecuritiesAmt string `xml:"SecuritiesAmt"`
		OtherAmt      string `xml:"OtherAmt"`
	} `xml:"GainOrLossGrp"`
	NetGainOrLossInvestmentsGrp struct {
		Text                         string `xml:",chardata"`
		TotalRevenueColumnAmt        string `xml:"TotalRevenueColumnAmt"`
		RelatedOrExemptFuncIncomeAmt string `xml:"RelatedOrExemptFuncIncomeAmt"`
	} `xml:"NetGainOrLossInvestmentsGrp"`
	TotalRevenueGrp struct {
		Text                         string `xml:",chardata"`
		TotalRevenueColumnAmt        string `xml:"TotalRevenueColumnAmt"`
		RelatedOrExemptFuncIncomeAmt string `xml:"RelatedOrExemptFuncIncomeAmt"`
		UnrelatedBusinessRevenueAmt  string `xml:"UnrelatedBusinessRevenueAmt"`
		ExclusionAmt                 string `xml:"ExclusionAmt"`
	} `xml:"TotalRevenueGrp"`
	CompCurrentOfcrDirectorsGrp struct {
		Text               string `xml:",chardata"`
		TotalAmt           string `xml:"TotalAmt"`
		ProgramServicesAmt string `xml:"ProgramServicesAmt"`
	} `xml:"CompCurrentOfcrDirectorsGrp"`
	OtherSalariesAndWagesGrp struct {
		Text                    string `xml:",chardata"`
		TotalAmt                string `xml:"TotalAmt"`
		ProgramServicesAmt      string `xml:"ProgramServicesAmt"`
		ManagementAndGeneralAmt string `xml:"ManagementAndGeneralAmt"`
		FundraisingAmt          string `xml:"FundraisingAmt"`
	} `xml:"OtherSalariesAndWagesGrp"`
	PensionPlanContributionsGrp struct {
		Text                    string `xml:",chardata"`
		TotalAmt                string `xml:"TotalAmt"`
		ProgramServicesAmt      string `xml:"ProgramServicesAmt"`
		ManagementAndGeneralAmt string `xml:"ManagementAndGeneralAmt"`
		FundraisingAmt          string `xml:"FundraisingAmt"`
	} `xml:"PensionPlanContributionsGrp"`
	OtherEmployeeBenefitsGrp struct {
		Text                    string `xml:",chardata"`
		TotalAmt                string `xml:"TotalAmt"`
		ProgramServicesAmt      string `xml:"ProgramServicesAmt"`
		ManagementAndGeneralAmt string `xml:"ManagementAndGeneralAmt"`
		FundraisingAmt          string `xml:"FundraisingAmt"`
	} `xml:"OtherEmployeeBenefitsGrp"`
	PayrollTaxesGrp struct {
		Text                    string `xml:",chardata"`
		TotalAmt                string `xml:"TotalAmt"`
		ProgramServicesAmt      string `xml:"ProgramServicesAmt"`
		ManagementAndGeneralAmt string `xml:"ManagementAndGeneralAmt"`
		FundraisingAmt          string `xml:"FundraisingAmt"`
	} `xml:"PayrollTaxesGrp"`
	FeesForServicesLegalGrp struct {
		Text               string `xml:",chardata"`
		TotalAmt           string `xml:"TotalAmt"`
		ProgramServicesAmt string `xml:"ProgramServicesAmt"`
	} `xml:"FeesForServicesLegalGrp"`
	FeesForServicesAccountingGrp struct {
		Text               string `xml:",chardata"`
		TotalAmt           string `xml:"TotalAmt"`
		ProgramServicesAmt string `xml:"ProgramServicesAmt"`
	} `xml:"FeesForServicesAccountingGrp"`
	FeesForServicesOtherGrp struct {
		Text               string `xml:",chardata"`
		TotalAmt           string `xml:"TotalAmt"`
		ProgramServicesAmt string `xml:"ProgramServicesAmt"`
	} `xml:"FeesForServicesOtherGrp"`
	AdvertisingGrp struct {
		Text                    string `xml:",chardata"`
		TotalAmt                string `xml:"TotalAmt"`
		ProgramServicesAmt      string `xml:"ProgramServicesAmt"`
		ManagementAndGeneralAmt string `xml:"ManagementAndGeneralAmt"`
		FundraisingAmt          string `xml:"FundraisingAmt"`
	} `xml:"AdvertisingGrp"`
	OfficeExpensesGrp struct {
		Text               string `xml:",chardata"`
		TotalAmt           string `xml:"TotalAmt"`
		ProgramServicesAmt string `xml:"ProgramServicesAmt"`
	} `xml:"OfficeExpensesGrp"`
	InformationTechnologyGrp struct {
		Text               string `xml:",chardata"`
		TotalAmt           string `xml:"TotalAmt"`
		ProgramServicesAmt string `xml:"ProgramServicesAmt"`
	} `xml:"InformationTechnologyGrp"`
	OccupancyGrp struct {
		Text                    string `xml:",chardata"`
		TotalAmt                string `xml:"TotalAmt"`
		ProgramServicesAmt      string `xml:"ProgramServicesAmt"`
		ManagementAndGeneralAmt string `xml:"ManagementAndGeneralAmt"`
		FundraisingAmt          string `xml:"FundraisingAmt"`
	} `xml:"OccupancyGrp"`
	TravelGrp struct {
		Text               string `xml:",chardata"`
		TotalAmt           string `xml:"TotalAmt"`
		ProgramServicesAmt string `xml:"ProgramServicesAmt"`
	} `xml:"TravelGrp"`
	ConferencesMeetingsGrp struct {
		Text               string `xml:",chardata"`
		TotalAmt           string `xml:"TotalAmt"`
		ProgramServicesAmt string `xml:"ProgramServicesAmt"`
	} `xml:"ConferencesMeetingsGrp"`
	DepreciationDepletionGrp struct {
		Text               string `xml:",chardata"`
		TotalAmt           string `xml:"TotalAmt"`
		ProgramServicesAmt string `xml:"ProgramServicesAmt"`
	} `xml:"DepreciationDepletionGrp"`
	InsuranceGrp struct {
		Text                    string `xml:",chardata"`
		TotalAmt                string `xml:"TotalAmt"`
		ProgramServicesAmt      string `xml:"ProgramServicesAmt"`
		ManagementAndGeneralAmt string `xml:"ManagementAndGeneralAmt"`
		FundraisingAmt          string `xml:"FundraisingAmt"`
	} `xml:"InsuranceGrp"`
	OtherExpensesGrp []struct {
		Text                    string `xml:",chardata"`
		Desc                    string `xml:"Desc"`
		TotalAmt                string `xml:"TotalAmt"`
		ProgramServicesAmt      string `xml:"ProgramServicesAmt"`
		ManagementAndGeneralAmt string `xml:"ManagementAndGeneralAmt"`
		FundraisingAmt          string `xml:"FundraisingAmt"`
	} `xml:"OtherExpensesGrp"`
	AllOtherExpensesGrp struct {
		Text                    string `xml:",chardata"`
		TotalAmt                string `xml:"TotalAmt"`
		ProgramServicesAmt      string `xml:"ProgramServicesAmt"`
		ManagementAndGeneralAmt string `xml:"ManagementAndGeneralAmt"`
		FundraisingAmt          string `xml:"FundraisingAmt"`
	} `xml:"AllOtherExpensesGrp"`
	TotalFunctionalExpensesGrp struct {
		Text                    string `xml:",chardata"`
		TotalAmt                string `xml:"TotalAmt"`
		ProgramServicesAmt      string `xml:"ProgramServicesAmt"`
		ManagementAndGeneralAmt string `xml:"ManagementAndGeneralAmt"`
		FundraisingAmt          string `xml:"FundraisingAmt"`
	} `xml:"TotalFunctionalExpensesGrp"`
	CashNonInterestBearingGrp struct {
		Text   string `xml:",chardata"`
		BOYAmt string `xml:"BOYAmt"`
		EOYAmt string `xml:"EOYAmt"`
	} `xml:"CashNonInterestBearingGrp"`
	PrepaidExpensesDefrdChargesGrp struct {
		Text   string `xml:",chardata"`
		BOYAmt string `xml:"BOYAmt"`
		EOYAmt string `xml:"EOYAmt"`
	} `xml:"PrepaidExpensesDefrdChargesGrp"`
	LandBldgEquipCostOrOtherBssAmt string `xml:"LandBldgEquipCostOrOtherBssAmt"`
	LandBldgEquipAccumDeprecAmt    string `xml:"LandBldgEquipAccumDeprecAmt"`
	LandBldgEquipBasisNetGrp       struct {
		Text   string `xml:",chardata"`
		BOYAmt string `xml:"BOYAmt"`
		EOYAmt string `xml:"EOYAmt"`
	} `xml:"LandBldgEquipBasisNetGrp"`
	OtherAssetsTotalGrp struct {
		Text   string `xml:",chardata"`
		BOYAmt string `xml:"BOYAmt"`
		EOYAmt string `xml:"EOYAmt"`
	} `xml:"OtherAssetsTotalGrp"`
	TotalAssetsGrp struct {
		Text   string `xml:",chardata"`
		BOYAmt string `xml:"BOYAmt"`
		EOYAmt string `xml:"EOYAmt"`
	} `xml:"TotalAssetsGrp"`
	TotalLiabilitiesGrp struct {
		Text   string `xml:",chardata"`
		BOYAmt string `xml:"BOYAmt"`
		EOYAmt string `xml:"EOYAmt"`
	} `xml:"TotalLiabilitiesGrp"`
	OrganizationFollowsSFAS117Ind string `xml:"OrganizationFollowsSFAS117Ind"`
	UnrestrictedNetAssetsGrp      struct {
		Text   string `xml:",chardata"`
		BOYAmt string `xml:"BOYAmt"`
		EOYAmt string `xml:"EOYAmt"`
	} `xml:"UnrestrictedNetAssetsGrp"`
	TotalNetAssetsFundBalanceGrp struct {
		Text   string `xml:",chardata"`
		BOYAmt string `xml:"BOYAmt"`
		EOYAmt string `xml:"EOYAmt"`
	} `xml:"TotalNetAssetsFundBalanceGrp"`
	TotLiabNetAssetsFundBalanceGrp struct {
		Text   string `xml:",chardata"`
		BOYAmt string `xml:"BOYAmt"`
		EOYAmt string `xml:"EOYAmt"`
	} `xml:"TotLiabNetAssetsFundBalanceGrp"`
	ReconcilationRevenueExpnssAmt string `xml:"ReconcilationRevenueExpnssAmt"`
	OtherChangesInNetAssetsAmt    string `xml:"OtherChangesInNetAssetsAmt"`
	MethodOfAccountingCashInd     string `xml:"MethodOfAccountingCashInd"`
	AccountantCompileOrReviewInd  string `xml:"AccountantCompileOrReviewInd"`
	FSAuditedInd                  string `xml:"FSAuditedInd"`
	FederalGrantAuditRequiredInd  string `xml:"FederalGrantAuditRequiredInd"`
}

type IRS990ReturnHeader struct {
	Text                string `xml:",chardata"`
	BinaryAttachmentCnt string `xml:"binaryAttachmentCnt,attr"`
	ReturnTs            string `xml:"ReturnTs"`
	TaxPeriodEndDt      string `xml:"TaxPeriodEndDt"`
	PreparerFirmGrp     struct {
		Text             string `xml:",chardata"`
		PreparerFirmEIN  string `xml:"PreparerFirmEIN"`
		PreparerFirmName struct {
			Text                 string `xml:",chardata"`
			BusinessNameLine1Txt string `xml:"BusinessNameLine1Txt"`
		} `xml:"PreparerFirmName"`
		PreparerUSAddress struct {
			Text                string `xml:",chardata"`
			AddressLine1Txt     string `xml:"AddressLine1Txt"`
			CityNm              string `xml:"CityNm"`
			StateAbbreviationCd string `xml:"StateAbbreviationCd"`
			ZIPCd               string `xml:"ZIPCd"`
		} `xml:"PreparerUSAddress"`
	} `xml:"PreparerFirmGrp"`
	ReturnTypeCd     string `xml:"ReturnTypeCd"`
	TaxPeriodBeginDt string `xml:"TaxPeriodBeginDt"`
	Filer            struct {
		Text         string `xml:",chardata"`
		EIN          string `xml:"EIN"`
		BusinessName struct {
			Text                 string `xml:",chardata"`
			BusinessNameLine1Txt string `xml:"BusinessNameLine1Txt"`
		} `xml:"BusinessName"`
		BusinessNameControlTxt string `xml:"BusinessNameControlTxt"`
		PhoneNum               string `xml:"PhoneNum"`
		USAddress              struct {
			Text                string `xml:",chardata"`
			AddressLine1Txt     string `xml:"AddressLine1Txt"`
			CityNm              string `xml:"CityNm"`
			StateAbbreviationCd string `xml:"StateAbbreviationCd"`
			ZIPCd               string `xml:"ZIPCd"`
		} `xml:"USAddress"`
	} `xml:"Filer"`
	BusinessOfficerGrp struct {
		Text                       string `xml:",chardata"`
		PersonNm                   string `xml:"PersonNm"`
		PersonTitleTxt             string `xml:"PersonTitleTxt"`
		PhoneNum                   string `xml:"PhoneNum"`
		SignatureDt                string `xml:"SignatureDt"`
		DiscussWithPaidPreparerInd string `xml:"DiscussWithPaidPreparerInd"`
	} `xml:"BusinessOfficerGrp"`
	PreparerPersonGrp struct {
		Text             string `xml:",chardata"`
		PreparerPersonNm string `xml:"PreparerPersonNm"`
		PTIN             string `xml:"PTIN"`
		PhoneNum         string `xml:"PhoneNum"`
	} `xml:"PreparerPersonGrp"`
	TaxYr   string `xml:"TaxYr"`
	BuildTS string `xml:"BuildTS"`
}

type IRS990ReturnData struct {
	Text            string               `xml:",chardata"`
	DocumentCnt     int                  `xml:"documentCnt,attr"`
	IRS990          *IRS990Data          `xml:"IRS990,omitempty"`
	IRS990ScheduleA *IRS990ScheduleAData `xml:"IRS990ScheduleA,omitempty"`
	IRS990ScheduleB *IRS990ScheduleBData `xml:"IRS990ScheduleB,omitempty"`
	IRS990ScheduleD *IRS990ScheduleDData `xml:"IRS990ScheduleD,omitempty"`
	IRS990ScheduleM *IRS990ScheduleMData `xml:"IRS990ScheduleM,omitempty"`
	IRS990ScheduleO *IRS990ScheduleOData `xml:"IRS990ScheduleO,omitempty"`
}

type Return1099 struct {
	XMLName        xml.Name            `xml:"Return"`
	Text           string              `xml:",chardata"`
	Xmlns          string              `xml:"xmlns,attr"`
	Xsi            string              `xml:"xsi,attr"`
	Version        string              `xml:"returnVersion,attr"`
	SchemaLocation string              `xml:"schemaLocation,attr"`
	ReturnHeader   *IRS990ReturnHeader `xml:"ReturnHeader"`
	ReturnData     *IRS990ReturnData   `xml:"ReturnData"`
}

// Parse parses the “Return1099” record from raw xml
func (r *Return1099) Parse(buf []byte) error {
	if err := xml.Unmarshal(buf, r); err != nil {
		return err
	}
	return nil
}

// Split xml files with a document
func (r *Return1099) InspectData() *ReturnInspectInfo {
	if r.ReturnHeader == nil || r.ReturnData == nil {
		return nil
	}

	var returnData []ReturnInspectData
	if r.ReturnData.IRS990ScheduleA != nil {
		data := IRS990ReturnData{DocumentCnt: 1}
		data.IRS990ScheduleA = r.ReturnData.IRS990ScheduleA
		returnData = append(returnData, ReturnInspectData{Data: data, DataType: "990ScheduleA"})
	}
	if r.ReturnData.IRS990ScheduleB != nil {
		data := IRS990ReturnData{DocumentCnt: 1}
		data.IRS990ScheduleB = r.ReturnData.IRS990ScheduleB
		returnData = append(returnData, ReturnInspectData{Data: data, DataType: "990ScheduleB"})
	}
	if r.ReturnData.IRS990ScheduleD != nil {
		data := IRS990ReturnData{DocumentCnt: 1}
		data.IRS990ScheduleD = r.ReturnData.IRS990ScheduleD
		returnData = append(returnData, ReturnInspectData{Data: data, DataType: "990ScheduleD"})
	}
	if r.ReturnData.IRS990ScheduleM != nil {
		data := IRS990ReturnData{DocumentCnt: 1}
		data.IRS990ScheduleM = r.ReturnData.IRS990ScheduleM
		returnData = append(returnData, ReturnInspectData{Data: data, DataType: "990ScheduleM"})
	}
	if r.ReturnData.IRS990ScheduleO != nil {
		data := IRS990ReturnData{DocumentCnt: 1}
		data.IRS990ScheduleO = r.ReturnData.IRS990ScheduleO
		returnData = append(returnData, ReturnInspectData{Data: data, DataType: "990ScheduleO"})
	}

	return &ReturnInspectInfo{Header: r.ReturnHeader, Data: returnData}
}

// ReturnYear returns year of return year
func (r *Return1099) ReturnYear() int {
	splits := strings.Split(r.Version, "v")
	if len(splits[0]) == 0 {
		return 0
	}
	year, err := strconv.Atoi(splits[0])
	if err != nil {
		return 0
	}
	return year
}

// ReturnYear returns year of return version
func (r *Return1099) ReturnVersion() string {
	return r.Version
}

// ReturnType returns type of return type
func (r *Return1099) ReturnType() string {
	return utils.IRS990ReturnTypeCode
}

// Converting the struct to String format.
func (r *Return1099) String() string {
	buf, err := xml.Marshal(r)
	if err != nil {
		return ""
	}
	buf, err = formatXML(buf)
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	return re.ReplaceAllString(string(buf), "")
}

func formatXML(data []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return b.Bytes(), nil
		}
		if err != nil {
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
}
