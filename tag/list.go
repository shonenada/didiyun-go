package tag

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
)

type ListTagRequest struct {
	Start     int `json:"start"`
	Limit     int `json:"limit"`
	Condition struct {
		Name string `json:"name"`
	} `json:"condition"`
}

type ListTagResponse struct {
	Errno  int      `json:"errno"`
	Errmsg string   `json:"errmsg"`
	Data   []string `json:"data"`
}

func (c *Client) ListTag(request *ListTagRequest) (*[]string, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.LIST_TAG_URL, data)
	ret := ListTagResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request: %s", ret.Errmsg)
	}
	return &ret.Data, nil
}
