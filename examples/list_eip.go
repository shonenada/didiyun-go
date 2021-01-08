package main

import (
	"fmt"
	"os"

	didiyun "didiyun-go"
	eip "didiyun-go/eip"
	. "didiyun-go/schema"
)

func PrettyPrintEip(data *[]EipInfo) {
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
