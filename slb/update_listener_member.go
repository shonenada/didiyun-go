package slb

import (
	"encoding/json"
	"fmt"

	. "didiyun-go/schema"
)

type UpdateListenerMemberRequest struct {
	Members []ListenerMember `json:"members"`
}

type UpdateListenerMemberResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) UpdateListenerMember(request *UpdateListenerMemberRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(UPDATE_LISTENER_MEMBER_SLB_URL, data)
	ret := UpdateListenerMemberResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
