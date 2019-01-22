package dc2

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type ListRequest struct {
	// 地域 ID
	RegionId string `json:"regionId"`

	// 可用区 ID，希望查询的可用区，不传则表示查询此地域下的所有可用区
	ZoneId string `json:"zondId"`

	// 查询 DC2 列表起始 index，从 0 开始
	Start int `json:"start"`

	// 查询 DC2 列表元素数量
	Limit int `json:"limit"`

	// 是否简化输出
	Simplify bool `json:"simplify"`

	// 查询 DC2 列表筛选条件
	condition Dc2Condition `json:"condition"`
}

type ListResponse struct {
	Errno     int       `json:"errno"`
	Errmsg    string    `json:"errmsg"`
	RequestId string    `json:"requestId"`
	Data      []Dc2Info `json:"data"`
}

type ListRequestBuilder struct {
	regionId  string
	zoneId    string
	start     int
	limit     int
	simplify  bool
	vpcUuid   string
	vpcUuids  []string
	dc2Name   string
	sgUuid    string
	dc2Uuids  []string
	sgExclude bool
	ip        string
	eip       string
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

func (b *ListRequestBuilder) SetSimplify(simplify bool) {
	b.simplify = simplify
}

func (b *ListRequestBuilder) SetVpcUuid(uuid string) {
	b.vpcUuid = uuid
}

func (b *ListRequestBuilder) SetVpcUuids(uuids []string) {
	b.vpcUuids = uuids
}

func (b *ListRequestBuilder) SetDc2Name(dc2name string) {
	b.dc2Name = dc2name
}

func (b *ListRequestBuilder) SetSgUuid(uuid string) {
	b.sgUuid = uuid
}

func (b *ListRequestBuilder) SetDc2Uuids(uuids []string) {
	b.dc2Uuids = uuids
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
		RegionId: b.regionId,
		ZoneId:   b.zoneId,
		Start:    start,
		Limit:    limit,
		Simplify: b.simplify,
		condition: Dc2Condition{
			VpcUuids:  b.vpcUuids,
			VpcUuid:   b.vpcUuid,
			SgUuid:    b.sgUuid,
			Dc2Uuids:  b.dc2Uuids,
			SgExclude: b.sgExclude,
			Ip:        b.ip,
			Eip:       b.eip,
		},
	}
}

func (c *Client) List(request *ListRequest) (*[]Dc2Info, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(LIST_DC2_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := ListResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
