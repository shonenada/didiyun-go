package slb

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type ListListenerMemberRequest struct {
	RegionId  string `json:"regionId"`
	Start     int    `json:"start"`
	Limit     int    `json:"limit"`
	Condition struct {
		PoolUuid string `json:"poolUuid"`
	} `json:"condition"`
}

type ListListenerMemberResponse struct {
	Errno     int              `json:"errno"`
	Errmsg    string           `json:"errmsg"`
	RequestId string           `json:"requestId"`
	Data      []ListenerMember `json:"data"`
}

func (c *Client) ListListenerMember(request *ListListenerMemberRequest) (*[]ListenerMember, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(LIST_LISTENER_MEMBER_SLB_URL, data)
	ret := ListListenerMemberResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
