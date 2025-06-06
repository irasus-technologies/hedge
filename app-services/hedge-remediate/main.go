/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package main

import (
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/transforms"
	"hedge/app-services/hedge-remediate/functions"
	"hedge/common/client"
	"hedge/common/config"
	comService "hedge/common/config"
	"hedge/common/dto"
	"hedge/common/telemetry"
	"os"
)

func main() {

	service, ok := pkg.NewAppServiceWithTargetType(client.HedgeRemediateServiceKey, &dto.Command{}) // Key used by Registry (Aka Consul))
	if !ok {
		os.Exit(-1)
	}
	lc := service.LoggingClient()

	subscribeTopic := "commands"
	node, err := comService.GetNode(service, "current")
	if err != nil {
		lc.Errorf("Could not get the node details and hence the nodeType, exiting")
		os.Exit(-1)
	}
	if !node.IsRemoteHost {
		subscribeTopic, err = comService.GetNodeTopicName("commands", service)
		if err != nil {
			lc.Errorf("Could not get topic name, exiting")
			os.Exit(-1)
		}
	}
	lc.Infof("Subscribing to topic: %s", subscribeTopic)

	hedge_event_publisher_trigger_url, err := service.GetAppSetting("EventPipelineTriggerURL")
	if err != nil {
		lc.Errorf("EventPipelineTriggerURL not configured, exiting")
		os.Exit(-1)
	}

	metricsManager, err := telemetry.NewMetricsManager(service, client.HedgeRemediateServiceName)
	if err != nil {
		lc.Error("Failed to create metrics manager. Returned error: ", err.Error())
		os.Exit(-1)
	}

	cmdExecutionService := functions.NewCommandExecutionService(service, client.HedgeRemediateServiceName, metricsManager.MetricsMgr, node.HostName)
	err = service.AddFunctionsPipelineForTopics("hedge-remediate-pipeline", []string{subscribeTopic, "commands"},
		cmdExecutionService.IdentifyCommandAndExecute,
		transforms.NewHTTPSender(hedge_event_publisher_trigger_url, "application/json", config.GetPersistOnError(service)).HTTPPost,
	)

	if err != nil {
		lc.Error("hedge-remediate functions pipeline failed: " + err.Error())
		os.Exit(-1)
	}

	metricsManager.Run()

	// Check and if required restore it: Probably only for testing, it sends the command to command topic
	/*	service.AddRoute("/api/v1/helix/remediation", func(writer http.ResponseWriter, req *http.Request) {
		// Direct API call for Command , Device Service publish to MQTT topic BMCCommandLog
		mqttService.RestRemediation(writer, req)
	}, http.MethodPost)*/

	err = service.Run()
	if err != nil {
		lc.Error("Run returned error: ", err.Error())
		os.Exit(-1)
	}

	os.Exit(0)
}
