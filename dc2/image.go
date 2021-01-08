package dc2

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type ListImageRequest struct {
	RegionId string `json:"regionId"`
}

type ListImageResponse struct {
	Errno     int         `json:"errno"`
	Errmsg    string      `json:"errmsg"`
	RequestId string      `json:"requestId"`
	Data      []ImageInfo `json:"data"`
}

func (c *Client) ListImage(request *ListImageRequest) (*[]ImageInfo, error) {
	data := map[string]string{
		"regionId": request.RegionId,
	}
	body, err := c.HTTPGet(LIST_IMAGE_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ListImageResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
