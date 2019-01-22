package eip

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type ChangeSizeRequest struct {
	RegionId string            `json:"regionId"`
	CouponId string            `json:"couponId,omitempty"`
	Ebs      []ChangeSizeInput `json:"ebs"`
}

type ChangeSizeInput struct {
	EbsUuid string `json:"ebsUuid"`
	Size    int64  `json:"size"`
}

type ChangeSizeResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) ChangeSize(request *ChangeSizeRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(CHANGE_SIZE_EBS_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := ChangeSizeResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}