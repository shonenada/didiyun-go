package vpc

import (
	"encoding/json"
	"fmt"

	. "didiyun-go/schema"
)

type DeleteSubnetRequest struct {
	RegionId string              `json:"regionId"`
	VpcUuid  []DeleteSubnetInput `json:"vpcUuid"`
	Subnet   []DeleteSubnetInput `json:"subnet"`
}

type DeleteSubnetInput struct {
	subnetUuid string `json:"subnetUuid"`
}

type DeleteSubnetResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) DeleteSubnet(request *DeleteSubnetRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(DELETE_SUBNET_VPC_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := DeleteSubnetResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
