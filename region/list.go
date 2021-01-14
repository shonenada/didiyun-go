package region

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shonenada/didiyun-go/api"
	"github.com/shonenada/didiyun-go/schema"
)

type ListRegionAndZoneCondition struct {
	Product string `json:"product"`
}

type ListRegionRequest struct {
	Condition ListRegionAndZoneCondition `json:"condition"`
}

type ListRegionResponse struct {
	Errno     int                   `json:"errno"`
	Errmsg    string                `json:"errmsg"`
	RequestId string                `json:"requestId"`
	Data      []schema.RegionDetail `json:"data"`
}

func (c *Client) ListRegion(product string) (*[]schema.RegionDetail, error) {
	httpClient := &http.Client{}
	reqBody := &ListRegionRequest{
		Condition: ListRegionAndZoneCondition{
			Product: product,
		},
	}

	reqStr, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Errorf("Failed to marshal body: %s", err)
	}

	req, err := http.NewRequest("POST", api.LIST_REGION_URL, bytes.NewBuffer([]byte(reqStr)))
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
	ret := ListRegionResponse{}
	json.Unmarshal([]byte(body), &ret)

	if ret.Errno != 0 {
		return nil, fmt.Errorf("Failed to request [%s]: %s", ret.RequestId, ret.Errmsg)
	}

	return &ret.Data, nil
}

func (c *Client) ListDc2Regions() (*[]schema.RegionDetail, error) {
	return c.ListRegion("dc2")
}

func (c *Client) ListEbsRegions() (*[]schema.RegionDetail, error) {
	return c.ListRegion("ebs")
}

func (c *Client) ListEipRegions() (*[]schema.RegionDetail, error) {
	return c.ListRegion("eip")
}

func (c *Client) ListSgRegions() (*[]schema.RegionDetail, error) {
	return c.ListRegion("sg")
}

func (c *Client) ListSnapRegions() (*[]schema.RegionDetail, error) {
	return c.ListRegion("snap")
}
func (c *Client) ListVpcRegions() (*[]schema.RegionDetail, error) {
	return c.ListRegion("vpc")
}
