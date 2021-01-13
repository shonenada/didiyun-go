package schema

type Period struct {
	AutoRenewCount int    `json:"autoRenewCnt"`
	IsAutoSwitch   bool   `json:"autoSwitch"`
	PayType        string `json:"payType"`
}

type ChargePeriod struct {
	Period Period `json:"period"`
}

type Charge struct {
	ChargePeriod  ChargePeriod `json:"chargeInfo"`
	CostThisMonth float64      `json:"costThisMonth"`
	EndTime       int64        `json:"endTime"`
}
