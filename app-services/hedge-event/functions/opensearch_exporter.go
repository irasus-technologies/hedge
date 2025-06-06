/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package functions

import (
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/util"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	"github.com/go-jose/go-jose/v3/json"
	"github.com/pkg/errors"
	"hedge/common/dto"
)

type OpenSearchExporterInterface interface {
	SaveEventToOpenSearch(ctx interfaces.AppFunctionContext, data interface{}) (bool, interface{})
}

type OpenSearchExporter struct {
	openSearchClient ElasticClientInterface
	lc               logger.LoggingClient
}

func NewOpenSearchExporter(service interfaces.ApplicationService) *OpenSearchExporter {
	openSearchExporter := new(OpenSearchExporter)
	openSearchExporter.openSearchClient = NewHedgeOpenSearchClient(service)
	openSearchExporter.lc = service.LoggingClient()
	return openSearchExporter
}

func (exporter OpenSearchExporter) SaveEventToOpenSearch(ctx interfaces.AppFunctionContext, data interface{}) (bool, interface{}) {

	updated := false
	if data == nil {
		// We didn't receive a result
		return updated, errors.New("no Data Received")
	}

	// Set up the request object.
	bytes, err := util.CoerceType(data)
	if err != nil {
		return false, errors.New("error while marshalling data: " + err.Error())
	}

	var event dto.HedgeEvent
	err = json.Unmarshal(bytes, &event)
	if err != nil {
		exporter.lc.Errorf("error unmarshaling: error: %v", err)
		return updated, errors.New("error while unmarshalling data: " + err.Error())
	}

	switch event.Status {
	// Closure of event should not overwrite the original data, Allow to close only when there is corresponding Open event

	case "Closed":
		// Get the original event and then update it with closed status
		exporter.lc.Debugf("about to search existing event with correlation_id:%s and status:Open", event.CorrelationId)
		existingOpenEvents, err := exporter.openSearchClient.SearchEvents("correlation_id:\"" + event.CorrelationId + "\" AND status:Open")

		if err != nil {
			exporter.lc.Errorf("error finding Open event with correlationId=%s", event.CorrelationId)
			return updated, errors.Wrapf(err, "error while searching Open event with correlationId: %s", event.CorrelationId)
		}

		exporter.lc.Infof("Updating status of existing open events to Closed, Number of open events: %d", len(existingOpenEvents))
		for _, closureEvent := range existingOpenEvents {
			closureEvent.Status = "Closed"
			err = exporter.openSearchClient.IndexEvent(closureEvent)
			if err != nil {
				exporter.lc.Errorf("Error updating event status to Closed with correlationId=%s", event.CorrelationId)
				continue
			} else {
				updated = true
			}
		}
		return updated, event

	default:
		err := exporter.openSearchClient.IndexEvent(&event)
		if err != nil {
			exporter.lc.Errorf("Error indexing event: %+v", event)
			return updated, err
		} else {
			return true, event
		}
	}

	return updated, event
}
