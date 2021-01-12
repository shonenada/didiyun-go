package dc2

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	. "github.com/shonenada/didiyun-go/schema"
)

type RebootInput struct {
	Dc2Uuid string `json:"dc2Uuid"`
}

type RebootRequest struct {
	RegionId string      `json:"regionId"`
	ZoneId   string      `json:"zoneId,omitempty"`
	Dc2      RebootInput `json:"dc2"`
}

type RebootResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) Reboot(request *RebootRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.REBOOT_DC2_URL, data)
	ret := RebootResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
