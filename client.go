package didiyun

import (
	"github.com/shonenada/didiyun-go/dc2"
	"github.com/shonenada/didiyun-go/region"
)

type Client struct {
	AccessToken string
}

func (c *Client) Region() *region.Client {
	return &region.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) Dc2() *dc2.Client {
	return &dc2.Client{
		AccessToken: c.AccessToken,
	}
}
