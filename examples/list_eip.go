package main

import (
	"fmt"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
	eip "github.com/shonenada/didiyun-go/eip"
)

func PrettyPrintEip(data *[]eip.EipInfo) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tIP: %s\tDC2: %s\n", i+1, e.Uuid, e.Ip, e.Dc2.Name)
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
