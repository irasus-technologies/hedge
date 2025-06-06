/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package config

type ServiceConfig struct {
	ExportConfig AppExportConfig
}

// AppExportConfig is example of service's custom structured configuration that is specified in the service's
// configuration.toml file and Configuration Provider (aka Consul), if enabled.
type AppExportConfig struct {
	BatchTimer     string
	BatchSize      int
	PersistOnError bool
}

// UpdateFromRaw updates the service's full configuration from raw data received from the Service Provider.
func (c *ServiceConfig) UpdateFromRaw(rawConfig interface{}) bool {
	configuration, ok := rawConfig.(*ServiceConfig)
	if !ok {
		return false //errors.New("unable to cast raw config to type 'ServiceConfig'")
	}
	*c = *configuration
	return true
}
