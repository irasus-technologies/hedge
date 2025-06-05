/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package dto

type DeviceInfo struct {
	Name        string
	Labels      []string
	ProfileName string
	Metrices    []string
}
