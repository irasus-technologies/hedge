/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package main

import (
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	pkg2 "hedge/app-services/hedge-metadata-notifier/pkg"
	"hedge/common/client"
	_ "net/http/pprof"
	"os"
)

func main() {

	service, ok := pkg.NewAppServiceWithTargetType(client.HedgeMetaDataNotifierServiceKey, &dtos.SystemEvent{})

	if !ok {
		os.Exit(-1)
	}
	lc := service.LoggingClient()

	if err := service.SetDefaultFunctionsPipeline(
		pkg2.NewNotifier(service).PublishToSupportNotification,
	); err != nil {
		lc.Errorf("SDK SetDefaultFunctionsPipeline failed: %v\n", err)
		os.Exit(-1)
	}

	if err := service.Run(); err != nil {
		lc.Errorf("Run returned error: %s", err.Error())
		os.Exit(-1)
	}

	os.Exit(0)
}
