package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/eip"
)

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := eip.CountRequest{
		RegionId: "gz",
	}

	if r, e := client.Eip().Count(&req); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Total EIP count: %d\n", r)
	}
}
