package schema

type VpcInfo struct {
	Job        Job    `json:"job"`
	Uuid       string `json:"vpcUuid"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	IsDefault  bool   `json:"isDefault"`
	Desc       string `json:"desc"`
	CIDR       string `json:"cidr"` // VPC 的网段
	Region     Region `json:"region"`
}

type SubnetInfo struct {
	Job        Job    `json:"job"`
	SubnetUuid string `json:"subnetUuid"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	IsDefault  bool   `json:"isDefault"`
	CIDR       string `json:"cidr"`
	Zone       Zone   `json:"zone"`
}
