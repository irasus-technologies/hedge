/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package service

import (
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
)

type AppServiceCreator interface {
	NewAppServiceWithTargetType(
		serviceKey string,
		targetType interface{},
	) (interfaces.ApplicationService, bool)
	NewAppService(serviceKey string) (interfaces.ApplicationService, bool)
}

type AppService struct{}

func (a *AppService) NewAppServiceWithTargetType(
	serviceKey string,
	targetType interface{},
) (interfaces.ApplicationService, bool) {
	return pkg.NewAppServiceWithTargetType(serviceKey, targetType)
}

func (a *AppService) NewAppService(serviceKey string) (interfaces.ApplicationService, bool) {
	return pkg.NewAppService(serviceKey)
}
