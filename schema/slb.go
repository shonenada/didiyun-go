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

type Flow struct {
	In  float32 `json:"float32"`
	Out float32 `json:"float32"`
}

type BEIP struct {
	BeipUuid   string `json:"beipUuid"`
	Ip         string `json:"ip"`
	CreateTime int    `json:"createTime"`
	UpdateTime int    `json:"updateTime"`
}

type EIP struct {
	Name           string `json:"name",omitempty`
	Bandwidth      string `json:"bandwidth"`
	ChargeWithFlow bool   `json:"chargeWithFlow"`
}

type Listener struct {
	SlbListenerUuid string           `json:"slbListenerUuid,omitempty"`
	Name            string           `json:"listener"`
	Algorithm       string           `json:"algorithm"`
	Protocol        string           `json:"protocol"`
	ListenerPort    int              `json:"listenerPort"`
	BackProtocol    string           `json:"backProtocol"`
	Monitor         Monitor          `json:"monitor"`
	Members         []ListenerMember `json:"members,omitempty"`
}
type Monitor struct {
	Protocol           string `json:"protocol"`
	Interval           int    `json:"interval"`
	Timeout            int    `json:"timeout"`
	HealthyThreshold   int    `json:"healthyThreshold"`
	UnhealthyThreshold int    `json:"unhealthyThreshold"`
}

type ListenerMember struct {
	Dc2Uuid       string    `json:"dc2Uuid",omitempty`
	SlbMemberUuid string    `json:"slbMemberUuid",omitempty`
	HealthState   string    `json:"healthState,omitempty"`
	UnhealthState string    `json:"unhealthState,omitempty"`
	Port          int       `json:"port",omitempty`
	Weight        int       `json:"weight",omitempty`
	CreateTime    int       `json:"createTime,omitempty"`
	UpdateTime    int       `json:"updateTime,omitempty"`
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

type SlbSpec struct {
	OfferingUuid string `json:"offeringUuid"`
}
