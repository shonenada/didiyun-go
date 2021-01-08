package slb

import (
	"encoding/json"
	"fmt"

	. "didiyun-go/schema"
)

type CreateListenerRequest struct {
	SlbUuid     string     `json:"slbUuid"`
	SlbListener []Listener `json:"slbListener"`
}

func (c *Client) CreateListener(request *CreateListenerRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(CREATE_LISTENER_SLB_URL, data)
	ret := ChangeNameResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
