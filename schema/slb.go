package schema

type SlbResponse struct {
	Job        Job    `json:"job"`
	SlbUuid    string `json:"slbUuid"`
	Name       string `json:"name"`
	Ip         string `json:"ip"`
	WafStatus  string `json:"wafStatus"`
	CreateTime int    `json:"createTime"`
	Updateime  int    `json:"updateTime"`
	BeIP       BEIP   `json:"beip"`
	Vpc        VPC    `json:"vpc"`
	Flow       Flow   `json:"flow"`
	Spec       Sppc   `json:"spec"`
	Region     Region `json:"region"`
}

type SlbCondition struct {
	SlbUuids []string `json:"slbUuids,omitempty"`
	VpcUuids []string `json:"vpcUuids,omitempty"`
	Beips    []string `json:"beips,omitempty"`
	Dc2Ips   []string `json:"dc2ips,omitempty"`
	Ips      []string `json:"ips,omitempty"`
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
