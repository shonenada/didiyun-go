package vpc

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
)

type CheckCIDROverlapRequest struct {
	RegionId string `json:"regionId"`
	Uuid     string `json:"vpcUuid"`
	CIDR     string `json:"cidr"`
}

type IsOverLapOutput struct {
	IsOverlap bool `json:"isOverLap"`
}

type CheckCIDROverlapResponse struct {
	Errno     int               `json:"errno"`
	Errmsg    string            `json:"errmsg"`
	RequestId string            `json:"requestId"`
	Data      []IsOverLapOutput `json:"data"`
}

func (c *Client) CheckCDIROverlap(request *CheckCIDROverlapRequest) (*[]IsOverLapOutput, error) {
	data := map[string]string{
		"regionId": request.RegionId,
		"vpcUuid":  request.Uuid,
		"cidr":     request.CIDR,
	}

	body, err := c.HTTPGet(api.CHECK_SUBNET_VPC_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := CheckCIDROverlapResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
