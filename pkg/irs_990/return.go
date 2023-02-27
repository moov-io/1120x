// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import (
	"encoding/xml"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/moov-io/1120x/pkg/utils"
)

type Return struct {
	Text           string `xml:",chardata"`
	Xmlns          string `xml:"xmlns,attr,omitempty" json:",omitempty"`
	Xsi            string `xml:"xsi,attr"`
	SchemaLocation string `xml:"schemaLocation,attr"`
	Version        string `xml:"returnVersion,attr"`

	ReturnHeader ReturnHeaderType `xml:"ReturnHeader"`
	ReturnData   ReturnData       `xml:"ReturnData"`
}

// Parse parses the “Return1099” record from raw xml
func (r *Return) Parse(buf []byte) error {
	if err := xml.Unmarshal(buf, r); err != nil {
		return err
	}
	return nil
}

type inspectStruct struct {
	Data interface{}
	Type string
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	//nolint:exhaustive
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

func generateReturnData(inspect inspectStruct) *utils.ReturnInspectData {
	switch inspect.Type {
	case utils.IRS990ScheduleA:
		value, _ := inspect.Data.(*IRS990ScheduleA)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleA: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleB:
		value, _ := inspect.Data.(*IRS990ScheduleB)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleB: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleC:
		value, _ := inspect.Data.(*IRS990ScheduleC)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleC: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleD:
		value, _ := inspect.Data.(*IRS990ScheduleD)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleD: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleE:
		value, _ := inspect.Data.(*IRS990ScheduleE)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleE: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleF:
		value, _ := inspect.Data.(*IRS990ScheduleF)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleF: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleG:
		value, _ := inspect.Data.(*IRS990ScheduleG)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleG: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleH:
		value, _ := inspect.Data.(*IRS990ScheduleH)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleH: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleI:
		value, _ := inspect.Data.(*IRS990ScheduleI)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleI: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleJ:
		value, _ := inspect.Data.(*IRS990ScheduleJ)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleJ: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleK:
		value, _ := inspect.Data.([]IRS990ScheduleK)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleK: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleL:
		value, _ := inspect.Data.(*IRS990ScheduleL)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleL: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleM:
		value, _ := inspect.Data.(*IRS990ScheduleM)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleM: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleN:
		value, _ := inspect.Data.(*IRS990ScheduleN)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleN: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleO:
		value, _ := inspect.Data.(*IRS990ScheduleO)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleO: value}, DataType: inspect.Type}
	case utils.IRS990ScheduleR:
		value, _ := inspect.Data.(*IRS990ScheduleR)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990ScheduleR: value}, DataType: inspect.Type}
	case utils.IRS990:
		value, _ := inspect.Data.(*IRS990)
		return &utils.ReturnInspectData{Data: ReturnData{DocumentCnt: 1, IRS990: value}, DataType: inspect.Type}
	}
	return nil
}

// Split xml files with a document
func (r *Return) InspectData() *utils.ReturnInspectInfo {
	var returnData []utils.ReturnInspectData
	inspects := []inspectStruct{
		{r.ReturnData.IRS990ScheduleA, utils.IRS990ScheduleA},
		{r.ReturnData.IRS990ScheduleB, utils.IRS990ScheduleB},
		{r.ReturnData.IRS990ScheduleC, utils.IRS990ScheduleC},
		{r.ReturnData.IRS990ScheduleD, utils.IRS990ScheduleD},
		{r.ReturnData.IRS990ScheduleE, utils.IRS990ScheduleE},
		{r.ReturnData.IRS990ScheduleF, utils.IRS990ScheduleF},
		{r.ReturnData.IRS990ScheduleG, utils.IRS990ScheduleG},
		{r.ReturnData.IRS990ScheduleH, utils.IRS990ScheduleH},
		{r.ReturnData.IRS990ScheduleI, utils.IRS990ScheduleI},
		{r.ReturnData.IRS990ScheduleJ, utils.IRS990ScheduleJ},
		{r.ReturnData.IRS990ScheduleK, utils.IRS990ScheduleK},
		{r.ReturnData.IRS990ScheduleL, utils.IRS990ScheduleL},
		{r.ReturnData.IRS990ScheduleM, utils.IRS990ScheduleM},
		{r.ReturnData.IRS990ScheduleN, utils.IRS990ScheduleN},
		{r.ReturnData.IRS990ScheduleO, utils.IRS990ScheduleO},
		{r.ReturnData.IRS990ScheduleR, utils.IRS990ScheduleR},
		{r.ReturnData.IRS990, utils.IRS990},
	}

	for _, ins := range inspects {
		if isNil(ins.Data) {
			continue
		}
		if d := generateReturnData(ins); d != nil {
			returnData = append(returnData, *d)
		}
	}

	if len(returnData) == 0 {
		return nil
	}

	return &utils.ReturnInspectInfo{Header: r.ReturnHeader, Data: returnData}
}

