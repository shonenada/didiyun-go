package sg

import (
	"encoding/json"
	"fmt"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

func (c *Client) GetByUuid(regionId string, uuid string) (*schema.Sg, error) {
	request := ListRequest{
		RegionId: regionId,
		Condition: SgCondition{
			SgUuids: []string{uuid},
		},
	}

	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal body: %s", err)
	}
	body, err := c.HTTPPost(api.LIST_SG_URL, data)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	ret := ListResponse{}
	json.Unmarshal(body, &ret)
	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}
	if len(ret.Data) == 0 {
		return nil, fmt.Errorf("Sg not found")
	}

	return &ret.Data[0], nil
}
