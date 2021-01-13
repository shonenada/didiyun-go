package schema

type SgInfo struct {
	Job         Job    `json:"job"`
	Uuid        string `json:"sgUuid"`
	Name        string `json:"name"`
	IsDefault   bool   `json:"isDefault"`
	Dc2Count    int    `json:"dcCnt"`
	SgRuleCount int    `json:"sgRuleCnt"`
	Region      Region `json:"region"`
	Vpc         SgVpc  `json:"vpc"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
}

type SgRule struct {
	Uuid        string   `json:"sgRuleUuid"`
	Type        string   `json:"type"`
	Protocol    string   `json:"protocol"`
	StartPort   int      `json:"startPort"`
	EndPort     int      `json:"endPort"`
	AllowedCIDR string   `json:"allowedCidr"`
	IsDefault   bool     `json:"isDefault"`
	Job         Job      `json:"job"`
	Sg          SimpleSg `json:"sg"`
	Vpc         SgVpc    `json:"vpc"`
	CreateTime  int64    `json:"createTime"`
	UpdateTime  int64    `json:"updateTime"`
}

type SgVpc struct {
	Uuid        string `json:"vpcUuid"`
	Name        string `json:"name"`
	IsDefault   bool   `json:"isDefault"`
	Description string `json:"desc"`
	CIDR        string `json:"cidr"`
	Region      Region `json:"region,omitempty"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
}

type SimpleSg struct {
	Uuid       string `json:"sgUuid"`
	Name       string `json:"name"`
	IsDefault  bool   `json:"isDefault"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}
