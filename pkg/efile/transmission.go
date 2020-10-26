// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

//////////////////////////////////////////////
//		Transmission File                   //
//  MIME Multi-part Content Header          //
//  MIME Part Boundary and Content Header   //
//    SOAP Envelope (SOAPEnvelope)          //
//       SOAP Header                        //
//       SOAP Body                          //
//  MIME Part Boundary                      //
//    SOAP Attachment (SOAPAttachment)      //
//  MIME Multi-part End Boundary            //
//////////////////////////////////////////////

package efile

// interface for irs e-file
type IrsTransmissionFile interface {
	Validate() error
	SOAPEnvelope() ([]byte, error)
	SOAPAttachment() ([]byte, error)
	Version() string
}

// interface for return
type IrsReturnFile interface {
	Validate() error
	ZipData() ([]byte, error)
	Version() string
}
