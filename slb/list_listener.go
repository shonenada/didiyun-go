package slb

import (
	"encoding/json"
	"fmt"

	. "didiyun-go/schema"
)

type ListListenerRequest struct {
	RegionId  string       `json:"regionId"`
	Start     int          `json:"start"`
	Limit     int          `json:"limit"`
	Condition SlbCondition `json:"condition"`
}

type ListListenerResponse struct {
	Errno     int                `json:"errno"`
	Errmsg    string             `json:"errmsg"`
	RequestId string             `json:"requestId"`
	Data      []ListenerResponse `json:"data"`
}

type ListenerResponse struct {
	Job            Job           `json:"job"`
	SlbListenrUuid string        `json:"slbListenerUuid"`
	Protocol       string        `json:"protocol"`
	ListenerPort   int           `json:"listenerPort"`
	BackProtocol   string        `json:"backProtocol"`
	MemberPorts    []int         `json:"memberPorts"`
	PoolUuid       string        `json:"poolUiid"`
	CreateTime     int           `json:"createTime"`
	UpdateTime     int           `json:"updateTime"`
	Algorithm      Algorithm     `json:"algorithm"`
	HealthStatus   HealthStatus  `json:"healthStatus"`
	Monitor        HealthMonitor `json:"monitor"`
}

type HealthStatus struct {
	HealthyMemberCnt int `json:"healthyMemberCnt"`
	TotalMemberCnt   int `json:"totalMemberCnt"`
}

type HealthMonitor struct {
	SlbHealthMonitorUuid string `json:"slbHealthMonitorUuid"`
	Protocol             string `json:"protocol"`
	Interval             int    `json:"interval"`
	Timeout              int    `json:"timeout"`
	UnhealthyThreshold   int    `json:"unthealthyThreshold"`
	HealthyThreshold     int    `json:"thealthyThreshold"`
}

func (c *Client) ListListener(request *ListListenerRequest) (*[]ListenerResponse, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(LIST_LISTENER_SLB_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ListListenerResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
