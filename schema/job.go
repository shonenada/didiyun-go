package schema

type Job struct {
	IsDone       bool    `json:"done"`
	Progress     float64 `json:"progress"`
	ResourceUuid string  `json:"resourceUuid"`
	Result       string  `json:"result"`
	IsSuccess    bool    `json:"success"`
	Type         string  `json:"type"`
	Uuid         string  `json:"jobUuid"`
}
