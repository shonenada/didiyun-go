package schema

type Ebs struct {
	Id         string   `json:"ebsId"`
	Name       string   `json:"name"`
	Uuid       string   `json:"ebsUuid"`
	DeviceName string   `json:"deviceName"`
	Tags       []string `json:"ebsTags"`
	Attr       string   `json:"attr"`
	Size       int64    `json:"size"`
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
