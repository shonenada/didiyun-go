package dc2

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type GetRequest struct {
	RegionId string `json:"regionId"`         // 地域 ID
	ZoneId   string `json:"zoneId,omitempty"` // 可用区 ID
	Uuid     string `json:"dc2Uuid"`          // 所查询的 DC2 的 UUID
}

type GetResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Dc2 `json:"data"`
}

func (c *Client) Get(request *GetRequest) (*schema.Dc2, error) {
	data := map[string]string{
		"regionId": request.RegionId,
		"zoneId":   request.ZoneId,
		"dc2Uuid":  request.Uuid,
	}
	body, err := c.HTTPGet(api.GET_DC2_URL, data)
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
