/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package data

type MetricResult struct {
	Metric map[string]interface{} `json:"metric,omitempty" codec:"metric,omitempty"`
	//Values  [][]interface{} `json:"values"`
	Values []interface{} `json:"values,omitempty" codec:"values,omitempty"`
}

type TimeSeriesData struct {
	ResultType string         `json:"resultType,omitempty" codec:"resultType,omitempty"`
	Result     []MetricResult `json:"result,omitempty" codec:"result,omitempty"`
}

type TimeSeriesResponse struct {
	Status string         `json:"status,omitempty" codec:"status,omitempty"`
	Data   TimeSeriesData `json:"data,omitempty" codec:"data,omitempty"`
}
