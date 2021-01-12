package schema

type EbsDc2 struct {
	Uuid       string `json:"dc2Uuid"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	OsType     string `json:"osType"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

type Ebs struct {
	Attr       string   `json:"attr"`
	CreateTime int64    `json:"createTime"`
	DeviceName string   `json:"deviceName"`
	EbsId      string   `json:"ebsId"`
	EbsUuid    string   `json:"ebsUuid"`
	EbsTags    []string `json:"ebsTags"`
	Name       string   `json:"name"`
	Region     Region   `json:"region"`
	Size       int64    `json:"size"`
	UpdateTime int64    `json:"updateTime"`
	Dc2        EbsDc2   `json:"dc2"`
}
