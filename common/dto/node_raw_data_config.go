/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package dto

type NodeRawDataConfig struct {
	SendRawData bool     `json:"sendRawData" codec:"sendRawData"`
	StartTime   int64    `json:"startTime"   codec:"startTime"`
	EndTime     int64    `json:"endTime"     codec:"endTime"`
	Node        NodeHost `json:"node"        codec:"node"` // nodeID is a key
}

type NodeHost struct {
	NodeID string `json:"nodeId" codec:"nodeId"`
	Host   string `json:"host"   codec:"host"`
}
