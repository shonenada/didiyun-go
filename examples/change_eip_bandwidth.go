package main

import (
	"fmt"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
	eip "github.com/shonenada/didiyun-go/eip"
	. "github.com/shonenada/didiyun-go/schema"
)

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	client := &didiyun.Client{
		AccessToken: accessToken,
	}

	targetUuid := ""
	targetBandWidth := 10

	request := eip.ChangeBandWidthRequest{
		RegionId: "gz",
		Eip: []eip.ChangeBandWidthInput{{
			EipUuid:        targetUuid,
			BandWidth:      targetBandWidth,
			ChargeWithFlow: false,
		}},
	}

	if r, e := client.Eip().ChangeBandWidth(&request); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Job UUID: %s\n", r.Uuid)
	}

}
