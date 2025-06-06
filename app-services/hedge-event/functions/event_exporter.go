/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package functions

import (
	"errors"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/util"
)

type EventExporter struct {
	eventExporter OpenSearchExporterInterface
}

func NewEventExporter(service interfaces.ApplicationService) *EventExporter {
	exp := new(EventExporter)
	exp.eventExporter = NewOpenSearchExporter(service)
	return exp
}

func (exp *EventExporter) StoreEvent(ctx interfaces.AppFunctionContext, data interface{}) (bool, interface{}) {

	lc := ctx.LoggingClient()
	lc.Info("saving event")

	if data == nil {
		lc.Error("no event data Received")
		return false, errors.New("no data Received")
	}

	ok, _ := exp.eventExporter.SaveEventToOpenSearch(ctx, data)
	if !ok {
		lc.Errorf("error saving to local elastic, payload %v", data)
		return false, nil
	}

	// Pass-thru when other event store like Helix BHOM is enabled
	// whether there is an error or not, return true if ADE export is enabled since it is the next step of the pipeline
	return true, data
}

func Print(ctx interfaces.AppFunctionContext, data interface{}) (bool, interface{}) {
	if data == nil {
		// We didn't receive a result
		return false, nil
	}

	dataToPrint, err := util.CoerceType(data)
	if err != nil {
		return false, err
	}
	lc := ctx.LoggingClient()
	lc.Infof("Read data from pipeline: %s", string(dataToPrint))
	return true, data
}
