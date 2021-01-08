package dc2

import (
	"encoding/json"
	"fmt"

	. "didiyun-go/schema"
)

type ChangePasswordInput struct {
	Dc2Uuid  string `json:"dc2Uuid"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	RegionId string                `json:"regionId"`
	ZoneId   string                `json:"zoneId,omitempty"`
	Dc2      []ChangePasswordInput `json:"dc2"`
}

type ChangePasswordResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) ChangePassword(request *ChangePasswordRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(CHANGE_PASSWORD_DC2_URL, data)
	ret := ChangePasswordResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
