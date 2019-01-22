package dc2

import (
	. "github.com/shonenada/didiyun-go/schema"
)

type ReinstallDc2Input struct {
	Dc2Uuid                 string   `json:"dc2Uuid"`
	ImgUuid                 string   `json:"imgUuid"`
	PubKeyUuids             []string `json:"pubKeyUuids"`
	Password                string   `json:"password"`
	ProSecurityAgentEnabled bool     `json:"proSecurityAgentEnabled"`
	MonitoringAgentEnabled  bool     `json:"monitoringAgentEnabled"`
}

type ReinstallDc2Request struct {
	RegionId string            `json:"regionId"`
	ZoneId   string            `json:"zoneId"`
	Dc2      ReinstallDc2Input `json:"dc2"`
}

type ReinstallDc2Response struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}
