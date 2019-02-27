package schema

type EbsSpec struct {
	DiskType    string `json:"diskType"`
	MaxDiskSize int64  `json:"maxDiskSize"`
	MinDiskSize int64  `json:"minDiskSize"`
	Model       string `json:"model"`
	Name        string `json:"name"`
	Size        string `json:"size"`
}

type EbsInfo struct {
	Attr       string   `json:"attr"` // EBS 属性（`Root` 为根盘，`Data` 为数据盘）
	CreateTime int64    `json:"createTime"`
	Dc2        Dc2      `json:"dc2"`
	EbsTags    []string `json:"ebsTags"`
	EbsUuid    string   `json:"ebsUuid"`
	Job        Job      `json:"job"`
	Name       string   `json:"name"`
	Region     Region   `json:"region"` // 区域信息
	Size       string   `json:"size"`
	Spec       EbsSpec  `json:"spec"`
	UpdateTime int64    `json:"updateTime"`
}
