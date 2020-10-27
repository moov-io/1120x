// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import "github.com/moov-io/1120x/pkg/utils"

type Return struct {
	ReturnHeader  ReturnHeaderType `xml:"ReturnHeader"`
	ReturnData    ReturnData       `xml:"ReturnData"`
	Xmlns         string           `xml:"xmlns,attr,omitempty" json:",omitempty"`
	ReturnVersion string           `xml:"returnVersion,attr"`
}

func (r Return) Validate() error {
	return utils.Validate(&r)
}

func (r *Return) Init() error {
	r.Xmlns = "http://www.irs.gov/efile"
	return nil
}

type ReturnData struct {
	IRS990                     IRS990                      `xml:"IRS990"`
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
