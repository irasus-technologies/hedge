/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package models

type ContentData struct {
	NodeType    string   `json:"nodeType,omitempty"`
	ContentDir  []string `json:"contentDir,omitempty"`
	TargetNodes []string `json:"targetNodes,omitempty"`
}

type KeyFieldTuple struct {
	Key   string
	Field string
}

type VaultSecretData struct {
	Key   string
	Value string
}
