package vpc

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	. "github.com/shonenada/didiyun-go/schema"
)

type SubnetChangeNameRequest struct {
	RegionId string                  `json:"regionId"`
	VpcUuid  string                  `json:"vpcUuid"`
	Subnet   []SubnetChangeNameInput `json:"subnet"`
}

type SubnetChangeNameInput struct {
	SubnetUuid string `json:"subnetUuid"`
	Name       string `json:"name"`
}

type SubnetChangeNameResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) SunbetChangeName(request *SubnetChangeNameRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.CHANGE_NAME_SUBNET_VPC_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := SubnetChangeNameResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
