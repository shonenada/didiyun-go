package dc2

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type Dc2Condition struct {
	Name      string   `json:"dc2Name"`
	Uuids     []string `json:"dc2Uuids"`
	Eip       string   `json:"eip"`
	Ip        string   `json:"ip"`
	SgUuid    string   `json:"sgUuid"`
	SgExclude bool     `json:"sgExclude"`
	VpcUuid   string   `json:"vpcUuid"`
	VpcUuids  []string `json:"vpcUuids"`
}

type ListRequest struct {
	RegionId   string       `json:"regionId"`
	ZoneId     string       `json:"zondId,omitempty"`
	Start      int          `json:"start"`
	Limit      int          `json:"limit"`
	IsSimplify bool         `json:"simplify,omitempty"`
	Condition  Dc2Condition `json:"condition"`
}

type ListResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Dc2 `json:"data"`
}

type ListRequestBuilder struct {
	regionId   string
	zoneId     string
	start      int
	limit      int
	isSimplify bool
	vpcUuid    string
	vpcUuids   []string
	name       string
	sgUuid     string
	uuids      []string
	sgExclude  bool
	ip         string
	eip        string
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

func (b *ListRequestBuilder) SetVpcUuid(uuid string) {
	b.vpcUuid = uuid
}

func (b *ListRequestBuilder) SetVpcUuids(uuids []string) {
	b.vpcUuids = uuids
}

func (b *ListRequestBuilder) SetName(name string) {
	b.name = name
}

func (b *ListRequestBuilder) SetSgUuid(uuid string) {
	b.sgUuid = uuid
}

func (b *ListRequestBuilder) SetUuids(uuids []string) {
	b.uuids = uuids
}

func (b *ListRequestBuilder) SetSgExclude(isExclude bool) {
	b.sgExclude = isExclude
}

func (b *ListRequestBuilder) SetIp(ip string) {
	b.ip = ip
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
		ZoneId:     b.zoneId,
		Start:      start,
		Limit:      limit,
		IsSimplify: b.isSimplify,
		Condition: Dc2Condition{
			Name:      b.name,
			VpcUuids:  b.vpcUuids,
			VpcUuid:   b.vpcUuid,
			SgUuid:    b.sgUuid,
			Uuids:     b.uuids,
			SgExclude: b.sgExclude,
			Ip:        b.ip,
			Eip:       b.eip,
		},
	}
}

func (c *Client) List(request *ListRequest) (*[]schema.Dc2, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.LIST_DC2_URL, data)
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
