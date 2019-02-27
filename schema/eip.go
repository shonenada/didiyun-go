package schema

type EipInfo struct {
	Charge     Charge   `json:"charge"`
	CreateTime string   `json:"createTime"` // 创建时间
	EipId      string   `json:"eipId"`
	EipTags    []string `json:"eipTags"`
	Ip         string   `json:"ip"` // EIP（DC2 公网 IP）
	Region     Region   `json:"region"`
	Spec       EipSpec  `json:"spec"`
	State      string   `json:"state"`
	Status     string   `json:"status"`
	UpdateTime string   `json:"updateTime"` // 更新时间
	Uuid       string   `json:"eipUuid"`    // EIP 唯一标识
}
