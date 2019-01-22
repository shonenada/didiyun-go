package dc2

import (
	. "github.com/shonenada/didiyun-go/schema"
)

type ListImageRequest struct {
	RegionId string `json:"regionId"`
}

type ListImageResponse struct {
	Errno     int         `json:"errno"`
	Errmsg    string      `json:"errmsg"`
	RequestId string      `json:"requestId"`
	Data      []ImageInfo `json:"data"`
}
