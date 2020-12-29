package slb

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type CreateListenerMemberRequest struct {
	PoolUuid string         `json:"poolUuid"`
	Members  []SimpleMember `json:"members"`
}

type CreateListenerMemberResponse struct {
	Errno     int    `json:"errno"`
	ErrMsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
}

func (c *Client) Delete(request *CreateListenerMemberRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(CREATE_LISTENER_MEMBER_SLB_URL, data)
	ret := CreateListenerMemberResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
