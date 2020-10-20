// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	// ErrUnknownReturnType is given when has unknown return type
	ErrUnknownReturnType = errors.New("has unknown return type")
	// ErrUnknownReturnType is given when failed to create tax return
	ErrFailedCreateTaxReturn = errors.New("failed to create tax return")
	// ErrEmptyXML is given when hasn't xml document
	ErrEmptyXML = errors.New("hasn't xml document")
)

var (
	IRS990ReturnTypeCode    = "990"
	DefaultValidateFunction = "Validate"
	IsValidateFunction      = "IsValid"
)

func validateCallback(r interface{}) error {
	method := reflect.ValueOf(r).MethodByName(IsValidateFunction)
	if method.IsValid() {
		response := method.Call(nil)
		if len(response) > 0 {
			err := response[0]
			if !err.IsNil() {
				return err.Interface().(error)
			}
		}
	}
	return nil
}

func validateCallbackByValue(data reflect.Value) error {
	method := data.MethodByName(DefaultValidateFunction)
	if method.IsValid() {
		response := method.Call(nil)
		if len(response) > 0 {
			err := response[0]
			if !err.IsNil() {
				return err.Interface().(error)
			}
		}
	}
	return nil
}

// to validate interface
func Validate(r interface{}) error {
	err := validateCallback(r)
	if err != nil {
		return err
	}

	fields := reflect.ValueOf(r).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fmt.Println(fields.Type().Field(i).Name)
		fieldData := fields.Field(i)
		kind := fieldData.Kind()
		if kind == reflect.Slice {
			for i := 0; i < fieldData.Len(); i++ {
				err = validateCallbackByValue(fieldData.Index(i))
				if err != nil {
					return err
				}
			}
		} else if kind == reflect.Map {
			for _, key := range fieldData.MapKeys() {
				err = validateCallbackByValue(fieldData.MapIndex(key))
				if err != nil {
					return err
				}
			}
		} else {
			err = validateCallbackByValue(fieldData)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
