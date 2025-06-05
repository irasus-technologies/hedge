/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package dto

type Discovery struct {
	Protocols map[string]string `json:"protocols"`
}

// type ProtocolProperties map[string]string
