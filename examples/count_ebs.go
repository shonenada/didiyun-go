package main

import (
	"fmt"
	"os"

	didiyun "didiyun-go"
	ebs "didiyun-go/ebs"
)

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := ebs.CountRequest{
		RegionId: "gz",
	}

	if r, e := client.Ebs().Count(&req); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Total: %d\n", r)
	}
}
