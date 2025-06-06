/*******************************************************************************
* Contributors: BMC Helix, Inc.
*
* (c) Copyright 2020-2025 BMC Helix, Inc.

* SPDX-License-Identifier: Apache-2.0
*******************************************************************************/

package functions

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/startup"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	"github.com/hashicorp/go-uuid"
	"github.com/opensearch-project/opensearch-go/v4"
	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/opensearch-project/opensearch-go/v4/opensearchutil"
	"github.com/pkg/errors"
	hedgedto "hedge/common/dto"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Refer https://github.com/opensearch-project/opensearch-go/blob/main/USER_GUIDE.md for example implementation code
type ElasticClientInterface interface {
	SearchEvents(luceneQuery string) ([]*hedgedto.HedgeEvent, error)
	Search(luceneQuery string, indexName string) (map[string]interface{}, error)
	IndexEvent(hedgeEvent *hedgedto.HedgeEvent) error
	Index(req opensearchapi.IndexReq) error
	ConvertToHedgeEvents(result map[string]interface{}) ([]*hedgedto.HedgeEvent, error)
	BuildSearchRequest(luceneQuery string, indexName string) opensearchapi.SearchReq
}

type HedgeOpenSearchClient struct {
	Client                    *opensearch.Client
	logger                    logger.LoggingClient
	PredictionTemplatePattern string
}

type BulkResponse struct {
	Status int  `json:"status" validate:"required"`
	Errors bool `json:"errors" validate:"required"`
}

const (
	ElasticEventIndexName = "event_index"
)

/*
New opensearch Client creation from configuration in configuration.toml
The following configuration is required
OpenSearchURL, secrets -- username, password
elastic index name to be passed as param:
event_index for events
remediate_index for remediation
*/

func NewHedgeOpenSearchClient(service interfaces.ApplicationService) *HedgeOpenSearchClient {

	var openSearchClient *HedgeOpenSearchClient
	var err error

	logger := service.LoggingClient()
	logger.Info("About to create the elastic client")

	startupTimer := startup.NewStartUpTimer("opensearch-client")
	for startupTimer.HasNotElapsed() {

		openSearchClient, err = createHedgeElasticClient(service)
		if err == nil {
			break
		}
		openSearchClient = nil
		fmt.Printf("Couldn't create opensearch client: %v", err.Error())
		startupTimer.SleepForInterval()
	}
	if openSearchClient == nil {
		fmt.Printf("Failed to create opensearch client in allotted time")
		os.Exit(1)
	}

	return openSearchClient
}

func createHedgeElasticClient(service interfaces.ApplicationService) (*HedgeOpenSearchClient, error) {

	logger := service.LoggingClient()

	var elasticURL []string
	var username, password string
	//var secrets map[string]string
	var err error
	var skipCertVerification bool // false by default

	const (
		OPENSEARCHURL = "OpenSearchURL"
		//TEMPLATEPATTERN      = "TemplatePattern"
		SECRETS              = "secrets"
		SKIPCERTVERIFICATION = "SkipCertVerification"
	)

	properties := []string{OPENSEARCHURL, SECRETS, SKIPCERTVERIFICATION}
	errorData := ""

	for _, p := range properties {
		switch p {
		case OPENSEARCHURL:
			elasticURL, err = service.GetAppSettingStrings(OPENSEARCHURL)
			errorData = "elasticURL"
		case SECRETS:
			/*
				// Below code doesn't work since we don't add opensearch secrets as part of install
					secrets, err = service.SecretProvider().GetSecret("opensearch", "username", "password")
				username = secrets["username"]
				password = secrets["password"]
				errorData = "secrets"*/
			// Below works mostly may be because security is disabled, need to fix this
			username = ""
			password = ""

		case SKIPCERTVERIFICATION:
			value, _ := service.GetAppSetting(SKIPCERTVERIFICATION)
			skipCertVerification, _ = strconv.ParseBool(value)
			logger.Infof("SkipCertVerification: %t", skipCertVerification)
		}
		if err != nil {
			logger.Errorf("Could not read %s, error: %s", errorData, err.Error())
			return nil, errors.Wrapf(err, "Could not read %s", errorData)
		}
	}
	logger.Infof("OpenSearch URL: %s, skip certificate verification: %t",
		elasticURL, skipCertVerification)
	logger.Infof("Read opensearch secrets")

	tp := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipCertVerification},
	}
	cfg := opensearch.Config{
		Addresses: elasticURL,
		Username:  username,
		Password:  password,
		Transport: tp,
	}

	openSearchClient, err := opensearch.NewClient(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("ERROR: Unable to create client: %v\n", err))
		return nil, errors.Wrapf(err, "failed to create elasticsearch client for %s", elasticURL)
	}

	return &HedgeOpenSearchClient{
		Client: openSearchClient,
		logger: logger,
	}, nil
}

