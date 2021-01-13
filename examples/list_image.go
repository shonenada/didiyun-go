package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/dc2"
	"github.com/shonenada/didiyun-go/schema"
)

func PrettyPrint(data *[]schema.Image) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tOS: %s(%s) - %s\tPaltform: %s\tScenes: %s\tType: %s\n", i+1, e.Uuid, e.Name, e.OsFamily, e.OsArch, e.OsVersion, e.Platform, e.Scenes, e.ImgType)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	request := dc2.ListImageRequest{
		RegionId: "gz",
	}

	if r, e := client.Dc2().ListImage(&request); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrint(r)
	}

}
