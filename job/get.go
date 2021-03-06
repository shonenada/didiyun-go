package job

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type ResultRequest struct {
	RegionId string `json:"regionId"`
	Uuids    string `json:"jobUuids"`
}

type ResultResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []schema.Job `json:"data"`
}

func (c *Client) GetResult(request *ResultRequest) (*[]schema.Job, error) {
	data := map[string]string{
		"regionId": request.RegionId,
		"jobUuids": request.Uuids,
	}
	body, err := c.HTTPGet(api.GET_RESULT_JOB_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ResultResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	return &ret.Data, nil
}
