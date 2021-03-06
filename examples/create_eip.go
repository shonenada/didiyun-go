package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/eip"
)

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	if len(os.Args) < 2 {
		log.Fatal("Region is not defined")
	}

	region := os.Args[1]

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	request := eip.CreateRequest{
		RegionId:         region,
		IsAutoContinue:   false,
		Count:            1,
		BandWidth:        1,
		IsChargeWithFlow: true,
		Tags:             []string{"Create By Didiyun-Go"},
	}
	if r, e := client.Eip().Create(&request); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Job UUID: %s\n", r.Uuid)
	}
}
