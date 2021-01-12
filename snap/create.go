package snap

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	. "github.com/shonenada/didiyun-go/schema"
)

type CreateRequest struct {
	RegionId string `json:"regionId"`
	Name     string `json:"snapName"`
	Dc2Uuid  string `json:"dc2Uuid"` // 根据 DC2 创建快照，与 ebsUuid 参数二选一
	EbsUuid  string `json:"ebsUuid"` // 根据 EBS 创建快照，与 dc2Uuid 参数二选一
}

type CreateResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) Create(request *CreateRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.CREATE_SNAP_URL, data)
	ret := CreateResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
