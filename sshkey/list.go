package sshkey

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type ListResponse struct {
	Errno     int                 `json:"errno"`
	Errmsg    string              `json:"errmsg"`
	RequestId string              `json:"requestId"`
	Data      []schema.SSHKeyInfo `json:"data"`
}

func (c *Client) List() (*[]schema.SSHKeyInfo, error) {
	body, err := c.HTTPGet(api.LIST_SSHKEY_URL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ListResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
