package main

type TimeSeriesResponse struct {
	DataResult struct {
		Entities map[string]string `json:"entities"`
	} `json:"dataResult"`
}
