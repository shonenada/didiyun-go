package schema

type Dc2Condition struct {
	Dc2Name   string   `json:"dc2Name"`   // 模糊查询 DC2 名字
	Dc2Uuids  []string `json:"dc2Uuids"`  // 查询包含在此 UUID 列表内的 DC2
	Eip       string   `json:"eip"`       // 精确匹配公网EIP
	Ip        string   `json:"ip"`        // 精确匹配内网IP
	SgUuid    string   `json:"sgUuid"`    // 查询此 SG 下的 DC2 列表
	SgExclude bool     `json:"sgExclude"` // 为 `true` 时表示查询不在此 SG 下的 DC2 列表
	VpcUuid   string   `json:"vpcUuid"`   // 查询指定 VPC 下的 DC2 列表
	VpcUuids  []string `json:"vpcUuids"`  // 查询多个 VPC 下的 DC2 列表
}

type Dc2Vpc struct {
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
	VpcId      string   `json:"vpcId"`
	VpcTags    []string `json:"vpcTags"`
	VpcUuid    string   `json:"vpcUuid"`
}

type Dc2Region struct {
	AreaName string `json:"areaName"` // 区域名称
	Id       string `json:"id"`       // region id
	Name     string `json:"name"`     // region name
	Zone     Zone   `json:"zone"`     // zone信息
}

type Dc2EbsSpec struct {
	DiskType    string `json:"diskType"`
	MaxDiskSize int64  `json:"maxDiskSize"`
	MinDiskSize int64  `json:"minDiskSize"`
	Model       string `json:"model"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
}

type Dc2EbsInfo struct {
	Attr       string     `json:"attr"` // EBS 属性（`Root` 为根盘，`Data` 为数据盘）
	CreateTime int64      `json:"createTime"`
	Dc2        Dc2        `json:"dc2"`
	EbsTags    []string   `json:"ebsTags"`
	EbsUuid    string     `json:"ebsUuid"`
	Job        Job        `json:"job"`
	Name       string     `json:"name"`
	Region     Region     `json:"region"` // 区域信息
	Size       int64      `json:"size"`
	Spec       Dc2EbsSpec `json:"spec"`
	UpdateTime int64      `json:"updateTime"`
}
type Dc2EipInfo struct {
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
type Dc2Info struct {
	Charge     Charge     `json:"charge"`
	CreateTime int64      `json:"createTime"` // DC2 创建时间
	Ebs        []EbsInfo  `json:"ebs"`        // 与 DC2 关联的 EBS 信息，没有 EBS 则没有该字段，如果是通用型 DC2，则必有这个字段，且根盘信息包含在内
	Eip        Dc2EipInfo `json:"eip"`        // 与 DC2 关联的 EIP 信息，没有 EIP 则没有该字段
	Id         string     `json:"dc2Id"`
	ImageUuid  string     `json:"imgUuid"`
	Ip         string     `json:"ip"`     // DC2 内网 IP
	Name       string     `json:"name"`   // DC2 名称
	Job        Job        `json:"job"`    // 此 DC2 正在进行的任务，若无任务则没有此字段
	OSType     string     `json:"osType"` // DC2 操作系统发行版及版本号
	Platform   string     `json:"platform"`
	Region     Dc2Region  `json:"region"` // region 信息
	Spec       Dc2Spec    `json:"spec"`
	Status     string     `json:"status"`     // DC2 状态
	Tags       []string   `json:"dc2Tags"`    // DC2 的 tags
	UpdateTime int64      `json:"updateTime"` // DC2 更新时间
	Uuid       string     `json:"dc2Uuid"`    // DC2 唯一标识
	Vpc        Dc2Vpc     `json:"vpc"`
}

type Dc2 struct {
	CreateTime int64  `json:"createTime"`
	Name       string `json:"name"`
	OSType     string `json:"osType"`
	Status     string `json:"status"`
	UpdateTime int64  `json:"updateTime"`
	Uuid       string `json:"dc2Uuid"`
}
