/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package redis

import "github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"

var loggingClient logger.LoggingClient

func newLoggingClient() logger.LoggingClient {
	if loggingClient == nil {
		return logger.NewClient("redis", "INFO")
	}
	return loggingClient
}
