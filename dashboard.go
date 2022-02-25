package main

type SubDashboard struct {
	ID        string
	Name      string
	ServiceID string
}

func GenMainDashboard(name string, serviceMetricID string, subDashboards []SubDashboard) *Dashboard {
	markdown := ""
	if len(subDashboards) > 0 {
		for _, subDashboard := range subDashboards {
			markdown = markdown + "[" + subDashboard.Name + "](#dashboard;id=" + subDashboard.ID + ")\n\n"
		}
	}
	navTile := NewMarkdownTile(
		"Markdown",
		Bounds{
			Top:    0,
			Left:   0,
			Width:  152,
			Height: 646,
		},
		markdown,
	)
	dashboard := &Dashboard{
		DashboardMetadata: DashboardMetadata{
			Name:            name,
			Shared:          false,
			Owner:           "reinhard.pilz@dynatrace.com",
			DashboardFilter: DashboardFilter{},
			Tags:            []string{"generated"},
			Popularity:      1,
		},
		Tiles: []Tile{
			{
				Name:       "Top List",
				TileType:   TileTypes.DataExplorer,
				Configured: true,
				Bounds: Bounds{
					Top:    0,
					Left:   798,
					Width:  266,
					Height: 646,
				},
				TileFilter: TileFilter{},
				CustomName: "Top List",
				Queries: Queries{
					{
						ID: "A",
						// Metric:           "calc:service.mainframecicstransactionresp.time",
						Metric:           serviceMetricID,
						SpaceAggregation: SpaceAggregations.Avg,
						TimeAggregation:  TimeAggregations.Default,
						SplitBy: []string{
							"CICS Txn Id",
							"dt.entity.service",
						},
						SortBy: "DESC",
						FilterBy: Filter{
							Operator:      FilterOperators.And,
							NestedFilters: Filters{},
							Criteria:      Criteria{},
						},
						Limit:   100,
						Enabled: true,
					},
				},
				VisualConfig: &VisualConfig{
					Type:   VisualConfigTypes.TopList,
					Global: Global{HideLegend: false},
					Rules: VisualConfigRules{
						{
							Matcher: "A:",
							Properties: Properties{
								Color: Colors.Default,
							},
							SeriesOverride: []string{},
						},
					},
					Axes: Axes{
						XAxis: XAxis{Visible: true},
						YAxes: []string{},
					},
					HeatmapSettings: HeatmapSettings{},
					Thresholds: []Threshold{
						{
							AxisTarget: AxisTargets.Left,
							ColumnID:   "Response Time",
							Rules: []ThresholdRule{
								{
									Color: "#7dc540",
								},
								{
									Color: "#f5d30f",
								},
								{
									Color: "#dc172a",
									Value: 10,
								},
							},
							QueryID: "",
							Visible: false,
						},
					},
					TableSettings: TableSettings{
						SsThresholdBackgroundAppliedToCell: false,
					},
					GraphChartSettings: GraphChartSettings{ConnectNulls: false},
					HoneycombSettings:  HoneycombSettings{},
				},
				QueriesSettings: &QueriesSettings{
					Resolution:         "",
					FoldTransformation: FoldTransformations.Total,
					FoldAggregation:    FoldAggregations.Max,
				},
			},
			navTile,
		},
	}
	if len(subDashboards) > 0 {
		idx := -1
		for _, subDashboard := range subDashboards {
			idx++
			serviceTile := Tile{
				Name:       "Service or request",
				TileType:   TileTypes.ServiceVersatile,
				Configured: true,
				Bounds: Bounds{
					Top:    idx * 152,
					Left:   152,
					Width:  646,
					Height: 152,
				},
				TileFilter:       TileFilter{},
				AssignedEntities: []string{subDashboard.ServiceID},
			}
			dashboard.Tiles = append(dashboard.Tiles, serviceTile)
		}
	}

	return dashboard
}

