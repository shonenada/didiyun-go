package ebs

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type ListRequest struct {
	RegionId  string        `json:"regionId"`
	ZoneId    string        `json:"zoneId,omitempty"`
	Start     int           `json:"start,omitempty"`
	Limit     int           `json:"limit,omitempty"`
	Condition ListCondition `json:"condition"`
}

type ListCondition struct {
	Dc2Uuids []string `json:"dc2Uuids"`
}

type ListResponse struct {
	Errno     int       `json:"errno"`
	Errmsg    string    `json:"errmsg"`
	RequestId string    `json:"requestId"`
	Data      []EbsInfo `json:"data"`
}

type EbsInfo struct {
	Job        Job      `json:"job"`
	EbsUuid    string   `json:"ebsUuid"`
	Name       string   `json:"name"`
	Attr       string   `json:"attr"` // EBS 属性（`Root` 为根盘，`Data` 为数据盘）
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
	EbsTags    []string `json:"ebsTags"`
	Dc2        Dc2      `json:"dc2"`
}

type Dc2 struct {
	Uuid       string `json:"dc2Uuid"`
	Name       string `json:"name"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
	Status     string `json:"status"`
	OSType     string `json:"osType"`
}

func (c *Client) List(request *ListRequest) (*[]EbsInfo, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(LIST_EBS_URL, data)
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
