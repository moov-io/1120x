// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"time"

	"github.com/moov-io/1120x/pkg/utils"
)

// Country Code Type Including 'US'
type AllCountriesType string

func (r AllCountriesType) Validate() error {
	return utils.Validate(&r)
}

// May be one of All States
type AllStatesCd string

func (r AllStatesCd) Validate() error {
	return utils.Validate(&r)
}

// Must match the pattern [A-Za-z0-9\(\)]*
type AlphaNumericAndParenthesesType string

func (r AlphaNumericAndParenthesesType) Validate() error {
	return utils.Validate(&r)
}

func (r AlphaNumericAndParenthesesType) IsValid() error {
	reg := regexp.MustCompile(`[A-Za-z0-9\(\)]*`)
	if !reg.MatchString(string(r)) {
		return errors.New("AlphaNumericAndParenthesesType is invalid")
	}
	return nil
}

// Must match the pattern [A-Za-z0-9]*
type AlphaNumericType string

func (r AlphaNumericType) Validate() error {
	return utils.Validate(&r)
}

func (r AlphaNumericType) IsValid() error {
	reg := regexp.MustCompile(`[A-Za-z0-9]*`)
	if !reg.MatchString(string(r)) {
		return errors.New("AlphaNumericType is invalid")
	}
	return nil
}

// Must match the pattern [A-Za-z0-9\-]+
type BankAccountNumberType string

func (r BankAccountNumberType) Validate() error {
	return utils.Validate(&r)
}

func (r BankAccountNumberType) IsValid() error {
	reg := regexp.MustCompile(`[A-Za-z0-9\-]+`)
	if !reg.MatchString(string(r)) {
		return errors.New("BankAccountNumberType is invalid")
	}
	return nil
}

// May be one of 1, 2
type BankAccountType string

func (r BankAccountType) Validate() error {
	return utils.Validate(&r)
}

