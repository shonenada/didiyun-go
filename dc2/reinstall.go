package dc2

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	. "github.com/shonenada/didiyun-go/schema"
)

type ReinstallParams struct {
	Uuid                      string   `json:"dc2Uuid"`
	ImgUuid                   string   `json:"imgUuid"`
	PubKeyUuids               []string `json:"pubKeyUuids"`
	Password                  string   `json:"password"`
	IsProSecurityAgentEnabled bool     `json:"proSecurityAgentEnabled,omitempty"`
	IsMonitoringAgentEnabled  bool     `json:"monitoringAgentEnabled,omitempty"`
}

type ReinstallRequest struct {
	RegionId string          `json:"regionId"`
	ZoneId   string          `json:"zoneId,omitempty"`
	Dc2      ReinstallParams `json:"dc2"`
}

type ReinstallResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) Reinstall(request *ReinstallRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.REINSTALL_DC2_URL, data)
	ret := ReinstallResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
