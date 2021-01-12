package tag

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	. "github.com/shonenada/didiyun-go/schema"
)

type UpdateResourceTagsInput struct {
	ResourceUuid string `json:"resourceUuid"`
	ResourceType string `json:"resourceType"`
}

type UpdateTagRequest struct {
	Resource []UpdateResourceTagsInput `json:"resource"`
	Tags     []string                  `json:"tags"`
}

type UpdateTagResponse struct {
	Errno  int      `json:"errno"`
	Errmsg string   `json:"errmsg"`
	Data   []string `json:"data"`
}

func (c *Client) UpdateTag(request *UpdateTagRequest) (*[]string, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.UPDATE_TAG_URL, data)
	ret := UpdateTagResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request: %s", ret.Errmsg)
	}
	return &ret.Data, nil
}
