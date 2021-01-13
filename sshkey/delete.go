package sshkey

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
)

type DeleteRequest struct {
	Uuid string `json:"pubKeyUuid"`
}

type DeleteResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []struct {
		PubKeyUuid string `json:"pubKeyUuid"`
	} `json:"data"`
}

func (c *Client) Delete(request *DeleteRequest) (string, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.DELETE_SSHKEY_URL, data)
	ret := DeleteResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return "", fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return ret.Data[0].PubKeyUuid, nil
}
