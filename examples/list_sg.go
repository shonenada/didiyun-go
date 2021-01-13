package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/schema"
	"github.com/shonenada/didiyun-go/sg"
)

func PrettyPrint(data *[]schema.Sg) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tRegion : %s\tVpc: %s\n", i+1, e.Uuid, e.Name, e.Region.Name, e.Vpc.Name)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	request := sg.ListRequest{
		RegionId: "gz",
	}

	if r, e := client.Sg().List(&request); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrint(r)
	}

}
