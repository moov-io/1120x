// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package irs_990

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"path/filepath"

	"github.com/moov-io/1120x/pkg/utils"
)

type Irs990File struct {
	XmlData  Return                 `xml:"ReturnXml"`
	Manifest *IRSSubmissionManifest `xml:"Manifest,omitempty" json:",omitempty"`
}

func (r Irs990File) Validate() error {
	return utils.Validate(&r)
}

func (r *Irs990File) ZipData() ([]byte, error) {
	if r.Manifest == nil {
		return nil, errors.New("manifest should not empty")
	}

	// Create a buffer to write our archive to.
	fileBuf := new(bytes.Buffer)

	xmlBuf, err := xml.Marshal(&r.XmlData)
	if err != nil {
		return nil, err
	}
	manifest, err := r.Manifest.XmlData()
	if err != nil {
		return nil, err
	}

	// Create a new zip archive.
	writer := zip.NewWriter(fileBuf)

	f, err := writer.Create(filepath.Join("xml", "submission.xml"))
	if err != nil {
		return nil, err
	}
	_, err = f.Write(xmlBuf)
	if err != nil {
		return nil, err
	}

	f, err = writer.Create(filepath.Join("manifest", "manifest.xml"))
	if err != nil {
		return nil, err
	}
	_, err = f.Write(manifest)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	return fileBuf.Bytes(), err
}

func (r Irs990File) Version() string {
	return r.XmlData.Version
}
