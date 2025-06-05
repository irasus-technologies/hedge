/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package functions

import (
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"hedge/common/dto"
)

type Executor interface {
	CommandExecutor(ctx interfaces.AppFunctionContext, command dto.Command) (bool, dto.CommandExecutionLog)
}
