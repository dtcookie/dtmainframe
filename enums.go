package main

type AxisTarget string

var AxisTargets = struct {
	Left AxisTarget
}{
	AxisTarget("LEFT"),
}

type Color string

var Colors = struct {
	Default Color
}{
	Color("DEFAULT"),
}

type VisualConfigType string

var VisualConfigTypes = struct {
	TopList     VisualConfigType
	SingleValue VisualConfigType
}{
	VisualConfigType("TOP_LIST"),
	VisualConfigType("SINGLE_VALUE"),
}

type FoldAggregation string

var FoldAggregations = struct {
	Max FoldAggregation
}{
	FoldAggregation("MAX"),
}

type FoldTransformation string

var FoldTransformations = struct {
	Total FoldTransformation
}{
	FoldTransformation("TOTAL"),
}

type SpaceAggregation string

var SpaceAggregations = struct {
	Avg SpaceAggregation
	MAx SpaceAggregation
}{
	SpaceAggregation("AVG"),
	SpaceAggregation("MAX"),
}

type TimeAggregation string

var TimeAggregations = struct {
	Default TimeAggregation
}{
	TimeAggregation("DEFAULT"),
}

type TileType string

var TileTypes = struct {
	ServiceVersatile TileType
	DataExplorer     TileType
	Markdown         TileType
}{
	TileType("SERVICE_VERSATILE"),
	TileType("DATA_EXPLORER"),
	TileType("MARKDOWN"),
}

type FilterOperator string

var FilterOperators = struct {
	And FilterOperator
	Or  FilterOperator
}{
	FilterOperator("AND"),
	FilterOperator("OR"),
}

type FilterType string

var FilterTypes = struct {
	ID   FilterType
	Name FilterType
}{
	FilterType("ID"),
	FilterType("NAME"),
}

/*
"APPLICATION"
"APPLICATIONS"
"APPLICATIONS_MOST_ACTIVE"
"APPLICATION_METHOD"
"APPLICATION_WORLDMAP"
"AWS"
"BOUNCE_RATE"
"CUSTOM_APPLICATION"
"CUSTOM_CHARTING"
"DATABASE"
"DATABASES_OVERVIEW"
"DCRUM_SERVICES"
"DEM_KEY_USER_ACTION"
"DEVICE_APPLICATION_METHOD"
"DOCKER"
"DTAQL"
"HEADER"
"HOST"
"HOSTS"
"LOG_ANALYTICS"
"MARKDOWN"
"MOBILE_APPLICATION"
"NETWORK"
"NETWORK_MEDIUM"
"OPENSTACK"
"OPENSTACK_AV_ZONE"
"OPENSTACK_PROJECT"
"OPEN_PROBLEMS"
"PROCESS_GROUPS_ONE"
"PURE_MODEL"
"RESOURCES"
"SCALABLE_LIST"
"SERVICES"
"SERVICE_VERSATILE"
"SESSION_METRICS"
"SLO"
"SYNTHETIC_HTTP_MONITOR"
"SYNTHETIC_SINGLE_EXT_TEST"
"SYNTHETIC_SINGLE_WEBCHECK"
"SYNTHETIC_TESTS"
"THIRD_PARTY_MOST_ACTIVE"
"UEM_ACTIVE_SESSIONS"
"UEM_CONVERSIONS_OVERALL"
"UEM_CONVERSIONS_PER_GOAL"
"UEM_JSERRORS_OVERALL"
"UEM_KEY_USER_ACTIONS"
"USERS"
"VIRTUALIZATION"
*/
