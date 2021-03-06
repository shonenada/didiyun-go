package sg

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type DetachRequest struct {
	RegionId string            `json:"regionId"`
	Sg       []DetachSgParams  `json:"sg"`
	Dc2      []DetachDc2Params `json:"dc2"`
}

type DetachSgParams struct {
	Uuid string `json:"SgUuid"`
}

type DetachDc2Params struct {
	Uuid string `json:"dc2Uuid"`
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
	body, err := c.HTTPPost(api.DEATCH_SG_URL, data)
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
