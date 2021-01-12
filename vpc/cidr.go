package vpc

// 获取 VPC 可用网段

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	. "github.com/shonenada/didiyun-go/schema"
)

type VPCAvailableCidr struct {
	AvailableCidr []string `json:"availableCidr"`
}

type CIDRResponse struct {
	Errno     int                `json:"errno"`
	Errmsg    string             `json:"errmsg"`
	RequestId string             `json:"requestId"`
	Data      []VPCAvailableCidr `json:"data"`
}

func (c *Client) CIDR() (*[]VPCAvailableCidr, error) {
	body, err := c.HTTPGet(api.CIDR_VPC_URL, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := CIDRResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
