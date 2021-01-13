package dc2

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
)

type CountRequest struct {
	RegionId    string   `json:"regionId"`
	ZoneId      string   `json:"zoneId,omitempty"`
	Name        string   `json:"dc2Name,omitempty"`
	SgUuid      string   `json:"sgUuid,omitempty"`
	IsSgExclude bool     `json:"sgExclude,omitempty"`
	Ip          string   `json:"ip,omitempty"`
	Eip         string   `json:"eip,omitempty"`
	Uuids       []string `json:"dc2Uuids,omitempty"`
	VpcUuids    []string `json:"vpcUuids,omitempty"`
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
		return -1, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.COUNT_DC2_URL, data)
	if err != nil {
		return -1, fmt.Errorf("Error: %s", err)
	}
	ret := CountResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return -1, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return ret.Data[0].TotalCount, nil
}
