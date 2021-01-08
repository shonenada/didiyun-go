package main

import (
	"fmt"
	"log"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
	ebs "github.com/shonenada/didiyun-go/ebs"
)

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	if len(os.Args) < 3 {
		log.Fatal("Region or Zone is not defined")
	}

	region := os.Args[1]
	zone := os.Args[2]

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	request := ebs.CreateRequest{
		RegionId:     region,
		ZoneId:       zone,
		AutoContinue: false,
		Count:        1,
		Size:         20,
		DiskType:     "HDD",
		EbsTags:      []string{"Create By Didiyun-Go"},
	}
	if r, e := client.Ebs().Create(&request); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Job UUID: %s\n", r.Uuid)
	}
}
