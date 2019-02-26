package schema

type Dc2Condition struct {
	VpcUuid   string   `json:"vpcUuid"`   // 查询指定 VPC 下的 DC2 列表
	VpcUuids  []string `json:"vpcUuids"`  // 查询多个 VPC 下的 DC2 列表
	Dc2Name   string   `json:"dc2Name"`   // 模糊查询 DC2 名字
	SgUuid    string   `json:"sgUuid"`    // 查询此 SG 下的 DC2 列表
	Dc2Uuids  []string `json:"dc2Uuids"`  // 查询包含在此 UUID 列表内的 DC2
	SgExclude bool     `json:"sgExclude"` // 为 `true` 时表示查询不在此 SG 下的 DC2 列表
	Ip        string   `json:"ip"`        // 精确匹配内网IP
	Eip       string   `json:"eip"`       // 精确匹配公网EIP
}

type Dc2Info struct {
	Job        Job      `json:"job"`        // 此 DC2 正在进行的任务，若无任务则没有此字段
	Uuid       string   `json:"dc2Uuid"`    // DC2 唯一标识
	Name       string   `json:"name"`       // DC2 名称
	CreateTime int64    `json:"createTime"` // DC2 创建时间
	UpdateTime int64    `json:"updateTime"` // DC2 更新时间
	Ip         string   `json:"ip"`         // DC2 内网 IP
	Tags       []string `json:"tags"`       // DC2 的 tags
	Status     string   `json:"status"`     // DC2 状态
	OSType     string   `json:"osType"`     // DC2 操作系统发行版及版本号
	Eip        Eip      `json:"eip"`        // 与 DC2 关联的 EIP 信息，没有 EIP 则没有该字段
	Ebs        []Ebs    `json:"ebs"`        // 与 DC2 关联的 EBS 信息，没有 EBS 则没有该字段，如果是通用型 DC2，则必有这个字段，且根盘信息包含在内
	Region     Region   `json:"region"`     // region 信息
}

type Dc2 struct {
	Uuid       string `json:"dc2Uuid"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Status     string `json:"status"`
	OSType     string `json:"osType"`
}
