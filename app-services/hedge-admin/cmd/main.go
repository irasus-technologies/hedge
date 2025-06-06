/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package main

import (
	"fmt"
	"os"

	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg"

	"hedge/app-services/hedge-admin/internal/config"
	"hedge/app-services/hedge-admin/internal/router"
	"hedge/common/client"
)

func main() {

	service, ok := pkg.NewAppService(client.HedgeAdminServiceKey)
	if !ok {
		fmt.Printf("Failed to start App Service: %s\n", client.HedgeAdminServiceKey)
		os.Exit(-1)
	}
	lc := service.LoggingClient()

	appConfig := config.NewAppConfig()
	appConfig.LoadAppConfigurations(service)

	r := router.NewRouter(service, appConfig)
	r.LoadRestRoutes()
	err := r.RegisterCurrentNode()
	if err != nil {
		lc.Error("RegisterCurrentNode returned error: %s", err.Error())
	} else {
		lc.Infof("Successfully added node to current node's database")
	}

	// 5) Lastly, we'll go ahead and tell the SDK to "start"
	err = service.Run()
	if err != nil {
		lc.Errorf("Run returned error: %v", err)
		os.Exit(-1)
	}

	// Do any required cleanup here

	os.Exit(0)

}
