/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package db

import (
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	localModel "hedge/app-services/hedge-admin/models"
	"hedge/common/db"
	"hedge/common/db/redis"
	"hedge/common/dto"
	hedgeErrors "hedge/common/errors"
)

type RedisDbClient interface {
	SaveProtocols(dsps dto.DeviceServiceProtocols) ([]string, hedgeErrors.HedgeError)
	SaveNode(node *dto.Node) (string, hedgeErrors.HedgeError)
	DeleteNode(nodeId string, keyFieldTuples []localModel.KeyFieldTuple) hedgeErrors.HedgeError
	GetNode(nodeName string) (*dto.Node, hedgeErrors.HedgeError)
	GetAllNodes() ([]dto.Node, hedgeErrors.HedgeError)
	SaveNodeGroup(parentGroup string, dbNodeGroup *dto.DBNodeGroup) (string, hedgeErrors.HedgeError)
	FindNodeKey(parentGroup string) (string, hedgeErrors.HedgeError)
	UpsertChildNodeGroups(parentNodeName string, childNodes []string) (string, hedgeErrors.HedgeError)
	GetNodeGroup(nodeGroupName string) (*dto.DBNodeGroup, hedgeErrors.HedgeError)
	GetDBNodeGroupMembers(nodeHashKey string) ([]dto.DBNodeGroup, hedgeErrors.HedgeError)
	DeleteNodeGroup(parentNodeGroupName string, field string) hedgeErrors.HedgeError
}

type DBClient redis.DBClient

func NewDBClient(service interfaces.ApplicationService) *DBClient {
	dbConfig := db.NewDatabaseConfig()
	dbConfig.LoadAppConfigurations(service)
	dbClient := redis.CreateDBClient(dbConfig)
	return (*DBClient)(dbClient)
}
