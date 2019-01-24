package sg

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type DeleteRequest struct {
	RegionId string        `json:"regionId"`
	Sg       []DeleteInput `json:"sg"`
}

type DeleteInput struct {
	SgUuid string `json:"sgUuid"`
}

type DeleteResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) Delete(request *DeleteRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(DELETE_SG_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := DeleteResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}