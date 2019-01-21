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

	regionClient := client.Region()

	if r, e := regionClient.ListDc2Regions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("DC2 Regions: ", r)
	}

	if r, e := regionClient.ListEbsRegions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("EBS Regions: ", r)
	}

	if r, e := regionClient.ListEipRegions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("EBS Regions: ", r)
	}

	if r, e := regionClient.ListSgRegions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("SG Regions: ", r)
	}

	if r, e := regionClient.ListSnapRegions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("SNAP Regions: ", r)
	}

	if r, e := regionClient.ListVpcRegions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("VPC Regions: ", r)
	}
}
