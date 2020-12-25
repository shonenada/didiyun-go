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

type EIP struct {
	Name           string `json:"name",omitempty`
	Bandwidth      string `json:"bandwidth"`
	ChargeWithFlow bool   `json:"chargeWithFlow"`
}

type Listener struct {
	Name         string   `json:"listener"`
	Algorithm    string   `json:"algorithm"`
	Protocol     string   `json:"protocol"`
	ListenerPort int      `json:"listenerPort"`
	BackProtocol string   `json:"backProtocol"`
	Monitor      Monitor  `json:"monitor"`
	Members      []Member `json:"members"`
}

type Monitor struct {
	Protocol           string `json:"protocol"`
	Interval           int    `json:"interval"`
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthyThreshold"`
	UnhealthyThreshold int    `json:"unhealthyThreshold"`
}

type Member struct {
	Dc2Uuid string `json:"dc2Uuid"`
	Weight  int    `json:"weight"`
	Port    int    `json:"port"`
}

type CreateResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}
