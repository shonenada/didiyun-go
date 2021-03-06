package ebs

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
)

type CountRequest struct {
	RegionId string   `json:"regionId"`
	ZoneId   string   `json:"zoneId,omitempty"`
	Uuids    []string `json:"dc2Uuids,omitempty"`
}

type CountResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []struct {
		TotalCount int `json:"totalCnt"`
	} `json:"data"`
}

func (c *Client) Count(request *CountRequest) (int, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.COUNT_EBS_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := CountResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return -1, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return ret.Data[0].TotalCount, nil
}
