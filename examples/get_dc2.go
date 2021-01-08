package main

import (
	"fmt"
	"log"
	"os"

	didiyun "didiyun-go"
	dc2 "didiyun-go/dc2"
	. "didiyun-go/schema"
)

func PrettyPrintDc2(data *Dc2Info) {
	fmt.Printf("- Uuid: %s\tName: %s\tIP: %s\tEip: %s\tEbs: %v", data.Uuid, data.Name, data.Ip, data.Eip.Ip, data.Ebs)
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
