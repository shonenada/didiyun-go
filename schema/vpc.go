package schema

type Vpc struct {
	Uuid        string `json:"vpcUuid"`
	Name        string `json:"name"`
	IsDefault   bool   `json:"isDefault"`
	Description string `json:"desc"`
	CIDR        string `json:"cidr"` // VPC 的网段
	Region      Region `json:"region"`
	Job         Job    `json:"job"`
	CreateTime  int64  `json:"createTime"`
	UpdateTime  int64  `json:"updateTime"`
}

type SubnetInfo struct {
	Uuid       string `json:"subnetUuid"`
	Name       string `json:"name"`
	IsDefault  bool   `json:"isDefault"`
	CIDR       string `json:"cidr"`
	Job        Job    `json:"job"`
	Zone       Zone   `json:"zone"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}