func (r BankAccountType) IsValid() error {
	for _, vv := range []string{
		"1", "2",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("BankAccountType is invalid")
}

// Must match the pattern [A-D]
type BondReferenceCd string

func (r BondReferenceCd) Validate() error {
	return utils.Validate(&r)
}

func (r BondReferenceCd) IsValid() error {
	reg := regexp.MustCompile(`[A-D]`)
	if !reg.MatchString(string(r)) {
		return errors.New("BondReferenceCd is invalid")
	}
	return nil
}

// Must match the pattern [0-9]{6}
type BusinessCd string

func (r BusinessCd) Validate() error {
	return utils.Validate(&r)
}

func (r BusinessCd) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{6}`)
	if !reg.MatchString(string(r)) {
		return errors.New("BusinessCd is invalid")
	}
	return nil
}

// Must match the pattern ([A-Z0-9\-]|&){1,4}
type BusinessNameControlType string

func (r BusinessNameControlType) Validate() error {
	return utils.Validate(&r)
}

func (r BusinessNameControlType) IsValid() error {
	reg := regexp.MustCompile(`([A-Z0-9\-]|&){1,4}`)
	if !reg.MatchString(string(r)) {
		return errors.New("BusinessNameControlType is invalid")
	}
	return nil
}

// Must match the pattern (([A-Za-z0-9#\-\(\)]|&|') ?)*([A-Za-z0-9#\-\(\)]|&|')
type BusinessNameLine1Type string

func (r BusinessNameLine1Type) Validate() error {
	return utils.Validate(&r)
}

func (r BusinessNameLine1Type) IsValid() error {
	reg := regexp.MustCompile(`(([A-Za-z0-9#\-\(\)]|&|') ?)*([A-Za-z0-9#\-\(\)]|&|')`)
	if !reg.MatchString(string(r)) {
		return errors.New("BusinessNameLine1Type is invalid")
	}
	return nil
}

// Must match the pattern (([A-Za-z0-9#/%\-\(\)]|&|') ?)*([A-Za-z0-9#/%\-\(\)]|&|')
type BusinessNameLine2Type string

func (r BusinessNameLine2Type) Validate() error {
	return utils.Validate(&r)
}

func (r BusinessNameLine2Type) IsValid() error {
	reg := regexp.MustCompile(`(([A-Za-z0-9#/%\-\(\)]|&|') ?)*([A-Za-z0-9#/%\-\(\)]|&|')`)
	if !reg.MatchString(string(r)) {
		return errors.New("BusinessNameLine2Type is invalid")
	}
	return nil
}

// Must match the pattern [0-9]{2}
type CHNAConductedYr int

func (r CHNAConductedYr) Validate() error {
	return utils.Validate(&r)
}

func (r CHNAConductedYr) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{2}`)
	if !reg.MatchString(fmt.Sprint(r)) {
		return errors.New("CHNAConductedYr is invalid")
	}
	return nil
}

// May be no more than 9 items long
type CUSIPNumberType string

func (r CUSIPNumberType) Validate() error {
	return utils.Validate(&r)
}

func (r CUSIPNumberType) IsValid() error {
	if len(r) > 9 {
		return errors.New("CUSIPNumberType is invalid")
	}
	return nil
}

// Must match the pattern [A-Z]{2}
type CheckDigitType string

func (r CheckDigitType) Validate() error {
	return utils.Validate(&r)
}

func (r CheckDigitType) IsValid() error {
	reg := regexp.MustCompile(`[A-Z]{2}`)
	if !reg.MatchString(string(r)) {
		return errors.New("CheckDigitType is invalid")
	}
	return nil
}

// May be one of X
type CheckboxType string

func (r CheckboxType) Validate() error {
	return utils.Validate(&r)
}

// Must match the pattern ([A-Za-z] ?)*[A-Za-z]
type CityType string

func (r CityType) Validate() error {
	return utils.Validate(&r)
}

func (r CityType) IsValid() error {
	reg := regexp.MustCompile(`([A-Za-z] ?)*[A-Za-z]`)
	if !reg.MatchString(string(r)) {
		return errors.New("CityType is invalid")
	}
	return nil
}

// May be one of English Free-File, Spanish Free-File, Free Fillable Forms, Free File VITA
type ConsortiumType string

func (r ConsortiumType) Validate() error {
	return utils.Validate(&r)
}

func (r ConsortiumType) IsValid() error {
	for _, vv := range []string{
		"English Free-File", "Spanish Free-File", "Free Fillable Forms", "Free File VITA",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("ConsortiumType is invalid")
}

// May be one of
// AF, AX, AL, AG, AQ, AN, AO, AV, AY, AC, AR, AM, AA, AT, AS, AU, AJ, BF, BA, FQ, BG, BB, BO, BE, BH, BN, BD, BT, BL,
// BK, BC, BV, BR, IO, VI, BX, BU, UV, BM, BY, CB, CM, CA, CV, CJ, CT, CD, CI, CH, KT, IP, CK, CO, CN, CF, CG, CW, CR,
// CS, IV, HR, CU, UC, CY, EZ, DA, DX, DJ, DO, DR, TT, EC, EG, ES, EK, ER, EN, ET, FK, FO, FM, FJ, FI, FR, FP, FS, GB,
// GA, GG, GM, GH, GI, GR, GL, GJ, GQ, GT, GK, GV, PU, GY, HA, HM, VT, HO, HK, HQ, HU, IC, IN, ID, IR, IZ, EI, IS, IT,
// JM, JN, JA, DQ, JE, JQ, JO, KZ, KE, KQ, KR, KN, KS, KV, KU, KG, LA, LG, LE, LT, LI, LY, LS, LH, LU, MC, MK, MA, MI,
// MY, MV, ML, MT, IM, RM, MR, MP, MX, MQ, MD, MN, MG, MJ, MH, MO, MZ, WA, NR, BQ, NP, NL, NC, NZ, NU, NG, NI, NE, NF,
// CQ, NO, MU, OC, PK, PS, LQ, PM, PP, PF, PA, PE, RP, PC, PL, PO, RQ, QA, RO, RS, RW, TB, RN, WS, SM, TP, SA, SG, RI,
// SE, SL, SN, NN, LO, SI, BP, SO, SF, SX, OD, SP, PG, CE, SH, SC, ST, SB, VC, SU, NS, SV, WZ, SW, SZ, SY, TW, TI, TZ,
// TH, TO, TL, TN, TD, TS, TU, TX, TK, TV, UG, UP, AE, UK, UY, UZ, NH, VE, VM, VQ, WQ, WF, WI, YM, ZA, ZI
type CountryType string

func (r CountryType) Validate() error {
	return utils.Validate(&r)
}

func (r CountryType) IsValid() error {
	for _, vv := range []string{
		"AF", "AX", "AL", "AG", "AQ", "AN", "AO", "AV", "AY", "AC", "AR", "AM", "AA", "AT", "AS", "AU", "AJ", "BF",
		"BA", "FQ", "BG", "BB", "BO", "BE", "BH", "BN", "BD", "BT", "BL", "BK", "BC", "BV", "BR", "IO", "VI", "BX",
		"BU", "UV", "BM", "BY", "CB", "CM", "CA", "CV", "CJ", "CT", "CD", "CI", "CH", "KT", "IP", "CK", "CO", "CN",
		"CF", "CG", "CW", "CR", "CS", "IV", "HR", "CU", "UC", "CY", "EZ", "DA", "DX", "DJ", "DO", "DR", "TT", "EC",
		"EG", "ES", "EK", "ER", "EN", "ET", "FK", "FO", "FM", "FJ", "FI", "FR", "FP", "FS", "GB", "GA", "GG", "GM",
		"GH", "GI", "GR", "GL", "GJ", "GQ", "GT", "GK", "GV", "PU", "GY", "HA", "HM", "VT", "HO", "HK", "HQ", "HU",
		"IC", "IN", "ID", "IR", "IZ", "EI", "IS", "IT", "JM", "JN", "JA", "DQ", "JE", "JQ", "JO", "KZ", "KE", "KQ",
		"KR", "KN", "KS", "KV", "KU", "KG", "LA", "LG", "LE", "LT", "LI", "LY", "LS", "LH", "LU", "MC", "MK", "MA",
		"MI", "MY", "MV", "ML", "MT", "IM", "RM", "MR", "MP", "MX", "MQ", "MD", "MN", "MG", "MJ", "MH", "MO", "MZ",
		"WA", "NR", "BQ", "NP", "NL", "NC", "NZ", "NU", "NG", "NI", "NE", "NF", "CQ", "NO", "MU", "OC", "PK", "PS",
		"LQ", "PM", "PP", "PF", "PA", "PE", "RP", "PC", "PL", "PO", "RQ", "QA", "RO", "RS", "RW", "TB", "RN", "WS",
		"SM", "TP", "SA", "SG", "RI", "SE", "SL", "SN", "NN", "LO", "SI", "BP", "SO", "SF", "SX", "OD", "SP", "PG",
		"CE", "SH", "SC", "ST", "SB", "VC", "SU", "NS", "SV", "WZ", "SW", "SZ", "SY", "TW", "TI", "TZ", "TH", "TO",
		"TL", "TN", "TD", "TS", "TU", "TX", "TK", "TV", "UG", "UP", "AE", "UK", "UY", "UZ", "NH", "VE", "VM", "VQ",
		"WQ", "WF", "WI", "YM", "ZA", "ZI",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("CountryType is invalid")
}

// Base type for a date
type DateType time.Time

func (r DateType) Validate() error {
	return utils.Validate(&r)
}

func (r *DateType) UnmarshalText(text []byte) error {
	return (*xsdDate)(r).UnmarshalText(text)
}
func (r DateType) MarshalText() ([]byte, error) {
	return xsdDate(r).MarshalText()
}

// 2-digit decimal typically used by a non-negative decimal amount field.
type DecimalNNType float64

func (r DecimalNNType) Validate() error {
	return utils.Validate(&r)
}

// 2-digit decimal typically used by a decimal amount field.
type DecimalType float64

func (r DecimalType) Validate() error {
	return utils.Validate(&r)
}

// May be one of HY, MQ, MM, S/L
type DepreciationConventionCodeType string

func (r DepreciationConventionCodeType) Validate() error {
	return utils.Validate(&r)
}

func (r DepreciationConventionCodeType) IsValid() error {
	for _, vv := range []string{
		"HY", "MQ", "MM", "S/L",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("DepreciationConventionCodeType is invalid")
}

// May be one of 200 DB, 150 DB, DB, S/L, Various
type DepreciationMethodCodeType string

func (r DepreciationMethodCodeType) Validate() error {
	return utils.Validate(&r)
}

func (r DepreciationMethodCodeType) IsValid() error {
	for _, vv := range []string{
		"200 DB", "150 DB", "DB", "S/L", "Various",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("DepreciationMethodCodeType is invalid")
}

// Must match the pattern [A-Fa-f0-9]{40}
type DeviceIdType string

func (r DeviceIdType) Validate() error {
	return utils.Validate(&r)
}

func (r DeviceIdType) IsValid() error {
	reg := regexp.MustCompile(`[A-Fa-f0-9]{40}`)
	if !reg.MatchString(string(r)) {
		return errors.New("DeviceIdType is invalid")
	}
	return nil
}

// May be one of N/A
type DirectControllingNACd string

func (r DirectControllingNACd) Validate() error {
	return utils.Validate(&r)
}

// May be one of PDF
type DocumentTypeCd string

func (r DocumentTypeCd) Validate() error {
	return utils.Validate(&r)
}

// Must match the pattern [0-9]{6}
type EFINType string

func (r EFINType) Validate() error {
	return utils.Validate(&r)
}

func (r EFINType) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{6}`)
	if !reg.MatchString(string(r)) {
		return errors.New("EFINType is invalid")
	}
	return nil
}

// Must match the pattern [0-9]{9}
type EINType string

func (r EINType) Validate() error {
	return utils.Validate(&r)
}

func (r EINType) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{9}`)
	if !reg.MatchString(string(r)) {
		return errors.New("EINType is invalid")
	}
	return nil
}

// Must match the pattern [0-9]{5}
type ETINType string

func (r ETINType) Validate() error {
	return utils.Validate(&r)
}

func (r ETINType) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{5}`)
	if !reg.MatchString(string(r)) {
		return errors.New("ETINType is invalid")
	}
	return nil
}

