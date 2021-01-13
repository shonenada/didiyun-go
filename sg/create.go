package sg

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type CreateRequest struct {
	RegionId string               `json:"regionId"`
	Name     string               `json:"name"`
	VpcUuid  string               `json:"vpcUuid,omitempty"`
	SgRule   []CreateSgRuleParams `json:"sgRule,omitempty"`
	Dc2      []Dc2CreateSgParams  `json:"dc2,omitempty"`
}

type CreateSgRuleParams struct {
	Type        string `json:"type"`
	Protocol    string `json:"protocol"`
	StartPort   int    `json:"startPort"`
	EndPort     int    `json:"endPort"`
	AllowedCidr string `json:"allowedCidr"`
}

type Dc2CreateSgParams struct {
	Uuid string `json:"dc2Uuid"`
}

type CreateResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Job `json:"data"`
}

func (c *Client) Create(request *CreateRequest) (*schema.Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.CREATE_SG_URL, data)
	ret := CreateResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
