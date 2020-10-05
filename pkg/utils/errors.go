package utils

import "errors"

var (
	// ErrUnknownReturnType is given when has unknown return type
	ErrUnknownReturnType = errors.New("has unknown return type")
	// ErrUnknownReturnType is given when failed to create tax return
	ErrFailedCreateTaxReturn = errors.New("failed to create tax return")
	// ErrEmptyXML is given when hasn't xml document
	ErrEmptyXML = errors.New("hasn't xml document")
)

var (
	IRS990ReturnTypeCode = "990"
)
