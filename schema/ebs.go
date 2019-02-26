package schema

type Ebs struct {
	Name       string `json:"name"`       // EBS 名称
	Uuid       string `json:"ebsUuid"`    // 唯一标识
	Attr       string `json:"attr"`       // 属性（`Root`为根盘，`Data`为数据盘）
	Region     Region `json:"region"`     // 区域信息
	CreateTime int64  `json:"createTime"` // 创建时间
	UpdateTime int64  `json:"updateTime"` // 更新时间
}

type EbsSpec struct {
	DiskType string `json:"diskType"`
	Model    string `json:"model"`
	Name     string `json:"name"`
	Size     string `json:"size"`
}

type EbsInfo struct {
	Job        Job      `json:"job"`
	EbsUuid    string   `json:"ebsUuid"`
	Name       string   `json:"name"`
	Attr       string   `json:"attr"` // EBS 属性（`Root` 为根盘，`Data` 为数据盘）
	Size       string   `json:"size"`
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
	EbsTags    []string `json:"ebsTags"`
	Dc2        Dc2      `json:"dc2"`
	Spec       EbsSpec  `json:"spec"`
}
