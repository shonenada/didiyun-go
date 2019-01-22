package dc2

type DeleteDc2Request struct {
	RegionId string           `json:"regionId"`
	ZoneId   string           `json:"zoneId"`
	Dc2      []DeleteDc2Input `json:"dc2"` // 要删除的DC2信息，一次不能超过20台
}

type DeleteDc2Input struct {
	Dc2Uuid   string `json:"dc2Uuid"`   // 需要删除规格的DC2的uuid
	DeleteEip bool   `json:"deleteEip"` // 是否同时删除DC2上绑定的EIP
	DeleteEbs bool   `json:"deleteEbs"` // 是否同时删除DC2上绑定的EBS
	IgnoreSlb bool   `json:"ignoreSlb"` // 是否忽略DC2上绑定的负载均衡
}

type DeleteDc2Response struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) DeleteDc2(reqBody *DeleteDc2Request) (*Job, error) {
	httpClient := &http.Client{}
	reqStr, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}

	req, err := http.NewRequest("POST", DELETE_DC2_URL, bytes.NewBuffer([]byte(reqStr)))
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

	ret := DeteleDc2Response{}
	json.Unmarshal([]byte(body), &ret)

	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}

	return &ret.Data[0], nil
}
