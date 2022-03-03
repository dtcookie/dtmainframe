package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type CICSServices map[string]*CICSService

func (me CICSServices) Fetch(client *Client) (bool, error) {
	faked := false
	var data []byte
	var err error
	if Exists("api_vi_entity_services.json") {
		faked = true
		data, err = ioutil.ReadFile("api_vi_entity_services.json")
	} else {
		data, err = client.GET("/api/v1/entity/services?relativeTime=3days&includeDetails=false", 200)
	}
	if err != nil {
		return faked, err
	}
	records := ServiceRecords{}
	err = json.Unmarshal(data, &records)
	if err != nil {
		return faked, err
	}
	for _, record := range records {
		if faked {
			if strings.ToUpper(record.ServiceType) == "METHOD" {
				cicsService := &CICSService{
					ID:          record.EntityID,
					Name:        record.DisplayName,
					KeyRequests: []string{},
				}
				err := cicsService.Populate(client)
				if err != nil {
					return faked, err
				}
				me[record.EntityID] = cicsService
			}
		} else {
			if strings.ToUpper(record.ServiceType) == "CICS" {
				cicsService := &CICSService{
					ID:          record.EntityID,
					Name:        record.DisplayName,
					KeyRequests: []string{},
				}
				err := cicsService.Populate(client)
				if err != nil {
					return faked, err
				}
				me[record.EntityID] = cicsService
			}
		}
	}
	return faked, nil
}

type CICSService struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	KeyRequests []string `json:"keyRequests"`
}

func (me *CICSService) Populate(client *Client) error {
	path := fmt.Sprintf("/api/v1/timeseries/com.dynatrace.builtin%%3Aservicemethod.responsetime?includeData=true&aggregationType=AVG&relativeTime=2hours&queryMode=TOTAL&entity=%s", me.ID)
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
