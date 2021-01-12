package schema

type EipInfo struct {
	Job    Job      `json:"job"`
	Id     string   `json:"eipId"`
	Uuid   string   `json:"eipUuid"`
	Ip     string   `json:"ip"`
	Tags   []string `json:"eipTags"`
	State  string   `json:"state"`
	Status string   `json:"status"`
	Region Region   `json:"region"`
	Dc2    struct {
		Dc2Uuid    string `json:"dc2Uuid"`
		Name       string `json:name"`
		Status     string `json:"status"`
		OsType     string `json:"osType"`
		CreateTime int64  `json:"createTime"`
		UpdateTime int64  `json:"updateTime"`
	} `json:"dc2"`
}

// EIP 规格
type EipSpec struct {
	Bandwidth         int64  `json:"babandwidth"` // 带宽
	CreateDate        string `json:"createDate"`
	ChargeType        string `json:"chargeType"` // 计费类型,`bandwidth` 为按带宽计费;`flow` 为按流量计费
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
