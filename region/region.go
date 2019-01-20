package region

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	AccessToken string
}

type ListRegionAndZoneCondition struct {
	Product string `json:"product"`
}

type ListRegionRequest struct {
	Condition ListRegionAndZoneCondition `json:"condition"`
}

type Region struct {
	AreaName string `json:"areaName"`
	Id       string `json:"id"`
	name     string `json:"name"`
	zone     string `json:"zone"`
}

type ListRegionResponse struct {
	Errno     int      `json:"errno"`
	Errmsg    string   `json:"errmsg"`
	RequestId string   `json:"requestId"`
	Data      []Region `json:"data"`
}

func (c *Client) ListRegion(product string) ListRegionResponse {
	httpClient := &http.Client{}
	reqBody := &ListRegionRequest{
		Condition: ListRegionAndZoneCondition{
			Product: product,
		},
	}

	reqStr, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Errorf("Fail to marshal body: %s", err)
	}

	req, err := http.NewRequest("POST", "https://open.didiyunapi.com/dicloud/i/region/list", bytes.NewBuffer([]byte(reqStr)))
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
	return ret
}

func (c *Client) ListDC2Regions() ListRegionResponse {
	return c.ListRegion("dc2")
}

func (c *Client) ListEBSRegions() ListRegionResponse {
	return c.ListRegion("ebs")
}

func (c *Client) ListEIPRegions() ListRegionResponse {
	return c.ListRegion("eip")
}

func (c *Client) ListSGRegions() ListRegionResponse {
	return c.ListRegion("sg")
}

func (c *Client) ListSNAPRegions() ListRegionResponse {
	return c.ListRegion("snap")
}
func (c *Client) ListVPCRegions() ListRegionResponse {
	return c.ListRegion("vpc")
}
