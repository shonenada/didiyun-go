package schema

type SnapDc2 struct {
	Uuid string `json:"dc2Uuid"`
	Name string `json:"name"`
}

type SnapEbs struct {
	Uuid     string `json:"ebsUuid"`
	Name     string `json:"name"`
	DiskType string `json:"diskType"`
}

type SnapInfo struct {
	Job           Job     `json:"job"`
	SnapUuid      string  `json:"snapUuid"`
	Name          string  `json:"name"`
	CanBeReverted bool    `json:"canBeReverted"` // 是否可以回滚
	IsGeneral     bool    `json:"isGeneral"`     // 是否通用系统(e1,g1,s1) DC2 快照
	CreateTime    int64   `json:"createTime"`
	UpdateTime    int64   `json:"updateTime"`
	Size          int64   `json:"size"`
	VolumeSize    int64   `json:"volumeSize"`
	Type          string  `json:"type"`
	Dc2           SnapDc2 `json:"dc2"`
	Ebs           SnapEbs `json:"ebs"`
	Region        Region  `json:"region"`
}
