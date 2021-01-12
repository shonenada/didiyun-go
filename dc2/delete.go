package dc2

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	. "github.com/shonenada/didiyun-go/schema"
)

type DeleteRequest struct {
	RegionId  string        `json:"regionId"`
	ZoneId    string        `json:"zoneId,omitempty"`
	Dc2       []DeleteInput `json:"dc2"`                 // 要删除的DC2信息，一次不能超过20台
	DeleteEip bool          `json:"deleteEip,omitempty"` // 是否同时删除 DC2 上绑定的 EIP
	DeleteEbs bool          `json:"deleteEbs,omitempty"` // 是否同时删除 DC2 上绑定的 EBS
	IgnoreSLB bool          `json:"ignoreSlb,omitempty"` // 是否忽略 DC2 上绑定的负载均衡
}

type DeleteInput struct {
	Dc2Uuid string `json:"dc2Uuid"` // 需要删除规格的 DC2 的uuid
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
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.DELETE_DC2_URL, data)
	ret := DeleteResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
