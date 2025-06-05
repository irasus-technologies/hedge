/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package dto

type MQTTmessage struct {
	MessageDefinition MessageDefinition `json:"messageDefinition"`
	Payload           interface{}       `json:"payload,omitempty"`
}

type MessageDefinition struct {
	TargetService string `json:"targetService"`
	TargetURL     string `json:"targetURL"`
	HTTPMethod    string `json:"httpMethod"`
	CorrelationID string `json:"correlationId"`
}
