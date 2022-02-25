package main

import "encoding/json"

type Records []Record

type Record struct {
	EntityID               string   `json:"entityId"`
	DisplayName            string   `json:"displayName"`
	DiscoveredName         string   `json:"discoveredName"`
	ServiceTechnologyTypes []string `json:"serviceTechnologyTypes"`
	ServiceType            string   `json:"serviceType"`
}

func (record Record) String() string {
	data, err := json.MarshalIndent(record, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}
