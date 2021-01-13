package slb

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type GetRequest struct {
	RegionId string `json:"regionId"`
	Uuid     string `json:"slbUuid"`
}

type GetResponse struct {
	Errno     int           `json:"errno"`
	Errmsg    string        `json:"errmsg"`
	RequestId string        `json:"requestId"`
	Data      []SlbResponse `json:"data"`
}

type SlbResponse struct {
	Job        schema.Job     `json:"job"`
	SlbUuid    string         `json:"slbUuid"`
	Name       string         `json:"name"`
	Ip         string         `json:"ip"`
	WafStatus  string         `json:"wafStatus"`
	CreateTime int            `json:"createTime"`
	Updateime  int            `json:"updateTime"`
	BeIP       schema.BEIP    `json:"beip"`
	Vpc        schema.VPC     `json:"vpc"`
	Flow       schema.Flow    `json:"flow"`
	Spec       schema.SlbSpec `json:"spec"`
	Region     schema.Region  `json:"region"`
}

func (c *Client) Get(request *GetRequest) (*[]SlbResponse, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.GET_SLB_URL, data)
	ret := GetResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
