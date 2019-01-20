package didiyun

import (
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
