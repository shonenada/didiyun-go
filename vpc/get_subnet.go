package vpc

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type GetSubnetRequest struct {
	RegionId   string `json:"regionId"`
	VpcUuid    string `json:"vpcUuid"`
	SubnetUuid string `json:"subnetUuid"`
}

type GetSubnetResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []SubnetInfo `json:"data"`
}

func (c *Client) GetSubnet(request *GetSubnetRequest) (*SubnetInfo, error) {
	data := map[string]string{
		"regionId":   request.RegionId,
		"vpcUuid":    request.VpcUuid,
		"subnetUuid": request.SubnetUuid,
	}
	body, err := c.HTTPGet(GET_SUBNET_VPC_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := GetSubnetResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
