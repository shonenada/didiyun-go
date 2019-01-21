package dc2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/shonenada/didiyun-go/schema"
)

type ListDc2Request struct {
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

type Dc2Response struct {
	// 此 DC2 正在进行的任务，若无任务则没有此字段
	Job Job `json:"job"`

	// DC2 唯一标识
	Uuid string `json:"dc2Uuid"`

	// DC2 名称
	Name string `json:"name"`

	// DC2 创建时间
	CreateTime int64 `json:"createTime"`

	// DC2 更新时间
	UpdateTime int64 `json:"updateTime"`

	// DC2 内网 IP
	Ip string `json:"ip"`

	// DC2 的 tags
	Tags []string `json:"tags"`

	// DC2 状态
	Status string `json:"status"`

	// DC2 操作系统发行版及版本号
	OSType string `json:"osType"`

	// 与 DC2 关联的 EIP 信息，没有 EIP 则没有该字段
	Eip Eip `json:"eip"`

	// 与 DC2 关联的 EBS 信息，没有 EBS 则没有该字段，如果是通用型 DC2，则必有这个字段，且根盘信息包含在内
	Ebs []Ebs `json:"ebs"`

	// region 信息
	Region Region `json:"region"`
}

type ListDc2Response struct {
	Errno     int           `json:"errno"`
	Errmsg    string        `json:"errmsg"`
	RequestId string        `json:"requestId"`
	Data      []Dc2Response `json:"data"`
}

type ListDc2RequestBuilder struct {
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

func (b *ListDc2RequestBuilder) SetRegionId(regionId string) {
	b.regionId = regionId
}

func (b *ListDc2RequestBuilder) SetZoneId(zoneId string) {
	b.zoneId = zoneId
}

func (b *ListDc2RequestBuilder) SetStart(start int) {
	b.start = start
}

func (b *ListDc2RequestBuilder) SetLimit(limit int) {
	b.limit = limit
}

func (b *ListDc2RequestBuilder) SetSimplify(simplify bool) {
	b.simplify = simplify
}

func (b *ListDc2RequestBuilder) SetVpcUuid(uuid string) {
	b.vpcUuid = uuid
}

func (b *ListDc2RequestBuilder) SetVpcUuids(uuids []string) {
	b.vpcUuids = uuids
}

func (b *ListDc2RequestBuilder) SetDc2Name(dc2name string) {
	b.dc2Name = dc2name
}

func (b *ListDc2RequestBuilder) SetSgUuid(uuid string) {
	b.sgUuid = uuid
}

func (b *ListDc2RequestBuilder) SetDc2Uuids(uuids []string) {
	b.dc2Uuids = uuids
}

func (b *ListDc2RequestBuilder) SetSgExclude(isExclude bool) {
	b.sgExclude = isExclude
}

func (b *ListDc2RequestBuilder) SetIp(ip string) {
	b.ip = ip
}

func (b *ListDc2RequestBuilder) SetEip(eip string) {
	b.eip = eip
}

func (b *ListDc2RequestBuilder) Build() ListDc2Request {
	start := 0
	if b.start >= 0 {
		start = b.start
	}

	limit := 20
	if b.limit > 0 {
		limit = b.limit
	}

	return ListDc2Request{
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

func (c *Client) ListDc2(reqBody *ListDc2Request) (*[]Dc2Response, error) {
	httpClient := &http.Client{}
	reqStr, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}

	req, err := http.NewRequest("POST", LIST_DC2_URL, bytes.NewBuffer([]byte(reqStr)))
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := ListDc2Response{}
	json.Unmarshal([]byte(body), &ret)

	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}

	return &ret.Data, nil
}
