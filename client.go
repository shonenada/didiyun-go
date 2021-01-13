package didiyun

import (
	"github.com/shonenada/didiyun-go/dc2"
	"github.com/shonenada/didiyun-go/ebs"
	"github.com/shonenada/didiyun-go/eip"
	"github.com/shonenada/didiyun-go/job"
	"github.com/shonenada/didiyun-go/monitor"
	"github.com/shonenada/didiyun-go/region"
	"github.com/shonenada/didiyun-go/sg"
	"github.com/shonenada/didiyun-go/slb"
	"github.com/shonenada/didiyun-go/snapshot"
	"github.com/shonenada/didiyun-go/sshkey"
	"github.com/shonenada/didiyun-go/tag"
	"github.com/shonenada/didiyun-go/vpc"
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

func (c *Client) Snapshot() *snapshot.Client {
	return &snapshot.Client{
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
