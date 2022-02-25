package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CICSServices map[string]*CICSService

func (me CICSServices) Fetch(client *Client) error {
	data, err := client.GET("/api/v1/entity/services?relativeTime=3days&includeDetails=false", 200)
	if err != nil {
		return err
	}
	records := ServiceRecords{}
	err = json.Unmarshal(data, &records)
	if err != nil {
		return err
	}
	for _, record := range records {
		if strings.ToUpper(record.ServiceType) == "CICS" {
			cicsService := &CICSService{
				ID:          record.EntityID,
				Name:        record.DisplayName,
				KeyRequests: []string{},
			}
			err := cicsService.Populate(client)
			if err != nil {
				return err
			}
			me[record.EntityID] = cicsService
		}
	}
	return nil
}

type CICSService struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	KeyRequests []string `json:"keyRequests"`
}

func (me *CICSService) Populate(client *Client) error {
	path := fmt.Sprintf("/api/v1/timeseries/com.dynatrace.builtin%%3Aservicemethod.responsetime?includeData=true&aggregationType=AVG&relativeTime=15mins&queryMode=TOTAL&entity=%s", me.ID)
	data, err := client.GET(path, 200)
	if err != nil {
		return err
	}
	response := TimeSeriesResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}
	me.KeyRequests = []string{}
	for k := range response.DataResult.Entities {
		me.KeyRequests = append(me.KeyRequests, k)
	}
	return nil
}
