package efile

import "github.com/moov-io/1120x/pkg/utils"

type Irs990TransmissionFile struct {
}

func (r Irs990TransmissionFile) Validate() error {
	return utils.Validate(&r)
}

func (r Irs990TransmissionFile) SOAPEnvelope() ([]byte, error) {
	return nil, nil
}

func (r Irs990TransmissionFile) SOAPAttachment() ([]byte, error) {
	return nil, nil
}
