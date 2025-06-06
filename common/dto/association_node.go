/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package dto

type AssociationNode struct {
	NodeType string `json:"nodeType"`
	NodeName string `json:"nodeName"`
}

func NewAssociationNode(nodeType string, nodeName string) *AssociationNode {
	node := new(AssociationNode)
	node.NodeName = nodeName
	node.NodeType = nodeType
	return node
}
