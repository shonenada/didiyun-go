package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/snapshot"
)

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := snapshot.CountRequest{
		RegionId: "gz",
	}

	if r, e := client.Snap().Count(&req); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Total snapshot count: %d\n", r)
	}
}
