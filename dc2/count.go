package dc2

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type CountRequest struct {
	RegionId  string   `json:"regionId"`
	ZoneId    string   `json:"zoneId"`
	VpcUuids  []string `json:"vpcUuids"`
	Dc2Name   string   `json:"dc2Name"`
	SgUuid    string   `json:"sgUuid"`
	Dc2Uuids  []string `json:"dc2Uuids"`
	SgExclude bool     `json:"sgExclude"`
	Ip        string   `json:"ip"`
	Eip       string   `json:"eip"`
}

type Dc2Count struct {
	totalCount int `json:"totalCnt"`
}

type CountResponse struct {
	Errno     int        `json:"errno"`
	Errmsg    string     `json:"errmsg"`
	RequestId string     `json:"requestId"`
	Data      []Dc2Count `json:"data"`
}

func (c *Client) Count(request *CountRequest) (int, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(COUNT_DC2_URL, data)
	ret := CountResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return 0, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return ret.Data[0].totalCount, nil
}
