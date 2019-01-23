package vpc

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type GetRequest struct {
	RegionId string `json:"regionId"`
	VpcUuid  string `json:"vpcUuid"`
}

type GetResponse struct {
	Errno     int       `json:"errno"`
	Errmsg    string    `json:"errmsg"`
	RequestId string    `json:"requestId"`
	Data      []VpcInfo `json:"data"`
}

func (c *Client) Get(request *GetRequest) (*EbsInfo, error) {
	data := map[string]string{
		"regionId": request.RegionId,
		"vpcUuid":  request.VpcUuid,
	}
	body, err := c.HTTPGet(GET_VPC_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := GetResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
