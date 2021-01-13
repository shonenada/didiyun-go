package schema

type Eip struct {
	Job    Job      `json:"job"`
	Id     string   `json:"eipId"`
	Uuid   string   `json:"eipUuid"`
	Ip     string   `json:"ip"`
	Tags   []string `json:"eipTags"`
	State  string   `json:"state"`
	Status string   `json:"status"`
	Region Region   `json:"region"`
	Dc2    EipDc2   `json:"dc2"`
}

type EipDc2 struct {
	Name       string `json:name"`
	Uuid       string `json:"dc2Uuid"`
	Status     string `json:"status"`
	OsType     string `json:"osType"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

type EipSpec struct {
	Bandwidth         int64  `json:"babandwidth"`
	CreateDate        string `json:"createDate"`
	ChargeType        string `json:"chargeType"`
	Description       string `json:"description"`
	InboundBandwidth  int64  `json:"inboundBandwidth"`
	LastOpDate        string `json:"lastOpDate"`
	Name              string `json:"name"`
	OfferingUuid      string `json:"offeringUuid"`
	OutboundBandwidth int64  `json:"outboundBandwidth"`
	PeerOfferingUuid  string `json:"peerOfferingUuid"`
	State             string `json:"state"`
	Type              string `json:"type"`
	Uuid              string `json:"uuid"`
}
