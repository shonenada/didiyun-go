package sg

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type ListRuleRequest struct {
	RegionId  string          `json:"regionId"`
	Start     int             `json:"start"`
	Limit     int             `json:"limit"`
	Condition SgRuleCondition `json:"condition"`
}

type ListRuleResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []SgRuleInfo `json:"data"`
}

type ListRuleRequestBuilder struct {
	regionId string
	start    int
	limit    int
	sgUuid   string
	dc2Uuid  string
	ruleType string
}

func (b *ListRuleRequestBuilder) SetRegionId(regionId string) {
	b.regionId = regionId
}

func (b *ListRuleRequestBuilder) SetStart(start int) {
	b.start = start
}

func (b *ListRuleRequestBuilder) SetLimit(limit int) {
	b.limit = limit
}

func (b *ListRuleRequestBuilder) SetSgUuid(sgUuid string) {
	b.sgUuid = sgUuid
}

func (b *ListRuleRequestBuilder) SetDc2Uuid(dc2Uuid string) {
	b.dc2Uuid = dc2Uuid
}

func (b *ListRuleRequestBuilder) SetType(ruleType string) {
	b.ruleType = ruleType
}

func (b *ListRuleRequestBuilder) Build() ListRuleRequest {
	start := 0
	if b.start >= 0 {
		start = b.start
	}

	limit := 20
	if b.limit > 0 {
		limit = b.limit
	}

	return ListRuleRequest{
		RegionId: b.regionId,
		Start:    start,
		Limit:    limit,
		Condition: SgRuleCondition{
			SgUuid:  b.sgUuid,
			Dc2Uuid: b.dc2Uuid,
			Type:    b.ruleType,
		},
	}
}

func (c *Client) ListRule(request *ListRuleRequest) (*[]SgRuleInfo, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(LIST_SG_RULE_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ListRuleResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
