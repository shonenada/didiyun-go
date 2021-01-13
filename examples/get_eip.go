package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/eip"
	"github.com/shonenada/didiyun-go/schema"
)

func PrettyPrintEip(data *schema.Eip) {
	fmt.Printf("Uuid: %s\tIP: %s\tDc2: %s\n", data.Uuid, data.Ip, data.Dc2.Name)
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
	req := eip.GetRequest{}
	req.RegionId = "gz"
	req.EipUuid = uuid

	if r, e := client.Eip().Get(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrintEip(r)
	}
}
