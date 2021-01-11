package schema

type EipInfo struct {
	Job    Job      `json:"job"`
	Id     string   `json:"eipId"`
	Uuid   string   `json:"eipUuid"`
	Ip     string   `json:"ip"`
	Tags   []string `json:"eipTags"`
	State  string   `json:"state"`
	Status string   `json:"status"`
	Region Region   `json:"region"`
	Dc2    struct {
		Dc2Uuid    string `json:"dc2Uuid"`
		Name       string `json:name"`
		Status     string `json:"status"`
		OsType     string `json:"osType"`
		CreateTime int64  `json:"createTime"`
		UpdateTime int64  `json:"updateTime"`
	} `json:"dc2"`
}
