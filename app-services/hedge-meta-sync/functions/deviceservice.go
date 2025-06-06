/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package functions

type DeviceService struct {
	ApiVersion string  `json:"apiVersion"`
	Service    Service `json:"service"`
}

type Service struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	AdminState  string   `json:"adminState"`
	Labels      []string `json:"labels"`
	BaseAddress string   `json:"baseAddress"`
}

type ServicesArray struct {
	ApiVersion string    `json:"apiVersion"`
	TotalCount int       `json:"totalCount"`
	Services   []Service `json:"services"`
}
