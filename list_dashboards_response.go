package main

import (
	"encoding/json"
	"fmt"
)

type DashboardStub struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type ListDashboardsResponse struct {
	Dashboards []DashboardStub `json:"dashboards"`
}

func ListDashboards(client *Client) ([]DashboardStub, error) {
	data, err := client.GET("/api/config/v1/dashboards?tags=generated", 200)
	if err != nil {
		return nil, err
	}
	response := ListDashboardsResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return response.Dashboards, nil
}

func FindDashboard(client *Client, name string) (string, error) {
	stubs, err := ListDashboards(client)
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

type CreateDashboardResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateDashboard(client *Client, dashboard *Dashboard) (string, error) {
	fmt.Println("Creating Dashboard '" + dashboard.DashboardMetadata.Name + "'")
	data, err := client.POST("/api/config/v1/dashboards", dashboard, 201)
	if err != nil {
		return "", err
	}
	var response CreateDashboardResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return "", err
	}
	return response.ID, nil
}

func UpdateDashboard(client *Client, id string, dashboard *Dashboard) error {
	fmt.Println("Updating Dashboard '" + dashboard.DashboardMetadata.Name + "'")
	dashboard.ID = id
	_, err := client.PUT(fmt.Sprintf("/api/config/v1/dashboards/%s", id), dashboard, 204)
	return err
}

func CreateOrUpdateDashboard(client *Client, dashboard *Dashboard) (string, error) {
	dashboardID, err := FindDashboard(client, dashboard.DashboardMetadata.Name)
	if err != nil {
		return "", err
	}
	if len(dashboardID) > 0 {
		dashboard.ID = dashboardID
		if err = UpdateDashboard(client, dashboardID, dashboard); err != nil {
			return "", err
		}
		return dashboardID, nil

	}
	return CreateDashboard(client, dashboard)
}
