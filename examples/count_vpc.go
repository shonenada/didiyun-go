package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/vpc"
)

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := vpc.CountRequest{
		RegionId: "gz",
	}

	if r, e := client.Vpc().Count(&req); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Total VPC: %d\n", r)
	}
}
