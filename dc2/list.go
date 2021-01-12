package dc2

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	. "github.com/shonenada/didiyun-go/schema"
)

type Dc2Condition struct {
	Name      string   `json:"dc2Name"`   // 模糊查询 DC2 名字
	Uuids     []string `json:"dc2Uuids"`  // 查询包含在此 UUID 列表内的 DC2
	Eip       string   `json:"eip"`       // 精确匹配公网EIP
	Ip        string   `json:"ip"`        // 精确匹配内网IP
	SgUuid    string   `json:"sgUuid"`    // 查询此 SG 下的 DC2 列表
	SgExclude bool     `json:"sgExclude"` // 为 `true` 时表示查询不在此 SG 下的 DC2 列表
	VpcUuid   string   `json:"vpcUuid"`   // 查询指定 VPC 下的 DC2 列表
	VpcUuids  []string `json:"vpcUuids"`  // 查询多个 VPC 下的 DC2 列表
}

type ListRequest struct {
	RegionId   string       `json:"regionId"`           // 地域 ID
	ZoneId     string       `json:"zondId,omitempty"`   // 可用区 ID，希望查询的可用区，不传则表示查询此地域下的所有可用区
	Start      int          `json:"start"`              // 查询 DC2 列表起始 index，从 0 开始
	Limit      int          `json:"limit"`              // 查询 DC2 列表元素数量
	IsSimplify bool         `json:"simplify,omitempty"` // 是否简化输出
	Condition  Dc2Condition `json:"condition"`          // 查询 DC2 列表筛选条件
}

type ListResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Dc2  `json:"data"`
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
		RegionId: b.regionId,
		ZoneId:   b.zoneId,
		Start:    start,
		Limit:    limit,
		Simplify: b.simplify,
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

func (c *Client) List(request *ListRequest) (*[]Dc2, error) {
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
