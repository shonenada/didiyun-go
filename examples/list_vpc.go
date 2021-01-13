package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/schema"
	"github.com/shonenada/didiyun-go/vpc"
)

func PrettyPrint(data *[]schema.Vpc) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tRegion: %s\tDefault?: %t\tCidr: %s\n", i+1, e.Uuid, e.Name, e.Region.Name, e.IsDefault, e.CIDR)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	request := vpc.ListRequest{
		RegionId: "gz",
	}
	if r, e := client.Vpc().List(&request); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrint(r)
	}
}
