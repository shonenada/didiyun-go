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

type Ebs struct {
	// EBS 名称
	Name string `json:"name"`

	// 唯一标识
	Uuid string `json:"ebsUuid"`

	// 属性（`Root`为根盘，`Data`为数据盘）
	Attr string `json:"attr"`

	// 区域信息
	Region Region `json:"region"`

	// 创建时间
	CreateTime int64 `json:"createTime"`

	// 更新时间
	UpdateTime int64 `json:"updateTime"`
}
