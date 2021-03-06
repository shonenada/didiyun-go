package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/job"
	"github.com/shonenada/didiyun-go/schema"
)

func PrettyPrintJob(data *[]schema.Job) {
	for i, e := range *data {
		fmt.Printf("[%d] Uuid: %s\tResourceUuid: %s\tType: %sProgress: %d\tResult: %s\n", i+1, e.Uuid, e.ResourceUuid, e.Type, e.Progress, e.Result)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")

	if len(os.Args) < 2 {
		log.Fatal("jobUuid is not defined")
	}

	uuid := os.Args[1]

	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	req := job.ResultRequest{}
	req.RegionId = "gz"
	req.Uuids = uuid

	if r, e := client.Job().GetResult(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrintJob(r)
	}
}
