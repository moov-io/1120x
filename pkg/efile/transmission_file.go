package efile

// interface for irs efile
type IrsTransmissionFile interface {
	Validate() error
	SOAPEnvelope() ([]byte, error)
	SOAPAttachment() ([]byte, error)
}

func NewIrs990TransmissionFile() IrsTransmissionFile {
	return nil
}
