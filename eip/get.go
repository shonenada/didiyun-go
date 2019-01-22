package eip

import (
	"encoding/json"

	"github.com/shonenada/didiyun-go/schame"
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

func (c *Client) List(request *GetRequest) (*[]EipInfo, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(LIST_EIP_URL, data)
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
