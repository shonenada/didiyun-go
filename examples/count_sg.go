package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/sg"
)

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := sg.CountRequest{
		RegionId: "gz",
	}

	if r, e := client.Sg().Count(&req); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Total SG count: %d\n", r)
	}
}
