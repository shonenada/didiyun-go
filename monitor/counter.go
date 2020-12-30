package monitor

import (
	. "github.com/shonenada/didiyun-go/schema"
)

type ResourceInput struct {
	ResourceType  string   `json:"resourceType"`
	ResourceUuids []string `json:"resourceUuids"`
	Metric        []string `json:"metric"`
}

type GetMonitorCounterRequest struct {
	RegionId  string     `json:"regionId"`
	Resources []Resource `json:"resources"`
}

type GetMonitorCounterResponse struct {
	Errno     int             `json:"errno"`
	Errmsg    string          `json:"errmsg"`
	RequestId string          `json:"requestId"`
	Data      []CounterOutput `json:"data"`
}

type CounterOutput struct {
	ResourceType string    `json:"resourceType"`
	ResourceUuid string    `json:"resourceUuid"`
	Alias        string    `json:"aliass"`
	Metric       string    `json:"merics"`
	MetricAlias  string    `json:"metricAlias"`
	Counters     []Counter `json:"counters"`
}

func (c *Client) GetCounter(request *GetMonitorCounterRequest) (*[]CounterOutput, error) {
}
