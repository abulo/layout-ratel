package server

type ServerConf struct {
	Enable                   *bool   `json:"Enable"`
	Host                     *string `json:"Host"`
	Port                     *int    `json:"Port"`
	Deployment               *string `json:"Deployment"`
	EnableMetric             *bool   `json:"EnableMetric"`
	EnableTrace              *bool   `json:"EnableTrace"`
	EnableSlowQuery          *bool   `json:"EnableSlowQuery"`
	ServiceAddress           *string `json:"ServiceAddress"`
	SlowQueryThresholdInMill *int    `json:"SlowQueryThresholdInMill"`
}
