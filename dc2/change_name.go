package dc2

import (
	. "github.com/shonenada/didiyun-go/schema"
)

type ChangeNameRequest struct {
	RegionId string            `json:"regionId"`
	ZoneId   string            `json:"zoneId"`
	Dc2      []ChangeNameInput `json:"dc2"`
}

type ChangeNameInput struct {
	Dc2Uuid string `json:"dc2Uuid"`
	Name    string `json:"name"`
}

type ChangeNameResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}
