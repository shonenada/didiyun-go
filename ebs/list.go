package ebs

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type ListRequest struct {
	RegionId  string        `json:"regionId"`
	ZoneId    string        `json:"zoneId,omitempty"`
	Start     int           `json:"start,omitempty"`
	Limit     int           `json:"limit,omitempty"`
	Condition ListCondition `json:"condition"`
}

type ListCondition struct {
	Dc2Uuids []string `json:"dc2Uuids"`
}

type ListResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Ebs `json:"data"`
}

func (c *Client) List(request *ListRequest) (*[]schema.Ebs, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.LIST_EBS_URL, data)
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
