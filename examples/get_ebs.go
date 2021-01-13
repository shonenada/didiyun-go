package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/ebs"
	"github.com/shonenada/didiyun-go/schema"
)

func PrettyPrint(data *schema.Ebs) {
	fmt.Printf("- Uuid: %s\tName: %s\tAttr: %s\n", data.Uuid, data.Name, data.Attr)
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
	req := ebs.GetRequest{}
	req.RegionId = "gz"
	req.EbsUuid = uuid

	if r, e := client.Ebs().Get(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrint(r)
	}
}
