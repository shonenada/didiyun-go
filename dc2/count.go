package dc2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/shonenada/didiyun-go/schema"
)

type CountDc2Request struct {
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

type CountDc2Response struct {
	Errno     int        `json:"errno"`
	Errmsg    string     `json:"errmsg"`
	RequestId string     `json:"requestId"`
	Data      []Dc2Count `json:"data"`
}

func (c *Client) CountDc2(reqBody *CountDc2Request) (int, error) {
	httpClient := &http.Client{}
	reqStr, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}

	req, err := http.NewRequest("POST", COUNT_DC2_URL, bytes.NewBuffer([]byte(reqStr)))
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	req.Header.Add("Content-Type", "application/json")
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
	ret := CountDc2Response{}
	json.Unmarshal([]byte(body), &ret)

	if ret.Errno != 0 {
		return 0, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}

	return ret.Data[0].totalCount, nil
}
