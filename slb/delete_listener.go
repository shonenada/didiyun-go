package slb

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type DeleteListenerRequest struct {
	SlbListener []struct {
		Uuid string `json:"slbListenerUuid"`
	} `json:"slbListener"`
}

type DeleteListenerResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Job `json:"data"`
}

func (c *Client) DeleteListener(request *DeleteListenerRequest) (*schema.Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.DELETE_LISTENER_SLB_URL, data)
	ret := DeleteListenerResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
