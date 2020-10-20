package efile

import "github.com/moov-io/1120x/pkg/irs_990"

// interface for irs efile
type IrsFileInterface interface {
	Validate() error
}

func NewReturnIrs990() IrsFileInterface {
	return &irs_990.Return{}
}
