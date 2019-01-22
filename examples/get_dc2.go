package main

import (
	"fmt"
	"log"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
	dc2 "github.com/shonenada/didiyun-go/dc2"
	. "github.com/shonenada/didiyun-go/schema"
)

func PrettyPrintDc2(data []Dc2Info) {
	for i, e := range data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tIP: %s\tEip: %s\nRegion: %s\n", i+1, e.Uuid, e.Name, e.Ip, e.Eip.Ip, e.Region.Name)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	if len(os.Args) < 2 {
		log.Fatal("dc2Uuid is not defined")
	}

	dc2Uuid := os.Args[1]

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := dc2.GetRequest{}
	req.RegionId = "gz"
	req.Dc2Uuid = dc2Uuid

	if r, e := client.Dc2().Get(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrintDc2(r)
	}
}