func (e *HedgeOpenSearchClient) SearchEvents(luceneQuery string) ([]*hedgedto.HedgeEvent, error) {
	// Print the response status and indexed document version.
	hits, err := e.Search(luceneQuery, ElasticEventIndexName)
	if err != nil {
		return nil, err
	}
	events, err := e.ConvertToHedgeEvents(hits)
	return events, err
}

func (e *HedgeOpenSearchClient) Search(luceneQuery string, indexName string) (map[string]interface{}, error) {
	searchReq := e.BuildSearchRequest(luceneQuery, indexName)

	// Perform the SearchEvents request.
	res, err := e.Client.Do(context.Background(), searchReq, nil)

	if err != nil {
		e.logger.Error(fmt.Sprintf("\"Error getting response: %v\n", err))
		return nil, err
	}

	defer res.Body.Close()

	if res.IsError() {
		e.logger.Error(fmt.Sprintf("[%s] Error retrieving the document\n", res.Status()))
		return nil, err
	}
	// Deserialize the response into a map.
	hits := make(map[string]interface{}, 1)
	if err := json.NewDecoder(res.Body).Decode(&hits); err != nil {
		e.logger.Warn(fmt.Sprintf("search: Error parsing the response body: %v\n", err))
		return nil, err
	}
	return hits, nil
}

func (e *HedgeOpenSearchClient) IndexEvent(hedgeEvent *hedgedto.HedgeEvent) error {
	e.logger.Debug("Adding to event database")
	currentTime := time.Now().UnixNano() / 1000000
	if hedgeEvent.Created == 0 {
		hedgeEvent.Created = currentTime // Don't overwrite createDate for existing docs
	}
	hedgeEvent.Modified = currentTime

	// If ID is empty, generate one
	if len(hedgeEvent.Id) == 0 {
		hedgeEvent.Id, _ = uuid.GenerateUUID()
	}

	req := opensearchapi.IndexReq{
		Index:      ElasticEventIndexName,
		DocumentID: hedgeEvent.Id,
		Body:       opensearchutil.NewJSONReader(hedgeEvent),
		Params: opensearchapi.IndexParams{
			Refresh: "true",
		},
	}

	e.logger.Debug(fmt.Sprintf("Document being saved=%s", hedgeEvent.Id))
	return e.Index(req)
}

func (e *HedgeOpenSearchClient) Index(req opensearchapi.IndexReq) error {
	// Perform the request with the client.
	res, err := e.Client.Do(context.Background(), req, nil)
	if err != nil {
		e.logger.Errorf("Error indexing the request, error response: %v", err)
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		bodyBytes, _ := io.ReadAll(res.Body)
		e.logger.Error(fmt.Sprintf("[%s] Error response from Elasticsearch: %s", res.Status(), string(bodyBytes)))
		e.logger.Error(fmt.Sprintf("[%s] Error indexing document ID=%s\n", res.Status(), req.DocumentID))
		return errors.New(fmt.Sprintf("Error indexing document: %s\n", res.String()))
	}
	// Deserialize the response into a map.
	r := make(map[string]interface{}, 1)

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		e.logger.Warn(fmt.Sprintf("IndexRequest: Error parsing the response body: %v\n", err))
		return err
	}
	// Print the response status and indexed document version.
	version := int(r["_version"].(float64))
	e.logger.Debugf("[%s] %s; version=%d", res.Status(), r["result"], version)
	return nil
}