func GenDashboard(name string, serviceID string, metricID string) *Dashboard {
	return &Dashboard{
		DashboardMetadata: DashboardMetadata{
			// Name:            "New RP MF",
			Name:            name,
			Shared:          false,
			Owner:           "reinhard.pilz@dynatrace.com",
			DashboardFilter: DashboardFilter{},
			Tags:            []string{"generated"},
			Popularity:      1,
		},
		Tiles: []Tile{
			NewTopList("Top list", serviceID, metricID, Bounds{
				Top:    0,
				Left:   0,
				Width:  266,
				Height: 684,
			}),
		},
	}
}

func NewMarkdownTile(name string, bounds Bounds, markdown string) Tile {
	return Tile{
		Name:       name,
		TileType:   TileTypes.Markdown,
		Configured: true,
		Bounds:     bounds,
		TileFilter: TileFilter{},
		Markdown:   &markdown,
	}
}

func NewServiceTile(name string, bounds Bounds, assignedEntities []string) Tile {
	return Tile{
		Name:             name,
		TileType:         TileTypes.ServiceVersatile,
		Configured:       true,
		Bounds:           bounds,
		TileFilter:       TileFilter{},
		AssignedEntities: assignedEntities,
	}
}

func NewTopList(name string, serviceID string, metricID string, bounds Bounds) Tile {
	return Tile{
		Name:       name,
		TileType:   TileTypes.DataExplorer,
		Configured: true,
		Bounds:     bounds,
		TileFilter: TileFilter{},
		CustomName: name,
		Queries: Queries{
			Query{
				ID: "A",
				// Metric:           "calc:service.mainframecicstransactionresp.time",
				Metric:           metricID,
				SpaceAggregation: SpaceAggregations.Avg,
				TimeAggregation:  TimeAggregations.Default,
				SplitBy:          []string{"CICS Txn Id"},
				SortBy:           "DESC",
				FilterBy: Filter{
					Operator: FilterOperators.And,
					NestedFilters: Filters{
						Filter{
							Filter:        "dt.entity.service",
							FilterType:    FilterTypes.ID,
							Operator:      FilterOperators.Or,
							NestedFilters: Filters{},
							Criteria: Criteria{
								Criterium{
									// Value:     "SERVICE-DF3F103EBBA71BBA",
									Value:     serviceID,
									Evaluator: "IN",
								},
							},
						},
					},
				},
				Limit:   100,
				Enabled: true,
			},
		},
		VisualConfig: &VisualConfig{
			Type:   VisualConfigTypes.TopList,
			Global: Global{HideLegend: false},
			Rules: VisualConfigRules{
				VisualConfigRule{
					Matcher:        "A",
					Properties:     Properties{Color: Colors.Default},
					SeriesOverride: []string{},
				},
			},
			Axes: Axes{
				XAxis: XAxis{Visible: true},
				YAxes: []string{},
			},
			HeatmapSettings: HeatmapSettings{},
			Thresholds: []Threshold{
				{
					AxisTarget: AxisTargets.Left,
					ColumnID:   "Response time",
					Rules: []ThresholdRule{
						{
							Color: "#7dc540",
						},
						{
							Color: "#f5d30f",
						},
						{
							Value: 10,
							Color: "#dc172a",
						},
					},
					QueryID: "",
					Visible: false,
				},
			},
			TableSettings: TableSettings{},
			GraphChartSettings: GraphChartSettings{
				ConnectNulls: false,
			},
			HoneycombSettings: HoneycombSettings{},
		},
		QueriesSettings: &QueriesSettings{
			Resolution:         "",
			FoldTransformation: FoldTransformations.Total,
			FoldAggregation:    FoldAggregations.Max,
		},
	}
}

type ManagementZone struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DashboardFilter struct {
	Timeframe      string         `json:"timeframe,omitempty"`
	ManagementZone ManagementZone `json:"managementZone,omitempty"`
}

type DashboardMetadata struct {
	Name            string          `json:"name"`
	Shared          bool            `json:"shared"`
	Owner           string          `json:"owner"`
	DashboardFilter DashboardFilter `json:"dashboardFilter"`
	Popularity      int             `json:"popularity"`
	Tags            []string        `json:"tags"`
}

type Metadata struct {
	ConfigurationVersions []int  `json:"configurationVersions"`
	ClusterVersion        string `json:"clusterVersion"`
}

