package dc2

import (
	. "github.com/shonenada/didiyun-go/schema"
)

type ChangePasswordInput struct {
	Dc2Uuid  string `json:"dc2Uuid"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	RegionId string                `json:"regionId"`
	ZoneId   string                `json:"zoneId"`
	Dc2      []ChangePasswordInput `json:"dc2"`
}

type ChangePasswordResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}
