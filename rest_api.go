package main

import (
	"encoding/json"
	"fmt"
)

// func FetchCICSServices(client *Client) (CICSServices, error) {
// 	data, err := client.GET("/api/v1/entity/services?relativeTime=3days&includeDetails=false", 200)
// 	if err != nil {
// 		return nil, err
// 	}
// 	records := ServiceRecords{}
// 	err = json.Unmarshal(data, &records)
// 	if err != nil {
// 		return nil, err
// 	}
// 	result := CICSServices{}
// 	for _, record := range records {
// 		if strings.ToUpper(record.ServiceType) == "CICS" {
// 			result[record.EntityID] = CICSService{
// 				ID:          record.EntityID,
// 				Name:        record.DisplayName,
// 				KeyRequests: []string{},
// 			}
// 		}
// 	}
// 	return result, nil
// }

func FetchKeyRequestIDs(client *Client, serviceID string) ([]string, error) {
	path := fmt.Sprintf("/api/v1/timeseries/com.dynatrace.builtin%%3Aservicemethod.responsetime?includeData=true&aggregationType=AVG&relativeTime=15mins&queryMode=TOTAL&entity=%s", serviceID)
	data, err := client.GET(path, 200)
	if err != nil {
		return nil, err
	}
	response := TimeSeriesResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	result := []string{}
	for k := range response.DataResult.Entities {
		result = append(result, k)
	}
	return result, nil
}

// func FetchCICSServices(client *Client) (CICSServices, error) {
// 	result := CICSServices{}
// 	cicsServiceIDs, err := FetchCICSServiceIDs(client)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, cicsServiceID := range cicsServiceIDs {
// 		keyRequestIDs, err := FetchKeyRequestIDs(client, cicsServiceID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		result = append(result, CICSService{
// 			ID:   cicsServiceID,
// 			Name: string,
// 		})
// 	}

// 	return result, nil
// }
