package dc2

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type GetDc2Request struct {
	// 地域 ID
	RegionId string `json:"regionId"`

	// 可用区 ID
	ZoneId string `json:"zoneId"`

	// 所查询的 DC2 的 UUID
	Dc2Uuid string `json:"dc2Uuid"`
}

type GetDc2Response struct {
	Errno     int       `json:"errno"`
	Errmsg    string    `json:"errmsg"`
	RequestId string    `json:"requestId"`
	Data      []Dc2Info `json:"data"`
}

func (c *Client) GetDc2(req *GetDc2Request) ([]Dc2Info, error) {
	data := map[string]string{
		"regionId": req.RegionId,
		"zoneId":   req.ZoneId,
		"dc2Uuid":  req.Dc2Uuid,
	}
	body, nil := c.HTTPGet(GET_DC2_URL, data)
	ret := GetDc2Response{}
	json.Unmarshal([]byte(body), &ret)
	if ret.Errno != 0 {
		fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return ret.Data, nil
}
