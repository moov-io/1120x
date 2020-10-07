// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package xmls

import (
	"encoding/xml"
	"regexp"
)

type FormApplicationProperties struct {
	Text                    string `xml:",chardata"`
	FormHeaderGeneric       string `xml:"FormHeaderGeneric"`
	FormHeaderOriginalData  string `xml:"FormHeaderOriginalData"`
	SystemMode              string `xml:"SystemMode"`
	PrintoutWindowResizable bool   `xml:"PrintoutWindowResizable"`
}

type FormView struct {
	Text                          string `xml:",chardata"`
	ShowReturnSummaryOnStartup    bool   `xml:"ShowReturnSummaryOnStartup"`
	DefaultZoomLevel              string `xml:"DefaultZoomLevel"`
	ReturnTreeMouseoverBgColor    string `xml:"ReturnTreeMouseoverBgColor"`
	ReturnTreeSelectedItemBgColor string `xml:"ReturnTreeSelectedItemBgColor"`
	ErrorFieldBgColor             string `xml:"ErrorFieldBgColor"`
	ChangedFieldBgColor           string `xml:"ChangedFieldBgColor"`
	TableHeaderBgColor            string `xml:"TableHeaderBgColor"`
	TableRow1BgColor              string `xml:"TableRow1BgColor"`
	TableRow2BgColor              string `xml:"TableRow2BgColor"`
	SingleReturnAutoDisplay       bool   `xml:"SingleReturnAutoDisplay"`
	StartUpDataStage              string `xml:"StartUpDataStage"`
}

type FormUserPreferences struct {
	Text string   `xml:",chardata"`
	View FormView `xml:"View"`
}

type FormParameters struct {
	Text              string `xml:",chardata"`
	Stage             string `xml:"Stage"`
	DLN               string `xml:"DLN"`
	DLNLatest         string `xml:"DLNLatest"`
	DLNChanged        string `xml:"DLNChanged"`
	TIN               string `xml:"TIN"`
	TINLatest         string `xml:"TINLatest"`
	TINChanged        string `xml:"TINChanged"`
	TaxpayerPrint     bool   `xml:"TaxpayerPrint"`
	Print             string `xml:"Print"`
	DocumentId        string `xml:"DocumentId"`
	ReturnVersion     string `xml:"ReturnVersion"`
	ZoomLevel         string `xml:"ZoomLevel"`
	DisplayName       string `xml:"DisplayName"`
	Regulation        string `xml:"Regulation"`
	Location          string `xml:"Location"`
	LocationSeq       string `xml:"LocationSeq"`
	SubmissionVersion string `xml:"SubmissionVersion"`
	SubmissionType    string `xml:"SubmissionType"`
}

type FormSubmissionHeaderAndDocument struct {
	Text               string      `xml:",chardata"`
	ReturnHeader       interface{} `xml:"ReturnHeader"`
	SubmissionDocument interface{} `xml:"SubmissionDocument"`
}

type FormData struct {
	XMLName                     xml.Name                        `xml:"AppData"`
	Text                        string                          `xml:",chardata"`
	ApplicationProperties       FormApplicationProperties       `xml:"ApplicationProperties"`
	UserPreferences             FormUserPreferences             `xml:"UserPreferences"`
	Parameters                  FormParameters                  `xml:"Parameters"`
	SubmissionHeaderAndDocument FormSubmissionHeaderAndDocument `xml:"SubmissionHeaderAndDocument"`
}

type XMLParameters struct {
	Parameters FormParameters
	Properties FormApplicationProperties
	View       FormView
}

// Converting the struct to String format.
func (r *FormData) String() string {
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

type FormDataCreator struct {
	SubmissionVersion string
	SubmissionYear    int
	SubmissionType    string
	Header            interface{}
	Data              []ReturnInspectData
}

func (r *FormDataCreator) GenerateXMLDocument(params *XMLParameters) []XMLDocument {
	var data []XMLDocument
	for _, returnData := range r.Data {
		formData := FormData{}
		if params != nil {
			formData.Parameters = params.Parameters
			formData.UserPreferences.View = params.View
			formData.ApplicationProperties = params.Properties
		}
		formData.Parameters.SubmissionVersion = r.SubmissionVersion
		formData.Parameters.SubmissionType = r.SubmissionType
		formData.SubmissionHeaderAndDocument.ReturnHeader = r.Header
		formData.SubmissionHeaderAndDocument.SubmissionDocument = returnData
		data = append(data, XMLDocument{XML: []byte(formData.String()), Type: returnData.DataType, Year: r.SubmissionYear})
	}
	return data
}
