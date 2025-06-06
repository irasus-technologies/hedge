/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package redis

import (
	"encoding/json"
	"errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	redis2 "github.com/gomodule/redigo/redis"
	"hedge/common/db"
	"hedge/common/db/redis"
	comModels "hedge/common/dto"
)

type EventDBClient redis.DBClient

type DBClient interface {
	GetDbClient(dbConfig *db.DatabaseConfig) DBClient
	SaveEvent(event comModels.HedgeEvent) error
	GetEventByCorrelationId(correlationId string, lc logger.LoggingClient) (*comModels.HedgeEvent, error)
	DeleteEvent(correlationId string) error
}

var DBClientImpl DBClient

func init() {
	DBClientImpl = &EventDBClient{}
}

func NewDBClient(dbConfig *db.DatabaseConfig) DBClient {
	return DBClientImpl.GetDbClient(dbConfig)
}

func (dbClient *EventDBClient) GetDbClient(dbConfig *db.DatabaseConfig) DBClient {
	dbc := redis.CreateDBClient(dbConfig)
	return (*EventDBClient)(dbc)
}

func (dbClient *EventDBClient) SaveEvent(event comModels.HedgeEvent) error {
	conn := dbClient.Pool.Get()
	defer conn.Close()
	evJson, _ := json.Marshal(event)
	_, err := conn.Do("HSET", db.OTEvent, event.CorrelationId, evJson)
	return err
}

func (dbClient *EventDBClient) GetEventByCorrelationId(correlationId string, lc logger.LoggingClient) (*comModels.HedgeEvent, error) {
	conn := dbClient.Pool.Get()
	defer conn.Close()
	var event comModels.HedgeEvent
	eventData, err := redis2.Bytes(conn.Do("HGET", db.OTEvent, correlationId))
	if err != nil && errors.Is(err, redis2.ErrNil) {
		lc.Infof("no existing event found for correlationId: %s, error: %s", correlationId, err.Error())
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	err = json.Unmarshal(eventData, &event)
	return &event, err

}

func (dbClient *EventDBClient) DeleteEvent(correlationId string) error {
	conn := dbClient.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("HDEL", db.OTEvent, correlationId)
	return err
}
