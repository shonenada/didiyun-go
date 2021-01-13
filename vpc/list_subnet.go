package vpc

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type ListSubnetRequest struct {
	RegionId  string          `json:"regionId"`
	ZoneId    string          `json:"zoneId"`
	Start     int             `json:"start"`
	Limit     int             `json:"limit"`
	Condition SubnetCondition `json:"condition"`
}

type SubnetCondition struct {
	Uuid string `json:"vpcUuid,omitempty"`
}

type ListSubnetResponse struct {
	Errno     int                 `json:"errno"`
	Errmsg    string              `json:"errmsg"`
	RequestId string              `json:"requestId"`
	Data      []schema.SubnetInfo `json:"data"`
}

func (c *Client) ListSubnet(request *ListSubnetRequest) (*[]schema.SubnetInfo, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.LIST_SUBNET_VPC_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ListSubnetResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
