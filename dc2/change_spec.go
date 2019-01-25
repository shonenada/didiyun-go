package dc2

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type ChangeSpecRequest struct {
	RegionId string            `json:"regionId"`
	ZoneId   string            `json:"zoneId,omitempty"`
	CouponId string            `json:"couponId,omitempty"`
	Dc2      []ChangeSpecInput `json:"dc2"`
}

type ChangeSpecResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

type ChangeSpecInput struct {
	Dc2Uuid  string `json:"dc2Uuid"`
	Dc2Model string `json:"dc2Model"`
}

func (c *Client) ChangeSpec(request *ChangeSpecRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(CHANGE_SPEC_DC2_URL, data)
	ret := ChangeSpecResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
