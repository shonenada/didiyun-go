package sg

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type AttachRequest struct {
	RegionId string            `json:"regionId"`
	Sg       []AttachSgParams  `json:"sg"`
	Dc2      []AttachDc2Params `json:"dc2"`
}

type AttachSgParams struct {
	SgUuid string `json:"sgUuid"`
}

type AttachDc2Params struct {
	Dc2Uuid string `json:"dc2Uuid"`
}

type AttachResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Job `json:"data"`
}

func (c *Client) Attach(request *AttachRequest) (*schema.Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.ATTACH_SG_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := AttachResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
