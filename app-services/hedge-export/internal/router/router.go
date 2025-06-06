/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.
 
* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package router

import (
	"hedge/app-services/hedge-export/internal/db/redis"
	"hedge/app-services/hedge-export/internal/dto"
	. "hedge/app-services/hedge-export/internal/service"
	"encoding/json"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	profileName = "profile"
	metricName  = "metric"
)

type Router struct {
	edgexSdk      interfaces.ApplicationService
	appConfig     *dto.AppConfig
	metricService *MetricConfigService
	MetricExport  *MetricExportService
}

func NewRouter(applicationService interfaces.ApplicationService, dbClient *redis.DBClient, appConfig *dto.AppConfig) *Router {
	router := new(Router)
	router.edgexSdk = applicationService
	router.metricService = NewMetricConfigService(applicationService, dbClient)
	router.appConfig = appConfig
	return router
}

func (r Router) LoadRestRoutes() {

	r.addMetricExportRoutes()

}

func (r Router) addMetricExportRoutes() {

	r.edgexSdk.AddCustomRoute("/api/v3/exportdata/metric/config", interfaces.Authenticated, func(c echo.Context) error {
		var metricConfig dto.MetricExportConfig
		err := json.NewDecoder(c.Request().Body).Decode(&metricConfig)
		if err != nil {
			http.Error(c.Response(), err.Error(), http.StatusBadRequest)
			c.Response().Header().Set("Content-Type", "application/json")
			return err
		}

		id, err := r.metricService.AddMetricExportData(metricConfig)
		if err != nil {
			http.Error(c.Response(), err.Error(), http.StatusBadRequest)
			c.Response().Header().Set("Content-Type", "application/json")
			return err
		}

		c.Response().WriteHeader(http.StatusCreated)
		bytes, _ := json.Marshal("New Metric Export Data stored successfully for Metric : " + id)
		c.Response().Write(bytes)
		c.Response().Header().Set("Content-Type", "application/json")
		return nil
	}, http.MethodPost)

	r.edgexSdk.AddCustomRoute("/api/v3/exportdata/metric/config", interfaces.Authenticated, func(c echo.Context) error {
		var metricConfig dto.MetricExportConfig
		err := json.NewDecoder(c.Request().Body).Decode(&metricConfig)
		if err != nil {
			http.Error(c.Response(), err.Error(), http.StatusBadRequest)
			c.Response().Header().Set("Content-Type", "application/json")
			return err
		}

		id, err := r.metricService.UpdateMetricExportData(metricConfig)
		if err != nil {
			http.Error(c.Response(), err.Error(), http.StatusBadRequest)
			c.Response().Header().Set("Content-Type", "application/json")
			return err
		}

		c.Response().WriteHeader(http.StatusCreated)
		bytes, _ := json.Marshal("Metric Export Data updated successfully for Metric : " + id)
		c.Response().Write(bytes)
		c.Response().Header().Set("Content-Type", "application/json")
		return nil
	}, http.MethodPut)

	r.edgexSdk.AddCustomRoute("/api/v3/exportdata/metric/config", interfaces.Authenticated, func(c echo.Context) error {
		data, err := r.metricService.GetAllMetricExportData()

		exportData := dto.ExportData{
			ExportFrequency:  r.appConfig.ExportBatchFrequency,
			MetricesToExport: data,
		}
		if err != nil {
			http.Error(c.Response(), err.Error(), http.StatusBadRequest)
			c.Response().Header().Set("Content-Type", "application/json")
			return err
		}

		c.Response().WriteHeader(http.StatusOK)
		bytes, _ := json.Marshal(exportData)
		c.Response().Write(bytes)
		c.Response().Header().Set("Content-Type", "application/json")
		return nil
	}, http.MethodGet)

	r.edgexSdk.AddCustomRoute("/api/v3/exportdata/metric/config/:"+profileName+"/:"+metricName, interfaces.Authenticated, func(c echo.Context) error {
		profile := c.Param(profileName)
		metric := c.Param(metricName)
		configName := metric + ":" + profile
		data, err := r.metricService.GetMetricExportData(configName)
		if err != nil {
			http.Error(c.Response(), err.Error(), http.StatusBadRequest)
			c.Response().Header().Set("Content-Type", "application/json")
			return err
		}

		c.Response().WriteHeader(http.StatusOK)
		bytes, _ := json.Marshal(data)
		c.Response().Write(bytes)
		c.Response().Header().Set("Content-Type", "application/json")
		return nil
	}, http.MethodGet)

	r.edgexSdk.AddCustomRoute("/api/v3/exportdata/metric/config/:"+profileName+"/:"+metricName, interfaces.Authenticated, func(c echo.Context) error {
		profile := c.Param(profileName)
		metric := c.Param(metricName)
		configName := metric + ":" + profile
		err := r.metricService.DeleteMetricExportData(configName)
		if err != nil {
			http.Error(c.Response(), err.Error(), http.StatusBadRequest)
			c.Response().Header().Set("Content-Type", "application/json")
			return err
		}

		c.Response().WriteHeader(http.StatusOK)
		bytes, _ := json.Marshal("Metric Export Data with Metric :" + metric + " deleted successfully")
		c.Response().Write(bytes)
		c.Response().Header().Set("Content-Type", "application/json")
		return nil
	}, http.MethodDelete)

	r.edgexSdk.AddCustomRoute("/api/v3/exportdata/metric/config/:exportFrequency", interfaces.Authenticated, func(c echo.Context) error {
		frequency := c.Param("exportFrequency")
		err := r.metricService.SetFrequencyExportData(frequency)
		if err != nil {
			http.Error(c.Response(), err.Error(), http.StatusBadRequest)
			c.Response().Header().Set("Content-Type", "application/json")
			return err
		}
		go r.MetricExport.ResetExportData()
		c.Response().WriteHeader(http.StatusOK)
		c.Response().Header().Set("Content-Type", "application/json")
		return nil
	}, http.MethodGet)
}
