package monitor

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type CounterInput struct {
	ResourceType string `json:"resourceType"`
	ResourceUuid string `json:"resourceUuid"`
	MonitorTags  string `json:"monitorTags"`
	Metric       string `json:"metric"`
	StartTime    int    `json:"startTime"`
	EndTime      int    `json:"EndTim"`
}

type GetMonitorDatatRequest struct {
	RegionId string         `json:"regionId"`
	Counter  []CounterInput `json:"jobUuids"`
}

type MetricOutput struct {
	ResourceType string `json:"resourceType"`
	MonitorTags  string `json:"monitorTags"`
	Metric       string `json:"metric"`
	MetricAlias  string `json:"metricAlias"`
	Values       []struct {
		Timestamp int64 `json:"timestamp"`
		Value     int64 `json:"value"`
	} `json:"values"`
}

type GetMonitorDataResponse struct {
	Errno     int            `json:"errno"`
	Errmsg    string         `json:"errmsg"`
	RequestId string         `json:"requestId"`
	Data      []MetricOutput `json:"data"`
}

func (c *Client) GetResult(request *GetMonitorDatatRequest) (*[]Job, error) {
	data, err := json.Marshal(request)
	body, err := c.HTTPGet(GET_MONITOR_DATA_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := GetMonitorDatatResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
