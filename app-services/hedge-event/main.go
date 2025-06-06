/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package main

import (
	"fmt"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg"
	"hedge/app-services/hedge-event/functions"
	"hedge/common/client"
	"os"
)

func main() {

	service, ok := pkg.NewAppServiceWithTargetType(client.HedgeEventServiceKey, new(interface{}))
	if !ok {
		fmt.Printf("Failed to start App Service: %s\n", client.HedgeEventServiceKey)
		os.Exit(-1)
	}
	lc := service.LoggingClient()

	victoriaMetricsURL, err := service.GetAppSetting("VictoriaMetricsURL")
	if err != nil {
		lc.Error(err.Error())
		os.Exit(-1)
	}
	lc.Debugf("Victoria Metrics/TS URL as in the application settings: %v", victoriaMetricsURL)

	exporter := functions.NewEventExporter(service)
	if exporter == nil {
		lc.Error("Could not create exporter, exiting")
		os.Exit(-1)
	}

	subscribeEventTopics, err := service.GetAppSettingStrings("SubscribeEventTopics")
	if err != nil {
		service.LoggingClient().Errorf("Error getting SubscribeEventTopic from configuration: %v", err)
	}

	// Pipeline for handling Event
	if err := service.AddFunctionsPipelineForTopics("event-pipeline", subscribeEventTopics,
		functions.Print,
		exporter.StoreEvent,
		functions.NewTSEventPublisher().BuildOTMetricEventAsResponseData,
	); err != nil {
		lc.Errorf("'event-pipeline' failed: %v\n", err)
		return
	}

	// 5) Lastly, we'll go ahead and tell the SDK to "start" and begin listening for events
	// to trigger the pipeline.
	err = service.Run()
	if err != nil {
		lc.Error("Run returned error: ", err.Error())
		os.Exit(-1)
	}

	// Do any required cleanup here

	os.Exit(0)
}
