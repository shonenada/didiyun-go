package main

import (
	"fmt"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
	dc2 "github.com/shonenada/didiyun-go/dc2"
)

func PrettyPrintDc2(data *[]dc2.Dc2Response) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tIP: %s\tEip: %s\nRegion: %s\n", i+1, e.Uuid, e.Name, e.Ip, e.Eip.Ip, e.Region.Name)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	builder := dc2.ListDc2RequestBuilder{}
	builder.SetRegionId("gz")
	req := builder.Build()

	if r, e := client.Dc2().ListDc2(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrintDc2(r)
	}

}
