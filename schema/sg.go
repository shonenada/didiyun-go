package schema

type SgCondition struct {
	SgUuids    []string `json:"sgUuid,omitempty"`
	VpcUuid    string   `json:"vpcUuid,omitempty"`
	Dc2Uuid    string   `json:"dc2Uuid,omitempty"`
	Dc2Exclude bool     `json:"dc2Exclude"` // 与 dc2Uuid 配合使用，不传或传 false，表示查询此 DC2 所绑定的 SG 列表，传 true 表示查询此 DC2 未绑定的 SG 列表
}

type SgInfo struct {
	Job         Job      `json:"job"`
	SgUuid      string   `json:"sgUuid"`
	Name        string   `json:"name"`
	CreateTime  int64    `json:"createTime"`
	UpdateTime  int64    `json:"updateTime"`
	IsDefault   bool     `json:"isDefault"`
	Dc2Count    int      `json:"dcCnt"`
	SgRuleCount int      `json:"sgRuleCnt"`
	Region      Region   `json:"region"`
	Vpc         VpcOuput `json:"vpc"`
}

type VpcOutput struct {
	VpcUuid    string `json:"vpcUuid"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	IsDefault  bool   `json:"isDefault"`
	Desc       string `json:"desc"`
	CIDR       string `json:"cidr"`
	Region     Region `json:"region,omitempty"`
}

type SgOutput struct {
	SgUuid     string `json:"sgUuid"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	IsDefault  bool   `json:"isDefault"`
}

type SgRuleCondiction struct {
	SgUuid  string `json:"sgUuid,omitempty"`
	Dc2Uuid string `json:"dc2Uuid,omitempty"`
	Type    string `json:"type,omitempty"` // 要查询的 SGRule 类型，"Ingress" 为入方向，"Egress" 为出方向
}

type SgRuleInfo struct {
	Job         Job       `json:"job"`
	SgRuleUuid  string    `json:"sgRuleUuid"`
	Type        string    `json:"type"`
	Protocol    string    `json:"protocol"`
	StartPort   int       `json:"startPort"`
	EndPort     int       `json:"endPort"`
	AllowedCIDR string    `json:"allowedCidr"`
	CreateTime  int64     `json:"createTime"`
	UpdateTime  int64     `json:"updateTime"`
	IsDefault   bool      `json:"isDefault"`
	Sg          SgOutput  `json:"sg"`
	Vpc         VpcOutput `json:"vpc"`
}

type CreateSgRuleInput struct {
	Type        string `json:"type"`
	Protocol    string `json:"protocol"`
	StartPort   int    `json:"startPort"`
	EndPort     int    `json:"endPort"`
	AllowedCidr string `json:"allowedCidr"`
}
