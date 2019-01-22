package dc2

import (
	. "github.com/shonenada/didiyun-go/schema"
)

type DeleteSSHKeyResponse struct {
	Errno     int                    `json:"errno"`
	Errmsg    string                 `json:"errmsg"`
	RequestId string                 `json:"requestId"`
	Data      []DeleteSSHKeyResponse `json:"data"`
}

type DeleteSSHKeyResponse struct {
	PubKeyUuid string `json:"pubKeyUuid"`
}
