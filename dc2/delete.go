package dc2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/shonenada/didiyun-go/schema"
)

type DeleteDc2Request struct {
	RegionId string           `json:"regionId"`
	ZoneId   string           `json:"zoneId"`
	Dc2      []DeleteDc2Input `json:"dc2"` // 要删除的DC2信息，一次不能超过20台
}

type DeleteDc2Input struct {
	Dc2Uuid   string `json:"dc2Uuid"`   // 需要删除规格的DC2的uuid
	DeleteEip bool   `json:"deleteEip"` // 是否同时删除DC2上绑定的EIP
	DeleteEbs bool   `json:"deleteEbs"` // 是否同时删除DC2上绑定的EBS
	IgnoreSlb bool   `json:"ignoreSlb"` // 是否忽略DC2上绑定的负载均衡
}

type DeleteDc2Response struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) DeleteDc2(request *DeleteDc2Request) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(DELETE_DC2_URL, data)
	ret := DeleteDc2Response{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
