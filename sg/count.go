package sg

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type CountRequest struct {
	RegionId   string   `json:"regionId"`
	SgUuids    []string `json:"sgUuids,omitempty"`
	VpcUuid    string   `json:"vpcUuid,omitempty"`
	Dc2Uuid    string   `json:"dc2Uuid,omitempty"`
	Dc2Exclude bool     `json:"dc2Exclude,omitempty"`
}

type CountResponse struct {
	Errno     int       `json:"errno"`
	Errmsg    string    `json:"errmsg"`
	RequestId string    `json:"requestId"`
	Data      []SgCount `json:"data"`
}

type SgCount struct {
	TotalCount int `json:"totalCnt"`
}

func (c *Client) Count(request *CountRequest) (int, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(COUNT_SG_URL, data)
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
