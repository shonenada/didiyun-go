package schema

type Dc2Condition struct {
	// 查询多个 VPC 下的 DC2 列表
	VpcUuids []string `json:"vpcUuids"`

	// 查询指定 VPC 下的 DC2 列表
	VpcUuid string `json:"vpcUuid"`

	// 模糊查询 DC2 名字
	Dc2Name string `json:"dc2Name"`

	// 查询此 SG 下的 DC2 列表
	SgUuid string `json:"sgUuid"`

	// 查询包含在此 UUID 列表内的 DC2
	Dc2Uuids []string `json:"dc2Uuids"`

	// 为 `true` 时表示查询不在此 SG 下的 DC2 列表
	SgExclude bool `json:"sgExclude"`

	// 精确匹配内网IP
	Ip string `json:"ip"`

	// 精确匹配公网EIP
	Eip string `json:"eip"`
}

type Dc2Info struct {
	// 此 DC2 正在进行的任务，若无任务则没有此字段
	Job Job `json:"job"`

	// DC2 唯一标识
	Uuid string `json:"dc2Uuid"`

	// DC2 名称
	Name string `json:"name"`

	// DC2 创建时间
	CreateTime int64 `json:"createTime"`

	// DC2 更新时间
	UpdateTime int64 `json:"updateTime"`

	// DC2 内网 IP
	Ip string `json:"ip"`

	// DC2 的 tags
	Tags []string `json:"tags"`

	// DC2 状态
	Status string `json:"status"`

	// DC2 操作系统发行版及版本号
	OSType string `json:"osType"`

	// 与 DC2 关联的 EIP 信息，没有 EIP 则没有该字段
	Eip Eip `json:"eip"`

	// 与 DC2 关联的 EBS 信息，没有 EBS 则没有该字段，如果是通用型 DC2，则必有这个字段，且根盘信息包含在内
	Ebs []Ebs `json:"ebs"`

	// region 信息
	Region Region `json:"region"`
}
