package eip

import (
	"encoding/json"
	"fmt"

	. "didiyun-go/schema"
)

type CreateRequest struct {
	RegionId       string   `json:"regionId"`
	AutoContinue   bool     `json:"autoContinue,omitempty"`
	PayPeriod      int      `json:"payPeriod,omitempty"`
	CouponId       string   `json:"couponId,omitempty"`
	Count          int      `json:"count,omitempty"`
	BandWidth      int      `json:bandwidth,omitempty"`
	ChargeWithFlow bool     `json:"chargeWithFlow,omitempty"`
	BindingUuid    string   `json:"bindingUuid,omitempty"`
	EipTags        []string `json:"EipTags"`
}

type CreateResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) Create(request *CreateRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(CREATE_EIP_URL, data)
	ret := CreateResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
