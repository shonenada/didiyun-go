package main

import (
	"fmt"
	"os"

	"github.com/shonenada/didiyun-go"
	"github.com/shonenada/didiyun-go/schema"
)

func PrettyPrintSSHKey(data *[]schema.SSHKey) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\n", i+1, e.Uuid, e.Name)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	if r, e := client.SSHKey().List(); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrintSSHKey(r)
	}
}