// ReturnYear returns year of return year
func (r *Return) ReturnYear() int {
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
func (r *Return) ReturnVersion() string {
	return r.Version
}

// ReturnType returns type of return type
func (r *Return) ReturnType() string {
	return utils.IRS990ReturnTypeCode
}

// Converting the struct to String format.
func (r *Return) String() string {
	buf, err := xml.Marshal(r)
	if err != nil {
		return ""
	}
	buf, err = utils.FormatXML(buf)
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`(?m)^\s*$[\r\n]*|[\r\n]+\s+\z`)
	return re.ReplaceAllString(string(buf), "")
}

func (r Return) Validate() error {
	return utils.Validate(&r)
}

func (r *Return) Init() error {
	r.Xmlns = "http://www.irs.gov/efile"
	r.SchemaLocation = "http://www.irs.gov/efile"
	r.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	return nil
}

type ReturnData struct {
	IRS990                     *IRS990                     `xml:"IRS990"`
	IRS990ScheduleA            *IRS990ScheduleA            `xml:"IRS990ScheduleA,omitempty" json:",omitempty"`
	IRS990ScheduleB            *IRS990ScheduleB            `xml:"IRS990ScheduleB,omitempty" json:",omitempty"`
	IRS990ScheduleC            *IRS990ScheduleC            `xml:"IRS990ScheduleC,omitempty" json:",omitempty"`
	IRS990ScheduleD            *IRS990ScheduleD            `xml:"IRS990ScheduleD,omitempty" json:",omitempty"`
	IRS990ScheduleE            *IRS990ScheduleE            `xml:"IRS990ScheduleE,omitempty" json:",omitempty"`
	IRS990ScheduleF            *IRS990ScheduleF            `xml:"IRS990ScheduleF,omitempty" json:",omitempty"`
	IRS990ScheduleG            *IRS990ScheduleG            `xml:"IRS990ScheduleG,omitempty" json:",omitempty"`
	IRS990ScheduleH            *IRS990ScheduleH            `xml:"IRS990ScheduleH,omitempty" json:",omitempty"`
	IRS990ScheduleI            *IRS990ScheduleI            `xml:"IRS990ScheduleI,omitempty" json:",omitempty"`
	IRS990ScheduleJ            *IRS990ScheduleJ            `xml:"IRS990ScheduleJ,omitempty" json:",omitempty"`
	IRS990ScheduleK            []IRS990ScheduleK           `xml:"IRS990ScheduleK,omitempty" json:",omitempty"`
	IRS990ScheduleL            *IRS990ScheduleL            `xml:"IRS990ScheduleL,omitempty" json:",omitempty"`
	IRS990ScheduleM            *IRS990ScheduleM            `xml:"IRS990ScheduleM,omitempty" json:",omitempty"`
	IRS990ScheduleN            *IRS990ScheduleN            `xml:"IRS990ScheduleN,omitempty" json:",omitempty"`
	IRS990ScheduleO            *IRS990ScheduleO            `xml:"IRS990ScheduleO,omitempty" json:",omitempty"`
	IRS990ScheduleR            *IRS990ScheduleR            `xml:"IRS990ScheduleR,omitempty" json:",omitempty"`
	AffiliateListing           *AffiliateListing           `xml:"AffiliateListing,omitempty" json:",omitempty"`
	ReasonableCauseExplanation *ReasonableCauseExplanation `xml:"ReasonableCauseExplanation,omitempty" json:",omitempty"`
	AffiliatedGroupAttachment  *AffiliatedGroupAttachment  `xml:"AffiliatedGroupAttachment,omitempty" json:",omitempty"`
	AffiliatedGroupSchedule    *AffiliatedGroupSchedule    `xml:"AffiliatedGroupSchedule,omitempty" json:",omitempty"`
	AveragingAttachment        *AveragingAttachment        `xml:"AveragingAttachment,omitempty" json:",omitempty"`
	RequestForCopyAttachment   *RequestForCopyAttachment   `xml:"RequestForCopyAttachment,omitempty" json:",omitempty"`
	BinaryAttachment           []BinaryAttachment          `xml:"BinaryAttachment,omitempty" json:",omitempty"`
	DocumentCnt                int                         `xml:"documentCnt,attr"`
}

