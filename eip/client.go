package eip

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	AccessToken string
}

func (c *Client) HTTPGet(URL string, params map[string]string) ([]byte, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", URL, nil)
	qs := req.URL.Query()
	for k, v := range params {
		qs.Add(k, v)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	req.URL.RawQuery = qs.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}

	return body, nil
}

func (c *Client) HTTPPost(URL string, data []byte) ([]byte, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error: %s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

	return body, nil
}
