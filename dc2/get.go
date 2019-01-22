package dc2

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type GetRequest struct {
	// 地域 ID
	RegionId string `json:"regionId"`

	// 可用区 ID
	ZoneId string `json:"zoneId,omitempty"`

	// 所查询的 DC2 的 UUID
	Dc2Uuid string `json:"dc2Uuid"`
}

type GetResponse struct {
	Errno     int       `json:"errno"`
	Errmsg    string    `json:"errmsg"`
	RequestId string    `json:"requestId"`
	Data      []Dc2Info `json:"data"`
}

func (c *Client) Get(request *GetRequest) (*Dc2Info, error) {
	data := map[string]string{
		"regionId": request.RegionId,
		"zoneId":   request.ZoneId,
		"dc2Uuid":  request.Dc2Uuid,
	}
	body, err := c.HTTPGet(GET_DC2_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := GetResponse{}
	json.Unmarshal([]byte(body), &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data[0], nil
}
