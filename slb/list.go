package slb

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type SlbCondition struct {
	SlbUuids []string `json:"slbUuids,omitempty"`
	VpcUuids []string `json:"vpcUuids,omitempty"`
	Beips    []string `json:"beips,omitempty"`
	Dc2Ips   []string `json:"dc2ips,omitempty"`
	Ips      []string `json:"ips,omitempty"`
}

type ListRequest struct {
	RegionId  string       `json:"regionId"`
	Start     int          `json:"start,omitempty"`
	Limit     int          `json:"limit,omitempty"`
	Condition SlbCondition `json:"condition,omitempty"`
}

type ListResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Slb `json:"data"`
}

func (c *Client) List(request *ListRequest) (*[]schema.Slb, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.LIST_SLB_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ListResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
