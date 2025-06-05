/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/


package service

import (
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/labstack/echo/v4"
	"hedge/common/config"
	"hedge/common/dto"
	srv "hedge/common/service"
	"strconv"
	"strings"
)

func devicesSummary(service interfaces.ApplicationService, devices []dtos.Device) []dto.DeviceSummary {
	var deviceSummary = make([]dto.DeviceSummary, 0)
	for _, childDev := range devices {
		var ds = dto.DeviceSummary{}

		nodeName, serviceName := srv.SplitDeviceServiceNames(childDev.ServiceName)
		node, err := config.GetNode(service, nodeName)
		if err != nil {
			service.LoggingClient().Error("Error getting node: %v", err)
		}
		ds.Id = childDev.Id
		ds.Name = childDev.Name
		ds.Node = node
		ds.OperatingState = childDev.OperatingState
		ds.ProfileName = childDev.ProfileName
		ds.Labels = childDev.Labels
		ds.DeviceService = serviceName
		if childDev.Location != nil {
			ds.Location = dto.ConvertLocation(childDev.Location, ds.Name)
		}

		deviceSummary = append(deviceSummary, ds)
	}

	return deviceSummary
}

func DeviceQuery(c echo.Context) *dto.Query {
	var q dto.Query

	device := c.QueryParam("name")
	if device != "" {
		q.Filter.Device = device
		q.Filter.Present = true
	}

	profile := c.QueryParam("profileName")
	if profile != "" {
		q.Filter.Profile = profile
		q.Filter.Present = true
	}

	service := c.QueryParam("service")
	if service != "" {
		q.Filter.Service = service
		q.Filter.Present = true
	}

	labels := c.QueryParam("labels")
	if labels != "" {
		q.Filter.Labels = labels
		q.Filter.Present = true
	}

	node := c.QueryParam("edgeNode")
	if node != "" {
		q.Filter.EdgeNode = node
	}

	sortBy := c.QueryParam("sortBy")
	if sortBy != "" {
		q.Filter.SortBy = sortBy
	}

	sortType := c.QueryParam("sortType")
	if sortType != "" {
		q.Filter.SortType = strings.ToLower(sortType)
	}

	page := c.QueryParam("page")
	switch page {
	case "":
		q.Page.Number = 1
	default:
		pg, err := strconv.Atoi(page)
		if err != nil || pg <= 0 {
			return &q
		}
		q.Page.Number = pg
	}

	pageSize := c.QueryParam("pageSize")
	switch pageSize {
	case "":
		q.Page.Size = 15
	default:
		ps, err := strconv.Atoi(pageSize)
		if err != nil || ps <= 0 {
			return &q
		}
		q.Page.Size = ps
	}

	return &q
}

func Intersection(a, b []string) []string {
	m := make(map[string]int)

	var out []string
	for _, item := range a {
		m[item] += 1
	}
	for _, item := range b {
		m[item] += 1
	}

	for k, v := range m {
		// check exist in both a and b
		if v >= 2 {
			out = append(out, k)
		}
	}
	return out
}
