/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package driver

const (
	SqlDropTable                          = "DROP TABLE IF EXISTS VIRTUAL_RESOURCE;"
	SqlCreateTable                        = "CREATE TABLE VIRTUAL_RESOURCE (DEVICE_NAME string, COMMAND_NAME String, DEVICE_RESOURCE_NAME string, ENABLE_RANDOMIZATION bool, DATA_TYPE String, VALUE string);"
	SqlSelect                             = "SELECT DEVICE_NAME, COMMAND_NAME, DEVICE_RESOURCE_NAME, ENABLE_RANDOMIZATION, DATA_TYPE, VALUE FROM VIRTUAL_RESOURCE where DEVICE_NAME==$1 and DEVICE_RESOURCE_NAME==$2;"
	SqlInsert                             = "INSERT INTO VIRTUAL_RESOURCE VALUES ($1, $2, $3, $4, $5, $6);"
	SqlUpdateRandomization                = "UPDATE VIRTUAL_RESOURCE SET ENABLE_RANDOMIZATION=$1 WHERE DEVICE_NAME==$2 AND DEVICE_RESOURCE_NAME==$3;"
	SqlUpdateValue                        = "UPDATE VIRTUAL_RESOURCE SET VALUE=$1 WHERE DEVICE_NAME==$2 AND DEVICE_RESOURCE_NAME==$3;"
	SqlUpdateValueAndDisableRandomization = "UPDATE VIRTUAL_RESOURCE SET VALUE=$1, ENABLE_RANDOMIZATION=false WHERE DEVICE_NAME==$2 AND DEVICE_RESOURCE_NAME==$3;"
	SqlDelete                             = "DELETE FROM VIRTUAL_RESOURCE WHERE DEVICE_NAME==$1"
)
