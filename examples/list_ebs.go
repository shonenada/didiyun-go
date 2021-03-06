package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/ebs"
	"github.com/shonenada/didiyun-go/schema"
)

func PrettyPrint(data *[]schema.Ebs) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tAttr: %s\tDc2: %s\n", i+1, e.Uuid, e.Name, e.Attr, e.Dc2.Name)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	request := ebs.ListRequest{
		RegionId: "gz",
	}

	if r, e := client.Ebs().List(&request); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrint(r)
	}

}
