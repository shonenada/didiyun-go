package sg

import (
	"encoding/json"
	"fmt"

	. "didiyun-go/schema"
)

type CreateRequest struct {
	RegionId string              `json:"regionId"`
	Name     string              `json:"name"`
	VpcUuid  string              `json:"vpcUuid,omitempty"`
	SgRule   []CreateSgRuleInput `json:"sgRule,omitempty"`
	Dc2      []Dc2CreateSgInput  `json:"dc2,omitempty"`
}

type Dc2CreateSgInput struct {
	Dc2Uuid string `json:"dc2Uuid"`
}

type CreateResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) Create(request *CreateRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(CREATE_SG_URL, data)
	ret := CreateResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
