package schema

type Region struct {
	AreaName string `json:"areaName"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Zone     []Zone `json:"zone"`
}

type Zone struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
