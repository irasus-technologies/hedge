/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package dto

type ExportData struct {
	ExportFrequency  string               `json:"exportFrequency"`
	MetricesToExport []MetricExportConfig `json:"metricesToExport"`
}

type MetricExportConfig struct {
	Profile      string   `json:"profile"`
	Metric       string   `json:"metric"`
	IsAllDevices bool     `json:"isAllDevices"`
	Devices      []string `json:"devices"`
}
