/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package service

import (
	"hedge/common/dto"
)

type DeviceServiceInter interface {
	LoadProfileAndLabels() *DeviceInfoService
	GetDeviceProfiles() []string
	GetDeviceLabels() []string
	GetDevicesByLabels(labels []string) []string
	GetDevicesByLabelsCriteriaOR(labels []string) []string
	GetDeviceToDeviceInfoMap() map[string]dto.DeviceInfo
	GetLabels() []string
	GetProfiles() []string
	GetDevicesByProfile(profile string) []string
	GetDevicesByLabel(label string) []string
	GetMetricsByDevices(devices []string) []string
	GetDeviceInfoMap() (deviceToDeviceInfoMap map[string]dto.DeviceInfo, metricToDeviceInfoMap map[string][]dto.DeviceInfo, err error)
	LoadDeviceInfoFromDB()
	ClearCache()
}