// May be one of AL, AK, AZ, AR, CA, CO, CT, DE, DC, FL, GA, HI, ID, IL, IN, IA, KS, KY, LA, ME, MD, MA, MI, MN, MS, MO,
// MT, NE, NV, NH, NJ, NM, NY, NC, ND, OH, OK, OR, PA, PR, RI, SC, SD, TN, TX, VI, UT, VT, VA, WA, WV, WI, WY
type FUTAStateCdType string

func (r FUTAStateCdType) Validate() error {
	return utils.Validate(&r)
}

func (r FUTAStateCdType) IsValid() error {
	for _, vv := range []string{
		"AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY",
		"LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH",
		"OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "VI", "UT", "VT", "VA", "WA", "WV", "WI", "WY",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("FUTAStateCdType is invalid")
}

// Must be at least 1 items long
type ForeignEntityReferenceIdNum string

func (r ForeignEntityReferenceIdNum) Validate() error {
	return utils.Validate(&r)
}

func (r ForeignEntityReferenceIdNum) IsValid() error {
	if len(r) > 1 {
		return errors.New("ForeignEntityReferenceIdNum is invalid")
	}
	return nil
}

// Must match the pattern [0-9]{1,30}
type ForeignPhoneNumberType string

func (r ForeignPhoneNumberType) Validate() error {
	return utils.Validate(&r)
}

func (r ForeignPhoneNumberType) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{1,30}`)
	if !reg.MatchString(string(r)) {
		return errors.New("ForeignPhoneNumberType is invalid")
	}
	return nil
}

// May be one of RIC
type ForeignRegulatedInvestmtCompCdType string

func (r ForeignRegulatedInvestmtCompCdType) Validate() error {
	return utils.Validate(&r)
}

// May be one of
// IRS, ALST, AKST, AZST, ARST, CAST, CNCT, COCT, COST, CTST, DECT, DEST, DCST, FLST, GAST, HIST, IDST, ILST, INST,
// IAST, KACT, KSST, KYST, LAST, LECT, LXCT, MEST, MDST, MAST, MIST, MNST, MSST, MOST, MTST, NEST, NVST, NHST, NJST,
// NMST, NYST, NCST, NDST, OCCT, OHST, OKST, ORCT, ORST, PAST, RIST, SCST, SDST, SLCT, TNST, TOCT, TXST, UTST, VTST,
// VAST, WAST, WVST, WIST, WYST, NYCT, PHCT
type GovernmentCodeType string

func (r GovernmentCodeType) Validate() error {
	return utils.Validate(&r)
}

func (r GovernmentCodeType) IsValid() error {
	for _, vv := range []string{
		"IRS", "ALST", "AKST", "AZST", "ARST", "CAST", "CNCT", "COCT", "COST", "CTST", "DECT", "DEST", "DCST", "FLST",
		"GAST", "HIST", "IDST", "ILST", "INST", "IAST", "KACT", "KSST", "KYST", "LAST", "LECT", "LXCT", "MEST", "MDST",
		"MAST", "MIST", "MNST", "MSST", "MOST", "MTST", "NEST", "NVST", "NHST", "NJST", "NMST", "NYST", "NCST", "NDST",
		"OCCT", "OHST", "OKST", "ORCT", "ORST", "PAST", "RIST", "SCST", "SDST", "SLCT", "TNST", "TOCT", "TXST", "UTST",
		"VTST", "VAST", "WAST", "WVST", "WIST", "WYST", "NYCT", "PHCT",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("GovernmentCodeType is invalid")
}

// Must match the pattern \d{4}
type GroupExemptionNum string

func (r GroupExemptionNum) Validate() error {
	return utils.Validate(&r)
}

func (r GroupExemptionNum) IsValid() error {
	reg := regexp.MustCompile(`\d{4}`)
	if !reg.MatchString(string(r)) {
		return errors.New("GroupExemptionNum is invalid")
	}
	return nil
}

// Must match the pattern [0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}
type IPv4Type string

func (r IPv4Type) Validate() error {
	return utils.Validate(&r)
}

func (r IPv4Type) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`)
	if !reg.MatchString(string(r)) {
		return errors.New("IPv4Type is invalid")
	}
	return nil
}

