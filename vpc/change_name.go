package vpc

import (
	"encoding/json"
	"fmt"

	. "didiyun-go/schema"
)

type ChangeNameRequest struct {
	RegionId string            `json:"regionId"`
	Vpc      []ChangeNameInput `json:"vpc"`
}

type ChangeNameInput struct {
	VpcUuid string `json:"vpcUuid"`
	Name    string `json:"name"`
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
	body, err := c.HTTPPost(CHANGE_NAME_VPC_URL, data)
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
