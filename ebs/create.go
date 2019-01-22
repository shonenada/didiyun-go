package ebs

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type CreateRequest struct {
	RegionId     string   `json:"regionId"`
	ZoneId       string   `json:"zondId"`
	AutoContinue bool     `json:"autoContinue,omitempty"`
	PayPeriod    int      `json:"payPeriod,omitempty"`
	Count        int      `json:"count,omitempty"`
	CouponId     string   `json:"couponId,omitempty"`
	Name         string   `json:"name,omitempty"`
	Size         int64    `json:"size"`
	DiskType     string   `json:"diskType"` // 创建的 EBS 类型（"HDD" 或 "SSD"）
	Dc2Uuid      string   `json:"dc2Uuid,omitempty"`
	SnapUuid     string   `json:"snapUuid,omitempty"`
	EbsTags      []string `json:"ebsTags,omitempty"`
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
	body, err := c.HTTPPost(CREATE_EBS_URL, data)
	ret := CreateResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}