// Must match the pattern
// [0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}
type IPv6Type string

func (r IPv6Type) Validate() error {
	return utils.Validate(&r)
}

func (r IPv6Type) IsValid() error {
	reg := regexp.MustCompile(`[0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}:[0-9A-F]{1,4}`)
	if !reg.MatchString(string(r)) {
		return errors.New("IPv6Type is invalid")
	}
	return nil
}

// Must match the pattern efile|.*
type IRSServiceCenterType string

func (r IRSServiceCenterType) Validate() error {
	return utils.Validate(&r)
}

func (r IRSServiceCenterType) IsValid() error {
	reg := regexp.MustCompile(`efile|.*`)
	if !reg.MatchString(string(r)) {
		return errors.New("IRSServiceCenterType is invalid")
	}
	return nil
}

// Must match the pattern [A-Z0-9]{6}
type ISPType string

func (r ISPType) Validate() error {
	return utils.Validate(&r)
}

func (r ISPType) IsValid() error {
	reg := regexp.MustCompile(`[A-Z0-9]{6}`)
	if !reg.MatchString(string(r)) {
		return errors.New("ISPType is invalid")
	}
	return nil
}

// Must match the pattern [A-Za-z0-9:\.\-]{1,30}
type IdType string

func (r IdType) Validate() error {
	return utils.Validate(&r)
}

