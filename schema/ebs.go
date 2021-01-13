package schema

type Ebs struct {
	Id         string   `json:"ebsId"`
	Uuid       string   `json:"ebsUuid"`
	Name       string   `json:"name"`
	DeviceName string   `json:"deviceName"`
	Attr       string   `json:"attr"`
	Size       int64    `json:"size"`
	Tags       []string `json:"ebsTags"`
	Dc2        EbsDc2   `json:"dc2"`
	Region     Region   `json:"region"`
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
}

type EbsDc2 struct {
	Uuid       string `json:"dc2Uuid"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	OsType     string `json:"osType"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}
