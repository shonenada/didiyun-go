package schema

type EipInfo struct {
	Charge     Charge   `json:"charge"`
	EipId      string   `json:"eipId"`
	EipTags    []string `json:"eipTags"`
	Uuid       string   `json:"eipUuid"`    // EIP 唯一标识
	Ip         string   `json:"ip"`         // EIP（DC2 公网 IP）
	CreateTime string   `json:"createTime"` // 创建时间
	UpdateTime string   `json:"updateTime"` // 更新时间
	State      string   `json:"state"`
	Status     string   `json:"status"`
	Spec       EipSpec  `json:"spec"`
	Region     Region   `json:"region"`
}
