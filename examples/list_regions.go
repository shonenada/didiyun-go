package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/schema"
)

func PretyPrint(data *[]schema.Region) {
	for i, e := range *data {
		fmt.Printf("- [%d] AreaName: %s\tId: %s\n", i+1, e.AreaName, e.Id)
		for j, z := range e.Zone {
			fmt.Printf("    Zone[%d]: %s (%s)\n", j, z.Name, z.Id)
		}
		fmt.Println("")
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}

	regionClient := client.Region()

	if r, e := regionClient.ListDc2Regions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("DC2 Regions: ")
		PretyPrint(r)
	}

	if r, e := regionClient.ListEbsRegions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("EBS Regions: ")
		PretyPrint(r)
	}

	if r, e := regionClient.ListEipRegions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("EIP Regions: ")
		PretyPrint(r)
	}

	if r, e := regionClient.ListSgRegions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("SG Regions: ")
		PretyPrint(r)
	}

	if r, e := regionClient.ListSnapRegions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("SNAP Regions: ")
		PretyPrint(r)
	}

	if r, e := regionClient.ListVpcRegions(); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("VPC Regions: ")
		PretyPrint(r)
	}
}
