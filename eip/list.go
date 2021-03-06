package eip

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type EipCondition struct {
	Uuids []string `json:"eipUuids,omitempty"`
	Eip   string   `json:"eip,omitempty"`
}

type ListRequest struct {
	RegionId   string       `json:"regionId"`
	Start      int          `json:"start"`
	Limit      int          `json:"limit"`
	IsSimplify bool         `json:"simplify,omitempty"`
	condition  EipCondition `json:"condition,omitempty"`
}

type ListResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Eip `json:"data"`
}

type ListRequestBuilder struct {
	regionId   string
	start      int
	limit      int
	isSimplify bool
	uuids      []string
	eip        string
}

func (b *ListRequestBuilder) SetRegionId(regionId string) {
	b.regionId = regionId
}

func (b *ListRequestBuilder) SetStart(start int) {
	b.start = start
}

func (b *ListRequestBuilder) SetLimit(limit int) {
	b.limit = limit
}

func (b *ListRequestBuilder) SetIsSimplify(isSimplify bool) {
	b.isSimplify = isSimplify
}

func (b *ListRequestBuilder) SetUuids(uuids []string) {
	b.uuids = uuids
}

func (b *ListRequestBuilder) SetEip(eip string) {
	b.eip = eip
}

func (b *ListRequestBuilder) Build() ListRequest {
	start := 0
	if b.start >= 0 {
		start = b.start
	}

	limit := 20
	if b.limit > 0 {
		limit = b.limit
	}

	return ListRequest{
		RegionId:   b.regionId,
		Start:      start,
		Limit:      limit,
		IsSimplify: b.isSimplify,
		condition: EipCondition{
			Uuids: b.uuids,
			Eip:   b.eip,
		},
	}
}

func (c *Client) List(request *ListRequest) (*[]schema.Eip, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.LIST_EIP_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ListResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
