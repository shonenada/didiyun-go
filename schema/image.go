package schema

type Image struct {
	Uuid              string           `json:"imgUuid"`
	Name              string           `json:"name"`
	Description       string           `json:"description"`
	OsArch            string           `json:"osArch" default:""`
	OsFamily          string           `json:"osFamily"`  // 操作系统发行版本
	OsVersion         string           `json:"osVersion"` // 操作系统版本号
	Platform          string           `json:"platform"`  //  镜像平台，`Linux` 或 `Windows`
	ImgType           string           `json:"type"`      // 镜像类型，`standard` 为标准镜像，`oneClick` 为一键部署镜像
	Scenes            []string         `json:"scenes"`    // 镜像使用场景，`base` 为基础镜像，`gpu` 为带GPU驱动的镜像
	Regions           []string         `json:"regions"`
	SupportedAccounts []string         `json:"supportedAccounts"`
	Requirement       ImageRequirement `json:"requirement"` // 镜像对于 DC2 配置的最小要求
}

type ImageRequirement struct {
	MinCpuNum     int      `json:"minCpuNum"`
	MinDiskSize   int      `json:"minDiskSize"`
	MinMemorySize int      `json:"minMemorySize"`
	Dc2Types      []string `json:"dc2Types"`
}
