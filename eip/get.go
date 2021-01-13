package eip

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type GetRequest struct {
	RegionId string `json:"regionId"`
	Uuid     string `json:"eipUuid"`
}

type GetResponse struct {
	Errno     int              `json:"errno"`
	Errmsg    string           `json:"errmsg"`
	RequestId string           `json:"requestId"`
	Data      []schema.EipInfo `json:"data"`
}

func (c *Client) Get(request *GetRequest) (*schema.EipInfo, error) {
	data := map[string]string{
		"regionId": request.RegionId,
		"eipUuid":  request.Uuid,
	}
	body, err := c.HTTPGet(api.GET_EIP_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := GetResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
