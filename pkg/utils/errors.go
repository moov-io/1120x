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
	debug       = false
)

var (
	IRS990          = "990"
	IRS990ScheduleA = "990ScheduleA"
	IRS990ScheduleB = "990ScheduleB"
	IRS990ScheduleC = "990ScheduleC"
	IRS990ScheduleD = "990ScheduleD"
	IRS990ScheduleE = "990ScheduleE"
	IRS990ScheduleF = "990ScheduleF"
	IRS990ScheduleG = "990ScheduleG"
	IRS990ScheduleH = "990ScheduleH"
	IRS990ScheduleI = "990ScheduleI"
	IRS990ScheduleJ = "990ScheduleJ"
	IRS990ScheduleK = "990ScheduleK"
	IRS990ScheduleL = "990ScheduleL"
	IRS990ScheduleM = "990ScheduleM"
	IRS990ScheduleN = "990ScheduleN"
	IRS990ScheduleO = "990ScheduleO"
	IRS990ScheduleR = "990ScheduleR"
)

var (
	IRS990ReturnTypeCode    = "990"
	DefaultValidateFunction = "Validate"
	IsValidateFunction      = "IsValid"
)

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
	var err error
	if debug {
		fmt.Println(reflect.ValueOf(r).Type())
		fmt.Println("==========")
	}
	fields := reflect.ValueOf(r).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldData := fields.Field(i)
		if debug {
			fmt.Println(fields.Type().Field(i).Name)
		}
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
		} else if kind == reflect.Ptr {
			if fieldData.Pointer() != 0 {
				err = validateCallbackByValue(fieldData)
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
