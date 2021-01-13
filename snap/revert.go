package snap

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type RevertRequest struct {
	RegionId   string         `json:"regionId"`
	Snap       []RevertParams `json:"snap"`
	IsStartDc2 bool           `json:"startDc2"` // 还原后是否需要同时启动 DC2
	IsStopDc2  bool           `json:"stopDc2"`  // 还原前，是否执行关闭 DC2
}

type RevertParams struct {
	Uuid string `json:"snapUuid"`
}

type RevertResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Job `json:"data"`
}

func (c *Client) Revert(request *RevertRequest) (*schema.Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.REVERT_SNAP_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := RevertResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
