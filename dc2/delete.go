package dc2

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type DeleteRequest struct {
	RegionId    string         `json:"regionId"`
	ZoneId      string         `json:"zoneId,omitempty"`
	Dc2         []DeleteParams `json:"dc2"`                 // 要删除的DC2信息，一次不能超过20台
	IsDeleteEip bool           `json:"deleteEip,omitempty"` // 是否同时删除 DC2 上绑定的 EIP
	IsDeleteEbs bool           `json:"deleteEbs,omitempty"` // 是否同时删除 DC2 上绑定的 EBS
	IsIgnoreSLB bool           `json:"ignoreSlb,omitempty"` // 是否忽略 DC2 上绑定的负载均衡
}

type DeleteParams struct {
	Uuid string `json:"dc2Uuid"` // 需要删除规格的 DC2 的uuid
}

type DeleteResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Job `json:"data"`
}

func (c *Client) Delete(request *DeleteRequest) (*schema.Job, error) {
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
