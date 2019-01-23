package snap

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type DeleteRequest struct {
	RegionId string        `json:"regionId"`
	Snap     []DeleteInput `json:"snap"`
}

type DeleteInput struct {
	SnapUuid string `json:"ebsUuid"`
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
	body, err := c.HTTPPost(DELETE_SNAP_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := DeleteResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