func (e *HedgeOpenSearchClient) ConvertToHedgeEvents(result map[string]interface{}) ([]*hedgedto.HedgeEvent, error) {

	events := make([]*hedgedto.HedgeEvent, 0)
	if result["hits"] == nil {
		return events, nil
	}
	hits := result["hits"].(map[string]interface{})
	if hits == nil || hits["hits"] == nil {
		e.logger.Info("No data found")
		return events, nil
	}

	hitsArray := hits["hits"].([]interface{})

	for i := range hitsArray {
		element := hitsArray[i].(map[string]interface{})
		hedgeEvent := new(hedgedto.HedgeEvent)
		hedgeEvent.Id = element["_id"].(string)
		events = append(events, hedgeEvent)

		var source map[string]interface{}
		ok := false
		if source, ok = element["_source"].(map[string]interface{}); !ok {
			return events, nil
		}

		*hedgeEvent = hedgedto.HedgeEvent{
			Id:             e.getString(element, "_id"),
			DeviceName:     e.getString(source, "device_name"),
			Class:          e.getString(source, "class"),
			EventType:      e.getString(source, "event_type"),
			Name:           e.getString(source, "name"),
			Msg:            e.getString(source, "msg"),
			SourceNode:     e.getString(source, "source_node"),
			Severity:       e.getString(source, "severity"),
			RelatedMetrics: e.getStringList(source, "related_metrics"),
			Thresholds:     e.getMap(source, "thresholds"),
			ActualValues:   e.getMap(source, "actual_values"),
			Unit:           e.getString(source, "unit"),
			Location:       e.getString(source, "location"),
			CorrelationId:  e.getString(source, "correlation_id"),
			Created:        e.getInt(source, "created"),
			Modified:       e.getInt(source, "modified"),
			Status:         e.getString(source, "status"),
			Profile:        e.getString(source, "profile"),
		}

		if source["remediations"] != nil && source["remediations"].([]interface{}) != nil {
			remediationsCollection := source["remediations"].([]interface{})
			remediations := make([]hedgedto.Remediation, len(remediationsCollection))
			for i, remediation := range remediationsCollection {
				remediationMap := remediation.(map[string]interface{})
				remediations[i] = hedgedto.Remediation{
					Id:           e.getString(remediationMap, "id"),
					Type:         e.getString(remediationMap, "type"),
					Summary:      e.getString(remediationMap, "summary"),
					Status:       e.getString(remediationMap, "status"),
					ErrorMessage: e.getString(remediationMap, "error_message"),
				}

			}
			hedgeEvent.Remediations = remediations
		}

		// if the status has not changed, retain the additional data that was present earlier
		// The rule keeps firing and rule output doesn't include additionalData
		additionalData, ok := source["additional_data"].(map[string]interface{})
		if ok {
			hedgeEvent.AdditionalData = make(map[string]string)
			for key, val := range additionalData {
				if _, ok := val.(string); ok {
					hedgeEvent.AdditionalData[key], _ = val.(string)
				}
			}
		}

	}

	return events, nil
}

/*
Utility method to build a SearchRequest
luceneQuery example:
correlationId: \"someCorreId\" AND _id:"\someid\"
*/
func (e *HedgeOpenSearchClient) BuildSearchRequest(luceneQuery string, indexName string) opensearchapi.SearchReq {
	//ctx := context.Background()
	index := []string{indexName}
	query := map[string]map[string]map[string]string{
		"query": {
			"query_string": {
				"query": luceneQuery,
			},
		},
	}

	searchReq := opensearchapi.SearchReq{
		Indices: index,
		Body:    opensearchutil.NewJSONReader(&query),
		Params: opensearchapi.SearchParams{
			Pretty: true,
		},
	}
	return searchReq
}

func (e *HedgeOpenSearchClient) getString(data map[string]interface{}, key string) string {
	if val, ok := data[key]; ok {
		return val.(string)
	}
	return ""
}

func (e *HedgeOpenSearchClient) getStringList(data map[string]interface{}, key string) []string {
	if val, ok := data[key]; ok && val != nil {
		if val2, ok := val.([]string); ok {
			return val2
		}
	}
	return []string{}
}

func (e *HedgeOpenSearchClient) getInt(data map[string]interface{}, key string) int64 {
	if val, ok := data[key]; ok && val != nil {
		if val2, ok := val.(int64); ok {
			return val2
		}
	}
	return 0
}

func (e *HedgeOpenSearchClient) getMap(data map[string]interface{}, key string) map[string]interface{} {
	if val, ok := data[key]; ok && val != nil {
		if val2, ok := val.(map[string]interface{}); ok {
			return val2
		}
	}
	return map[string]interface{}{}
}
