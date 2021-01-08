package slb

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type CountRequest struct {
	VpcUuids []string `json:"vpcUuids"`
}

type CountResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []struct {
		TotalCnt int `json:"totalCnt"`
	} `json:"data"`
}

func (c *Client) Count(request *CountRequest) (int, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return -1, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(COUNT_SLB_URL, data)
	ret := CountResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return -1, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return ret.Data[0].TotalCnt, nil
}
