package snap

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type RevertRequest struct {
	RegionId string        `json:"regionId"`
	Snap     []RevertInput `json:"snap"`
	StartDc2 bool          `json:"startDc2"` // 还原后是否需要同时启动 DC2
	StopDc2  bool          `json:"stopDc2"`  // 还原前，是否执行关闭 DC2
}

type RevertInput struct {
	SnapUuid string `json:"snapUuid"`
}

type RevertResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) Revert(request *RevertRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(REVERT_SNAP_URL, data)
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
