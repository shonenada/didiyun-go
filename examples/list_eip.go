package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/eip"
	"github.com/shonenada/didiyun-go/schema"
)

func PrettyPrintEip(data *[]schema.Eip) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tIP: %s\t\n", i+1, e.Uuid, e.Ip)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	builder := eip.ListRequestBuilder{}
	builder.SetRegionId("gz")
	req := builder.Build()
	if r, e := client.Eip().List(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrintEip(r)
	}

}
