/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package main

import (
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg"
	"hedge/app-services/hedge-device-extensions/internal/router"
	"hedge/common/client"
	"hedge/common/telemetry"
	"os"
)

var metricsManager *telemetry.MetricsManager

func main() {

	service, ok := pkg.NewAppService(client.HedgeDeviceExtnsServiceKey) // Key used by Registry (Aka Consul))
	if !ok {
		os.Exit(-1)
	}
	lc := service.LoggingClient()

	var err error
	metricsManager, err = telemetry.NewMetricsManager(service, client.HedgeDeviceExtnsServiceName)
	if err != nil {
		lc.Errorf("Failed to create metrics manager. Returned error: %v", err)
		os.Exit(-1)
	}

	router.NewRouter(service, client.HedgeDeviceExtnsServiceName, metricsManager.MetricsMgr).LoadRestRoutes()

	dtr := router.NewDtRouter(service)
	if dtr == nil {
		lc.Errorf("failed to create DtRouter instance")
		os.Exit(-1)
	}
	dtr.LoadDtRoutes()

	metricsManager.Run()

	// 5) Lastly, we'll go ahead and tell the SDK to "start"
	err = service.Run()
	if err != nil {
		lc.Error("Run returned error: ", err.Error())
		os.Exit(-1)
	}

	// Do any required cleanup here

	os.Exit(0)

}
