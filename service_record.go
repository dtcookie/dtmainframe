package main

type ServiceRecords []ServiceRecord

type ServiceRecord struct {
	EntityID    string `json:"entityId"`
	DisplayName string `json:"displayName"`
	ServiceType string `json:"serviceType"`
}
