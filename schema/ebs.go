package schema

type EbsSpec struct {
	DiskType    string `json:"diskType"`
	Model       string `json:"model"`
	Name        string `json:"name"`
	MaxDiskSize int64  `json:"maxDiskSize"`
	MinDiskSize int64  `json:"minDiskSize"`
	Size        string `json:"size"`
}

type EbsInfo struct {
	Job        Job      `json:"job"`
	EbsUuid    string   `json:"ebsUuid"`
	Name       string   `json:"name"`
	Attr       string   `json:"attr"` // EBS 属性（`Root` 为根盘，`Data` 为数据盘）
	Size       string   `json:"size"`
	CreateTime int64    `json:"createTime"`
	Region     Region   `json:"region"` // 区域信息
	UpdateTime int64    `json:"updateTime"`
	EbsTags    []string `json:"ebsTags"`
	Dc2        Dc2      `json:"dc2"`
	Spec       EbsSpec  `json:"spec"`
}
