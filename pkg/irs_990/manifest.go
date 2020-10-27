// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import (
	"encoding/xml"

	"github.com/moov-io/1120x/pkg/utils"
)

type IRSSubmissionManifest struct {
	SubmissionId            SubmissionIdType        `xml:"SubmissionId"`
	EFIN                    EFINType                `xml:"EFIN"`
	TaxYr                   *YearType               `xml:"TaxYr,omitempty" json:",omitempty"`
	GovernmentCd            GovernmentCodeType      `xml:"GovernmentCd"`
	FederalSubmissionTypeCd FederalSubmissionTypeCd `xml:"FederalSubmissionTypeCd"`
	TaxPeriodBeginDt        *DateType               `xml:"TaxPeriodBeginDt,omitempty" json:",omitempty"`
	TaxPeriodEndDt          *DateType               `xml:"TaxPeriodEndDt,omitempty" json:",omitempty"`
	TIN                     EINType                 `xml:"TIN"`
	Xmlns                   string                  `xml:"xmlns,attr,omitempty" json:",omitempty"`
}

func (r IRSSubmissionManifest) Validate() error {
	return utils.Validate(&r)
}

func (r *IRSSubmissionManifest) Init() error {
	r.Xmlns = "http://www.irs.gov/efile"
	return nil
}

func (r *IRSSubmissionManifest) XmlData() ([]byte, error) {
	return xml.Marshal(r)
}

func (r *IRSSubmissionManifest) SubmissionIdentifier() SubmissionIdType {
	return r.SubmissionId
}

func (r *IRSSubmissionManifest) SetSubmissionIdentifier(id SubmissionIdType) {
	r.SubmissionId = id
}

type StateSubmissionManifest struct {
	SubmissionId           SubmissionIdType        `xml:"SubmissionId"`
	EFIN                   EFINType                `xml:"EFIN"`
	TaxYr                  YearType                `xml:"TaxYr"`
	GovernmentCd           GovernmentCodeType      `xml:"GovernmentCd"`
	StateSubmissionTyp     StateSubmissionTyp      `xml:"StateSubmissionTyp"`
	SubmissionCategoryCd   SubmissionCategoryType  `xml:"SubmissionCategoryCd"`
	FederalEIN             FederalEIN              `xml:"FederalEIN"`
	BusinessNameControlTxt BusinessNameControlType `xml:"BusinessNameControlTxt"`
	PrimarySSN             SSNType                 `xml:"PrimarySSN"`
	PrimaryNameControlTxt  PersonNameControlType   `xml:"PrimaryNameControlTxt"`
	SpouseSSN              SSNType                 `xml:"SpouseSSN"`
	SpouseNameControlTxt   PersonNameControlType   `xml:"SpouseNameControlTxt"`
	TempId                 TempIdType              `xml:"TempId"`
	IRSSubmissionId        *SubmissionIdType       `xml:"IRSSubmissionId,omitempty" json:",omitempty"`
	StateSchemaVersionNum  string                  `xml:"StateSchemaVersionNum,omitempty" json:",omitempty"`
	Xmlns                  string                  `xml:"xmlns,attr,omitempty" json:",omitempty"`
}

func (r StateSubmissionManifest) Validate() error {
	return utils.Validate(&r)
}

func (r *StateSubmissionManifest) Init() error {
	r.Xmlns = "http://www.irs.gov/efile"
	return nil
}

func (r *StateSubmissionManifest) XmlData() ([]byte, error) {
	return xml.Marshal(r)
}

func (r *StateSubmissionManifest) SubmissionIdentifier() SubmissionIdType {
	return r.SubmissionId
}

func (r *StateSubmissionManifest) SetSubmissionIdentifier(id SubmissionIdType) {
	r.SubmissionId = id
}
