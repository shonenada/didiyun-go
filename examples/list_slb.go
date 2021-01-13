package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/schema"
	"github.com/shonenada/didiyun-go/slb"
)

func PrettyPrint(data *[]schema.Slb) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tI$ : %s\tRegion: %s\n", i+1, e.Uuid, e.Name, e.Ip, e.Region.Name)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	request := slb.ListRequest{
		RegionId: "gz",
	}

	if r, e := client.Slb().List(&request); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrint(r)
	}

}
