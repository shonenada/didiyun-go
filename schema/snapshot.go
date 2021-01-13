package schema

type Snapshot struct {
	Uuid            string      `json:"snapUuid"`
	Name            string      `json:"name"`
	IsCanBeReverted bool        `json:"canBeReverted"`
	IsGeneral       bool        `json:"isGeneral"`
	Size            int64       `json:"size"`
	VolumeSize      int64       `json:"volumeSize"`
	Type            string      `json:"type"`
	Dc2             SnapshotDc2 `json:"dc2"`
	Ebs             SnapshotEbs `json:"ebs"`
	Region          Region      `json:"region"`
	Job             Job         `json:"job"`
	CreateTime      int64       `json:"createTime"`
	UpdateTime      int64       `json:"updateTime"`
}

type SnapshotDc2 struct {
	Uuid string `json:"dc2Uuid"`
	Name string `json:"name"`
}

type SnapshotEbs struct {
	Uuid     string `json:"ebsUuid"`
	Name     string `json:"name"`
	DiskType string `json:"diskType"`
}
