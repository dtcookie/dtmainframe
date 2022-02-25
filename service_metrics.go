package main

import (
	"encoding/json"
	"fmt"
)

func falsePtr() *bool {
	f := false
	return &f
}

func strPtr(s string) *string {
	return &s
}

func NewMainframeCICSTransactionRespTime() *ServiceMetric {
	return &ServiceMetric{
		TSMMetricKey: strPtr("calc:service.mainframecicstransactionresp.time"),
		Name:         "Mainframe CICS Transaction Resp. Time",
		Enabled:      true,
		MetricDefinition: MetricDefinition{
			Metric:           "RESPONSE_TIME",
			RequestAttribute: nil,
		},
		Unit:            Units.MicroSecond,
		UnitDisplayName: "",
		EntityId:        nil,
		ManagementZones: []string{},
		Conditions: []Condition{
			{
				Attribute: "CICS_TRANSACTION_ID",
				ComparisonInfo: ComparisonInfo{
					Type:          ComparisonInfoTypes.String,
					Comparison:    Comparisons.Exists,
					Value:         nil,
					Values:        nil,
					Negate:        false,
					CaseSensitive: falsePtr(),
				},
			},
			{
				Attribute: "SERVICE_TYPE",
				ComparisonInfo: ComparisonInfo{
					Type:       ComparisonInfoTypes.ServiceType,
					Comparison: Comparisons.Equals,
					Value:      "CICS_SERVICE",
					Values:     nil,
					Negate:     false,
				},
			},
		},
		DimensionDefinition: DimensionDefinition{
			Name:            "CICS Txn Id",
			Dimension:       "{CICS:TransactionId}",
			Placeholders:    []string{},
			TopX:            10,
			TopXDirection:   TopXDirections.Descending,
			TopXAggregation: TopXAggregations.Sum,
		},
	}
}

type ServiceMetricStub struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ListServiceMetricsResponse struct {
	Values []ServiceMetricStub `json:"values"`
}

func ListServiceMetrics(client *Client) ([]ServiceMetricStub, error) {
	data, err := client.GET("/api/config/v1/calculatedMetrics/service", 200)
	if err != nil {
		return nil, err
	}
	response := ListServiceMetricsResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response.Values, nil
}

func CreateOrUpdateServiceMetric(client *Client, serviceMetric *ServiceMetric) (string, error) {
	serviceMetricID, err := FindServiceMetric(client, serviceMetric.Name)
	if err != nil {
		return "", err
	}
	if len(serviceMetricID) > 0 {
		err = UpdateServiceMetric(client, serviceMetricID, serviceMetric)
		if err != nil {
			return "", err
		}
		return serviceMetricID, nil
	}
	return CreateServiceMetric(client, serviceMetric)
}

func FindServiceMetric(client *Client, name string) (string, error) {
	stubs, err := ListServiceMetrics(client)
	if err != nil {
		return "", err
	}
	for _, stub := range stubs {
		if stub.Name == name {
			return stub.ID, nil
		}
	}
	return "", nil
}

func CreateServiceMetric(client *Client, serviceMetric *ServiceMetric) (string, error) {
	fmt.Println("Creating Service Metric '" + serviceMetric.Name + "'")
	data, err := client.POST("/api/config/v1/calculatedMetrics/service", serviceMetric, 201)
	if err != nil {
		return "", err
	}
	var response CreateServiceMetricResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return "", err
	}
	return response.ID, nil
}

func UpdateServiceMetric(client *Client, id string, serviceMetric *ServiceMetric) error {
	fmt.Println("Updating Service Metric '" + serviceMetric.Name + "'")
	// tsmMetricKey := serviceMetric.TSMMetricKey
	// serviceMetric.TSMMetricKey = nil
	_, err := client.PUT(fmt.Sprintf("/api/config/v1/calculatedMetrics/service/%s", id), serviceMetric, 204)
	// serviceMetric.TSMMetricKey = tsmMetricKey
	return err
}

type CreateServiceMetricResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MetricDefinition struct {
	Metric           string  `json:"metric"`
	RequestAttribute *string `json:"requestAttribute"`
}

type Unit string

var Units = struct {
	MicroSecond Unit
}{
	Unit("MICRO_SECOND"),
}

type Condition struct {
	Attribute      string         `json:"attribute"`
	ComparisonInfo ComparisonInfo `json:"comparisonInfo"`
}

type ServiceMetric struct {
	TSMMetricKey        *string             `json:"tsmMetricKey"`
	Name                string              `json:"name"`
	Enabled             bool                `json:"enabled"`
	MetricDefinition    MetricDefinition    `json:"metricDefinition"`
	Unit                Unit                `json:"unit"`
	UnitDisplayName     string              `json:"unitDisplayName"`
	EntityId            *string             `json:"entityId"`
	ManagementZones     []string            `json:"managementZones"`
	Conditions          []Condition         `json:"conditions"`
	DimensionDefinition DimensionDefinition `json:"dimensionDefinition"`
}

type DimensionDefinition struct {
	Name            string          `json:"name"`
	Dimension       string          `json:"dimension"`
	Placeholders    []string        `json:"placeholders"`
	TopX            int             `json:"topX"`
	TopXDirection   TopXDirection   `json:"topXDirection"`
	TopXAggregation TopXAggregation `json:"topXAggregation"`
}

type TopXAggregation string

var TopXAggregations = struct {
	Sum TopXAggregation
}{
	TopXAggregation("SUM"),
}

type TopXDirection string

var TopXDirections = struct {
	Descending TopXDirection
}{
	TopXDirection("DESCENDING"),
}

type ComparisonInfoType string

var ComparisonInfoTypes = struct {
	String      ComparisonInfoType
	ServiceType ComparisonInfoType
}{
	ComparisonInfoType("STRING"),
	ComparisonInfoType("SERVICE_TYPE"),
}

type Comparison string

var Comparisons = struct {
	Exists Comparison
	Equals Comparison
}{
	Comparison("EXISTS"),
	Comparison("EQUALS"),
}

type ComparisonInfo struct {
	Type          ComparisonInfoType `json:"type"`
	Comparison    Comparison         `json:"comparison"`
	Value         interface{}        `json:"value"`
	Values        []interface{}      `json:"values"`
	Negate        bool               `json:"negate"`
	CaseSensitive *bool              `json:"caseSensitive,omitempty"`
}
