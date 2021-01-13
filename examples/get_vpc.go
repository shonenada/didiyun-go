package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/schema"
	"github.com/shonenada/didiyun-go/vpc"
)

func PrettyPrint(data *schema.Vpc) {
	fmt.Printf("- Uuid: %s\tName: %s\tIsDefault: %l\tCidr: %s\tRegion: %s\n", data.Uuid, data.Name, data.IsDefault, data.CIDR, data.Region.Name)
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	if len(os.Args) < 2 {
		log.Fatal("vpcUuid is not defined")
	}

	uuid := os.Args[1]

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := vpc.GetRequest{}
	req.RegionId = "gz"
	req.Uuid = uuid

	if r, e := client.Vpc().Get(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrint(r)
	}
}
