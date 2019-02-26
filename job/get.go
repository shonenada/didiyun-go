package job

import (
	"encoding/json"
	"fmt"

	. "github.com/shonenada/didiyun-go/schema"
)

type ResultRequest struct {
	RegionId string `json:"regionId"`
	JobUuids string `json:"jobUuids"`
}

type ResultResponse struct {
	Errno     int    `json:"errno"`
	Errmsg    string `json:"errmsg"`
	RequestId string `json:"requestId"`
	Data      []Job  `json:"data"`
}

func (c *Client) GetResult(request *ResultRequest) (*[]Job, error) {
	data := map[string]string{
		"regionId": request.RegionId,
		"jobUuids": request.JobUuids,
	}
	body, err := c.HTTPGet(GET_RESULT_JOB_URL, data)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}
	ret := ResultResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
