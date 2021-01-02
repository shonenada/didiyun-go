package tag

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type AddResourceTagsInput struct {
	ResourceUuid string `json:"resourceUuid"`
	ResourceType string `json:"resourceType"`
}

type AddTagRequest struct {
	Resource []AddResourceTagsInput `json:"resource"`
	Tags     []string               `json:"tags"`
}

type AddTagResponse struct {
	Errno  int      `json:"errno"`
	Errmsg string   `json:"errmsg"`
	Data   []string `json:"data"`
}

func (c *Client) AddTag(request *AddTagRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(ADD_TAG_URL, data)
	ret := AddTagResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request: %s", ret.Errmsg)
	}
	return &ret.Data[0], nil
}
