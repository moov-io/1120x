// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import "github.com/moov-io/1120x/pkg/utils"

type IRS990 struct {
	SpecialConditionDesc           []string                            `xml:"SpecialConditionDesc,omitempty" json:",omitempty"`
	AddressChangeInd               CheckboxType                        `xml:"AddressChangeInd,omitempty" json:",omitempty"`
	InitialReturnInd               CheckboxType                        `xml:"InitialReturnInd,omitempty" json:",omitempty"`
	FinalReturnInd                 CheckboxType                        `xml:"FinalReturnInd,omitempty" json:",omitempty"`
	AmendedReturnInd               CheckboxType                        `xml:"AmendedReturnInd,omitempty" json:",omitempty"`
	DoingBusinessAsName            *BusinessNameType                   `xml:"DoingBusinessAsName,omitempty" json:",omitempty"`
	PrincipalOfficerNm             PersonNameType                      `xml:"PrincipalOfficerNm"`
	PrincipalOfcrBusinessName      BusinessNameType                    `xml:"PrincipalOfcrBusinessName"`
	USAddress                      USAddressType                       `xml:"USAddress"`
	ForeignAddress                 ForeignAddressType                  `xml:"ForeignAddress"`
	GrossReceiptsAmt               int                                 `xml:"GrossReceiptsAmt"`
	GroupReturnForAffiliatesInd    bool                                `xml:"GroupReturnForAffiliatesInd"`
	AllAffiliatesIncludedInd       *AllAffiliatesIncludedInd           `xml:"AllAffiliatesIncludedInd,omitempty" json:",omitempty"`
	GroupExemptionNum              GroupExemptionNum                   `xml:"GroupExemptionNum,omitempty" json:",omitempty"`
	Organization501c3Ind           CheckboxType                        `xml:"Organization501c3Ind"`
	Organization501cInd            Organization501cInd                 `xml:"Organization501cInd"`
	Organization4947a1NotPFInd     CheckboxType                        `xml:"Organization4947a1NotPFInd"`
	Organization527Ind             CheckboxType                        `xml:"Organization527Ind"`
	WebsiteAddressTxt              string                              `xml:"WebsiteAddressTxt,omitempty" json:",omitempty"`
	TypeOfOrganizationCorpInd      CheckboxType                        `xml:"TypeOfOrganizationCorpInd,omitempty" json:",omitempty"`
	TypeOfOrganizationTrustInd     CheckboxType                        `xml:"TypeOfOrganizationTrustInd,omitempty" json:",omitempty"`
	TypeOfOrganizationAssocInd     CheckboxType                        `xml:"TypeOfOrganizationAssocInd,omitempty" json:",omitempty"`
	TypeOfOrganizationOtherInd     CheckboxType                        `xml:"TypeOfOrganizationOtherInd,omitempty" json:",omitempty"`
	OtherOrganizationDsc           string                              `xml:"OtherOrganizationDsc,omitempty" json:",omitempty"`
	FormationYr                    YearType                            `xml:"FormationYr,omitempty" json:",omitempty"`
	LegalDomicileStateCd           StateType                           `xml:"LegalDomicileStateCd"`
	LegalDomicileCountryCd         CountryType                         `xml:"LegalDomicileCountryCd"`
	ActivityOrMissionDesc          string                              `xml:"ActivityOrMissionDesc"`
	ContractTerminationInd         CheckboxType                        `xml:"ContractTerminationInd,omitempty" json:",omitempty"`
	VotingMembersGoverningBodyCnt  int                                 `xml:"VotingMembersGoverningBodyCnt"`
	VotingMembersIndependentCnt    int                                 `xml:"VotingMembersIndependentCnt"`
	TotalEmployeeCnt               int                                 `xml:"TotalEmployeeCnt"`
	TotalVolunteersCnt             int                                 `xml:"TotalVolunteersCnt,omitempty" json:",omitempty"`
	TotalGrossUBIAmt               int                                 `xml:"TotalGrossUBIAmt"`
	NetUnrelatedBusTxblIncmAmt     int                                 `xml:"NetUnrelatedBusTxblIncmAmt,omitempty" json:",omitempty"`
	PYContributionsGrantsAmt       int                                 `xml:"PYContributionsGrantsAmt,omitempty" json:",omitempty"`
	CYContributionsGrantsAmt       int                                 `xml:"CYContributionsGrantsAmt"`
	PYProgramServiceRevenueAmt     int                                 `xml:"PYProgramServiceRevenueAmt,omitempty" json:",omitempty"`
	CYProgramServiceRevenueAmt     int                                 `xml:"CYProgramServiceRevenueAmt"`
	PYInvestmentIncomeAmt          int                                 `xml:"PYInvestmentIncomeAmt,omitempty" json:",omitempty"`
	CYInvestmentIncomeAmt          int                                 `xml:"CYInvestmentIncomeAmt"`
	PYOtherRevenueAmt              int                                 `xml:"PYOtherRevenueAmt,omitempty" json:",omitempty"`
	CYOtherRevenueAmt              int                                 `xml:"CYOtherRevenueAmt"`
	PYTotalRevenueAmt              int                                 `xml:"PYTotalRevenueAmt,omitempty" json:",omitempty"`
	CYTotalRevenueAmt              int                                 `xml:"CYTotalRevenueAmt"`
	PYGrantsAndSimilarPaidAmt      int                                 `xml:"PYGrantsAndSimilarPaidAmt,omitempty" json:",omitempty"`
	CYGrantsAndSimilarPaidAmt      int                                 `xml:"CYGrantsAndSimilarPaidAmt"`
	PYBenefitsPaidToMembersAmt     int                                 `xml:"PYBenefitsPaidToMembersAmt,omitempty" json:",omitempty"`
	CYBenefitsPaidToMembersAmt     int                                 `xml:"CYBenefitsPaidToMembersAmt"`
	PYSalariesCompEmpBnftPaidAmt   int                                 `xml:"PYSalariesCompEmpBnftPaidAmt,omitempty" json:",omitempty"`
	CYSalariesCompEmpBnftPaidAmt   int                                 `xml:"CYSalariesCompEmpBnftPaidAmt"`
	PYTotalProfFndrsngExpnsAmt     int                                 `xml:"PYTotalProfFndrsngExpnsAmt,omitempty" json:",omitempty"`
	CYTotalProfFndrsngExpnsAmt     int                                 `xml:"CYTotalProfFndrsngExpnsAmt"`
	CYTotalFundraisingExpenseAmt   int                                 `xml:"CYTotalFundraisingExpenseAmt"`
	PYOtherExpensesAmt             int                                 `xml:"PYOtherExpensesAmt,omitempty" json:",omitempty"`
	CYOtherExpensesAmt             int                                 `xml:"CYOtherExpensesAmt"`
	PYTotalExpensesAmt             int                                 `xml:"PYTotalExpensesAmt,omitempty" json:",omitempty"`
	CYTotalExpensesAmt             int                                 `xml:"CYTotalExpensesAmt"`
	PYRevenuesLessExpensesAmt      int                                 `xml:"PYRevenuesLessExpensesAmt,omitempty" json:",omitempty"`
	CYRevenuesLessExpensesAmt      int                                 `xml:"CYRevenuesLessExpensesAmt"`
	TotalAssetsBOYAmt              int                                 `xml:"TotalAssetsBOYAmt,omitempty" json:",omitempty"`
	TotalAssetsEOYAmt              int                                 `xml:"TotalAssetsEOYAmt"`
	TotalLiabilitiesBOYAmt         int                                 `xml:"TotalLiabilitiesBOYAmt,omitempty" json:",omitempty"`
	TotalLiabilitiesEOYAmt         int                                 `xml:"TotalLiabilitiesEOYAmt"`
	NetAssetsOrFundBalancesBOYAmt  int                                 `xml:"NetAssetsOrFundBalancesBOYAmt,omitempty" json:",omitempty"`
	NetAssetsOrFundBalancesEOYAmt  int                                 `xml:"NetAssetsOrFundBalancesEOYAmt"`
	InfoInScheduleOPartIIIInd      CheckboxType                        `xml:"InfoInScheduleOPartIIIInd,omitempty" json:",omitempty"`
	MissionDesc                    string                              `xml:"MissionDesc,omitempty" json:",omitempty"`
	SignificantNewProgramSrvcInd   bool                                `xml:"SignificantNewProgramSrvcInd,omitempty" json:",omitempty"`
	SignificantChangeInd           bool                                `xml:"SignificantChangeInd,omitempty" json:",omitempty"`
	ActivityCd                     int                                 `xml:"ActivityCd,omitempty" json:",omitempty"`
	ExpenseAmt                     int                                 `xml:"ExpenseAmt,omitempty" json:",omitempty"`
	GrantAmt                       int                                 `xml:"GrantAmt,omitempty" json:",omitempty"`
	RevenueAmt                     int                                 `xml:"RevenueAmt,omitempty" json:",omitempty"`
	Desc                           string                              `xml:"Desc"`
	ProgSrvcAccomActy2Grp          *ProgSrvcAccomplishmentActyGrpType  `xml:"ProgSrvcAccomActy2Grp,omitempty" json:",omitempty"`
	ProgSrvcAccomActy3Grp          *ProgSrvcAccomplishmentActyGrpType  `xml:"ProgSrvcAccomActy3Grp,omitempty" json:",omitempty"`
	ProgSrvcAccomActyOtherGrp      []ProgSrvcAccomplishmentActyGrpType `xml:"ProgSrvcAccomActyOtherGrp,omitempty" json:",omitempty"`
	TotalOtherProgSrvcExpenseAmt   int                                 `xml:"TotalOtherProgSrvcExpenseAmt,omitempty" json:",omitempty"`
	TotalOtherProgSrvcGrantAmt     int                                 `xml:"TotalOtherProgSrvcGrantAmt,omitempty" json:",omitempty"`
	TotalOtherProgSrvcRevenueAmt   int                                 `xml:"TotalOtherProgSrvcRevenueAmt,omitempty" json:",omitempty"`
	TotalProgramServiceExpensesAmt int                                 `xml:"TotalProgramServiceExpensesAmt,omitempty" json:",omitempty"`
	DescribedInSection501c3Ind     DescribedInSection501c3Ind          `xml:"DescribedInSection501c3Ind"`
	ScheduleBRequiredInd           ScheduleBRequiredInd                `xml:"ScheduleBRequiredInd"`
	PoliticalCampaignActyInd       PoliticalCampaignActyInd            `xml:"PoliticalCampaignActyInd"`
	LobbyingActivitiesInd          *LobbyingActivitiesInd              `xml:"LobbyingActivitiesInd,omitempty" json:",omitempty"`
	SubjectToProxyTaxInd           *SubjectToProxyTaxInd               `xml:"SubjectToProxyTaxInd,omitempty" json:",omitempty"`
	DonorAdvisedFundInd            DonorAdvisedFundInd                 `xml:"DonorAdvisedFundInd"`
	ConservationEasementsInd       ConservationEasementsInd            `xml:"ConservationEasementsInd"`
	CollectionsOfArtInd            CollectionsOfArtInd                 `xml:"CollectionsOfArtInd"`
	CreditCounselingInd            CreditCounselingInd                 `xml:"CreditCounselingInd"`
	TempOrPermanentEndowmentsInd   TempOrPermanentEndowmentsInd        `xml:"TempOrPermanentEndowmentsInd"`
	ReportLandBuildingEquipmentInd ReportLandBuildingEquipmentInd      `xml:"ReportLandBuildingEquipmentInd"`
	ReportInvestmentsOtherSecInd   ReportInvestmentsOtherSecInd        `xml:"ReportInvestmentsOtherSecInd"`
	ReportProgramRelatedInvstInd   ReportProgramRelatedInvstInd        `xml:"ReportProgramRelatedInvstInd"`
	ReportOtherAssetsInd           ReportOtherAssetsInd                `xml:"ReportOtherAssetsInd"`
	ReportOtherLiabilitiesInd      ReportOtherLiabilitiesInd           `xml:"ReportOtherLiabilitiesInd"`
	IncludeFIN48FootnoteInd        IncludeFIN48FootnoteInd             `xml:"IncludeFIN48FootnoteInd"`
	IndependentAuditFinclStmtInd   IndependentAuditFinclStmtInd        `xml:"IndependentAuditFinclStmtInd"`
	ConsolidatedAuditFinclStmtInd  ConsolidatedAuditFinclStmtInd       `xml:"ConsolidatedAuditFinclStmtInd"`
	SchoolOperatingInd             SchoolOperatingInd                  `xml:"SchoolOperatingInd"`
	ForeignOfficeInd               bool                                `xml:"ForeignOfficeInd"`
	ForeignActivitiesInd           ForeignActivitiesInd                `xml:"ForeignActivitiesInd"`
	MoreThan5000KToOrgInd          MoreThan5000KToOrgInd               `xml:"MoreThan5000KToOrgInd"`
	MoreThan5000KToIndividualsInd  MoreThan5000KToIndividualsInd       `xml:"MoreThan5000KToIndividualsInd"`
	ProfessionalFundraisingInd     ProfessionalFundraisingInd          `xml:"ProfessionalFundraisingInd"`
	FundraisingActivitiesInd       FundraisingActivitiesInd            `xml:"FundraisingActivitiesInd"`
	GamingActivitiesInd            GamingActivitiesInd                 `xml:"GamingActivitiesInd"`
	OperateHospitalInd             OperateHospitalInd                  `xml:"OperateHospitalInd"`
	AuditedFinancialStmtAttInd     *AuditedFinancialStmtAttInd         `xml:"AuditedFinancialStmtAttInd,omitempty" json:",omitempty"`
	GrantsToOrganizationsInd       GrantsToOrganizationsInd            `xml:"GrantsToOrganizationsInd"`
	GrantsToIndividualsInd         GrantsToIndividualsInd              `xml:"GrantsToIndividualsInd"`
	ScheduleJRequiredInd           ScheduleJRequiredInd                `xml:"ScheduleJRequiredInd"`
	TaxExemptBondsInd              TaxExemptBondsInd                   `xml:"TaxExemptBondsInd"`
	InvestTaxExemptBondsInd        bool                                `xml:"InvestTaxExemptBondsInd,omitempty" json:",omitempty"`
	EscrowAccountInd               bool                                `xml:"EscrowAccountInd,omitempty" json:",omitempty"`
	OnBehalfOfIssuerInd            bool                                `xml:"OnBehalfOfIssuerInd,omitempty" json:",omitempty"`
	EngagedInExcessBenefitTransInd *EngagedInExcessBenefitTransInd     `xml:"EngagedInExcessBenefitTransInd,omitempty" json:",omitempty"`
	PYExcessBenefitTransInd        *PYExcessBenefitTransInd            `xml:"PYExcessBenefitTransInd,omitempty" json:",omitempty"`
	LoanOutstandingInd             LoanOutstandingInd                  `xml:"LoanOutstandingInd"`
	GrantToRelatedPersonInd        GrantToRelatedPersonInd             `xml:"GrantToRelatedPersonInd"`
	BusinessRlnWithOrgMemInd       BusinessRlnWithOrgMemInd            `xml:"BusinessRlnWithOrgMemInd"`
	BusinessRlnWithFamMemInd       BusinessRlnWithFamMemInd            `xml:"BusinessRlnWithFamMemInd"`
	BusinessRlnWithOfficerEntInd   BusinessRlnWithOfficerEntInd        `xml:"BusinessRlnWithOfficerEntInd"`
	DeductibleNonCashContriInd     DeductibleNonCashContriInd          `xml:"DeductibleNonCashContriInd"`
	DeductibleArtContributionInd   DeductibleArtContributionInd        `xml:"DeductibleArtContributionInd"`
	TerminateOperationsInd         TerminateOperationsInd              `xml:"TerminateOperationsInd"`
	PartialLiquidationInd          PartialLiquidationInd               `xml:"PartialLiquidationInd"`
	DisregardedEntityInd           DisregardedEntityInd                `xml:"DisregardedEntityInd"`
	RelatedEntityInd               RelatedEntityInd                    `xml:"RelatedEntityInd"`
	RelatedOrganizationCtrlEntInd  bool                                `xml:"RelatedOrganizationCtrlEntInd"`
	TransactionWithControlEntInd   *TransactionWithControlEntInd       `xml:"TransactionWithControlEntInd,omitempty" json:",omitempty"`
	TrnsfrExmptNonChrtblRltdOrgInd *TrnsfrExmptNonChrtblRltdOrgInd     `xml:"TrnsfrExmptNonChrtblRltdOrgInd,omitempty" json:",omitempty"`
	ActivitiesConductedPrtshpInd   *ActivitiesConductedPrtshpInd       `xml:"ActivitiesConductedPrtshpInd"`
	ScheduleORequiredInd           bool                                `xml:"ScheduleORequiredInd"`
	InfoInScheduleOPartVInd        CheckboxType                        `xml:"InfoInScheduleOPartVInd,omitempty" json:",omitempty"`
	IRPDocumentCnt                 int                                 `xml:"IRPDocumentCnt"`
	IRPDocumentW2GCnt              int                                 `xml:"IRPDocumentW2GCnt"`
	BackupWthldComplianceInd       bool                                `xml:"BackupWthldComplianceInd,omitempty" json:",omitempty"`
	EmployeeCnt                    int                                 `xml:"EmployeeCnt"`
	EmploymentTaxReturnsFiledInd   bool                                `xml:"EmploymentTaxReturnsFiledInd,omitempty" json:",omitempty"`
	UnrelatedBusIncmOverLimitInd   bool                                `xml:"UnrelatedBusIncmOverLimitInd"`
	Form990TFiledInd               bool                                `xml:"Form990TFiledInd,omitempty" json:",omitempty"`
	ForeignFinancialAccountInd     bool                                `xml:"ForeignFinancialAccountInd"`
	ForeignCountryCd               []CountryType                       `xml:"ForeignCountryCd,omitempty" json:",omitempty"`
	ProhibitedTaxShelterTransInd   bool                                `xml:"ProhibitedTaxShelterTransInd"`
	TaxablePartyNotificationInd    bool                                `xml:"TaxablePartyNotificationInd"`
	Form8886TFiledInd              bool                                `xml:"Form8886TFiledInd,omitempty" json:",omitempty"`
	NondeductibleContributionsInd  bool                                `xml:"NondeductibleContributionsInd"`
	NondeductibleContriDisclInd    bool                                `xml:"NondeductibleContriDisclInd,omitempty" json:",omitempty"`
	QuidProQuoContributionsInd     bool                                `xml:"QuidProQuoContributionsInd,omitempty" json:",omitempty"`
	QuidProQuoContriDisclInd       bool                                `xml:"QuidProQuoContriDisclInd,omitempty" json:",omitempty"`
	Form8282PropertyDisposedOfInd  bool                                `xml:"Form8282PropertyDisposedOfInd,omitempty" json:",omitempty"`
	Form8282FiledCnt               int                                 `xml:"Form8282FiledCnt,omitempty" json:",omitempty"`
	RcvFndsToPayPrsnlBnftCntrctInd bool                                `xml:"RcvFndsToPayPrsnlBnftCntrctInd,omitempty" json:",omitempty"`
	PayPremiumsPrsnlBnftCntrctInd  bool                                `xml:"PayPremiumsPrsnlBnftCntrctInd,omitempty" json:",omitempty"`
	Form8899Filedind               bool                                `xml:"Form8899Filedind,omitempty" json:",omitempty"`
	Form1098CFiledInd              bool                                `xml:"Form1098CFiledInd,omitempty" json:",omitempty"`
	DAFExcessBusinessHoldingsInd   bool                                `xml:"DAFExcessBusinessHoldingsInd,omitempty" json:",omitempty"`
	TaxableDistributionsInd        bool                                `xml:"TaxableDistributionsInd,omitempty" json:",omitempty"`
	DistributionToDonorInd         bool                                `xml:"DistributionToDonorInd,omitempty" json:",omitempty"`
	InitiationFeesAndCapContriAmt  int                                 `xml:"InitiationFeesAndCapContriAmt,omitempty" json:",omitempty"`
	GrossReceiptsForPublicUseAmt   int                                 `xml:"GrossReceiptsForPublicUseAmt,omitempty" json:",omitempty"`
	MembersAndShrGrossIncomeAmt    int                                 `xml:"MembersAndShrGrossIncomeAmt,omitempty" json:",omitempty"`
	OtherSourcesGrossIncomeAmt     int                                 `xml:"OtherSourcesGrossIncomeAmt,omitempty" json:",omitempty"`
	OrgFiledInLieuOfForm1041Ind    bool                                `xml:"OrgFiledInLieuOfForm1041Ind,omitempty" json:",omitempty"`
	TaxExemptInterestAmt           int                                 `xml:"TaxExemptInterestAmt,omitempty" json:",omitempty"`
	LicensedMoreThanOneStateInd    bool                                `xml:"LicensedMoreThanOneStateInd,omitempty" json:",omitempty"`
	StateRequiredReservesAmt       int                                 `xml:"StateRequiredReservesAmt,omitempty" json:",omitempty"`
	ReservesMaintainedAmt          int                                 `xml:"ReservesMaintainedAmt,omitempty" json:",omitempty"`
	IndoorTanningServicesInd       bool                                `xml:"IndoorTanningServicesInd"`
	Form720FiledInd                bool                                `xml:"Form720FiledInd,omitempty" json:",omitempty"`
	InfoInScheduleOPartVIInd       CheckboxType                        `xml:"InfoInScheduleOPartVIInd,omitempty" json:",omitempty"`
	GoverningBodyVotingMembersCnt  int                                 `xml:"GoverningBodyVotingMembersCnt"`
	IndependentVotingMemberCnt     int                                 `xml:"IndependentVotingMemberCnt"`
	FamilyOrBusinessRlnInd         bool                                `xml:"FamilyOrBusinessRlnInd"`
	DelegationOfMgmtDutiesInd      bool                                `xml:"DelegationOfMgmtDutiesInd"`
	ChangeToOrgDocumentsInd        bool                                `xml:"ChangeToOrgDocumentsInd"`
	MaterialDiversionOrMisuseInd   bool                                `xml:"MaterialDiversionOrMisuseInd"`
	MembersOrStockholdersInd       bool                                `xml:"MembersOrStockholdersInd"`
	ElectionOfBoardMembersInd      bool                                `xml:"ElectionOfBoardMembersInd"`
	DecisionsSubjectToApprovaInd   bool                                `xml:"DecisionsSubjectToApprovaInd"`
	MinutesOfGoverningBodyInd      bool                                `xml:"MinutesOfGoverningBodyInd"`
	MinutesOfCommitteesInd         bool                                `xml:"MinutesOfCommitteesInd,omitempty" json:",omitempty"`
	OfficerMailingAddressInd       bool                                `xml:"OfficerMailingAddressInd"`
	LocalChaptersInd               bool                                `xml:"LocalChaptersInd"`
	PoliciesReferenceChaptersInd   bool                                `xml:"PoliciesReferenceChaptersInd,omitempty" json:",omitempty"`
	Form990ProvidedToGvrnBodyInd   bool                                `xml:"Form990ProvidedToGvrnBodyInd"`
	ConflictOfInterestPolicyInd    bool                                `xml:"ConflictOfInterestPolicyInd"`
	AnnualDisclosureCoveredPrsnInd bool                                `xml:"AnnualDisclosureCoveredPrsnInd,omitempty" json:",omitempty"`
	RegularMonitoringEnfrcInd      bool                                `xml:"RegularMonitoringEnfrcInd,omitempty" json:",omitempty"`
	WhistleblowerPolicyInd         bool                                `xml:"WhistleblowerPolicyInd"`
	DocumentRetentionPolicyInd     bool                                `xml:"DocumentRetentionPolicyInd"`
	CompensationProcessCEOInd      bool                                `xml:"CompensationProcessCEOInd,omitempty" json:",omitempty"`
	CompensationProcessOtherInd    bool                                `xml:"CompensationProcessOtherInd,omitempty" json:",omitempty"`
	InvestmentInJointVentureInd    bool                                `xml:"InvestmentInJointVentureInd"`
	WrittenPolicyOrProcedureInd    bool                                `xml:"WrittenPolicyOrProcedureInd,omitempty" json:",omitempty"`
	StatesWhereCopyOfReturnIsFldCd []StateType                         `xml:"StatesWhereCopyOfReturnIsFldCd,omitempty" json:",omitempty"`
	OwnWebsiteInd                  CheckboxType                        `xml:"OwnWebsiteInd,omitempty" json:",omitempty"`
	OtherWebsiteInd                CheckboxType                        `xml:"OtherWebsiteInd,omitempty" json:",omitempty"`
	UponRequestInd                 CheckboxType                        `xml:"UponRequestInd,omitempty" json:",omitempty"`
	OtherInd                       CheckboxType                        `xml:"OtherInd,omitempty" json:",omitempty"`
	BooksInCareOfDetail            BooksInCareOfDetail                 `xml:"BooksInCareOfDetail"`
	InfoInScheduleOPartVIIInd      CheckboxType                        `xml:"InfoInScheduleOPartVIIInd,omitempty" json:",omitempty"`
	NoListedPersonsCompensatedInd  CheckboxType                        `xml:"NoListedPersonsCompensatedInd,omitempty" json:",omitempty"`
	Form990PartVIISectionAGrp      []Form990PartVIISectionAGrp         `xml:"Form990PartVIISectionAGrp"`
	TotalReportableCompFromOrgAmt  int                                 `xml:"TotalReportableCompFromOrgAmt,omitempty" json:",omitempty"`
	TotReportableCompRltdOrgAmt    int                                 `xml:"TotReportableCompRltdOrgAmt,omitempty" json:",omitempty"`
	TotalOtherCompensationAmt      int                                 `xml:"TotalOtherCompensationAmt,omitempty" json:",omitempty"`
	IndivRcvdGreaterThan100KCnt    int                                 `xml:"IndivRcvdGreaterThan100KCnt,omitempty" json:",omitempty"`
	FormerOfcrEmployeesListedInd   bool                                `xml:"FormerOfcrEmployeesListedInd"`
	TotalCompGreaterThan150KInd    bool                                `xml:"TotalCompGreaterThan150KInd"`
	CompensationFromOtherSrcsInd   bool                                `xml:"CompensationFromOtherSrcsInd"`
	ContractorCompensationGrp      []Form990PartVIIGroup1Type          `xml:"ContractorCompensationGrp,omitempty" json:",omitempty"`
	CntrctRcvdGreaterThan100KCnt   int                                 `xml:"CntrctRcvdGreaterThan100KCnt,omitempty" json:",omitempty"`
	InfoInScheduleOPartVIIIInd     CheckboxType                        `xml:"InfoInScheduleOPartVIIIInd,omitempty" json:",omitempty"`
	FederatedCampaignsAmt          int                                 `xml:"FederatedCampaignsAmt,omitempty" json:",omitempty"`
	MembershipDuesAmt              int                                 `xml:"MembershipDuesAmt,omitempty" json:",omitempty"`
	FundraisingAmt                 int                                 `xml:"FundraisingAmt,omitempty" json:",omitempty"`
	RelatedOrganizationsAmt        int                                 `xml:"RelatedOrganizationsAmt,omitempty" json:",omitempty"`
	GovernmentGrantsAmt            int                                 `xml:"GovernmentGrantsAmt,omitempty" json:",omitempty"`
	AllOtherContributionsAmt       int                                 `xml:"AllOtherContributionsAmt,omitempty" json:",omitempty"`
	NoncashContributionsAmt        int                                 `xml:"NoncashContributionsAmt,omitempty" json:",omitempty"`
	TotalContributionsAmt          int                                 `xml:"TotalContributionsAmt,omitempty" json:",omitempty"`
	ProgramServiceRevenueGrp       []Form990PartVIIIGroup2Type         `xml:"ProgramServiceRevenueGrp,omitempty" json:",omitempty"`
	TotalOthProgramServiceRevGrp   *Form990PartVIIIGroup3Type          `xml:"TotalOthProgramServiceRevGrp,omitempty" json:",omitempty"`
	TotalProgramServiceRevenueAmt  int                                 `xml:"TotalProgramServiceRevenueAmt,omitempty" json:",omitempty"`
	InvestmentIncomeGrp            *Form990PartVIIIGroup3Type          `xml:"InvestmentIncomeGrp,omitempty" json:",omitempty"`
	IncmFromInvestBondProceedsGrp  *Form990PartVIIIGroup3Type          `xml:"IncmFromInvestBondProceedsGrp,omitempty" json:",omitempty"`
	RoyaltiesRevenueGrp            *Form990PartVIIIGroup3Type          `xml:"RoyaltiesRevenueGrp,omitempty" json:",omitempty"`
	GrossRentsGrp                  *Form990PartVIIIGroup4Type          `xml:"GrossRentsGrp,omitempty" json:",omitempty"`
	LessRentalExpensesGrp          *Form990PartVIIIGroup4Type          `xml:"LessRentalExpensesGrp,omitempty" json:",omitempty"`
	RentalIncomeOrLossGrp          *Form990PartVIIIGroup4Type          `xml:"RentalIncomeOrLossGrp,omitempty" json:",omitempty"`
	NetRentalIncomeOrLossGrp       *Form990PartVIIIGroup3Type          `xml:"NetRentalIncomeOrLossGrp,omitempty" json:",omitempty"`
	GrossAmountSalesAssetsGrp      *Form990PartVIIIGroup5Type          `xml:"GrossAmountSalesAssetsGrp,omitempty" json:",omitempty"`
	LessCostOthBasisSalesExpnssGrp *Form990PartVIIIGroup5Type          `xml:"LessCostOthBasisSalesExpnssGrp,omitempty" json:",omitempty"`
	GainOrLossGrp                  *Form990PartVIIIGroup5Type          `xml:"GainOrLossGrp,omitempty" json:",omitempty"`
	NetGainOrLossInvestmentsGrp    *Form990PartVIIIGroup3Type          `xml:"NetGainOrLossInvestmentsGrp,omitempty" json:",omitempty"`
	FundraisingGrossIncomeAmt      int                                 `xml:"FundraisingGrossIncomeAmt,omitempty" json:",omitempty"`
	ContriRptFundraisingEventAmt   int                                 `xml:"ContriRptFundraisingEventAmt,omitempty" json:",omitempty"`
	FundraisingDirectExpensesAmt   int                                 `xml:"FundraisingDirectExpensesAmt,omitempty" json:",omitempty"`
	NetIncmFromFundraisingEvtGrp   *Form990PartVIIIGroup7Type          `xml:"NetIncmFromFundraisingEvtGrp,omitempty" json:",omitempty"`
	GamingGrossIncomeAmt           int                                 `xml:"GamingGrossIncomeAmt,omitempty" json:",omitempty"`
	GamingDirectExpensesAmt        int                                 `xml:"GamingDirectExpensesAmt,omitempty" json:",omitempty"`
	NetIncomeFromGamingGrp         *Form990PartVIIIGroup3Type          `xml:"NetIncomeFromGamingGrp,omitempty" json:",omitempty"`
	GrossSalesOfInventoryAmt       int                                 `xml:"GrossSalesOfInventoryAmt,omitempty" json:",omitempty"`
	CostOfGoodsSoldAmt             int                                 `xml:"CostOfGoodsSoldAmt,omitempty" json:",omitempty"`
	NetIncomeOrLossGrp             *Form990PartVIIIGroup3Type          `xml:"NetIncomeOrLossGrp,omitempty" json:",omitempty"`
	OtherRevenueMiscGrp            []Form990PartVIIIGroup2Type         `xml:"OtherRevenueMiscGrp,omitempty" json:",omitempty"`
	MiscellaneousRevenueGrp        *Form990PartVIIIGroup3Type          `xml:"MiscellaneousRevenueGrp,omitempty" json:",omitempty"`
	OtherRevenueTotalAmt           int                                 `xml:"OtherRevenueTotalAmt,omitempty" json:",omitempty"`
	TotalRevenueGrp                Form990PartVIIIGroup6Type           `xml:"TotalRevenueGrp"`
	InfoInScheduleOPartIXInd       CheckboxType                        `xml:"InfoInScheduleOPartIXInd,omitempty" json:",omitempty"`
	GrantsToDomesticOrgsGrp        *Form990PartIXGroup1Type            `xml:"GrantsToDomesticOrgsGrp,omitempty" json:",omitempty"`
	GrantsToDomesticIndividualsGrp *Form990PartIXGroup1Type            `xml:"GrantsToDomesticIndividualsGrp,omitempty" json:",omitempty"`
	ForeignGrantsGrp               *Form990PartIXGroup1Type            `xml:"ForeignGrantsGrp,omitempty" json:",omitempty"`
	BenefitsToMembersGrp           *Form990PartIXGroup1Type            `xml:"BenefitsToMembersGrp,omitempty" json:",omitempty"`
	CompCurrentOfcrDirectorsGrp    *Form990PartIXGroup2Type            `xml:"CompCurrentOfcrDirectorsGrp,omitempty" json:",omitempty"`
	CompDisqualPersonsGrp          *Form990PartIXGroup2Type            `xml:"CompDisqualPersonsGrp,omitempty" json:",omitempty"`
	OtherSalariesAndWagesGrp       *Form990PartIXGroup2Type            `xml:"OtherSalariesAndWagesGrp,omitempty" json:",omitempty"`
	PensionPlanContributionsGrp    *Form990PartIXGroup2Type            `xml:"PensionPlanContributionsGrp,omitempty" json:",omitempty"`
	OtherEmployeeBenefitsGrp       *Form990PartIXGroup2Type            `xml:"OtherEmployeeBenefitsGrp,omitempty" json:",omitempty"`
	PayrollTaxesGrp                *Form990PartIXGroup2Type            `xml:"PayrollTaxesGrp,omitempty" json:",omitempty"`
	FeesForServicesManagementGrp   *Form990PartIXGroup2Type            `xml:"FeesForServicesManagementGrp,omitempty" json:",omitempty"`
	FeesForServicesLegalGrp        *Form990PartIXGroup2Type            `xml:"FeesForServicesLegalGrp,omitempty" json:",omitempty"`
	FeesForServicesAccountingGrp   *Form990PartIXGroup2Type            `xml:"FeesForServicesAccountingGrp,omitempty" json:",omitempty"`
	FeesForServicesLobbyingGrp     *Form990PartIXGroup2Type            `xml:"FeesForServicesLobbyingGrp,omitempty" json:",omitempty"`
	FeesForServicesProfFundraising *FeesForServicesProfFundraising     `xml:"FeesForServicesProfFundraising,omitempty" json:",omitempty"`
	FeesForSrvcInvstMgmntFeesGrp   *Form990PartIXGroup2Type            `xml:"FeesForSrvcInvstMgmntFeesGrp,omitempty" json:",omitempty"`
	FeesForServicesOtherGrp        *Form990PartIXGroup2Type            `xml:"FeesForServicesOtherGrp,omitempty" json:",omitempty"`
	AdvertisingGrp                 *Form990PartIXGroup2Type            `xml:"AdvertisingGrp,omitempty" json:",omitempty"`
	OfficeExpensesGrp              *Form990PartIXGroup2Type            `xml:"OfficeExpensesGrp,omitempty" json:",omitempty"`
	InformationTechnologyGrp       *Form990PartIXGroup2Type            `xml:"InformationTechnologyGrp,omitempty" json:",omitempty"`
	RoyaltiesGrp                   *Form990PartIXGroup2Type            `xml:"RoyaltiesGrp,omitempty" json:",omitempty"`
	OccupancyGrp                   *Form990PartIXGroup2Type            `xml:"OccupancyGrp,omitempty" json:",omitempty"`
	TravelGrp                      *Form990PartIXGroup2Type            `xml:"TravelGrp,omitempty" json:",omitempty"`
	PymtTravelEntrtnmntPubOfclGrp  *Form990PartIXGroup2Type            `xml:"PymtTravelEntrtnmntPubOfclGrp,omitempty" json:",omitempty"`
	ConferencesMeetingsGrp         *Form990PartIXGroup2Type            `xml:"ConferencesMeetingsGrp,omitempty" json:",omitempty"`
	InterestGrp                    *Form990PartIXGroup2Type            `xml:"InterestGrp,omitempty" json:",omitempty"`
	PaymentsToAffiliatesGrp        *Form990PartIXGroup2Type            `xml:"PaymentsToAffiliatesGrp,omitempty" json:",omitempty"`
	DepreciationDepletionGrp       *Form990PartIXGroup2Type            `xml:"DepreciationDepletionGrp,omitempty" json:",omitempty"`
	InsuranceGrp                   *Form990PartIXGroup2Type            `xml:"InsuranceGrp,omitempty" json:",omitempty"`
	OtherExpensesGrp               []Form990PartIXGroup4Type           `xml:"OtherExpensesGrp,omitempty" json:",omitempty"`
	AllOtherExpensesGrp            *Form990PartIXGroup2Type            `xml:"AllOtherExpensesGrp,omitempty" json:",omitempty"`
	TotalFunctionalExpensesGrp     Form990PartIXGroup3Type             `xml:"TotalFunctionalExpensesGrp"`
	JointCostsInd                  CheckboxType                        `xml:"JointCostsInd,omitempty" json:",omitempty"`
	TotalJointCostsGrp             *Form990PartIXGroup2Type            `xml:"TotalJointCostsGrp,omitempty" json:",omitempty"`
	InfoInScheduleOPartXInd        CheckboxType                        `xml:"InfoInScheduleOPartXInd,omitempty" json:",omitempty"`
	CashNonInterestBearingGrp      *Form990PartXGroup1Type             `xml:"CashNonInterestBearingGrp,omitempty" json:",omitempty"`
	SavingsAndTempCashInvstGrp     *Form990PartXGroup1Type             `xml:"SavingsAndTempCashInvstGrp,omitempty" json:",omitempty"`
	PledgesAndGrantsReceivableGrp  *Form990PartXGroup1Type             `xml:"PledgesAndGrantsReceivableGrp,omitempty" json:",omitempty"`
	AccountsReceivableGrp          *Form990PartXGroup1Type             `xml:"AccountsReceivableGrp,omitempty" json:",omitempty"`
	ReceivablesFromOfficersEtcGrp  *Form990PartXGroup1Type             `xml:"ReceivablesFromOfficersEtcGrp,omitempty" json:",omitempty"`
	RcvblFromDisqualifiedPrsnGrp   *Form990PartXGroup1Type             `xml:"RcvblFromDisqualifiedPrsnGrp,omitempty" json:",omitempty"`
	OthNotesLoansReceivableNetGrp  *Form990PartXGroup1Type             `xml:"OthNotesLoansReceivableNetGrp,omitempty" json:",omitempty"`
	InventoriesForSaleOrUseGrp     *Form990PartXGroup1Type             `xml:"InventoriesForSaleOrUseGrp,omitempty" json:",omitempty"`
	PrepaidExpensesDefrdChargesGrp *Form990PartXGroup1Type             `xml:"PrepaidExpensesDefrdChargesGrp,omitempty" json:",omitempty"`
	LandBldgEquipCostOrOtherBssAmt int                                 `xml:"LandBldgEquipCostOrOtherBssAmt,omitempty" json:",omitempty"`
	LandBldgEquipAccumDeprecAmt    int                                 `xml:"LandBldgEquipAccumDeprecAmt,omitempty" json:",omitempty"`
	LandBldgEquipBasisNetGrp       *Form990PartXGroup1Type             `xml:"LandBldgEquipBasisNetGrp,omitempty" json:",omitempty"`
	InvestmentsPubTradedSecGrp     *Form990PartXGroup1Type             `xml:"InvestmentsPubTradedSecGrp,omitempty" json:",omitempty"`
	InvestmentsOtherSecuritiesGrp  *Form990PartXGroup1Type             `xml:"InvestmentsOtherSecuritiesGrp,omitempty" json:",omitempty"`
	InvestmentsProgramRelatedGrp   *Form990PartXGroup1Type             `xml:"InvestmentsProgramRelatedGrp,omitempty" json:",omitempty"`
	IntangibleAssetsGrp            *Form990PartXGroup1Type             `xml:"IntangibleAssetsGrp,omitempty" json:",omitempty"`
	OtherAssetsTotalGrp            *Form990PartXGroup1Type             `xml:"OtherAssetsTotalGrp,omitempty" json:",omitempty"`
	TotalAssetsGrp                 Form990PartXGroup2Type              `xml:"TotalAssetsGrp"`
	AccountsPayableAccrExpnssGrp   *Form990PartXGroup1Type             `xml:"AccountsPayableAccrExpnssGrp,omitempty" json:",omitempty"`
	GrantsPayableGrp               *Form990PartXGroup1Type             `xml:"GrantsPayableGrp,omitempty" json:",omitempty"`
	DeferredRevenueGrp             *Form990PartXGroup1Type             `xml:"DeferredRevenueGrp,omitempty" json:",omitempty"`
	TaxExemptBondLiabilitiesGrp    *Form990PartXGroup1Type             `xml:"TaxExemptBondLiabilitiesGrp,omitempty" json:",omitempty"`
	EscrowAccountLiabilityGrp      *Form990PartXGroup1Type             `xml:"EscrowAccountLiabilityGrp,omitempty" json:",omitempty"`
	LoansFromOfficersDirectorsGrp  *Form990PartXGroup1Type             `xml:"LoansFromOfficersDirectorsGrp,omitempty" json:",omitempty"`
	MortgNotesPyblScrdInvstPropGrp *Form990PartXGroup1Type             `xml:"MortgNotesPyblScrdInvstPropGrp,omitempty" json:",omitempty"`
	UnsecuredNotesLoansPayableGrp  *Form990PartXGroup1Type             `xml:"UnsecuredNotesLoansPayableGrp,omitempty" json:",omitempty"`
	OtherLiabilitiesGrp            *Form990PartXGroup1Type             `xml:"OtherLiabilitiesGrp,omitempty" json:",omitempty"`
	TotalLiabilitiesGrp            Form990PartXGroup2Type              `xml:"TotalLiabilitiesGrp"`
	OrganizationFollowsSFAS117Ind  CheckboxType                        `xml:"OrganizationFollowsSFAS117Ind,omitempty" json:",omitempty"`
	UnrestrictedNetAssetsGrp       *Form990PartXGroup1Type             `xml:"UnrestrictedNetAssetsGrp,omitempty" json:",omitempty"`
	TemporarilyRstrNetAssetsGrp    *Form990PartXGroup1Type             `xml:"TemporarilyRstrNetAssetsGrp,omitempty" json:",omitempty"`
	PermanentlyRstrNetAssetsGrp    *Form990PartXGroup1Type             `xml:"PermanentlyRstrNetAssetsGrp,omitempty" json:",omitempty"`
	OrgDoesNotFollowSFAS117Ind     CheckboxType                        `xml:"OrgDoesNotFollowSFAS117Ind,omitempty" json:",omitempty"`
	CapStkTrPrinCurrentFundsGrp    *Form990PartXGroup1Type             `xml:"CapStkTrPrinCurrentFundsGrp,omitempty" json:",omitempty"`
	PdInCapSrplsLandBldgEqpFundGrp *Form990PartXGroup1Type             `xml:"PdInCapSrplsLandBldgEqpFundGrp,omitempty" json:",omitempty"`
	RtnEarnEndowmentIncmOthFndsGrp *Form990PartXGroup1Type             `xml:"RtnEarnEndowmentIncmOthFndsGrp,omitempty" json:",omitempty"`
	TotalNetAssetsFundBalanceGrp   Form990PartXGroup2Type              `xml:"TotalNetAssetsFundBalanceGrp"`
	TotLiabNetAssetsFundBalanceGrp Form990PartXGroup2Type              `xml:"TotLiabNetAssetsFundBalanceGrp"`
	InfoInScheduleOPartXIInd       CheckboxType                        `xml:"InfoInScheduleOPartXIInd,omitempty" json:",omitempty"`
	ReconcilationRevenueExpnssAmt  int                                 `xml:"ReconcilationRevenueExpnssAmt"`
	NetUnrlzdGainsLossesInvstAmt   int                                 `xml:"NetUnrlzdGainsLossesInvstAmt,omitempty" json:",omitempty"`
	DonatedServicesAndUseFcltsAmt  int                                 `xml:"DonatedServicesAndUseFcltsAmt,omitempty" json:",omitempty"`
	InvestmentExpenseAmt           int                                 `xml:"InvestmentExpenseAmt,omitempty" json:",omitempty"`
	PriorPeriodAdjustmentsAmt      int                                 `xml:"PriorPeriodAdjustmentsAmt,omitempty" json:",omitempty"`
	OtherChangesInNetAssetsAmt     int                                 `xml:"OtherChangesInNetAssetsAmt,omitempty" json:",omitempty"`
	InfoInScheduleOPartXIIInd      CheckboxType                        `xml:"InfoInScheduleOPartXIIInd,omitempty" json:",omitempty"`
	MethodOfAccountingCashInd      CheckboxType                        `xml:"MethodOfAccountingCashInd"`
	MethodOfAccountingAccrualInd   CheckboxType                        `xml:"MethodOfAccountingAccrualInd"`
	MethodOfAccountingOtherInd     MethodOfAccountingOtherInd          `xml:"MethodOfAccountingOtherInd"`
	AccountantCompileOrReviewInd   bool                                `xml:"AccountantCompileOrReviewInd"`
	AcctCompileOrReviewBasisGrp    *FinancialStatementType             `xml:"AcctCompileOrReviewBasisGrp,omitempty" json:",omitempty"`
	FSAuditedInd                   bool                                `xml:"FSAuditedInd"`
	FSAuditedBasisGrp              *FinancialStatementType             `xml:"FSAuditedBasisGrp,omitempty" json:",omitempty"`
	AuditCommitteeInd              bool                                `xml:"AuditCommitteeInd,omitempty" json:",omitempty"`
	FederalGrantAuditRequiredInd   bool                                `xml:"FederalGrantAuditRequiredInd,omitempty" json:",omitempty"`
	FederalGrantAuditPerformedInd  bool                                `xml:"FederalGrantAuditPerformedInd,omitempty" json:",omitempty"`
	DocumentId                     IdType                              `xml:"documentId,attr"`
	SoftwareId                     *SoftwareIdType                     `xml:"softwareId,attr,omitempty" json:",omitempty"`
	SoftwareVersionNum             string                              `xml:"softwareVersionNum,attr,omitempty" json:",omitempty"`
	DocumentName                   string                              `xml:"documentName,attr,omitempty" json:",omitempty"`
	ReferenceDocumentId            IdListType                          `xml:"referenceDocumentId,attr,omitempty" json:",omitempty"`
	ReferenceDocumentName          string                              `xml:"referenceDocumentName,attr,omitempty" json:",omitempty"`
}

