/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package twin

func NewSimulationDefinition(mlModelName, twinDefinitionName string) *SimulationDefinition {
	return &SimulationDefinition{
		RunType:              OneTimeRun,
		RunControl:           UserDriven,
		TwinDefinitionName:   twinDefinitionName,
		SimulationParameters: make([]ExternalParameter, 0),
	}
}
