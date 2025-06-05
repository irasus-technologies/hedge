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
	redis2 "hedge/app-services/hedge-device-extensions/pkg/db/redis"
	dto2 "hedge/common/dto"
	hedgeErrors "hedge/common/errors"
)

type MetaDataService interface {
	AddAssociation(nodeA string, nodeAType string, associationNodes []dto2.AssociationNode) (string, hedgeErrors.HedgeError)
	GetAssociation(nodeA string) ([]dto2.AssociationNode, hedgeErrors.HedgeError)
	DeleteAssociation(nodeA string, nodeAType string) hedgeErrors.HedgeError
	//GetGraphData(nodeA string, level int64) (graphItem models.GraphItem, err hedgeErrors.HedgeError)
	UpdateAssociation(nodeA string, nodeAType string, associationNodes []dto2.AssociationNode, forceCreate bool) (string, hedgeErrors.HedgeError)

	AddLocation(location dto2.Location) (locationMap map[string]string, err hedgeErrors.HedgeError)
	// Country is mandatory, state and city are optional, Will be used in UI lookup & for map dashboard
	GetLocations(country string, state string, city string) (locationsMap []map[string]string, err hedgeErrors.HedgeError)
	GetLocation(locationId string) (location dto2.Location, err hedgeErrors.HedgeError)

	GetMetricsForProfile(profile string) ([]string, hedgeErrors.HedgeError)

	GetProtocolsForService(service string) (map[string]string, hedgeErrors.HedgeError)

	GetDeviceDetails(deviceName string) (dtos.Device, string, hedgeErrors.HedgeError)
	GetDevices(query *dto2.Query) ([]dto2.DeviceSummary, dto2.Page, hedgeErrors.HedgeError)

	CreateCompleteProfile(profileName string, profileObject dto2.ProfileObject) (string, hedgeErrors.HedgeError)
	UpdateCompleteProfile(profileName string, profileObject dto2.ProfileObject) (string, hedgeErrors.HedgeError)
	GetCompleteProfile(profileName string) (dto2.ProfileObject, hedgeErrors.HedgeError)
	DeleteCompleteProfile(profileName string) hedgeErrors.HedgeError
	CreateCompleteDevice(deviceName string, deviceObject dto2.DeviceObject) (string, hedgeErrors.HedgeError)
	UpdateCompleteDevice(deviceName string, deviceObject dto2.DeviceObject) (string, hedgeErrors.HedgeError)
	GetCompleteDevice(deviceName string, metrics string, service interfaces.ApplicationService) (dto2.DeviceObject, hedgeErrors.HedgeError)
	DeleteCompleteDevice(deviceName string) hedgeErrors.HedgeError

	AddDeviceExtensionsInProfile(profileName string, deviceExtns []dto2.DeviceExtension) (string, hedgeErrors.HedgeError)
	UpdateDeviceExtensionsInProfile(profileName string, deviceExtns []dto2.DeviceExtension, forceCreate bool) (string, hedgeErrors.HedgeError)
	GetDeviceExtensionInProfile(profileName string) ([]dto2.DeviceExtension, hedgeErrors.HedgeError)
	DeleteDeviceExtensionInProfile(profileName string) hedgeErrors.HedgeError
	AddDeviceExtension(deviceName string, deviceExts []dto2.DeviceExt) (string, hedgeErrors.HedgeError)
	UpdateDeviceExtension(deviceName string, deviceExts []dto2.DeviceExt, forceCreate bool) (string, hedgeErrors.HedgeError)
	GetDeviceExtension(deviceName string) ([]dto2.DeviceExtResp, hedgeErrors.HedgeError)
	DeleteAllDeviceExtension(deviceName string) hedgeErrors.HedgeError
	DeleteDeviceExtension(deviceName string, attToDelete []string) hedgeErrors.HedgeError

	UpsertProfileContextualAttributes(profileName string, contextualAttrs []string) hedgeErrors.HedgeError
	GetProfileContextualAttributes(profileName string) ([]string, hedgeErrors.HedgeError)
	DeleteProfileContextualAttributes(profileName string) hedgeErrors.HedgeError
	UpsertDeviceContextualAttributes(deviceName string, contextualData map[string]interface{}) hedgeErrors.HedgeError
	GetDeviceContextualAttributes(deviceName string) (map[string]interface{}, hedgeErrors.HedgeError)
	DeleteDeviceContextualAttributes(deviceName string) hedgeErrors.HedgeError

	// Get Attribues and Profiles in kind of tabular form to support the UI
	GetProfileMetaDataSummary(profileNames []string) ([]dto2.ProfileSummary, hedgeErrors.HedgeError)
	GetRelatedProfiles(profileNames []string) ([]string, hedgeErrors.HedgeError)
	GetAttributesGroupedByProfiles(profileNames []string) (map[string][]string, hedgeErrors.HedgeError)

	//Get SQL MetaData
	GetSQLMetaData(query string) (dto2.SQLMetaData, hedgeErrors.HedgeError)

	// Downsampling Config
	UpsertDownsamplingConfig(profileName string, downsamplingConfig *dto2.DownsamplingConfig) hedgeErrors.HedgeError
	GetDownsamplingConfig(profileName string) (*dto2.DownsamplingConfig, hedgeErrors.HedgeError)
	DeleteDownsamplingConfig(profileName string) hedgeErrors.HedgeError
	UpsertAggregateDefinition(profileName string, downsamplingConfig dto2.DownsamplingConfig) hedgeErrors.HedgeError

	//NodeHost Raw Data Config
	UpsertNodeRawDataConfigs(nodeRawDataConfigs []dto2.NodeRawDataConfig) hedgeErrors.HedgeError
	GetNodeRawDataConfigs(nodeIDs []string) (map[string]*dto2.NodeRawDataConfig, hedgeErrors.HedgeError)

	GetDbClient() redis2.DeviceExtDBClientInterface

	GetAllNodes() ([]dto2.Node, hedgeErrors.HedgeError)
}
