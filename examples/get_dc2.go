package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/dc2"
	"github.com/shonenada/didiyun-go/schema"
)

func PrettyPrintDc2(data *schema.Dc2) {
	fmt.Printf("- Uuid: %s\tName: %s\tIP: %s\tEip: %s\tEbs: %v", data.Uuid, data.Name, data.Ip, data.Eip.Ip, data.Ebs)
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	if len(os.Args) < 2 {
		log.Fatal("uuid is not defined")
	}

	uuid := os.Args[1]

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := dc2.GetRequest{}
	req.RegionId = "gz"
	req.Uuid = uuid

	if r, e := client.Dc2().Get(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrintDc2(r)
	}
}