func (r ReturnData) Validate() error {
	return utils.Validate(&r)
}

// Content model for the 990/990-EZ/990-PF Return Header
type ReturnHeaderType struct {
	ReturnTs                    TimestampType      `xml:"ReturnTs"`
	TaxPeriodEndDt              DateType           `xml:"TaxPeriodEndDt"`
	DisasterReliefTxt           string             `xml:"DisasterReliefTxt,omitempty" json:",omitempty"`
	ISPNum                      *ISPType           `xml:"ISPNum,omitempty" json:",omitempty"`
	PreparerFirmGrp             *PreparerFirmGrp   `xml:"PreparerFirmGrp,omitempty" json:",omitempty"`
	SoftwareId                  SoftwareIdType     `xml:"SoftwareId"`
	SoftwareVersionNum          string             `xml:"SoftwareVersionNum,omitempty" json:",omitempty"`
	MultSoftwarePackagesUsedInd bool               `xml:"MultSoftwarePackagesUsedInd"`
	OriginatorGrp               OriginatorGrp      `xml:"OriginatorGrp"`
	PINEnteredByCd              *PINEnteredByCd    `xml:"PINEnteredByCd,omitempty" json:",omitempty"`
	SignatureOptionCd           *SignatureOptionCd `xml:"SignatureOptionCd,omitempty" json:",omitempty"`
	ReturnTypeCd                ReturnTypeCd       `xml:"ReturnTypeCd"`
	TaxPeriodBeginDt            DateType           `xml:"TaxPeriodBeginDt"`
	Filer                       Filer              `xml:"Filer"`
	BusinessOfficerGrp          BusinessOfficerGrp `xml:"BusinessOfficerGrp"`
	PreparerPersonGrp           *PreparerPersonGrp `xml:"PreparerPersonGrp,omitempty" json:",omitempty"`
	IPAddress                   *IPAddressType     `xml:"IPAddress,omitempty" json:",omitempty"`
	IPDt                        *DateType          `xml:"IPDt,omitempty" json:",omitempty"`
	IPTm                        *TimeType          `xml:"IPTm,omitempty" json:",omitempty"`
	IPTimezoneCd                *TimezoneType      `xml:"IPTimezoneCd,omitempty" json:",omitempty"`
	DeviceId                    *DeviceIdType      `xml:"DeviceId,omitempty" json:",omitempty"`
	TaxYr                       YearType           `xml:"TaxYr"`
	BinaryAttachmentCnt         int                `xml:"binaryAttachmentCnt,attr"`
}

func (r ReturnHeaderType) Validate() error {
	return utils.Validate(&r)
}
