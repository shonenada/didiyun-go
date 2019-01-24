package sg

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type DetachRequest struct {
	RegionId string           `json:"regionId"`
	Sg       []DetachSgInput  `json:"sg"`
	Dc2      []DetachDc2Input `json:"dc2"`
}

type DetachSgInput struct {
	SgUuid string `json:"SgUuid"`
}

type DetachDc2Input struct {
	Dc2Uuid string `json:"dc2Uuid"`
}

type DetachResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) Detach(request *DetachRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(DEATCH_SG_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := DetachResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