// Content model for Form 990
type IRS990Type struct {
	SpecialConditionDesc           []string                            `xml:"SpecialConditionDesc,omitempty" json:",omitempty"`
	AddressChangeInd               CheckboxType                        `xml:"AddressChangeInd,omitempty" json:",omitempty"`
	InitialReturnInd               CheckboxType                        `xml:"InitialReturnInd,omitempty" json:",omitempty"`
	FinalReturnInd                 CheckboxType                        `xml:"FinalReturnInd,omitempty" json:",omitempty"`
	AmendedReturnInd               CheckboxType                        `xml:"AmendedReturnInd,omitempty" json:",omitempty"`
	DoingBusinessAsName            *BusinessNameType                   `xml:"DoingBusinessAsName,omitempty" json:",omitempty"`
	PrincipalOfficerNm             PersonNameType                      `xml:"PrincipalOfficerNm"`
	PrincipalOfcrBusinessName      BusinessNameType                    `xml:"PrincipalOfcrBusinessName"`
	USAddress                      USAddressType                       `xml:"USAddress"`
	ForeignAddress                 ForeignAddressType                  `xml:"ForeignAddress"`
	GrossReceiptsAmt               int                                 `xml:"GrossReceiptsAmt"`
	GroupReturnForAffiliatesInd    bool                                `xml:"GroupReturnForAffiliatesInd"`
	AllAffiliatesIncludedInd       *AllAffiliatesIncludedInd           `xml:"AllAffiliatesIncludedInd,omitempty" json:",omitempty"`
	GroupExemptionNum              GroupExemptionNum                   `xml:"GroupExemptionNum,omitempty" json:",omitempty"`
	Organization501c3Ind           CheckboxType                        `xml:"Organization501c3Ind"`
	Organization501cInd            Organization501cInd                 `xml:"Organization501cInd"`
	Organization4947a1NotPFInd     CheckboxType                        `xml:"Organization4947a1NotPFInd"`
	Organization527Ind             CheckboxType                        `xml:"Organization527Ind"`
	WebsiteAddressTxt              string                              `xml:"WebsiteAddressTxt,omitempty" json:",omitempty"`
	TypeOfOrganizationCorpInd      CheckboxType                        `xml:"TypeOfOrganizationCorpInd,omitempty" json:",omitempty"`
	TypeOfOrganizationTrustInd     CheckboxType                        `xml:"TypeOfOrganizationTrustInd,omitempty" json:",omitempty"`
	TypeOfOrganizationAssocInd     CheckboxType                        `xml:"TypeOfOrganizationAssocInd,omitempty" json:",omitempty"`
	TypeOfOrganizationOtherInd     CheckboxType                        `xml:"TypeOfOrganizationOtherInd,omitempty" json:",omitempty"`
	OtherOrganizationDsc           string                              `xml:"OtherOrganizationDsc,omitempty" json:",omitempty"`
	FormationYr                    YearType                            `xml:"FormationYr,omitempty" json:",omitempty"`
	LegalDomicileStateCd           StateType                           `xml:"LegalDomicileStateCd"`
	LegalDomicileCountryCd         CountryType                         `xml:"LegalDomicileCountryCd"`
	ActivityOrMissionDesc          string                              `xml:"ActivityOrMissionDesc"`
	ContractTerminationInd         CheckboxType                        `xml:"ContractTerminationInd,omitempty" json:",omitempty"`
	VotingMembersGoverningBodyCnt  int                                 `xml:"VotingMembersGoverningBodyCnt"`
	VotingMembersIndependentCnt    int                                 `xml:"VotingMembersIndependentCnt"`
	TotalEmployeeCnt               int                                 `xml:"TotalEmployeeCnt"`
	TotalVolunteersCnt             int                                 `xml:"TotalVolunteersCnt,omitempty" json:",omitempty"`
	TotalGrossUBIAmt               int                                 `xml:"TotalGrossUBIAmt"`
	NetUnrelatedBusTxblIncmAmt     int                                 `xml:"NetUnrelatedBusTxblIncmAmt,omitempty" json:",omitempty"`
	PYContributionsGrantsAmt       int                                 `xml:"PYContributionsGrantsAmt,omitempty" json:",omitempty"`
	CYContributionsGrantsAmt       int                                 `xml:"CYContributionsGrantsAmt"`
	PYProgramServiceRevenueAmt     int                                 `xml:"PYProgramServiceRevenueAmt,omitempty" json:",omitempty"`
	CYProgramServiceRevenueAmt     int                                 `xml:"CYProgramServiceRevenueAmt"`
	PYInvestmentIncomeAmt          int                                 `xml:"PYInvestmentIncomeAmt,omitempty" json:",omitempty"`
	CYInvestmentIncomeAmt          int                                 `xml:"CYInvestmentIncomeAmt"`
	PYOtherRevenueAmt              int                                 `xml:"PYOtherRevenueAmt,omitempty" json:",omitempty"`
	CYOtherRevenueAmt              int                                 `xml:"CYOtherRevenueAmt"`
	PYTotalRevenueAmt              int                                 `xml:"PYTotalRevenueAmt,omitempty" json:",omitempty"`
	CYTotalRevenueAmt              int                                 `xml:"CYTotalRevenueAmt"`
	PYGrantsAndSimilarPaidAmt      int                                 `xml:"PYGrantsAndSimilarPaidAmt,omitempty" json:",omitempty"`
	CYGrantsAndSimilarPaidAmt      int                                 `xml:"CYGrantsAndSimilarPaidAmt"`
	PYBenefitsPaidToMembersAmt     int                                 `xml:"PYBenefitsPaidToMembersAmt,omitempty" json:",omitempty"`
	CYBenefitsPaidToMembersAmt     int                                 `xml:"CYBenefitsPaidToMembersAmt"`
	PYSalariesCompEmpBnftPaidAmt   int                                 `xml:"PYSalariesCompEmpBnftPaidAmt,omitempty" json:",omitempty"`
	CYSalariesCompEmpBnftPaidAmt   int                                 `xml:"CYSalariesCompEmpBnftPaidAmt"`
	PYTotalProfFndrsngExpnsAmt     int                                 `xml:"PYTotalProfFndrsngExpnsAmt,omitempty" json:",omitempty"`
	CYTotalProfFndrsngExpnsAmt     int                                 `xml:"CYTotalProfFndrsngExpnsAmt"`
	CYTotalFundraisingExpenseAmt   int                                 `xml:"CYTotalFundraisingExpenseAmt"`
	PYOtherExpensesAmt             int                                 `xml:"PYOtherExpensesAmt,omitempty" json:",omitempty"`
	CYOtherExpensesAmt             int                                 `xml:"CYOtherExpensesAmt"`
	PYTotalExpensesAmt             int                                 `xml:"PYTotalExpensesAmt,omitempty" json:",omitempty"`
	CYTotalExpensesAmt             int                                 `xml:"CYTotalExpensesAmt"`
	PYRevenuesLessExpensesAmt      int                                 `xml:"PYRevenuesLessExpensesAmt,omitempty" json:",omitempty"`
	CYRevenuesLessExpensesAmt      int                                 `xml:"CYRevenuesLessExpensesAmt"`
	TotalAssetsBOYAmt              int                                 `xml:"TotalAssetsBOYAmt,omitempty" json:",omitempty"`
	TotalAssetsEOYAmt              int                                 `xml:"TotalAssetsEOYAmt"`
	TotalLiabilitiesBOYAmt         int                                 `xml:"TotalLiabilitiesBOYAmt,omitempty" json:",omitempty"`
	TotalLiabilitiesEOYAmt         int                                 `xml:"TotalLiabilitiesEOYAmt"`
	NetAssetsOrFundBalancesBOYAmt  int                                 `xml:"NetAssetsOrFundBalancesBOYAmt,omitempty" json:",omitempty"`
	NetAssetsOrFundBalancesEOYAmt  int                                 `xml:"NetAssetsOrFundBalancesEOYAmt"`
	InfoInScheduleOPartIIIInd      CheckboxType                        `xml:"InfoInScheduleOPartIIIInd,omitempty" json:",omitempty"`
	MissionDesc                    string                              `xml:"MissionDesc,omitempty" json:",omitempty"`
	SignificantNewProgramSrvcInd   bool                                `xml:"SignificantNewProgramSrvcInd,omitempty" json:",omitempty"`
	SignificantChangeInd           bool                                `xml:"SignificantChangeInd,omitempty" json:",omitempty"`
	ActivityCd                     int                                 `xml:"ActivityCd,omitempty" json:",omitempty"`
	ExpenseAmt                     int                                 `xml:"ExpenseAmt,omitempty" json:",omitempty"`
	GrantAmt                       int                                 `xml:"GrantAmt,omitempty" json:",omitempty"`
	RevenueAmt                     int                                 `xml:"RevenueAmt,omitempty" json:",omitempty"`
	Desc                           string                              `xml:"Desc"`
	ProgSrvcAccomActy2Grp          *ProgSrvcAccomplishmentActyGrpType  `xml:"ProgSrvcAccomActy2Grp,omitempty" json:",omitempty"`
	ProgSrvcAccomActy3Grp          *ProgSrvcAccomplishmentActyGrpType  `xml:"ProgSrvcAccomActy3Grp,omitempty" json:",omitempty"`
	ProgSrvcAccomActyOtherGrp      []ProgSrvcAccomplishmentActyGrpType `xml:"ProgSrvcAccomActyOtherGrp,omitempty" json:",omitempty"`
	TotalOtherProgSrvcExpenseAmt   int                                 `xml:"TotalOtherProgSrvcExpenseAmt,omitempty" json:",omitempty"`
	TotalOtherProgSrvcGrantAmt     int                                 `xml:"TotalOtherProgSrvcGrantAmt,omitempty" json:",omitempty"`
	TotalOtherProgSrvcRevenueAmt   int                                 `xml:"TotalOtherProgSrvcRevenueAmt,omitempty" json:",omitempty"`
	TotalProgramServiceExpensesAmt int                                 `xml:"TotalProgramServiceExpensesAmt,omitempty" json:",omitempty"`
	DescribedInSection501c3Ind     DescribedInSection501c3Ind          `xml:"DescribedInSection501c3Ind"`
	ScheduleBRequiredInd           ScheduleBRequiredInd                `xml:"ScheduleBRequiredInd"`
	PoliticalCampaignActyInd       PoliticalCampaignActyInd            `xml:"PoliticalCampaignActyInd"`
	LobbyingActivitiesInd          *LobbyingActivitiesInd              `xml:"LobbyingActivitiesInd,omitempty" json:",omitempty"`
	SubjectToProxyTaxInd           *SubjectToProxyTaxInd               `xml:"SubjectToProxyTaxInd,omitempty" json:",omitempty"`
	DonorAdvisedFundInd            DonorAdvisedFundInd                 `xml:"DonorAdvisedFundInd"`
	ConservationEasementsInd       ConservationEasementsInd            `xml:"ConservationEasementsInd"`
	CollectionsOfArtInd            CollectionsOfArtInd                 `xml:"CollectionsOfArtInd"`
	CreditCounselingInd            CreditCounselingInd                 `xml:"CreditCounselingInd"`
	TempOrPermanentEndowmentsInd   TempOrPermanentEndowmentsInd        `xml:"TempOrPermanentEndowmentsInd"`
	ReportLandBuildingEquipmentInd ReportLandBuildingEquipmentInd      `xml:"ReportLandBuildingEquipmentInd"`
	ReportInvestmentsOtherSecInd   ReportInvestmentsOtherSecInd        `xml:"ReportInvestmentsOtherSecInd"`
	ReportProgramRelatedInvstInd   ReportProgramRelatedInvstInd        `xml:"ReportProgramRelatedInvstInd"`
	ReportOtherAssetsInd           ReportOtherAssetsInd                `xml:"ReportOtherAssetsInd"`
	ReportOtherLiabilitiesInd      ReportOtherLiabilitiesInd           `xml:"ReportOtherLiabilitiesInd"`
	IncludeFIN48FootnoteInd        IncludeFIN48FootnoteInd             `xml:"IncludeFIN48FootnoteInd"`
	IndependentAuditFinclStmtInd   IndependentAuditFinclStmtInd        `xml:"IndependentAuditFinclStmtInd"`
	ConsolidatedAuditFinclStmtInd  ConsolidatedAuditFinclStmtInd       `xml:"ConsolidatedAuditFinclStmtInd"`
	SchoolOperatingInd             SchoolOperatingInd                  `xml:"SchoolOperatingInd"`
	ForeignOfficeInd               bool                                `xml:"ForeignOfficeInd"`
	ForeignActivitiesInd           ForeignActivitiesInd                `xml:"ForeignActivitiesInd"`
	MoreThan5000KToOrgInd          MoreThan5000KToOrgInd               `xml:"MoreThan5000KToOrgInd"`
	MoreThan5000KToIndividualsInd  MoreThan5000KToIndividualsInd       `xml:"MoreThan5000KToIndividualsInd"`
	ProfessionalFundraisingInd     ProfessionalFundraisingInd          `xml:"ProfessionalFundraisingInd"`
	FundraisingActivitiesInd       FundraisingActivitiesInd            `xml:"FundraisingActivitiesInd"`
	GamingActivitiesInd            GamingActivitiesInd                 `xml:"GamingActivitiesInd"`
	OperateHospitalInd             OperateHospitalInd                  `xml:"OperateHospitalInd"`
	AuditedFinancialStmtAttInd     *AuditedFinancialStmtAttInd         `xml:"AuditedFinancialStmtAttInd,omitempty" json:",omitempty"`
	GrantsToOrganizationsInd       GrantsToOrganizationsInd            `xml:"GrantsToOrganizationsInd"`
	GrantsToIndividualsInd         GrantsToIndividualsInd              `xml:"GrantsToIndividualsInd"`
	ScheduleJRequiredInd           ScheduleJRequiredInd                `xml:"ScheduleJRequiredInd"`
	TaxExemptBondsInd              TaxExemptBondsInd                   `xml:"TaxExemptBondsInd"`
	InvestTaxExemptBondsInd        bool                                `xml:"InvestTaxExemptBondsInd,omitempty" json:",omitempty"`
	EscrowAccountInd               bool                                `xml:"EscrowAccountInd,omitempty" json:",omitempty"`
	OnBehalfOfIssuerInd            bool                                `xml:"OnBehalfOfIssuerInd,omitempty" json:",omitempty"`
	EngagedInExcessBenefitTransInd *EngagedInExcessBenefitTransInd     `xml:"EngagedInExcessBenefitTransInd,omitempty" json:",omitempty"`
	PYExcessBenefitTransInd        *PYExcessBenefitTransInd            `xml:"PYExcessBenefitTransInd,omitempty" json:",omitempty"`
	LoanOutstandingInd             LoanOutstandingInd                  `xml:"LoanOutstandingInd"`
	GrantToRelatedPersonInd        GrantToRelatedPersonInd             `xml:"GrantToRelatedPersonInd"`
	BusinessRlnWithOrgMemInd       BusinessRlnWithOrgMemInd            `xml:"BusinessRlnWithOrgMemInd"`
	BusinessRlnWithFamMemInd       BusinessRlnWithFamMemInd            `xml:"BusinessRlnWithFamMemInd"`
	BusinessRlnWithOfficerEntInd   BusinessRlnWithOfficerEntInd        `xml:"BusinessRlnWithOfficerEntInd"`
	DeductibleNonCashContriInd     DeductibleNonCashContriInd          `xml:"DeductibleNonCashContriInd"`
	DeductibleArtContributionInd   DeductibleArtContributionInd        `xml:"DeductibleArtContributionInd"`
	TerminateOperationsInd         TerminateOperationsInd              `xml:"TerminateOperationsInd"`
	PartialLiquidationInd          PartialLiquidationInd               `xml:"PartialLiquidationInd"`
	DisregardedEntityInd           DisregardedEntityInd                `xml:"DisregardedEntityInd"`
	RelatedEntityInd               RelatedEntityInd                    `xml:"RelatedEntityInd"`
	RelatedOrganizationCtrlEntInd  bool                                `xml:"RelatedOrganizationCtrlEntInd"`
	TransactionWithControlEntInd   *TransactionWithControlEntInd       `xml:"TransactionWithControlEntInd,omitempty" json:",omitempty"`
	TrnsfrExmptNonChrtblRltdOrgInd *TrnsfrExmptNonChrtblRltdOrgInd     `xml:"TrnsfrExmptNonChrtblRltdOrgInd,omitempty" json:",omitempty"`
	ActivitiesConductedPrtshpInd   *ActivitiesConductedPrtshpInd       `xml:"ActivitiesConductedPrtshpInd"`
	ScheduleORequiredInd           bool                                `xml:"ScheduleORequiredInd"`
	InfoInScheduleOPartVInd        CheckboxType                        `xml:"InfoInScheduleOPartVInd,omitempty" json:",omitempty"`
	IRPDocumentCnt                 int                                 `xml:"IRPDocumentCnt"`
	IRPDocumentW2GCnt              int                                 `xml:"IRPDocumentW2GCnt"`
	BackupWthldComplianceInd       bool                                `xml:"BackupWthldComplianceInd,omitempty" json:",omitempty"`
	EmployeeCnt                    int                                 `xml:"EmployeeCnt"`
	EmploymentTaxReturnsFiledInd   bool                                `xml:"EmploymentTaxReturnsFiledInd,omitempty" json:",omitempty"`
	UnrelatedBusIncmOverLimitInd   bool                                `xml:"UnrelatedBusIncmOverLimitInd"`
	Form990TFiledInd               bool                                `xml:"Form990TFiledInd,omitempty" json:",omitempty"`
	ForeignFinancialAccountInd     bool                                `xml:"ForeignFinancialAccountInd"`
	ForeignCountryCd               []CountryType                       `xml:"ForeignCountryCd,omitempty" json:",omitempty"`
	ProhibitedTaxShelterTransInd   bool                                `xml:"ProhibitedTaxShelterTransInd"`
	TaxablePartyNotificationInd    bool                                `xml:"TaxablePartyNotificationInd"`
	Form8886TFiledInd              bool                                `xml:"Form8886TFiledInd,omitempty" json:",omitempty"`
	NondeductibleContributionsInd  bool                                `xml:"NondeductibleContributionsInd"`
	NondeductibleContriDisclInd    bool                                `xml:"NondeductibleContriDisclInd,omitempty" json:",omitempty"`
	QuidProQuoContributionsInd     bool                                `xml:"QuidProQuoContributionsInd,omitempty" json:",omitempty"`
	QuidProQuoContriDisclInd       bool                                `xml:"QuidProQuoContriDisclInd,omitempty" json:",omitempty"`
	Form8282PropertyDisposedOfInd  bool                                `xml:"Form8282PropertyDisposedOfInd,omitempty" json:",omitempty"`
	Form8282FiledCnt               int                                 `xml:"Form8282FiledCnt,omitempty" json:",omitempty"`
	RcvFndsToPayPrsnlBnftCntrctInd bool                                `xml:"RcvFndsToPayPrsnlBnftCntrctInd,omitempty" json:",omitempty"`
	PayPremiumsPrsnlBnftCntrctInd  bool                                `xml:"PayPremiumsPrsnlBnftCntrctInd,omitempty" json:",omitempty"`
	Form8899Filedind               bool                                `xml:"Form8899Filedind,omitempty" json:",omitempty"`
	Form1098CFiledInd              bool                                `xml:"Form1098CFiledInd,omitempty" json:",omitempty"`
	DAFExcessBusinessHoldingsInd   bool                                `xml:"DAFExcessBusinessHoldingsInd,omitempty" json:",omitempty"`
	TaxableDistributionsInd        bool                                `xml:"TaxableDistributionsInd,omitempty" json:",omitempty"`
	DistributionToDonorInd         bool                                `xml:"DistributionToDonorInd,omitempty" json:",omitempty"`
	InitiationFeesAndCapContriAmt  int                                 `xml:"InitiationFeesAndCapContriAmt,omitempty" json:",omitempty"`
	GrossReceiptsForPublicUseAmt   int                                 `xml:"GrossReceiptsForPublicUseAmt,omitempty" json:",omitempty"`
	MembersAndShrGrossIncomeAmt    int                                 `xml:"MembersAndShrGrossIncomeAmt,omitempty" json:",omitempty"`
	OtherSourcesGrossIncomeAmt     int                                 `xml:"OtherSourcesGrossIncomeAmt,omitempty" json:",omitempty"`
	OrgFiledInLieuOfForm1041Ind    bool                                `xml:"OrgFiledInLieuOfForm1041Ind,omitempty" json:",omitempty"`
	TaxExemptInterestAmt           int                                 `xml:"TaxExemptInterestAmt,omitempty" json:",omitempty"`
	LicensedMoreThanOneStateInd    bool                                `xml:"LicensedMoreThanOneStateInd,omitempty" json:",omitempty"`
	StateRequiredReservesAmt       int                                 `xml:"StateRequiredReservesAmt,omitempty" json:",omitempty"`
	ReservesMaintainedAmt          int                                 `xml:"ReservesMaintainedAmt,omitempty" json:",omitempty"`
	IndoorTanningServicesInd       bool                                `xml:"IndoorTanningServicesInd"`
	Form720FiledInd                bool                                `xml:"Form720FiledInd,omitempty" json:",omitempty"`
	InfoInScheduleOPartVIInd       CheckboxType                        `xml:"InfoInScheduleOPartVIInd,omitempty" json:",omitempty"`
	GoverningBodyVotingMembersCnt  int                                 `xml:"GoverningBodyVotingMembersCnt"`
	IndependentVotingMemberCnt     int                                 `xml:"IndependentVotingMemberCnt"`
	FamilyOrBusinessRlnInd         bool                                `xml:"FamilyOrBusinessRlnInd"`
	DelegationOfMgmtDutiesInd      bool                                `xml:"DelegationOfMgmtDutiesInd"`
	ChangeToOrgDocumentsInd        bool                                `xml:"ChangeToOrgDocumentsInd"`
	MaterialDiversionOrMisuseInd   bool                                `xml:"MaterialDiversionOrMisuseInd"`
	MembersOrStockholdersInd       bool                                `xml:"MembersOrStockholdersInd"`
	ElectionOfBoardMembersInd      bool                                `xml:"ElectionOfBoardMembersInd"`
	DecisionsSubjectToApprovaInd   bool                                `xml:"DecisionsSubjectToApprovaInd"`
	MinutesOfGoverningBodyInd      bool                                `xml:"MinutesOfGoverningBodyInd"`
	MinutesOfCommitteesInd         bool                                `xml:"MinutesOfCommitteesInd,omitempty" json:",omitempty"`
	OfficerMailingAddressInd       bool                                `xml:"OfficerMailingAddressInd"`
	LocalChaptersInd               bool                                `xml:"LocalChaptersInd"`
	PoliciesReferenceChaptersInd   bool                                `xml:"PoliciesReferenceChaptersInd,omitempty" json:",omitempty"`
	Form990ProvidedToGvrnBodyInd   bool                                `xml:"Form990ProvidedToGvrnBodyInd"`
	ConflictOfInterestPolicyInd    bool                                `xml:"ConflictOfInterestPolicyInd"`
	AnnualDisclosureCoveredPrsnInd bool                                `xml:"AnnualDisclosureCoveredPrsnInd,omitempty" json:",omitempty"`
	RegularMonitoringEnfrcInd      bool                                `xml:"RegularMonitoringEnfrcInd,omitempty" json:",omitempty"`
	WhistleblowerPolicyInd         bool                                `xml:"WhistleblowerPolicyInd"`
	DocumentRetentionPolicyInd     bool                                `xml:"DocumentRetentionPolicyInd"`
	CompensationProcessCEOInd      bool                                `xml:"CompensationProcessCEOInd,omitempty" json:",omitempty"`
	CompensationProcessOtherInd    bool                                `xml:"CompensationProcessOtherInd,omitempty" json:",omitempty"`
	InvestmentInJointVentureInd    bool                                `xml:"InvestmentInJointVentureInd"`
	WrittenPolicyOrProcedureInd    bool                                `xml:"WrittenPolicyOrProcedureInd,omitempty" json:",omitempty"`
	StatesWhereCopyOfReturnIsFldCd []StateType                         `xml:"StatesWhereCopyOfReturnIsFldCd,omitempty" json:",omitempty"`
	OwnWebsiteInd                  CheckboxType                        `xml:"OwnWebsiteInd,omitempty" json:",omitempty"`
	OtherWebsiteInd                CheckboxType                        `xml:"OtherWebsiteInd,omitempty" json:",omitempty"`
	UponRequestInd                 CheckboxType                        `xml:"UponRequestInd,omitempty" json:",omitempty"`
	OtherInd                       CheckboxType                        `xml:"OtherInd,omitempty" json:",omitempty"`
	BooksInCareOfDetail            BooksInCareOfDetail                 `xml:"BooksInCareOfDetail"`
	InfoInScheduleOPartVIIInd      CheckboxType                        `xml:"InfoInScheduleOPartVIIInd,omitempty" json:",omitempty"`
	NoListedPersonsCompensatedInd  CheckboxType                        `xml:"NoListedPersonsCompensatedInd,omitempty" json:",omitempty"`
	Form990PartVIISectionAGrp      []Form990PartVIISectionAGrp         `xml:"Form990PartVIISectionAGrp"`
	TotalReportableCompFromOrgAmt  int                                 `xml:"TotalReportableCompFromOrgAmt,omitempty" json:",omitempty"`
	TotReportableCompRltdOrgAmt    int                                 `xml:"TotReportableCompRltdOrgAmt,omitempty" json:",omitempty"`
	TotalOtherCompensationAmt      int                                 `xml:"TotalOtherCompensationAmt,omitempty" json:",omitempty"`
	IndivRcvdGreaterThan100KCnt    int                                 `xml:"IndivRcvdGreaterThan100KCnt,omitempty" json:",omitempty"`
	FormerOfcrEmployeesListedInd   bool                                `xml:"FormerOfcrEmployeesListedInd"`
	TotalCompGreaterThan150KInd    bool                                `xml:"TotalCompGreaterThan150KInd"`
	CompensationFromOtherSrcsInd   bool                                `xml:"CompensationFromOtherSrcsInd"`
	ContractorCompensationGrp      []Form990PartVIIGroup1Type          `xml:"ContractorCompensationGrp,omitempty" json:",omitempty"`
	CntrctRcvdGreaterThan100KCnt   int                                 `xml:"CntrctRcvdGreaterThan100KCnt,omitempty" json:",omitempty"`
	InfoInScheduleOPartVIIIInd     CheckboxType                        `xml:"InfoInScheduleOPartVIIIInd,omitempty" json:",omitempty"`
	FederatedCampaignsAmt          int                                 `xml:"FederatedCampaignsAmt,omitempty" json:",omitempty"`
	MembershipDuesAmt              int                                 `xml:"MembershipDuesAmt,omitempty" json:",omitempty"`
	FundraisingAmt                 int                                 `xml:"FundraisingAmt,omitempty" json:",omitempty"`
	RelatedOrganizationsAmt        int                                 `xml:"RelatedOrganizationsAmt,omitempty" json:",omitempty"`
	GovernmentGrantsAmt            int                                 `xml:"GovernmentGrantsAmt,omitempty" json:",omitempty"`
	AllOtherContributionsAmt       int                                 `xml:"AllOtherContributionsAmt,omitempty" json:",omitempty"`
	NoncashContributionsAmt        int                                 `xml:"NoncashContributionsAmt,omitempty" json:",omitempty"`
	TotalContributionsAmt          int                                 `xml:"TotalContributionsAmt,omitempty" json:",omitempty"`
	ProgramServiceRevenueGrp       []Form990PartVIIIGroup2Type         `xml:"ProgramServiceRevenueGrp,omitempty" json:",omitempty"`
	TotalOthProgramServiceRevGrp   *Form990PartVIIIGroup3Type          `xml:"TotalOthProgramServiceRevGrp,omitempty" json:",omitempty"`
	TotalProgramServiceRevenueAmt  int                                 `xml:"TotalProgramServiceRevenueAmt,omitempty" json:",omitempty"`
	InvestmentIncomeGrp            *Form990PartVIIIGroup3Type          `xml:"InvestmentIncomeGrp,omitempty" json:",omitempty"`
	IncmFromInvestBondProceedsGrp  *Form990PartVIIIGroup3Type          `xml:"IncmFromInvestBondProceedsGrp,omitempty" json:",omitempty"`
	RoyaltiesRevenueGrp            *Form990PartVIIIGroup3Type          `xml:"RoyaltiesRevenueGrp,omitempty" json:",omitempty"`
	GrossRentsGrp                  *Form990PartVIIIGroup4Type          `xml:"GrossRentsGrp,omitempty" json:",omitempty"`
	LessRentalExpensesGrp          *Form990PartVIIIGroup4Type          `xml:"LessRentalExpensesGrp,omitempty" json:",omitempty"`
	RentalIncomeOrLossGrp          *Form990PartVIIIGroup4Type          `xml:"RentalIncomeOrLossGrp,omitempty" json:",omitempty"`
	NetRentalIncomeOrLossGrp       *Form990PartVIIIGroup3Type          `xml:"NetRentalIncomeOrLossGrp,omitempty" json:",omitempty"`
	GrossAmountSalesAssetsGrp      *Form990PartVIIIGroup5Type          `xml:"GrossAmountSalesAssetsGrp,omitempty" json:",omitempty"`
	LessCostOthBasisSalesExpnssGrp *Form990PartVIIIGroup5Type          `xml:"LessCostOthBasisSalesExpnssGrp,omitempty" json:",omitempty"`
	GainOrLossGrp                  *Form990PartVIIIGroup5Type          `xml:"GainOrLossGrp,omitempty" json:",omitempty"`
	NetGainOrLossInvestmentsGrp    *Form990PartVIIIGroup3Type          `xml:"NetGainOrLossInvestmentsGrp,omitempty" json:",omitempty"`
	FundraisingGrossIncomeAmt      int                                 `xml:"FundraisingGrossIncomeAmt,omitempty" json:",omitempty"`
	ContriRptFundraisingEventAmt   int                                 `xml:"ContriRptFundraisingEventAmt,omitempty" json:",omitempty"`
	FundraisingDirectExpensesAmt   int                                 `xml:"FundraisingDirectExpensesAmt,omitempty" json:",omitempty"`
	NetIncmFromFundraisingEvtGrp   *Form990PartVIIIGroup7Type          `xml:"NetIncmFromFundraisingEvtGrp,omitempty" json:",omitempty"`
	GamingGrossIncomeAmt           int                                 `xml:"GamingGrossIncomeAmt,omitempty" json:",omitempty"`
	GamingDirectExpensesAmt        int                                 `xml:"GamingDirectExpensesAmt,omitempty" json:",omitempty"`
	NetIncomeFromGamingGrp         *Form990PartVIIIGroup3Type          `xml:"NetIncomeFromGamingGrp,omitempty" json:",omitempty"`
	GrossSalesOfInventoryAmt       int                                 `xml:"GrossSalesOfInventoryAmt,omitempty" json:",omitempty"`
	CostOfGoodsSoldAmt             int                                 `xml:"CostOfGoodsSoldAmt,omitempty" json:",omitempty"`
	NetIncomeOrLossGrp             *Form990PartVIIIGroup3Type          `xml:"NetIncomeOrLossGrp,omitempty" json:",omitempty"`
	OtherRevenueMiscGrp            []Form990PartVIIIGroup2Type         `xml:"OtherRevenueMiscGrp,omitempty" json:",omitempty"`
	MiscellaneousRevenueGrp        *Form990PartVIIIGroup3Type          `xml:"MiscellaneousRevenueGrp,omitempty" json:",omitempty"`
	OtherRevenueTotalAmt           int                                 `xml:"OtherRevenueTotalAmt,omitempty" json:",omitempty"`
	TotalRevenueGrp                Form990PartVIIIGroup6Type           `xml:"TotalRevenueGrp"`
	InfoInScheduleOPartIXInd       CheckboxType                        `xml:"InfoInScheduleOPartIXInd,omitempty" json:",omitempty"`
	GrantsToDomesticOrgsGrp        *Form990PartIXGroup1Type            `xml:"GrantsToDomesticOrgsGrp,omitempty" json:",omitempty"`
	GrantsToDomesticIndividualsGrp *Form990PartIXGroup1Type            `xml:"GrantsToDomesticIndividualsGrp,omitempty" json:",omitempty"`
	ForeignGrantsGrp               *Form990PartIXGroup1Type            `xml:"ForeignGrantsGrp,omitempty" json:",omitempty"`
	BenefitsToMembersGrp           *Form990PartIXGroup1Type            `xml:"BenefitsToMembersGrp,omitempty" json:",omitempty"`
	CompCurrentOfcrDirectorsGrp    *Form990PartIXGroup2Type            `xml:"CompCurrentOfcrDirectorsGrp,omitempty" json:",omitempty"`
	CompDisqualPersonsGrp          *Form990PartIXGroup2Type            `xml:"CompDisqualPersonsGrp,omitempty" json:",omitempty"`
	OtherSalariesAndWagesGrp       *Form990PartIXGroup2Type            `xml:"OtherSalariesAndWagesGrp,omitempty" json:",omitempty"`
	PensionPlanContributionsGrp    *Form990PartIXGroup2Type            `xml:"PensionPlanContributionsGrp,omitempty" json:",omitempty"`
	OtherEmployeeBenefitsGrp       *Form990PartIXGroup2Type            `xml:"OtherEmployeeBenefitsGrp,omitempty" json:",omitempty"`
	PayrollTaxesGrp                *Form990PartIXGroup2Type            `xml:"PayrollTaxesGrp,omitempty" json:",omitempty"`
	FeesForServicesManagementGrp   *Form990PartIXGroup2Type            `xml:"FeesForServicesManagementGrp,omitempty" json:",omitempty"`
	FeesForServicesLegalGrp        *Form990PartIXGroup2Type            `xml:"FeesForServicesLegalGrp,omitempty" json:",omitempty"`
	FeesForServicesAccountingGrp   *Form990PartIXGroup2Type            `xml:"FeesForServicesAccountingGrp,omitempty" json:",omitempty"`
	FeesForServicesLobbyingGrp     *Form990PartIXGroup2Type            `xml:"FeesForServicesLobbyingGrp,omitempty" json:",omitempty"`
	FeesForServicesProfFundraising *FeesForServicesProfFundraising     `xml:"FeesForServicesProfFundraising,omitempty" json:",omitempty"`
	FeesForSrvcInvstMgmntFeesGrp   *Form990PartIXGroup2Type            `xml:"FeesForSrvcInvstMgmntFeesGrp,omitempty" json:",omitempty"`
	FeesForServicesOtherGrp        *Form990PartIXGroup2Type            `xml:"FeesForServicesOtherGrp,omitempty" json:",omitempty"`
	AdvertisingGrp                 *Form990PartIXGroup2Type            `xml:"AdvertisingGrp,omitempty" json:",omitempty"`
	OfficeExpensesGrp              *Form990PartIXGroup2Type            `xml:"OfficeExpensesGrp,omitempty" json:",omitempty"`
	InformationTechnologyGrp       *Form990PartIXGroup2Type            `xml:"InformationTechnologyGrp,omitempty" json:",omitempty"`
	RoyaltiesGrp                   *Form990PartIXGroup2Type            `xml:"RoyaltiesGrp,omitempty" json:",omitempty"`
	OccupancyGrp                   *Form990PartIXGroup2Type            `xml:"OccupancyGrp,omitempty" json:",omitempty"`
	TravelGrp                      *Form990PartIXGroup2Type            `xml:"TravelGrp,omitempty" json:",omitempty"`
	PymtTravelEntrtnmntPubOfclGrp  *Form990PartIXGroup2Type            `xml:"PymtTravelEntrtnmntPubOfclGrp,omitempty" json:",omitempty"`
	ConferencesMeetingsGrp         *Form990PartIXGroup2Type            `xml:"ConferencesMeetingsGrp,omitempty" json:",omitempty"`
	InterestGrp                    *Form990PartIXGroup2Type            `xml:"InterestGrp,omitempty" json:",omitempty"`
	PaymentsToAffiliatesGrp        *Form990PartIXGroup2Type            `xml:"PaymentsToAffiliatesGrp,omitempty" json:",omitempty"`
	DepreciationDepletionGrp       *Form990PartIXGroup2Type            `xml:"DepreciationDepletionGrp,omitempty" json:",omitempty"`
	InsuranceGrp                   *Form990PartIXGroup2Type            `xml:"InsuranceGrp,omitempty" json:",omitempty"`
	OtherExpensesGrp               []Form990PartIXGroup4Type           `xml:"OtherExpensesGrp,omitempty" json:",omitempty"`
	AllOtherExpensesGrp            *Form990PartIXGroup2Type            `xml:"AllOtherExpensesGrp,omitempty" json:",omitempty"`
	TotalFunctionalExpensesGrp     Form990PartIXGroup3Type             `xml:"TotalFunctionalExpensesGrp"`
	JointCostsInd                  CheckboxType                        `xml:"JointCostsInd,omitempty" json:",omitempty"`
	TotalJointCostsGrp             *Form990PartIXGroup2Type            `xml:"TotalJointCostsGrp,omitempty" json:",omitempty"`
	InfoInScheduleOPartXInd        CheckboxType                        `xml:"InfoInScheduleOPartXInd,omitempty" json:",omitempty"`
	CashNonInterestBearingGrp      *Form990PartXGroup1Type             `xml:"CashNonInterestBearingGrp,omitempty" json:",omitempty"`
	SavingsAndTempCashInvstGrp     *Form990PartXGroup1Type             `xml:"SavingsAndTempCashInvstGrp,omitempty" json:",omitempty"`
	PledgesAndGrantsReceivableGrp  *Form990PartXGroup1Type             `xml:"PledgesAndGrantsReceivableGrp,omitempty" json:",omitempty"`
	AccountsReceivableGrp          *Form990PartXGroup1Type             `xml:"AccountsReceivableGrp,omitempty" json:",omitempty"`
	ReceivablesFromOfficersEtcGrp  *Form990PartXGroup1Type             `xml:"ReceivablesFromOfficersEtcGrp,omitempty" json:",omitempty"`
	RcvblFromDisqualifiedPrsnGrp   *Form990PartXGroup1Type             `xml:"RcvblFromDisqualifiedPrsnGrp,omitempty" json:",omitempty"`
	OthNotesLoansReceivableNetGrp  *Form990PartXGroup1Type             `xml:"OthNotesLoansReceivableNetGrp,omitempty" json:",omitempty"`
	InventoriesForSaleOrUseGrp     *Form990PartXGroup1Type             `xml:"InventoriesForSaleOrUseGrp,omitempty" json:",omitempty"`
	PrepaidExpensesDefrdChargesGrp *Form990PartXGroup1Type             `xml:"PrepaidExpensesDefrdChargesGrp,omitempty" json:",omitempty"`
	LandBldgEquipCostOrOtherBssAmt int                                 `xml:"LandBldgEquipCostOrOtherBssAmt,omitempty" json:",omitempty"`
	LandBldgEquipAccumDeprecAmt    int                                 `xml:"LandBldgEquipAccumDeprecAmt,omitempty" json:",omitempty"`
	LandBldgEquipBasisNetGrp       *Form990PartXGroup1Type             `xml:"LandBldgEquipBasisNetGrp,omitempty" json:",omitempty"`
	InvestmentsPubTradedSecGrp     *Form990PartXGroup1Type             `xml:"InvestmentsPubTradedSecGrp,omitempty" json:",omitempty"`
	InvestmentsOtherSecuritiesGrp  *Form990PartXGroup1Type             `xml:"InvestmentsOtherSecuritiesGrp,omitempty" json:",omitempty"`
	InvestmentsProgramRelatedGrp   *Form990PartXGroup1Type             `xml:"InvestmentsProgramRelatedGrp,omitempty" json:",omitempty"`
	IntangibleAssetsGrp            *Form990PartXGroup1Type             `xml:"IntangibleAssetsGrp,omitempty" json:",omitempty"`
	OtherAssetsTotalGrp            *Form990PartXGroup1Type             `xml:"OtherAssetsTotalGrp,omitempty" json:",omitempty"`
	TotalAssetsGrp                 Form990PartXGroup2Type              `xml:"TotalAssetsGrp"`
	AccountsPayableAccrExpnssGrp   *Form990PartXGroup1Type             `xml:"AccountsPayableAccrExpnssGrp,omitempty" json:",omitempty"`
	GrantsPayableGrp               *Form990PartXGroup1Type             `xml:"GrantsPayableGrp,omitempty" json:",omitempty"`
	DeferredRevenueGrp             *Form990PartXGroup1Type             `xml:"DeferredRevenueGrp,omitempty" json:",omitempty"`
	TaxExemptBondLiabilitiesGrp    *Form990PartXGroup1Type             `xml:"TaxExemptBondLiabilitiesGrp,omitempty" json:",omitempty"`
	EscrowAccountLiabilityGrp      *Form990PartXGroup1Type             `xml:"EscrowAccountLiabilityGrp,omitempty" json:",omitempty"`
	LoansFromOfficersDirectorsGrp  *Form990PartXGroup1Type             `xml:"LoansFromOfficersDirectorsGrp,omitempty" json:",omitempty"`
	MortgNotesPyblScrdInvstPropGrp *Form990PartXGroup1Type             `xml:"MortgNotesPyblScrdInvstPropGrp,omitempty" json:",omitempty"`
	UnsecuredNotesLoansPayableGrp  *Form990PartXGroup1Type             `xml:"UnsecuredNotesLoansPayableGrp,omitempty" json:",omitempty"`
	OtherLiabilitiesGrp            *Form990PartXGroup1Type             `xml:"OtherLiabilitiesGrp,omitempty" json:",omitempty"`
	TotalLiabilitiesGrp            Form990PartXGroup2Type              `xml:"TotalLiabilitiesGrp"`
	OrganizationFollowsSFAS117Ind  CheckboxType                        `xml:"OrganizationFollowsSFAS117Ind,omitempty" json:",omitempty"`
	UnrestrictedNetAssetsGrp       *Form990PartXGroup1Type             `xml:"UnrestrictedNetAssetsGrp,omitempty" json:",omitempty"`
	TemporarilyRstrNetAssetsGrp    *Form990PartXGroup1Type             `xml:"TemporarilyRstrNetAssetsGrp,omitempty" json:",omitempty"`
	PermanentlyRstrNetAssetsGrp    *Form990PartXGroup1Type             `xml:"PermanentlyRstrNetAssetsGrp,omitempty" json:",omitempty"`
	OrgDoesNotFollowSFAS117Ind     CheckboxType                        `xml:"OrgDoesNotFollowSFAS117Ind,omitempty" json:",omitempty"`
	CapStkTrPrinCurrentFundsGrp    *Form990PartXGroup1Type             `xml:"CapStkTrPrinCurrentFundsGrp,omitempty" json:",omitempty"`
	PdInCapSrplsLandBldgEqpFundGrp *Form990PartXGroup1Type             `xml:"PdInCapSrplsLandBldgEqpFundGrp,omitempty" json:",omitempty"`
	RtnEarnEndowmentIncmOthFndsGrp *Form990PartXGroup1Type             `xml:"RtnEarnEndowmentIncmOthFndsGrp,omitempty" json:",omitempty"`
	TotalNetAssetsFundBalanceGrp   Form990PartXGroup2Type              `xml:"TotalNetAssetsFundBalanceGrp"`
	TotLiabNetAssetsFundBalanceGrp Form990PartXGroup2Type              `xml:"TotLiabNetAssetsFundBalanceGrp"`
	InfoInScheduleOPartXIInd       CheckboxType                        `xml:"InfoInScheduleOPartXIInd,omitempty" json:",omitempty"`
	ReconcilationRevenueExpnssAmt  int                                 `xml:"ReconcilationRevenueExpnssAmt"`
	NetUnrlzdGainsLossesInvstAmt   int                                 `xml:"NetUnrlzdGainsLossesInvstAmt,omitempty" json:",omitempty"`
	DonatedServicesAndUseFcltsAmt  int                                 `xml:"DonatedServicesAndUseFcltsAmt,omitempty" json:",omitempty"`
	InvestmentExpenseAmt           int                                 `xml:"InvestmentExpenseAmt,omitempty" json:",omitempty"`
	PriorPeriodAdjustmentsAmt      int                                 `xml:"PriorPeriodAdjustmentsAmt,omitempty" json:",omitempty"`
	OtherChangesInNetAssetsAmt     int                                 `xml:"OtherChangesInNetAssetsAmt,omitempty" json:",omitempty"`
	InfoInScheduleOPartXIIInd      CheckboxType                        `xml:"InfoInScheduleOPartXIIInd,omitempty" json:",omitempty"`
	MethodOfAccountingCashInd      CheckboxType                        `xml:"MethodOfAccountingCashInd"`
	MethodOfAccountingAccrualInd   CheckboxType                        `xml:"MethodOfAccountingAccrualInd"`
	MethodOfAccountingOtherInd     MethodOfAccountingOtherInd          `xml:"MethodOfAccountingOtherInd"`
	AccountantCompileOrReviewInd   bool                                `xml:"AccountantCompileOrReviewInd"`
	AcctCompileOrReviewBasisGrp    *FinancialStatementType             `xml:"AcctCompileOrReviewBasisGrp,omitempty" json:",omitempty"`
	FSAuditedInd                   bool                                `xml:"FSAuditedInd"`
	FSAuditedBasisGrp              *FinancialStatementType             `xml:"FSAuditedBasisGrp,omitempty" json:",omitempty"`
	AuditCommitteeInd              bool                                `xml:"AuditCommitteeInd,omitempty" json:",omitempty"`
	FederalGrantAuditRequiredInd   bool                                `xml:"FederalGrantAuditRequiredInd,omitempty" json:",omitempty"`
	FederalGrantAuditPerformedInd  bool                                `xml:"FederalGrantAuditPerformedInd,omitempty" json:",omitempty"`
}

func (r IRS990) Validate() error {
	return utils.Validate(&r)
}

func (r IRS990Type) Validate() error {
	return utils.Validate(&r)
}