func (r IdType) IsValid() error {
	reg := regexp.MustCompile(`[A-Za-z0-9:\.\-]{1,30}`)
	if !reg.MatchString(string(r)) {
		return errors.New("IdType is invalid")
	}
	return nil
}

// Must match the pattern [0-9]{2}
type ImplementationStrategyAdptYr int

func (r ImplementationStrategyAdptYr) Validate() error {
	return utils.Validate(&r)
}

func (r ImplementationStrategyAdptYr) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{2}`)
	if !reg.MatchString(fmt.Sprint(r)) {
		return errors.New("ImplementationStrategyAdptYr is invalid")
	}
	return nil
}

// Must match the pattern (% )(([A-Za-z0-9#/%\-\(\)]|&|') ?)*([A-Za-z0-9#/%\-\(\)]|&|')
type InCareOfNameType string

func (r InCareOfNameType) Validate() error {
	return utils.Validate(&r)
}

func (r InCareOfNameType) IsValid() error {
	reg := regexp.MustCompile(`(% )(([A-Za-z0-9#/%\-\(\)]|&|') ?)*([A-Za-z0-9#/%\-\(\)]|&|')`)
	if !reg.MatchString(string(r)) {
		return errors.New("InCareOfNameType is invalid")
	}
	return nil
}

// Base type for a non-negative integer
type IntegerNNType int

func (r IntegerNNType) Validate() error {
	return utils.Validate(&r)
}

func (r IntegerNNType) IsValid() error {
	if r < 0 {
		return errors.New("IntegerNNType is invalid")
	}
	return nil
}

// Base type for an integer
type IntegerType int

func (r IntegerType) Validate() error {
	return utils.Validate(&r)
}

// May be one of cash, accrual, hybrid
type MethodOfAccountingType string

func (r MethodOfAccountingType) Validate() error {
	return utils.Validate(&r)
}

func (r MethodOfAccountingType) IsValid() error {
	for _, vv := range []string{
		"cash", "accrual", "hybrid",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("MethodOfAccountingType is invalid")
}

// May be one of C, F
type MethodValuationCd string

func (r MethodValuationCd) Validate() error {
	return utils.Validate(&r)
}

func (r MethodValuationCd) IsValid() error {
	for _, vv := range []string{
		"C", "F",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("MethodValuationCd is invalid")
}

// Month and day type in the format of --MM-DD
type MonthDayType time.Time

func (r MonthDayType) Validate() error {
	return utils.Validate(&r)
}

func (r *MonthDayType) UnmarshalText(text []byte) error {
	return (*xsdGMonthDay)(r).UnmarshalText(text)
}

func (r MonthDayType) MarshalText() ([]byte, error) {
	return xsdGMonthDay(r).MarshalText()
}

// Month type in the format of --MM
type MonthType time.Time

func (r MonthType) Validate() error {
	return utils.Validate(&r)
}

func (r *MonthType) UnmarshalText(text []byte) error {
	return (*xsdGMonth)(r).UnmarshalText(text)
}

func (r MonthType) MarshalText() ([]byte, error) {
	return xsdGMonth(r).MarshalText()
}

// Must match the pattern [A-Za-z]( |<)?(([A-Za-z#\-]|&)( |<)?)*([A-Za-z#\-]|&)
type NameLine1Type string

func (r NameLine1Type) Validate() error {
	return utils.Validate(&r)
}

func (r NameLine1Type) IsValid() error {
	reg := regexp.MustCompile(`[A-Za-z]( |<)?(([A-Za-z#\-]|&)( |<)?)*([A-Za-z#\-]|&)`)
	if !reg.MatchString(string(r)) {
		return errors.New("NameLine1Type is invalid")
	}
	return nil
}

// Must match the pattern [0-9]*
type NumericType string

func (r NumericType) Validate() error {
	return utils.Validate(&r)
}

func (r NumericType) IsValid() error {
	reg := regexp.MustCompile(`[0-9]*`)
	if !reg.MatchString(string(r)) {
		return errors.New("NumericType is invalid")
	}
	return nil
}

// Must match the pattern [2-9]|1[0-9]|2[02-7]
type Organization501cTypeTxt string

func (r Organization501cTypeTxt) Validate() error {
	return utils.Validate(&r)
}

func (r Organization501cTypeTxt) IsValid() error {
	reg := regexp.MustCompile(`[2-9]|1[0-9]|2[02-7]`)
	if !reg.MatchString(string(r)) {
		return errors.New("Organization501cTypeTxt is invalid")
	}
	return nil
}

// May be one of 1, 2, 3, 4, 5, 6, 7, 8, 9
type OrganizationTypeCd string

func (r OrganizationTypeCd) Validate() error {
	return utils.Validate(&r)
}

func (r OrganizationTypeCd) IsValid() error {
	for _, vv := range []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("OrganizationTypeCd is invalid")
}

// May be one of ERO, OnlineFiler, ReportingAgent, IRSAgent, FinancialAgent, LargeTaxpayer
type OriginatorType string

func (r OriginatorType) Validate() error {
	return utils.Validate(&r)
}

func (r OriginatorType) IsValid() error {
	for _, vv := range []string{
		"ERO", "OnlineFiler", "ReportingAgent", "IRSAgent", "FinancialAgent", "LargeTaxpayer",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("OriginatorType is invalid")
}

// May be one of Practitioner, Self-Select Practitioner, Self-Select On-Line
type PINCodeType string

func (r PINCodeType) Validate() error {
	return utils.Validate(&r)
}

func (r PINCodeType) IsValid() error {
	for _, vv := range []string{
		"Practitioner", "Self-Select Practitioner", "Self-Select On-Line",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("PINCodeType is invalid")
}

// May be one of Taxpayer, ERO
type PINEnteredByCd string

func (r PINEnteredByCd) Validate() error {
	return utils.Validate(&r)
}

func (r PINEnteredByCd) IsValid() error {
	for _, vv := range []string{
		"Taxpayer", "ERO",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("PINEnteredByCd is invalid")
}

// May be one of Taxpayer, ERO
type PINEnteredByType string

func (r PINEnteredByType) Validate() error {
	return utils.Validate(&r)
}

func (r PINEnteredByType) IsValid() error {
	for _, vv := range []string{
		"Taxpayer", "ERO",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("PINEnteredByType is invalid")
}

// Must match the pattern [0-9]{5}
type PINType string

func (r PINType) Validate() error {
	return utils.Validate(&r)
}

func (r PINType) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{5}`)
	if !reg.MatchString(string(r)) {
		return errors.New("PINType is invalid")
	}
	return nil
}

// Must match the pattern P[0-9]{8}
type PTINType string

func (r PTINType) Validate() error {
	return utils.Validate(&r)
}

func (r PTINType) IsValid() error {
	reg := regexp.MustCompile(`P[0-9]{8}`)
	if !reg.MatchString(string(r)) {
		return errors.New("PTINType is invalid")
	}
	return nil
}

// May be one of Partners Page English, Partners Page Spanish
type PartnersPageFilingType string

func (r PartnersPageFilingType) Validate() error {
	return utils.Validate(&r)
}

// Must match the pattern ([A-Za-z\-] ?)*[A-Za-z\-]
type PersonFirstNameType string

func (r PersonFirstNameType) Validate() error {
	return utils.Validate(&r)
}

func (r PersonFirstNameType) IsValid() error {
	reg := regexp.MustCompile(`([A-Za-z\-] ?)*[A-Za-z\-]`)
	if !reg.MatchString(string(r)) {
		return errors.New("PersonFirstNameType is invalid")
	}
	return nil
}

// Must match the pattern ([A-Za-z\-] ?)*[A-Za-z\-]
type PersonLastNameType string

func (r PersonLastNameType) Validate() error {
	return utils.Validate(&r)
}

func (r PersonLastNameType) IsValid() error {
	reg := regexp.MustCompile(`([A-Za-z\-] ?)*[A-Za-z\-]`)
	if !reg.MatchString(string(r)) {
		return errors.New("PersonLastNameType is invalid")
	}
	return nil
}

// Must match the pattern [A-Z][A-Z\- ]{0,3}
type PersonNameControlType string

func (r PersonNameControlType) Validate() error {
	return utils.Validate(&r)
}

func (r PersonNameControlType) IsValid() error {
	reg := regexp.MustCompile(`[A-Z][A-Z\- ]{0,3}`)
	if !reg.MatchString(string(r)) {
		return errors.New("PersonNameControlType is invalid")
	}
	return nil
}

// Must match the pattern ([A-Za-z0-9'\-] ?)*[A-Za-z0-9'\-]
type PersonNameType string

func (r PersonNameType) Validate() error {
	return utils.Validate(&r)
}

func (r PersonNameType) IsValid() error {
	reg := regexp.MustCompile(`([A-Za-z0-9'\-] ?)*[A-Za-z0-9'\-]`)
	if !reg.MatchString(string(r)) {
		return errors.New("PersonNameType is invalid")
	}
	return nil
}

// Must match the pattern ([!-~] ?)*[!-~]
type PersonTitleType string

func (r PersonTitleType) Validate() error {
	return utils.Validate(&r)
}

func (r PersonTitleType) IsValid() error {
	reg := regexp.MustCompile(`([!-~] ?)*[!-~]`)
	if !reg.MatchString(string(r)) {
		return errors.New("PersonTitleType is invalid")
	}
	return nil
}

// Must match the pattern [0-9]{10}
type PhoneNumberType string

func (r PhoneNumberType) Validate() error {
	return utils.Validate(&r)
}

func (r PhoneNumberType) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{10}`)
	if !reg.MatchString(string(r)) {
		return errors.New("PhoneNumberType is invalid")
	}
	return nil
}

// The end date of a calendar quarter.
type QuarterEndDateType time.Time

func (t QuarterEndDateType) Validate() error {
	return utils.Validate(&t)
}

func (t *QuarterEndDateType) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t QuarterEndDateType) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

// A fraction between 0 and 1 that allows up to 5 decimal places
type RatioType float64

func (r RatioType) Validate() error {
	return utils.Validate(&r)
}

// Must match the pattern [A-Z0-9]{1,20}
type RegistrationNumType string

func (r RegistrationNumType) Validate() error {
	return utils.Validate(&r)
}

func (r RegistrationNumType) IsValid() error {
	reg := regexp.MustCompile(`[A-Z0-9]{1,20}`)
	if !reg.MatchString(string(r)) {
		return errors.New("RegistrationNumType is invalid")
	}
	return nil
}

// May be one of 990, 990EZ, 990PF
type ReturnTypeCd string

func (r ReturnTypeCd) Validate() error {
	return utils.Validate(&r)
}

func (r ReturnTypeCd) IsValid() error {
	for _, vv := range []string{
		"990", "990EZ", "990PF",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("ReturnTypeCd is invalid")
}

// Must match the pattern (01|02|03|04|05|06|07|08|09|10|11|12|21|22|23|24|25|26|27|28|29|30|31|32)[0-9]{7}
type RoutingTransitNumberType string

func (r RoutingTransitNumberType) Validate() error {
	return utils.Validate(&r)
}

func (r RoutingTransitNumberType) IsValid() error {
	reg := regexp.MustCompile(`(01|02|03|04|05|06|07|08|09|10|11|12|21|22|23|24|25|26|27|28|29|30|31|32)[0-9]{7}`)
	if !reg.MatchString(string(r)) {
		return errors.New("RoutingTransitNumberType is invalid")
	}
	return nil
}

// Must match the pattern [0-9]{9}
type SSNType string

func (r SSNType) Validate() error {
	return utils.Validate(&r)
}

func (r SSNType) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{9}`)
	if !reg.MatchString(string(r)) {
		return errors.New("SSNType is invalid")
	}
	return nil
}

// Must match the pattern S[0-9]{8}
type STINType string

func (r STINType) Validate() error {
	return utils.Validate(&r)
}

func (r STINType) IsValid() error {
	reg := regexp.MustCompile(`S[0-9]{8}`)
	if !reg.MatchString(string(r)) {
		return errors.New("STINType is invalid")
	}
	return nil
}

// May be one of PIN Number, Binary Attachment 8453 Signature Document
type SignatureOptionCd string

func (r SignatureOptionCd) Validate() error {
	return utils.Validate(&r)
}

// Must match the pattern [0-9]{10}
type SignatureType string

func (r SignatureType) Validate() error {
	return utils.Validate(&r)
}

func (r SignatureType) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{10}`)
	if !reg.MatchString(string(r)) {
		return errors.New("SignatureType is invalid")
	}
	return nil
}

// Must match the pattern [0-9]{8}
type SoftwareIdType string

func (r SoftwareIdType) Validate() error {
	return utils.Validate(&r)
}

// May be one of
// AL, AK, AS, AZ, AR, CA, CO, MP, CT, DE, DC, FM, FL, GA, GU, HI, ID, IL, IN, IA, KS, KY, LA, ME, MH, MD, MA, MI, MN,
// MS, MO, MT, NE, NV, NH, NJ, NM, NY, NC, ND, OH, OK, OR, PW, PA, PR, RI, SC, SD, TN, TX, VI, UT, VT, VA, WA, WV, WI,
// WY, AA, AE, AP
type StateType string

func (r StateType) Validate() error {
	return utils.Validate(&r)
}

func (r StateType) IsValid() error {
	for _, vv := range []string{
		"AL", "AK", "AS", "AZ", "AR", "CA", "CO", "MP", "CT", "DE", "DC", "FM", "FL", "GA", "GU", "HI", "ID", "IL",
		"IN", "IA", "KS", "KY", "LA", "ME", "MH", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ",
		"NM", "NY", "NC", "ND", "OH", "OK", "OR", "PW", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "VI", "UT", "VT",
		"VA", "WA", "WV", "WI", "WY", "AA", "AE", "AP",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("StateType is invalid")
}

// Must match the pattern [A-Za-z0-9]( ?[A-Za-z0-9\-/])*
type StreetAddressType string

func (r StreetAddressType) Validate() error {
	return utils.Validate(&r)
}

func (r StreetAddressType) IsValid() error {
	reg := regexp.MustCompile(`[A-Za-z0-9]( ?[A-Za-z0-9\-/])*`)
	if !reg.MatchString(string(r)) {
		return errors.New("StreetAddressType is invalid")
	}
	return nil
}

// Base type for a string
type StringType string

func (r StringType) Validate() error {
	return utils.Validate(&r)
}

// May be one of VARIOUS
type StringVARIOUSType string

func (r StringVARIOUSType) Validate() error {
	return utils.Validate(&r)
}

// Must match the pattern (MA[0-9]{7})|([0-9]{11})
type TaxShelterRegistrationType string

func (r TaxShelterRegistrationType) Validate() error {
	return utils.Validate(&r)
}

func (r TaxShelterRegistrationType) IsValid() error {
	reg := regexp.MustCompile(`(MA[0-9]{7})|([0-9]{11})`)
	if !reg.MatchString(string(r)) {
		return errors.New("TaxShelterRegistrationType is invalid")
	}
	return nil
}

// Must match the pattern [0-9][0-9](01|02|03|04|05|06|07|08|09|10|11|12)
type TaxYearEndMonthDtType string

func (r TaxYearEndMonthDtType) Validate() error {
	return utils.Validate(&r)
}

func (r TaxYearEndMonthDtType) IsValid() error {
	reg := regexp.MustCompile(`[0-9][0-9](01|02|03|04|05|06|07|08|09|10|11|12)`)
	if !reg.MatchString(string(r)) {
		return errors.New("TaxYearEndMonthDtType is invalid")
	}
	return nil
}

// Must match the pattern ([!-~£§ÁÉÍÑÓ×ÚÜáéíñóúü] ?)*[!-~£§ÁÉÍÑÓ×ÚÜáéíñóúü]
type TextType string

func (r TextType) Validate() error {
	return utils.Validate(&r)
}

// Base type for time
type TimeType time.Time

func (t TimeType) Validate() error {
	return utils.Validate(&t)
}

func (t *TimeType) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t TimeType) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

// Base type for a date and time stamp
type TimestampType time.Time

func (t TimestampType) Validate() error {
	return utils.Validate(&t)
}

func (t *TimestampType) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t TimestampType) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// May be one of US, ES, ED, CS, CD, MS, MD, PS, PD, AS, AD, HS, HD
type TimezoneType string

func (r TimezoneType) Validate() error {
	return utils.Validate(&r)
}

func (r TimezoneType) IsValid() error {
	for _, vv := range []string{
		"US", "ES", "ED", "CS", "CD", "MS", "MD", "PS", "PD", "AS", "AD", "HS", "HD",
	} {
		if reflect.DeepEqual(r, vv) {
			return nil
		}
	}
	return errors.New("TimezoneType is invalid")
}

// Must match the pattern [A-HJ-NPR-Z0-9]{1,17}|[A-HJ-NPR-Z0-9]{19}
type VINType string

func (r VINType) Validate() error {
	return utils.Validate(&r)
}

func (r VINType) IsValid() error {
	reg := regexp.MustCompile(`[A-HJ-NPR-Z0-9]{1,17}|[A-HJ-NPR-Z0-9]{19}`)
	if !reg.MatchString(string(r)) {
		return errors.New("VINType is invalid")
	}
	return nil
}

// Year and month type in the format of YYYY-MM
type YearMonthType time.Time

func (r YearMonthType) Validate() error {
	return utils.Validate(&r)
}

func (r *YearMonthType) UnmarshalText(text []byte) error {
	return (*xsdGYearMonth)(r).UnmarshalText(text)
}
func (r YearMonthType) MarshalText() ([]byte, error) {
	return xsdGYearMonth(r).MarshalText()
}

// Base type for a 4-digit year
type YearType time.Time

func (r YearType) Validate() error {
	return utils.Validate(&r)
}

func (r *YearType) UnmarshalText(text []byte) error {
	return (*xsdGYear)(r).UnmarshalText(text)
}
func (r YearType) MarshalText() ([]byte, error) {
	return xsdGYear(r).MarshalText()
}

// Must match the pattern [0-9]{5}(([0-9]{4})|([0-9]{7}))?
type ZIPCodeType string

func (r ZIPCodeType) Validate() error {
	return utils.Validate(&r)
}

func (r ZIPCodeType) IsValid() error {
	reg := regexp.MustCompile(`[0-9]{5}(([0-9]{4})|([0-9]{7}))?`)
	if !reg.MatchString(string(r)) {
		return errors.New("ZIPCodeType is invalid")
	}
	return nil
}

type IdListType []string

func (x *IdListType) MarshalText() ([]byte, error) {
	result := make([][]byte, 0, len(*x))
	for _, v := range *x {
		result = append(result, []byte(v))
	}
	return bytes.Join(result, []byte(" ")), nil
}
func (x *IdListType) UnmarshalText(text []byte) error {
	for _, v := range bytes.Fields(text) {
		*x = append(*x, string(v))
	}
	return nil
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdGMonth time.Time

func (t *xsdGMonth) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "--01")
}
func (t xsdGMonth) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("--01")), nil
}
func (t xsdGMonth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGMonth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdGMonthDay time.Time

func (t *xsdGMonthDay) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "--01-02")
}
func (t xsdGMonthDay) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("--01-02")), nil
}
func (t xsdGMonthDay) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGMonthDay) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdGYear time.Time

func (t *xsdGYear) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006")
}
func (t xsdGYear) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006")), nil
}
func (t xsdGYear) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYear) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdGYearMonth time.Time

func (t *xsdGYearMonth) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01")), nil
}
func (t xsdGYearMonth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYearMonth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("15:04:05.999999999")), nil
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
