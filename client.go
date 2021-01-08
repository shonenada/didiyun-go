package didiyun

import (
	"didiyun-go/dc2"
	"didiyun-go/ebs"
	"didiyun-go/eip"
	"didiyun-go/job"
	"didiyun-go/monitor"
	"didiyun-go/region"
	"didiyun-go/sg"
	"didiyun-go/slb"
	"didiyun-go/snap"
	"didiyun-go/sshkey"
	"didiyun-go/tag"
	"didiyun-go/vpc"
)

type Client struct {
	AccessToken string
}

func (c *Client) Dc2() *dc2.Client {
	return &dc2.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) Ebs() *ebs.Client {
	return &ebs.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) Eip() *eip.Client {
	return &eip.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) Job() *job.Client {
	return &job.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) Monitor() *monitor.Client {
	return &monitor.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) Region() *region.Client {
	return &region.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) Sg() *sg.Client {
	return &sg.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) Slb() *slb.Client {
	return &slb.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) Snap() *snap.Client {
	return &snap.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) SSHKey() *sshkey.Client {
	return &sshkey.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) Tag() *tag.Client {
	return &tag.Client{
		AccessToken: c.AccessToken,
	}
}

func (c *Client) Vpc() *vpc.Client {
	return &vpc.Client{
		AccessToken: c.AccessToken,
	}
}
