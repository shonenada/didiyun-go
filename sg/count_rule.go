package sg

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
)

type CountRuleRequest struct {
	RegionId string `json:"regionId"`
	Uuid     string `json:"sgUuid,omitempty"`
	Dc2Uuid  string `json:"dc2Uuid,omitempty"`
	Type     string `json:"type"`
}

type CountRuleResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []struct {
		TotalCount int `json:"totalCnt"`
	} `json:"data"`
}

func (c *Client) CountRule(request *CountRuleRequest) (int, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return -1, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.COUNT_SG_RULE_URL, data)
	if err != nil {
		return -1, fmt.Errorf("Error: %s", err)
	}
	ret := CountRuleResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return -1, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return ret.Data[0].TotalCount, nil
}
