package dc2

import (
	"encoding/json"
	"fmt"

	. "didiyun-go/schema"
)

type StopInput struct {
	Dc2Uuid string `json:"dc2Uuid"`
}

type StopRequest struct {
	RegionId string    `json:"regionId"`
	ZoneId   string    `json:"zoneId,omitempty"`
	Dc2      StopInput `json:"dc2"`
}

type StopResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) Stop(request *StopRequest) (*Job, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(STOP_DC2_URL, data)
	ret := StopResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
