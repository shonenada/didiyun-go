package dc2

import (
	. "github.com/shonenada/didiyun-go/schema"
)

type StartDc2Input struct {
	Dc2Uuid string `json:"dc2Uuid"`
}

type StartDc2Request struct {
	RegionId string        `json:"regionId"`
	ZoneId   string        `json:"zoneId"`
	Dc2      StartDc2Input `json:"dc2"`
}

type StartDc2Response struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}
