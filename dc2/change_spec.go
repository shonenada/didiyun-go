package dc2

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type ChangeSpecRequest struct {
	RegionId string             `json:"regionId"`
	ZoneId   string             `json:"zoneId,omitempty"`
	CouponId string             `json:"couponId,omitempty"`
	Dc2      []ChangeSpecParams `json:"dc2"`
}

type ChangeSpecParams struct {
	Uuid  string `json:"dc2Uuid"`
	Model string `json:"dc2Model"`
}

type ChangeSpecResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Job `json:"data"`
}

func (c *Client) ChangeSpec(request *ChangeSpecRequest) (*schema.Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.CHANGE_SPEC_DC2_URL, data)
	ret := ChangeSpecResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
