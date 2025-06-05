/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package redis

import (
	"hedge/common/db"
	"hedge/common/db/redis"
)

type DBClient redis.DBClient

func NewDBClient(dbConfig *db.DatabaseConfig) *DBClient {
	dbClient := redis.CreateDBClient(dbConfig)
	return (*DBClient)(dbClient)
}
