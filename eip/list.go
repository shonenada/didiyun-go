package eip

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type EipCondition struct {
	Uuids []string `json:"eipUuids,omitempty"`
	Eip   string   `json:"eip,omitempty"`
}

type Dc2Info struct {
	Uuid       string `json:"dc2Uuid"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Status     string `json:"status"`
	OsType     string `json:"osType"`
}

type ListRequest struct {
	RegionId  string       `json:"regionId"`
	Start     int          `json:"start"`
	Limit     int          `json:"limit"`
	Simplify  bool         `json:"simplify,omitempty"`
	condition EipCondition `json:"condition,omitempty"`
}

type ListResponse struct {
	Errno     int       `json:"errno"`
	Errmsg    string    `json:"errmsg"`
	RequestId string    `json:"requestId"`
	Data      []EipInfo `json:"data"`
}

type ListRequestBuilder struct {
	regionId string
	start    int
	limit    int
	simplify bool
	uuids    []string
	eip      string
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

func (b *ListRequestBuilder) SetSimplify(isSimplify bool) {
	b.simplify = isSimplify
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
		RegionId: b.regionId,
		Start:    start,
		Limit:    limit,
		Simplify: b.simplify,
		condition: EipCondition{
			Uuids: b.uuids,
			Eip:   b.eip,
		},
	}
}

func (c *Client) List(request *ListRequest) (*[]EipInfo, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(LIST_EIP_URL, data)
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
