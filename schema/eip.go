package schema

type Eip struct {
	Uuid       string `json:"eipUuid"`    // EIP 唯一标识
	Ip         string `json:"ip"`         // EIP（DC2 公网 IP）
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}

type EipInfo struct {
	Job        Job      `json:"json"`
	Uuid       string   `json:"eipUuid"`
	Ip         string   `json:"ip"`
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
	EipTags    []string `json:"eipTags"`
	Dc2        Dc2Info  `json:"dc2"`
}
