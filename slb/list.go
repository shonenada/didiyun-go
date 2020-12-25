package slb

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type ListRequest struct {
	RegionId  string       `json:"regionId"`
	Start     int          `json:"start,omitempty"`
	Limit     int          `json:"limit,omitempty"`
	Condition SlbCondition `json:"condition,omitempty"`
}

type ListResponse struct {
	Errno     int           `json:"errno"`
	Errmsg    string        `json:"errmsg"`
	RequestId string        `json:"requestId"`
	Data      []SlbResponse `json:"data"`
}

type SlbResponse struct {
	Job        Job     `json:"job"`
	SlbUuid    string  `json:"slbUuid"`
	Name       string  `json:"name"`
	Ip         string  `json:"ip"`
	WafStatus  string  `json:"wafStatus"`
	CreateTime int     `json:"createTime"`
	UpdateTime int     `json:"updateTime"`
	Beip       BEIP    `json:"beip"`
	vpc        VPC     `json:"vpc"`
	flow       Flow    `json:"flow"`
	Spec       SlbSpec `json:"spec"`
	Region     Region  `json:"region"`
}

func (c *Client) List(request *ListRequest) (*[]SlbRe, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(LIST_SLB_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := ListResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
