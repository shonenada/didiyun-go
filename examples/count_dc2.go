package main

import (
	"fmt"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
	dc2 "github.com/shonenada/didiyun-go/dc2"
)

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := dc2.CountRequest{
		RegionId: "gz",
	}

	if r, e := client.Dc2().Count(&req); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Total: %d\n", r)
	}
}
