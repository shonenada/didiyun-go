package main

import (
	"fmt"
	"os"

	didiyun "didiyun-go"
	. "didiyun-go/schema"
)

func PrettyPrintSSHKeyInfo(data *[]SSHKeyInfo) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\n", i+1, e.PubKeyUuid, e.Name)
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
		PrettyPrintSSHKeyInfo(r)
	}
}
