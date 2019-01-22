package schema

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
