package vpc

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type CountSubnetRequest struct {
	RegionId string `json:"regionId"`
	ZoneId   string `json:"zoneId,omitempty"`
	VpcUuid  string `json:"vpcUuid"`
}

type CountSubnetResponse struct {
	Errno     int           `json:"errno"`
	Errmsg    string        `json:"errmsg"`
	RequestId string        `json:"requestId"`
	Data      []SubnetCount `json:"data"`
}

type SubnetCount struct {
	TotalCount int `json:"totalCnt"`
}

func (c *Client) CountSubnet(request *CountSubnetRequest) (int, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(COUNT_SUBNET_VPC_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := CountSubnetResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return -1, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return ret.Data[0].TotalCount, nil
}
