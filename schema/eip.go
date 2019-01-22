package schema

type Eip struct {
	// EIP 唯一标识
	Uuid string `json:"eipUuid"`

	// EIP（DC2 公网 IP）
	Ip string `json:"ip"`

	// 创建时间
	CreateTime string `json:"createTime"`

	// 更新时间
	UpdateTime string `json:"updateTime"`
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
