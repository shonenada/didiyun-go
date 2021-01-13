package sshkey

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type CreateRequest struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type CreateResponse struct {
	Errno     int             `json:"errno"`
	Errmsg    string          `json:"errmsg"`
	RequestId string          `json:"requestId"`
	Data      []schema.SSHKey `json:"data"`
}

func (c *Client) Create(request *CreateRequest) (*schema.SSHKey, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.CREATE_SSHKEY_URL, data)
	ret := CreateResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
