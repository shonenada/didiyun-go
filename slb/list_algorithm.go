package slb

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type ListAlgorithmResponse struct {
	Errno     int         `json:"errno"`
	Errmsg    string      `json:"errmsg"`
	RequestId string      `json:"RequestId"`
	Data      []Algorithm `json:"data"`
}

type Algorithm struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (c *Client) ListAlgorithm() (*[]Algorithm, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPGet(LIST_ALGORITHM_SLB_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := ListAlgorithmResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
