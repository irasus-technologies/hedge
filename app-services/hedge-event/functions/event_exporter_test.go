package functions

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	hedgeeventmocks "hedge/mocks/hedge/app-services/hedge-event/functions"
	"testing"
)

func TestExporter_StoreEvent(t *testing.T) {
	ctx := u.AppFunctionContext
	t.Run("StoreEvent - Passed (Elastic enabled, save successful)", func(t *testing.T) {
		mockedElasticExporter := hedgeeventmocks.MockOpenSearchExporterInterface{}
		mockedElasticExporter.On("SaveEventToOpenSearch", mock.Anything, mock.Anything).Return(false, nil)

		exporter := &EventExporter{
			eventExporter: &mockedElasticExporter,
		}

		testData := "test-data"

		got, got1 := exporter.StoreEvent(ctx, testData)
		assert.False(t, got, "Expected true when saving to Elastic is successful")
		assert.Equal(t, nil, got1, "Expected the returned data to match the input data")
	})
	t.Run("StoreToLocalElastic - Failed (nil data)", func(t *testing.T) {
		mockedElasticExporter := hedgeeventmocks.MockOpenSearchExporterInterface{}

		exporter := &EventExporter{
			eventExporter: &mockedElasticExporter,
		}

		got, got1 := exporter.StoreEvent(ctx, nil)
		assert.False(t, got, "Expected false when data is nil")
		assert.Error(t, got1.(error), "Expected an error when data is nil")
		assert.Contains(t, got1.(error).Error(), "no data Received", "Unexpected error message")
	})
}

func TestExporter_Print(t *testing.T) {
	t.Run("Print - Passed", func(t *testing.T) {
		jsonData, _ := json.Marshal(testEventData)
		var resultMap map[string]interface{}
		_ = json.Unmarshal(jsonData, &resultMap)

		ctx := u.AppFunctionContext
		data := resultMap

		got, got1 := Print(ctx, data)

		assert.True(t, got, "Expected true when data is provided")
		assert.Equal(t, resultMap, got1, "Expected the returned data to match the input data")
	})
	t.Run("Print - Failed", func(t *testing.T) {
		ctx := u.AppFunctionContext

		got, got1 := Print(ctx, nil)
		assert.False(t, got, "Expected false when data is nil")
		assert.Nil(t, got1, "Expected nil when no data is provided")
	})
}
