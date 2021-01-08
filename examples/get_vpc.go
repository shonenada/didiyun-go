package main

import (
	"fmt"
	"log"
	"os"

	didiyun "didiyun-go"
	. "didiyun-go/schema"
	vpc "didiyun-go/vpc"
)

func PrettyPrint(data *VpcInfo) {
	fmt.Printf("- Uuid: %s\tName: %s\tIsDefault: %l\tCidr: %s\tRegion: %s\n", data.Uuid, data.Name, data.IsDefault, data.CIDR, data.Region.Name)
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	if len(os.Args) < 2 {
		log.Fatal("vpcUuid is not defined")
	}

	vpcUuid := os.Args[1]

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := vpc.GetRequest{}
	req.RegionId = "gz"
	req.VpcUuid = vpcUuid

	if r, e := client.Vpc().Get(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrint(r)
	}
}
