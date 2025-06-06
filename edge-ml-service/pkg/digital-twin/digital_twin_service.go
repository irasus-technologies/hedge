/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package digital_twin

import (
	logLib "github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	//	"hedge/app-services/hedge-digital-twin/pkg/db"
	"hedge/common/db/redis"
	redis2 "hedge/edge-ml-service/pkg/db/redis"
	"hedge/edge-ml-service/pkg/dto/twin"
)

type DigitalTwinService struct {
	*twin.DigitalTwinDefinition
	*twin.SimulationDefinition
	LoggingClient logLib.LoggingClient
	DBLayer       redis2.TwinDB
}

type Status int

const (
	OK Status = iota
	NotFound
	AlreadyExist
	BadRequest
	InternalError
)

type DigitalTwinServiceResponse struct {
	ErrorMsg string
	Reason   string
	Name     string
	Keys     []string
	Status
	*twin.DigitalTwinDefinition
	*twin.SimulationDefinition
}

func NewDigitalTwinService(logClient logLib.LoggingClient, dbClient *redis.DBClient) *DigitalTwinService {
	return &DigitalTwinService{LoggingClient: logClient, DBLayer: redis2.NewDBLayer(dbClient, logClient)}
}
