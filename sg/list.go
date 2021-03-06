package sg

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type SgCondition struct {
	SgUuids    []string `json:"sgUuid,omitempty"`
	VpcUuid    string   `json:"vpcUuid,omitempty"`
	Dc2Uuid    string   `json:"dc2Uuid,omitempty"`
	Dc2Exclude bool     `json:"dc2Exclude"`
}

type ListRequest struct {
	RegionId  string      `json:"regionId"`
	Start     int         `json:"start,omitempty"`
	Limit     int         `json:"limit,omitempty"`
	Condition SgCondition `json:"condition,omitempty"`
}

type ListResponse struct {
	Errno     int         `json:"errno"`
	Errmsg    string      `json:"errmsg"`
	RequestId string      `json:"requestId"`
	Data      []schema.Sg `json:"data"`
}

type ListRequestBuilder struct {
	regionId   string
	start      int
	limit      int
	sgUuids    []string
	vpcUuid    string
	dc2Uuid    string
	dc2Exclude bool
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

func (b *ListRequestBuilder) SetSgUuids(sgUuids []string) {
	b.sgUuids = sgUuids
}

func (b *ListRequestBuilder) SetVpcUuid(uuid string) {
	b.vpcUuid = uuid
}

func (b *ListRequestBuilder) SetDc2Uuid(dc2Uuid string) {
	b.dc2Uuid = dc2Uuid
}

func (b *ListRequestBuilder) SetDc2Exclude(isExclude bool) {
	b.dc2Exclude = isExclude
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
		RegionId: b.regionId,
		Start:    start,
		Limit:    limit,
		Condition: SgCondition{
			SgUuids:    b.sgUuids,
			VpcUuid:    b.vpcUuid,
			Dc2Uuid:    b.dc2Uuid,
			Dc2Exclude: b.dc2Exclude,
		},
	}
}

func (c *Client) List(request *ListRequest) (*[]schema.Sg, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.LIST_SG_URL, data)
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
