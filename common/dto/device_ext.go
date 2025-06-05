/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package dto

type DeviceExtResp struct {
	Field       string `json:"field,omitempty" codec:"field,omitempty"`
	Value       string `json:"value" codec:"value"`
	Default     string `json:"default,omitempty" codec:"default,omitempty"`
	IsMandatory bool   `json:"isMandatory,omitempty" codec:"isMandatory,omitempty"`
}

type DeviceExt struct {
	Field string `json:"field,omitempty" codec:"field,omitempty"`
	Value string `json:"value" codec:"value"`
}
