/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package config

import (
	"errors"
	"hedge/common/dto"
)

// This file contains example of custom configuration that can be loaded from the service's configuration.toml
// and/or the Configuration Provider, aka Consul (if enabled).
// For more details see https://docs.edgexfoundry.org/2.0/microservices/device/Ch-DeviceServices/#custom-configuration

// Example structured custom configuration types. Must be wrapped with an outer struct with
// single element that matches the top level custom configuration element in your configuration.toml file,
// 'SimpleCustom' in this example.
type VirtualDeviceConfiguration struct {
	VirtualDeviceConfig VirtualDeviceConfig
	Hedge               dto.HedgeSvcConfig
}

// SimpleCustomConfig is example of service's custom structured configuration that is specified in the service's
// configuration.toml file and Configuration Provider (aka Consul), if enabled.
type VirtualDeviceConfig struct {
	ServiceName         string
	DeviceNamePrefix    string
	MaxDiscoveredDevice int64
	Writable            VirtialDeviceWritable
}

// SimpleWritable defines the service's custom configuration writable section, i.e. can be updated from Consul
type VirtialDeviceWritable struct {
	DiscoverSleepDurationSecs int64
}

// UpdateFromRaw updates the service's full configuration from raw data received from
// the Service Provider.
func (sw *VirtualDeviceConfiguration) UpdateFromRaw(rawConfig interface{}) bool {
	configuration, ok := rawConfig.(*VirtualDeviceConfiguration)
	if !ok {
		return false //errors.New("unable to cast raw config to type 'ServiceConfig'")
	}

	*sw = *configuration

	return true
}

// Validate ensures your custom configuration has proper values.
// Example of validating the sample custom configuration
func (scc *VirtualDeviceConfig) Validate() error {
	if len(scc.DeviceNamePrefix) == 0 {
		return errors.New("VirtualDeviceConfig.DeviceNamePrefix configuration setting can not be blank")
	}

	if scc.MaxDiscoveredDevice < 0 {
		return errors.New("VirtualDeviceConfig.MaxDiscoveredDevice configuration setting can not be less than 0")
	}

	if scc.Writable.DiscoverSleepDurationSecs < 10 {
		return errors.New("VirtualDeviceConfig.Writable.DiscoverSleepDurationSecs configuration setting must be 10 or greater")
	}

	return nil
}
