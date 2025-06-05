package functions

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hedge/common/dto"
	clientmocks "hedge/mocks/hedge/common/client"
	"testing"
)

/*func TestNewElasticExporter(t *testing.T) {
	elasticEventExporter := NewOpenSearchExporter(u.AppService)
	assert.Nil(t, elasticEventExporter)
}*/

func TestElasticExporter_SaveEventToElastic(t *testing.T) {
	t.Run("SaveEventToOpenSearch - Passed (Index successful)", func(t *testing.T) {
		mockedElasticClient := clientmocks.MockElasticClientInterface{}
		mockedElasticClient.On("IndexEvent", mock.Anything).Return(nil)

		elasticEvent := OpenSearchExporter{
			openSearchClient: &mockedElasticClient,
			lc:               u.AppService.LoggingClient(),
		}

		ok, gotData := elasticEvent.SaveEventToOpenSearch(u.AppFunctionContext, testEventData)
		assert.True(t, ok, "Expected SaveEventToOpenSearch to return true for valid data")
		assert.Equal(t, testEventData, gotData, "Expected returned data to match input")
	})
	t.Run("SaveEventToOpenSearch - Passed (Search successful)", func(t *testing.T) {
		mockedElasticClient := clientmocks.MockElasticClientInterface{}
		mockedElasticClient.On("SearchEvents", mock.Anything).Return([]*dto.HedgeEvent{&testEventData1}, nil)
		mockedElasticClient.On("IndexEvent", mock.Anything).Return(nil)

		elasticEvent := OpenSearchExporter{
			openSearchClient: &mockedElasticClient,
			lc:               u.AppService.LoggingClient(),
		}

		ok, gotData := elasticEvent.SaveEventToOpenSearch(u.AppFunctionContext, testEventData1)
		assert.True(t, ok, "Expected SaveEventToOpenSearch to return true for search success")
		assert.Equal(t, testEventData1, gotData, "Expected returned data to match input")
	})
	t.Run("SaveEventToOpenSearch - Failed (Search failed)", func(t *testing.T) {
		mockedElasticClient := clientmocks.MockElasticClientInterface{}
		mockedElasticClient.On("SearchEvents", mock.Anything).Return(nil, testError)

		elasticEvent := OpenSearchExporter{
			openSearchClient: &mockedElasticClient,
			lc:               u.AppService.LoggingClient(),
		}

		ok, gotData := elasticEvent.SaveEventToOpenSearch(u.AppFunctionContext, testEventData1)
		assert.False(t, ok, "Expected SaveEventToOpenSearch to return false for nil data")
		assert.Error(t, gotData.(error), "Expected an error for nil data")
		assert.Contains(t, gotData.(error).Error(), "error while searching Open event with correlationId: : "+testError.Error(), "Unexpected error message for nil data")
	})
	t.Run("SaveEventToOpenSearch - Failed (Nil data)", func(t *testing.T) {
		mockedElasticClient := clientmocks.MockElasticClientInterface{}

		elasticEvent := OpenSearchExporter{
			openSearchClient: &mockedElasticClient,
			lc:               u.AppService.LoggingClient(),
		}

		ok, gotData := elasticEvent.SaveEventToOpenSearch(u.AppFunctionContext, nil)
		assert.False(t, ok, "Expected SaveEventToOpenSearch to return false for nil data")
		assert.Error(t, gotData.(error), "Expected an error for nil data")
		assert.Contains(t, gotData.(error).Error(), "no Data Received", "Unexpected error message for nil data")
	})
	t.Run("SaveEventToOpenSearch - Failed (Invalid data)", func(t *testing.T) {
		mockedElasticClient := clientmocks.MockElasticClientInterface{}

		elasticEvent := OpenSearchExporter{
			openSearchClient: &mockedElasticClient,
			lc:               u.AppService.LoggingClient(),
		}

		invalidData := []byte("wrong data")

		ok, gotData := elasticEvent.SaveEventToOpenSearch(u.AppFunctionContext, invalidData)
		assert.False(t, ok, "Expected SaveEventToOpenSearch to return false for invalid data type")
		assert.Error(t, gotData.(error), "Expected an error for invalid data type")
		assert.Contains(t, gotData.(error).Error(), "error while unmarshalling data", "Unexpected error message for invalid data type")
	})
	t.Run("SaveEventToOpenSearch - Failed (Index failed)", func(t *testing.T) {
		mockedElasticClient := clientmocks.MockElasticClientInterface{}
		mockedElasticClient.On("IndexEvent", mock.Anything).Return(testError)

		elasticEvent := OpenSearchExporter{
			openSearchClient: &mockedElasticClient,
			lc:               u.AppService.LoggingClient(),
		}

		ok, gotData := elasticEvent.SaveEventToOpenSearch(u.AppFunctionContext, testEventData)
		assert.False(t, ok, "Expected SaveEventToOpenSearch to return false for index failure")
		assert.Error(t, gotData.(error), "Expected an error for index failure")
		assert.Contains(t, gotData.(error).Error(), "dummy error", "Unexpected error message for index failure")
	})
	t.Run("SaveEventToOpenSearch - Failed (Unmarshalling error)", func(t *testing.T) {
		mockedElasticClient := clientmocks.MockElasticClientInterface{}
		mockedElasticClient.On("IndexEvent", mock.Anything).Return(testError)

		elasticEvent := OpenSearchExporter{
			openSearchClient: &mockedElasticClient,
			lc:               u.AppService.LoggingClient(),
		}

		testData := "valid data"

		ok, gotData := elasticEvent.SaveEventToOpenSearch(u.AppFunctionContext, testData)
		assert.False(t, ok, "Expected SaveEventToOpenSearch to return false for index failure")
		assert.Error(t, gotData.(error), "Expected an error for index failure")
		assert.Contains(t, gotData.(error).Error(), "error while unmarshalling data", "Unexpected error message for index failure")
	})
	t.Run("SaveMLPredictionToElastic - Failed (Marshalling error)", func(t *testing.T) {
		mockedElasticClient := &clientmocks.MockElasticClientInterface{}

		elasticEvent := OpenSearchExporter{
			openSearchClient: mockedElasticClient,
			lc:               u.AppService.LoggingClient(),
		}

		testData := make(chan int) // Invalid data type for marshalling

		ok, gotData := elasticEvent.SaveEventToOpenSearch(u.AppFunctionContext, testData)
		assert.False(t, ok, "Expected SaveEventToOpenSearch to return false for index failure")
		assert.Error(t, gotData.(error), "Expected an error for index failure")
		assert.Contains(t, gotData.(error).Error(), "error while marshalling data", "Unexpected error message for index failure")
	})
}
