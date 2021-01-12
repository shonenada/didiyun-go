package vpc

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	. "github.com/shonenada/didiyun-go/schema"
)

type CreateSubnetRequest struct {
	RegionId string              `json:"regionId"`
	VpcUuid  string              `json:"vpcUuid"`
	Subnet   []CreateSubnetInput `json:"subnet"`
}

type CreateSubnetInput struct {
	Name   string `json:"name"`
	CIDR   string `json:"cidr"` // Subnet 网段，格式如 "10.0.0.0/16"
	ZoneId string `json:"zoneId"`
}

type CreateSubnetResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) CreateSubnet(request *CreateSubnetRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.CREATE_SUBNET_VPC_URL, data)
	ret := CreateSubnetResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
