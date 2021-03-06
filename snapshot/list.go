package snapshot

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type SnapCondition struct {
	Name    string `json:"snapName,omitempty"`
	Dc2Uuid string `json:"dc2Uuid,omitempty"`
	EbsUuid string `json:"ebsUuid,omitempty"`
}

type ListRequest struct {
	RegionId   string        `json:"regionId"`
	ZoneId     string        `json:"zoneId,omitempty"`
	Start      int           `json:"start"`
	Limit      int           `json:"limit"`
	IsSimplify bool          `json:"simplify,omitempty"`
	condition  SnapCondition `json:"condition,omitempty"`
}

type ListResponse struct {
	Errno     int               `json:"errno"`
	Errmsg    string            `json:"errmsg"`
	RequestId string            `json:"requestId"`
	Data      []schema.Snapshot `json:"data"`
}

type ListRequestBuilder struct {
	regionId   string
	zoneId     string
	start      int
	limit      int
	isSimplify bool
	dc2Uuid    string
	ebsUuid    string
	name       string
}

func (b *ListRequestBuilder) SetRegionId(regionId string) {
	b.regionId = regionId
}

func (b *ListRequestBuilder) SetZoneId(zoneId string) {
	b.zoneId = zoneId
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

func (b *ListRequestBuilder) SetDc2Uuids(uuid string) {
	b.dc2Uuid = uuid
}

func (b *ListRequestBuilder) SetEbsUuid(uuid string) {
	b.ebsUuid = uuid
}

func (b *ListRequestBuilder) SetSnapName(name string) {
	b.name = name
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
		ZoneId:     b.zoneId,
		Start:      start,
		Limit:      limit,
		IsSimplify: b.isSimplify,
		condition: SnapCondition{
			Dc2Uuid: b.dc2Uuid,
			EbsUuid: b.ebsUuid,
			Name:    b.name,
		},
	}
}

func (c *Client) List(request *ListRequest) ([]schema.Snapshot, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.LIST_SNAP_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ListResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return ret.Data, nil
}
