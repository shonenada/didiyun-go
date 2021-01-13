package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/dc2"
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
		fmt.Printf("Total DC2 count: %d\n", r)
	}
}
