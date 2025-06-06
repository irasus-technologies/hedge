/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package dto

import "fmt"

const (
	ErrorTypeFatal        = "Fatal"
	ErrorTypeError        = "Error"
	ErrorTypeValidation   = "Validation Error"
	ErrorTypeInfo         = "Info"
	ErrorTypeDebug        = "Debug"
	ErrorTypeUnauthorized = "Unauthorized"
)

type ErrorDetail struct {
	ErrorType    string
	ErrorMessage string
}

func (err *ErrorDetail) Error() string {
	return fmt.Sprintf("ErrorType: %s, Error Message: %s", err.ErrorType, err.ErrorMessage)
}

type Response struct {
	Data    interface{}
	Status  int
	Error   []ErrorDetail
	Message string
}
