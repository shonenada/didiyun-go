package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/dc2"
	"github.com/shonenada/didiyun-go/schema"
)

func PrettyPrintDc2(data *[]schema.Dc2) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tIP: %s\tEip: %s\tRegion: %s\tZone:%s \n", i+1, e.Uuid, e.Name, e.Ip, e.Eip.Ip, e.Region.Name, e.Region.Zone.Name)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	builder := dc2.ListRequestBuilder{}
	builder.SetRegionId("gz")
	builder.SetLimit(50)
	req := builder.Build()

	if r, e := client.Dc2().List(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrintDc2(r)
	}

}
