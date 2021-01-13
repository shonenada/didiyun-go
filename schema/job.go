package schema

type Job struct {
	Uuid         string  `json:"jobUuid"`
	Type         string  `json:"type"`
	ResourceUuid string  `json:"resourceUuid"`
	IsDone       bool    `json:"done"`
	Progress     float64 `json:"progress"`
	Result       string  `json:"result"`
	IsSuccess    bool    `json:"success"`
}
