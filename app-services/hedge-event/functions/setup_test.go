package functions

import (
	"errors"
	"hedge/common/dto"
	"hedge/mocks/hedge/common/infrastructure/interfaces/utils"
)

var (
	u              *utils.HedgeMockUtils
	testEventData  dto.HedgeEvent
	testEventData1 dto.HedgeEvent
	testError      = errors.New("dummy error")
)

func init() {
	u = utils.NewApplicationServiceMock(map[string]string{"ElasticURL": "", "ElasticUser": "User", "ElasticPassword": "pass"})

	testEventData = dto.HedgeEvent{
		Id:             "f1c5f0e8-6b64-48ad-91b7-c5981b5ca3b9",
		Status:         "Open",
		Thresholds:     map[string]interface{}{"Threshold": float64(0)},
		ActualValues:   map[string]interface{}{"ActualValue": float64(0)},
		AdditionalData: map[string]string{"key": "value"},
		CorrelationId:  "123456",
		Remediations: []dto.Remediation{
			{
				Id:      "123",
				Status:  "OK",
				Type:    "type",
				Summary: "summary",
			},
		},
		IsNewEvent:       false,
		IsRemediationTxn: false,
	}

	testEventData1 = dto.HedgeEvent{
		Id:               "f1c5f0e8-6b64-48ad-91b7-c5981b5ca3b1",
		Status:           "Closed",
		IsNewEvent:       true,
		CorrelationId:    "",
		Priority:         "MEDIUM",
		Remediations:     []dto.Remediation{{Id: "123", Type: "type", Summary: "summary", Status: "OK"}},
		IsRemediationTxn: true,
	}
}
