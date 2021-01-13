package schema

type SgInfo struct {
	Job         Job       `json:"job"`
	SgUuid      string    `json:"sgUuid"`
	Name        string    `json:"name"`
	CreateTime  int64     `json:"createTime"`
	UpdateTime  int64     `json:"updateTime"`
	IsDefault   bool      `json:"isDefault"`
	Dc2Count    int       `json:"dcCnt"`
	SgRuleCount int       `json:"sgRuleCnt"`
	Region      Region    `json:"region"`
	Vpc         VpcOutput `json:"vpc"`
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
