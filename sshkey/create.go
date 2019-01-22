package dc2

import (
	. "github.com/shonenada/didiyun-go/schema"
)

type CreateSSHKeyResponse struct {
	Errno     int          `json:"errno"`
	Errmsg    string       `json:"errmsg"`
	RequestId string       `json:"requestId"`
	Data      []SSHKeyInfo `json:"data"`
}
