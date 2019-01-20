package main

import (
	"fmt"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
)

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}

	fmt.Println("DC2 Regions: ", client.Region().ListDC2Regions())
	fmt.Println("EBS Regions: ", client.Region().ListEBSRegions())
	fmt.Println("EIP Regions: ", client.Region().ListEIPRegions())
	fmt.Println("SG Regions: ", client.Region().ListSGRegions())
	fmt.Println("SNAP Regions: ", client.Region().ListSNAPRegions())
	fmt.Println("VPC Regions: ", client.Region().ListVPCRegions())
}
