package slb

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type CreateRequest struct {
	RegionId     string     `json:"regionId"`
	ZondId       string     `json:"zoneId"`
	AutoContinue bool       `json:"autoContinue"`
	PayPeriod    int        `json:"payPeriod"`
	Count        int        `json:"count"`
	CouponId     string     `json:"couponId",omitempty`
	Name         string     `json:"name"`
	VpcUuid      string     `json:"vpUuid"`
	AddressType  string     `json:"addressType"`
	Eip          EIP        `json:"eip"`
	Listeners    []Listener `json:"listeners"`
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
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(CREATE_SLB_URL, data)
	ret := CreateResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
