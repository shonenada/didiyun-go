package ebs

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type DetachRequest struct {
	RegionId string         `json:"regionId"`
	Ebs      []DetachParams `json:"ebs"`
}

type DetachParams struct {
	Uuid string `json:"ebsUuid"`
}

type DetachResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Job `json:"data"`
}

func (c *Client) Detach(request *DetachRequest) (*schema.Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.DEATCH_EBS_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := DetachResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
