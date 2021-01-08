package main

import (
	"fmt"
	"os"

	didiyun "didiyun-go"
	. "didiyun-go/schema"
	slb "didiyun-go/slb"
)

func PrettyPrint(data *[]SlbResponse) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tI$ : %s\tRegion: %s\n", i+1, e.SlbUuid, e.Name, e.Ip, e.Region.Name)
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
