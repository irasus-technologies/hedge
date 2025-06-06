/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.

* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package functions

import (
	"encoding/json"
	"errors"
	comModels "hedge/common/dto"

	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/util"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"hedge/common/client"
)

type TSEventPublisher struct {
}

func NewTSEventPublisher() *TSEventPublisher {
	tsEventHandler := new(TSEventPublisher)
	return tsEventHandler
}

func (bm *TSEventPublisher) BuildOTMetricEventAsResponseData(ctx interfaces.AppFunctionContext, data interface{}) (bool, interface{}) {
	if data == nil {
		return false, errors.New("no Metric data received")
	}

	ctx.LoggingClient().Debug("Transforming HedgeEvent to Prometheus TS format")

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return false, errors.New("could not marshal data to bytes")
	}
	var hedgeEvent comModels.HedgeEvent
	err = json.Unmarshal(jsonBytes, &hedgeEvent)
	if err != nil {
		return false, errors.New("invalid data type received, expected HedgeEvent")
	}

	profile := hedgeEvent.Profile
	if profile == "" {
		// Some dummy profile
		profile = "OT_EVENT_PROFILE"
	}
	edgexEvent := dtos.NewEvent(profile, hedgeEvent.DeviceName, client.MetricEvent)
	var eventStatus int64
	if hedgeEvent.Status != "Closed" {
		eventStatus = 1
	} else {
		eventStatus = 0
	}

	edgexEvent.AddSimpleReading(client.MetricEvent, common.ValueTypeInt64, eventStatus)
	// Add labels/ tags
	tags := make(map[string]interface{})
	tags[client.LabelEventSummary] = hedgeEvent.Msg
	tags[client.LabelCorrelationId] = hedgeEvent.CorrelationId
	tags[client.LabelEventType] = hedgeEvent.EventType
	tags[client.LabelNodeName] = hedgeEvent.SourceNode
	edgexEvent.Tags = tags
	// Set the edgexEvent in Response so this is published to hedge/metrics towards end of the pipeline
	metricBytes, _ := util.CoerceType(edgexEvent)
	ctx.SetResponseData(metricBytes)
	ctx.SetResponseContentType("application/json")

	return true, data

}
