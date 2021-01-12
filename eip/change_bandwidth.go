package eip

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	. "github.com/shonenada/didiyun-go/schema"
)

type ChangeBandWidthRequest struct {
	RegionId string                  `json:"regionId"`
	Eip      []ChangeBandWidthParams `json:"eip"`
	CouponId string                  `json:"couponId,omitempty"`
}

type ChangeBandWidthParams struct {
	Uuid           string `json:"eipUuid"`
	BandWidth      int    `json:"bandWidth"`
	ChargeWithFlow bool   `json:"chargeWithFlow,omitempty"`
}

type ChangeBandWidthResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) ChangeBandWidth(request *ChangeBandWidthRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.CHANGE_BANDWIDTH_EIP_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ChangeBandWidthResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
