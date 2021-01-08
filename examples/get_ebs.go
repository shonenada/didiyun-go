package main

import (
	"fmt"
	"log"
	"os"

	didiyun "didiyun-go"
	ebs "didiyun-go/ebs"
	. "didiyun-go/schema"
)

func PrettyPrint(data *EbsInfo) {
	fmt.Printf("- Uuid: %s\tName: %s\tAttr: %s\n", data.EbsUuid, data.Name, data.Attr)
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	if len(os.Args) < 2 {
		log.Fatal("ebsUuid is not defined")
	}

	ebsUuid := os.Args[1]

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := ebs.GetRequest{}
	req.RegionId = "gz"
	req.EbsUuid = ebsUuid

	if r, e := client.Ebs().Get(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrint(r)
	}
}
