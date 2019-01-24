package main

import (
	"fmt"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
	. "github.com/shonenada/didiyun-go/schema"
	sg "github.com/shonenada/didiyun-go/sg"
)

func PrettyPrint(data *[]SgInfo) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tRegion : %s\tVpc: %s\n", i+1, e.SgUuid, e.Name, e.Region.Name, e.Vpc.Name)
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
