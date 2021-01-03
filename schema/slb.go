package schema

type SlbResponse struct {
	Job        Job     `json:"job"`
	SlbUuid    string  `json:"slbUuid"`
	Name       string  `json:"name"`
	Ip         string  `json:"ip"`
	WafStatus  string  `json:"wafStatus"`
	CreateTime int     `json:"createTime"`
	Updateime  int     `json:"updateTime"`
	BeIP       BEIP    `json:"beip"`
	Vpc        VPC     `json:"vpc"`
	Flow       Flow    `json:"flow"`
	Spec       SlbSpec `json:"spec"`
	Region     Region  `json:"region"`
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
	SlbListenerUuid string         `json:"slbListenerUuid,omitempty"`
	Name            string         `json:"listener"`
	Algorithm       string         `json:"algorithm"`
	Protocol        string         `json:"protocol"`
	ListenerPort    int            `json:"listenerPort"`
	BackProtocol    string         `json:"backProtocol"`
	Monitor         Monitor        `json:"monitor"`
	Members         []SimpleMember `json:"members,omitempty"`
}

type Monitor struct {
	Protocol           string `json:"protocol"`
	Interval           int    `json:"interval"`
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthyThreshold"`
	UnhealthyThreshold int    `json:"unhealthyThreshold"`
}

type SimpleMember struct {
	Dc2Uuid string `json:"dc2Uuid"`
	Weight  int    `json:"weight"`
	Port    int    `json:"port"`
}

type ListenerMember struct {
	SlbMemberUuid string    `json:"slbMemberUuid"`
	HealthState   string    `json:"healthState,omitempty"`
	unhealthState string    `json:"unhealthState,omitempty"`
	port          int       `json:"port"`
	weight        int       `json:"weight"`
	createTime    int       `json:"createTime,omitempty"`
	updateTime    int       `json:"updateTime,omitempty"`
	Dc2           MemberDc2 `json:"dc2,omitempty"`
}

type MemberDc2 struct {
	Dc2Uuid    string `json:"dc2Uuid"`
	Name       string `json:"name"`
	Ip         string `json:"ip"`
	createTime int    `json:"createTime"`
	updateTime int    `json:"updateTime"`
}

type Algorithm struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
