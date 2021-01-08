package slb

import (
	"encoding/json"
	"fmt"

	. "didiyun-go/schema"
)

type DeleteListenerRequest struct {
	SlbListener []struct {
		SlbListenerUuid string `json:"slbListenerUuid"`
	} `json:"slbListener"`
}

type DeleteListenerResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) DeleteListener(request *DeleteListenerRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(DELETE_LISTENER_SLB_URL, data)
	ret := DeleteListenerResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
