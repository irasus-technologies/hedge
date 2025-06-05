/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package redis

import (
	"encoding/json"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	"github.com/go-redsync/redsync/v4"
	redis2 "github.com/gomodule/redigo/redis"
	"hedge/common/db"
	"hedge/common/db/redis"
	bmcmodel "hedge/common/dto"
	hedgeErrors "hedge/common/errors"
)

type DBClient struct {
	client *redis.DBClient
}

var RemediateDBClientImpl RemediateDBClientInterface

type RemediateDBClientInterface interface {
	redis.CommonRedisDBInterface
	SaveRemediateEvent(event bmcmodel.HedgeEvent) error
	GetRemediateEventByCorrelationId(correlationId string, lc logger.LoggingClient) (*bmcmodel.HedgeEvent, error)
	DeleteRemediateEvent(correlationId string) error
	GetDbClient(dbConfig *db.DatabaseConfig) RemediateDBClientInterface
}

func init() {
	RemediateDBClientImpl = &DBClient{}
}

func (rc *DBClient) GetDbClient(dbConfig *db.DatabaseConfig) RemediateDBClientInterface {
	dbc := redis.CreateDBClient(dbConfig)
	return &DBClient{client: dbc}
}

func NewDBClient(dbConfig *db.DatabaseConfig) RemediateDBClientInterface {
	return RemediateDBClientImpl.GetDbClient(dbConfig)
}

func (rc *DBClient) SaveRemediateEvent(event bmcmodel.HedgeEvent) error {
	conn := rc.client.Pool.Get()
	defer conn.Close()
	evJson, _ := json.Marshal(event)
	_, err := conn.Do("HSET", db.OTRemediation, event.CorrelationId, evJson)
	return err
}

func (rc *DBClient) GetRemediateEventByCorrelationId(correlationId string, lc logger.LoggingClient) (*bmcmodel.HedgeEvent, error) {
	conn := rc.client.Pool.Get()
	defer conn.Close()
	var event bmcmodel.HedgeEvent
	eventData, err := redis2.Bytes(conn.Do("HGET", db.OTRemediation, correlationId))
	if err != nil && err == redis2.ErrNil {
		lc.Infof("no existing event found for correlationId: %v", correlationId, err.Error())
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	err = json.Unmarshal(eventData, &event)
	return &event, err

}

func (rc *DBClient) DeleteRemediateEvent(correlationId string) error {
	conn := rc.client.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("HDEL", db.OTRemediation, correlationId)
	return err
}

func (rc *DBClient) IncrMetricCounterBy(key string, value int64) (int64, hedgeErrors.HedgeError) {
	return rc.client.IncrMetricCounterBy(key, value)
}

func (rc *DBClient) SetMetricCounter(key string, value int64) hedgeErrors.HedgeError {
	return rc.client.SetMetricCounter(key, value)
}

func (rc *DBClient) GetMetricCounter(key string) (int64, hedgeErrors.HedgeError) {
	return rc.client.GetMetricCounter(key)
}

func (rc *DBClient) AcquireRedisLock(name string) (*redsync.Mutex, hedgeErrors.HedgeError) {
	return rc.client.AcquireRedisLock(name)
}
