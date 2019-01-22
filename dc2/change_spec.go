package dc2

import (
	. "github.com/shonenada/didiyun-go/schema"
)

type ChangeSpecRequest struct {
	RegionId string            `json:"regionId"`
	ZoneId   string            `json:"zoneId"`
	CouponId string            `json:"couponId"`
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
