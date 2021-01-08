package main

import (
	"fmt"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
	snap "github.com/shonenada/didiyun-go/snap"
)

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := snap.CountRequest{
		RegionId: "gz",
	}

	if r, e := client.Snap().Count(&req); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Total: %d\n", r)
	}
}
