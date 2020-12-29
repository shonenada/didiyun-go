package slb

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type UpdateListenerRequest struct {
	SlbListener []Listener `json:"listener"`
}

type UpdateListenrResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) UpdateListener(request *UpdateListenerRequest) (*[]SlbRe, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(UPDATE_LISTENER_SLB_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := UpdateListenerResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
