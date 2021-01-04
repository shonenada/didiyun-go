package tag

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type DeleteResourceTagsInput struct {
	ResourceUuid string `json:"resourceUuid"`
	ResourceType string `json:"resourceType"`
}

type DeleteTagRequest struct {
	Resource []DeleteResourceTagsInput `json:"resource"`
	Tags     []string                  `json:"tags"`
}

type DeleteTagResponse struct {
	Errno  int      `json:"errno"`
	Errmsg string   `json:"errmsg"`
	Data   []string `json:"data"`
}

func (c *Client) DeleteTag(request *DeleteTagRequest) (*[]string, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(DELETE_TAG_URL, data)
	ret := DeleteTagResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request: %s", ret.Errmsg)
	}
	return &ret.Data, nil
}
