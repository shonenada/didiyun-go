package dc2

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type GetRequest struct {
	RegionId string `json:"regionId"`
	ZoneId   string `json:"zoneId,omitempty"`
	Uuid     string `json:"dc2Uuid"`
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
