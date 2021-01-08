package main

import (
	"fmt"
	"os"

	didiyun "didiyun-go"
	vpc "didiyun-go/vpc"
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
		fmt.Printf("Total: %d\n", r)
	}
}
