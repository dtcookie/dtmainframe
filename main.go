package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("dtmainframe v2.0.0")
	demoLiveBaseURL := os.Getenv("DT_URL")
	if len(demoLiveBaseURL) == 0 {
		fmt.Println("You need to set the environment variable 'DT_URL'.")
		fmt.Println("For a Managed Cluster that would look like 'https://<dynatrace-server/e/<environment-id>'")
		fmt.Println("For a SaaS Cluster it would look like 'http://<environment-id>.live.dynatrace.com")
		os.Exit(0)
	}
	demoLIveApiToken := os.Getenv("DT_API_TOKEN")
	if len(demoLIveApiToken) == 0 {
		fmt.Println("You need to set the environment variable 'DT_API_TOKEN'.")
		fmt.Println("This token requires the permissions: 'Access problem and event feed, metrics, and topology', 'Read configuration', and 'Write configuration'.")
		os.Exit(0)
	}
	demoLIveClient := NewClient(new(Config), demoLiveBaseURL, New(demoLIveApiToken))
	cicsServices := CICSServices{}
	faked, err := cicsServices.Fetch(demoLIveClient)
	if err != nil {
		panic(err)
	}

	client := demoLIveClient
	var serviceMetric *ServiceMetric
	if faked {
		serviceMetric = NewIterationTransactionRespTime()
	} else {
		serviceMetric = NewMainframeCICSTransactionRespTime()
	}
	serviceMetricIID, err := CreateOrUpdateServiceMetric(client, serviceMetric)
	if err != nil {
		panic(err)
	}

	subDashboards := []SubDashboard{}

	for _, cicsService := range cicsServices {
		cicsServiceID := cicsService.ID
		dashboardName := "Mainframe - " + cicsService.Name
		dashboard := GenDashboard(dashboardName, cicsServiceID, serviceMetricIID)
		if len(cicsService.KeyRequests) > 0 {
			idx := 0
			for _, keyRequestID := range cicsService.KeyRequests {
				idx++
				newTile := NewServiceTile("Service or request", Bounds{
					Top:    idx * 228,
					Left:   266,
					Width:  684,
					Height: 228,
				}, []string{
					keyRequestID,
					cicsServiceID,
				})
				dashboard.Tiles = append(dashboard.Tiles, newTile)
			}
		} else {
			dashboard.Tiles = append(dashboard.Tiles, Tile{
				Name:       "NoKeyRequests",
				TileType:   TileTypes.Markdown,
				Configured: true,
				Bounds:     Bounds{Top: 228, Left: 266, Width: 684, Height: 38},
				TileFilter: TileFilter{},
				Markdown:   strPtr("## There are currently no Key Requests configured"),
			})
		}

		dashboardID, err := CreateOrUpdateDashboard(client, dashboard)
		subDashboards = append(subDashboards, SubDashboard{
			ID:        dashboardID,
			Name:      dashboardName,
			ServiceID: cicsServiceID,
		})
		if err != nil {
			panic(err)
		}
	}

	mainDashboard := GenMainDashboard("Mainframe Dashboard", serviceMetricIID, subDashboards)
	_, err = CreateOrUpdateDashboard(client, mainDashboard)
	if err != nil {
		panic(err)
	}

}
