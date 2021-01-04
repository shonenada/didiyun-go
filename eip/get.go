package eip

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type GetRequest struct {
	RegionId string `json:"regionId"`
	EipUuid  string `json:"eipUuid"`
}

type GetResponse struct {
	Errno     int       `json:"errno"`
	Errmsg    string    `json:"errmsg"`
	RequestId string    `json:"requestId"`
	Data      []EipInfo `json:"data"`
}

func (c *Client) Get(request *GetRequest) (*EipInfo, error) {
	data := map[string]string{
		"regionId": request.RegionId,
		"eipUuid":  request.EipUuid,
	}
	body, err := c.HTTPGet(GET_EIP_URL, data)
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