type Dashboard struct {
	ID string `json:"id,omitempty"`
	// Metadata          Metadata          `json:"metadata"`
	DashboardMetadata DashboardMetadata `json:"dashboardMetadata"`
	Tiles             Tiles             `json:"tiles"`
}

type Tiles []Tile

type Bounds struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type TileFilter struct {
}

type Tile struct {
	Name             string           `json:"name"`
	TileType         TileType         `json:"tileType"`
	Configured       bool             `json:"configured"`
	Bounds           Bounds           `json:"bounds"`
	TileFilter       TileFilter       `json:"tileFilter"`
	AssignedEntities []string         `json:"assignedEntities,omitempty"`
	CustomName       string           `json:"customName,omitempty"`
	Queries          Queries          `json:"queries,omitempty"`
	VisualConfig     *VisualConfig    `json:"visualConfig,omitempty"`
	QueriesSettings  *QueriesSettings `json:"queriesSettings,omitempty"`
	Markdown         *string          `json:"markdown,omitempty"`
}

type QueriesSettings struct {
	Resolution         string             `json:"resolution"`
	FoldTransformation FoldTransformation `json:"foldTransformation"`
	FoldAggregation    FoldAggregation    `json:"foldAggregation"`
}

type Global struct {
	HideLegend bool `json:"hideLegend"`
}

type XAxis struct {
	Visible bool `json:"visible"`
}

type Axes struct {
	XAxis XAxis    `json:"xAxis"`
	YAxes []string `json:"yAxes"`
}

type HeatmapSettings struct {
}

type TableSettings struct {
	SsThresholdBackgroundAppliedToCell bool `json:"isThresholdBackgroundAppliedToCell"`
}
type GraphChartSettings struct {
	ConnectNulls bool `json:"connectNulls"`
}

type HoneycombSettings struct {
}

type VisualConfig struct {
	Type               VisualConfigType   `json:"type"`
	Global             Global             `json:"global"`
	Rules              VisualConfigRules  `json:"rules"`
	Axes               Axes               `json:"axes"`
	HeatmapSettings    HeatmapSettings    `json:"heatmapSettings"`
	Thresholds         []Threshold        `json:"thresholds"`
	TableSettings      TableSettings      `json:"tableSettings"`
	GraphChartSettings GraphChartSettings `json:"graphChartSettings"`
	HoneycombSettings  HoneycombSettings  `json:"honeycombSettings"`
}

type Threshold struct {
	AxisTarget AxisTarget      `json:"axisTarget"`
	ColumnID   string          `json:"columnId"`
	Rules      []ThresholdRule `json:"rules"`
	QueryID    string          `json:"queryId"`
	Visible    bool            `json:"visible"`
}

type ThresholdRule struct {
	Value int    `json:"value,omitempty"`
	Color string `json:"color"`
}

type VisualConfigRules []VisualConfigRule

type Properties struct {
	Color Color `json:"color"`
}

type VisualConfigRule struct {
	Matcher        string     `json:"matcher"`
	Properties     Properties `json:"properties"`
	SeriesOverride []string   `json:"seriesOverrides"`
}

type Queries []Query

type Query struct {
	ID               string           `json:"id"`
	Metric           string           `json:"metric"`
	SpaceAggregation SpaceAggregation `json:"spaceAggregation"`
	TimeAggregation  TimeAggregation  `json:"timeAggregation"`
	SplitBy          []string         `json:"splitBy"`
	SortBy           string           `json:"sortBy"`
	FilterBy         Filter           `json:"filterBy"`
	Limit            int              `json:"limit"`
	Enabled          bool             `json:"enabled"`
}

type Filters []Filter

type Filter struct {
	Operator        FilterOperator `json:"filterOperator"`
	Filter          string         `json:"filter,omitempty"`
	FilterType      FilterType     `json:"filterType,omitempty"`
	EntityAttribute string         `json:"entityAttribute,omitempty"`
	NestedFilters   Filters        `json:"nestedFilters"`
	Criteria        Criteria       `json:"criteria,omitempty"`
}

type Criteria []Criterium

type Criterium struct {
	Value     string `json:"value"`
	Evaluator string `json:"evaluator"`
}
