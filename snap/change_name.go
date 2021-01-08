package snap

import (
	"encoding/json"
	"fmt"

	. "didiyun-go/schema"
)

type ChangeNameRequest struct {
	RegionId string            `json:"regionId"`
	Snap     []ChangeNameInput `json:"snap"`
}

type ChangeNameInput struct {
	SnapUuid string `json:"snapUuid"`
	Name     string `json:"name"`
}

type ChangeNameResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) ChangeName(request *ChangeNameRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(CHANGE_NAME_SNAP_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ChangeNameResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
