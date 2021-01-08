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

func (c *Client) ListAlgorithm() (*[]Algorithm, error) {
	body, err := c.HTTPGet(LIST_ALGORITHM_SLB_URL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ListAlgorithmResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
