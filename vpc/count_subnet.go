package vpc

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
)

type CountSubnetRequest struct {
	RegionId string `json:"regionId"`
	ZoneId   string `json:"zoneId,omitempty"`
	Uuid     string `json:"vpcUuid"`
}

type CountSubnetResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []struct {
		TotalCount int `json:"totalCnt"`
	} `json:"data"`
}

func (c *Client) CountSubnet(request *CountSubnetRequest) (int, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return -1, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.COUNT_SUBNET_VPC_URL, data)
	if err != nil {
		return -1, fmt.Errorf("Error: %s", err)
	}
	ret := CountSubnetResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return -1, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return ret.Data[0].TotalCount, nil
}
