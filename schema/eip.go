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
