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

type ListDc2Response struct {
	Errno     int       `json:"errno"`
	Errmsg    string    `json:"errmsg"`
	RequestId string    `json:"requestId"`
	Data      []Dc2Info `json:"data"`
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

func (c *Client) ListDc2(reqBody *ListDc2Request) (*[]Dc2Info, error) {
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
