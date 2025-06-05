/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package router

import (
	"github.com/labstack/echo/v4"
)

type Result map[string]interface{}
type Response struct {
	Result `json:"result"`
}

func (r *Router) SendError(errorMsg, status string, httpStatus int, c echo.Context) error {

	r.edgexSvc.LoggingClient().Error(errorMsg)
	response := Response{
		Result: Result{
			"status": status,
		},
	}
	return c.JSON(httpStatus, response)
}
