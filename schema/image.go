package schema

type Image struct {
	Uuid              string           `json:"imgUuid"`
	Name              string           `json:"name"`
	Description       string           `json:"description"`
	OsArch            string           `json:"osArch" default:""`
	OsFamily          string           `json:"osFamily"`
	OsVersion         string           `json:"osVersion"`
	Platform          string           `json:"platform"`
	ImgType           string           `json:"type"`
	Scenes            []string         `json:"scenes"`
	Regions           []string         `json:"regions"`
	SupportedAccounts []string         `json:"supportedAccounts"`
	Requirement       ImageRequirement `json:"requirement"`
}

type ImageRequirement struct {
	MinCpuNum     int      `json:"minCpuNum"`
	MinDiskSize   int      `json:"minDiskSize"`
	MinMemorySize int      `json:"minMemorySize"`
	Dc2Types      []string `json:"dc2Types"`
}
