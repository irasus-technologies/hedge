/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package dto

type SQLMetaData struct {
	DeviceAlias  string   `json:"deviceAlias,omitempty" codec:"deviceAlias,omitempty"`
	ProfileAlias string   `json:"profileAlias,omitempty" codec:"profileAlias,omitempty"`
	ProfileName  string   `json:"profileName,omitempty" codec:"profileName,omitempty"`
	MetricNames  []string `json:"metricName,omitempty" codec:"metricName,omitempty"`
}
