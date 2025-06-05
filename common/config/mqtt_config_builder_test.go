/*******************************************************************************
* Contributors: BMC Software, Inc. - BMC Helix Edge
*
* (c) Copyright 2020-2025 BMC Software, Inc.
*******************************************************************************/

package config

import (
	"hedge/mocks/hedge/common/infrastructure/interfaces/utils"
	"testing"
)

func TestBuildMQTTSecretConfig(t *testing.T) {
	//var service interfaces.ApplicationService

	hedgeMockUtils := utils.NewApplicationServiceMock(nil)
	hedgeMockUtils.InitMQTTSettings()
	mqttConfig, err := BuildMQTTSecretConfig(hedgeMockUtils.AppService, "events", "clientId001")

	if err != nil {
		t.Errorf("BuildMQTTSecretConfig failed, err:%s", err.Error())
	}

	if mqttConfig.Topic != "hedge/events" {
		t.Errorf("got %s, expected events", mqttConfig.Topic)
	}
}
