package eip

import (
	"encoding/json"

	"github.com/shonenada/didiyun-go/schame"
)

type EipCondition struct {
	Uuids []string `json:"eipUuids"`
	Eip   string   `json:"eip"`
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
	Simplify  bool         `json:"simplify"`
	condition EipCondition `json:"condition"`
}

type EipInfo struct {
	Job        Job      `json:"json"`
	Uuid       string   `json:"eipUuid"`
	Ip         string   `json:"ip"`
	CreateTime int64    `json:"createTime"`
	UpdateTime int64    `json:"updateTime"`
	EipTags    []string `json:"eipTags"`
	Dc2        Dc2Info  `json:"dc2"`
}

type ListResponse struct {
	Errno     int       `json:"errno"`
	Errmsg    string    `json:"errmsg"`
	RequestId string    `json:"requestId"`
	Data      []EipInfo `json:"data"`
}

func (c *Client) List(request) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(LIST_EIP_URL, data)
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
