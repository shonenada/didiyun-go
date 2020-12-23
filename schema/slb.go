package schema

type SlbCondition struct {
	SlbUuids []string `json:"slbUuids,omitempty"`
	VpcUuids []string `json:"vpcUuids,omitempty"`
	Beips    []string `json:"beips,omitempty"`
	Dc2Ips   []string `json:"dc2ips,omitempty"`
	Ips      []string `json:"ips,omitempty"`
}
