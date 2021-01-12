package monitor

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	. "github.com/shonenada/didiyun-go/schema"
)

const METRIC_CPU_UTIL = "cpu.util"
const METRIC_DISK_READ = "disk.read"
const METRIC_DISK_WRITE = "disk.write"
const METRIC_NET_IN = "net.in"
const METRIC_NET_OUT = "net.out"
const METRIC_RXBYTES = "rxbytes"
const METRIC_RXPKTS = "rxpkts"
const METRIC_TXBYTES = "txbytes"
const METRIC_TXPKTS = "txpkts"

type ResourceInput struct {
	ResourceType  string   `json:"resourceType"`
	ResourceUuids []string `json:"resourceUuids"`
	Metric        []string `json:"metric"`
}

type GetMonitorCounterRequest struct {
	RegionId  string          `json:"regionId"`
	Resources []ResourceInput `json:"resources"`
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
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.GET_MONITOR_COUNTER_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := GetMonitorCounterResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
