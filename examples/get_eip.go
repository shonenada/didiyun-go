package main

import (
	"fmt"
	"log"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
	eip "github.com/shonenada/didiyun-go/eip"
	. "github.com/shonenada/didiyun-go/schema"
)

func PrettyPrintEip(data *EipInfo) {
	fmt.Printf("Uuid: %s\tIP: %s\tDc2: %s\n", data.Uuid, data.Ip, data.Dc2.Name)
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	if len(os.Args) < 2 {
		log.Fatal("eipUuid is not defined")
	}

	eipUuid := os.Args[1]

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := eip.GetRequest{}
	req.RegionId = "gz"
	req.EipUuid = eipUuid

	if r, e := client.Eip().Get(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrintEip(r)
	}
}
