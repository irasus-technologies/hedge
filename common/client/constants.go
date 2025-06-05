/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.

* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package client

// Constants related to how services identify themselves in the Service Registry
const (
	ServiceKeyHedgePrefix = "app-hedge-"

	// ServiceNames
	HedgeDataEnrichmentServiceName = "hedge-data-enrichment"
	HedgeAdminServiceName          = "hedge-admin"
	HedgeDeviceExtnsServiceName    = "hedge-device-extensions"
	HedgeRemediateServiceName      = "hedge-remediate"
	HedgeMLManagementServiceName   = "hedge-ml-management"

	// ServiceKeys - note that the service key should start with app- for appservices and with device- for deviceservices
	HedgeDataEnrichmentServiceKey   = "app-hedge-data-enrichment"
	HedgeAdminServiceKey            = "app-hedge-admin"
	HedgeDeviceExtnsServiceKey      = "app-hedge-device-extensions"
	HedgeEventServiceKey            = "app-hedge-event"
	HedgeEventPublisherServiceKey   = "app-hedge-event-publisher"
	HedgeExportServiceKey           = "app-hedge-export"
	HedgeRemediateServiceKey        = "app-hedge-remediate"
	HedgeUserAppMgmtServiceKey      = "app-hedge-user-app-mgmt"
	HedgeMetaSyncServiceKey         = "app-hedge-meta-sync"
	HedgeNatsProxyServiceKey        = "app-hedge-nats-proxy"
	HedgeMLManagementServiceKey     = "app-hedge-ml-management"
	HedgeMLBrokerServiceKey         = "app-hedge-ml-broker"
	HedgeMLEdgeAgentServiceKey      = "app-hedge-ml-edge-agent"
	HedgeMLSandboxServiceKey        = "app-hedge-ml-sandbox"
	DeviceVirtual                   = "hedge-device-virtual"
	HedgeMetaDataNotifierServiceKey = "app-hedge-metadata-notifier"
	DigitalTwinServiceKey           = "app-hedge-digital-twin"
	SimulationPipelineServiceKey    = "app-simulation-pipeline"
)

const (
	CommandUpdateTicketId = "UpdateTicketId"
	CommandName           = "Command"
)

const (
	MetricEvent       = "IoTEvent"
	MetricTicket      = "Ticket"
	LabelDeviceName   = "device"
	LabelDeviceLabels = "tags"
	LabelProfileName  = "profile"
	LabelNodeName     = "host"
	// LabelLocation            = "site"
	LabelCorrelationId = "correlation_id"
	LabelEventType     = "type"
	LabelTicketType    = "type"
	LabelEventSummary  = "summary"
	LabelTicketSummary = "summary"
	LabelNameSpace     = "namespace"
	LabelNewRequest    = "NewRequest"
	LabelCloseRequest  = "CloseRequest"
	LabelDeviceCommand = "DeviceCommand"
	LabelStatus        = "status"
	LabelCommandType   = "commandType"
	StatusSuccess      = "Success"
	StatusFail         = "Failed"
	StatusClosed       = "Closed"
)
