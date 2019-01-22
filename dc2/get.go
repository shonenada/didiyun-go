package dc2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

func (c *Client) GetDc2(reqBody *GetDc2Request) (*[]Dc2Info, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", GET_DC2_URL, nil)

	qs := req.URL.Query()
	qs.Add("regionId", reqBody.RegionId)
	qs.Add("zoneId", reqBody.ZoneId)
	qs.Add("dc2Uuid", reqBody.Dc2Uuid)
	req.URL.RawQuery = qs.Encode()
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := GetDc2Response{}
	json.Unmarshal([]byte(body), &ret)

	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}

	return &ret.Data, nil

}
