package eip

import (
	"encoding/json"
	"fmt"
)

type ChangeBandWidthRequest struct {
	RegionId string                 `json:"regionId"`
	Eip      []ChangeBandWidthInput `json:"eip"`
	CouponId string                 `json:"couponId"`
}

type ChangeBandWidthInput struct {
	EipUuid        string `json:"eipUuid"`
	BandWidth      int    `json:"bandWidth"`
	ChargeWithFlow bool   `json:"chargeWithFlow"`
}

type ChangeBandWidthhResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) ChangeBandWidth(request *ChangeBandWidthRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(CHANGE_BANDWIDTH_EIP_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := ChangeBandWidthResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return ret.Data[0], nil
}
