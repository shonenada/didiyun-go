package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/schema"
	"github.com/shonenada/didiyun-go/snapshot"
)

func PrettyPrintEip(data []schema.Snapshot) {
	for i, e := range data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tDc2: %s\n", i+1, e.Uuid, e.Name, e.Dc2.Name)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	builder := snapshot.ListRequestBuilder{}
	builder.SetRegionId("gz")
	req := builder.Build()
	if r, e := client.Snapshot().List(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrintEip(r)
	}

}
