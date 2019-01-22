package dc2

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type DeleteRequest struct {
	RegionId string        `json:"regionId"`
	ZoneId   string        `json:"zoneId"`
	Dc2      []DeleteInput `json:"dc2"` // 要删除的DC2信息，一次不能超过20台
}

type DeleteInput struct {
	Dc2Uuid   string `json:"dc2Uuid"`   // 需要删除规格的DC2的uuid
	DeleteEip bool   `json:"deleteEip"` // 是否同时删除DC2上绑定的EIP
	DeleteEbs bool   `json:"deleteEbs"` // 是否同时删除DC2上绑定的EBS
	IgnoreSlb bool   `json:"ignoreSlb"` // 是否忽略DC2上绑定的负载均衡
}

type DeleteResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) Delete(request *DeleteRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(DELETE_DC2_URL, data)
	ret := DeleteResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
