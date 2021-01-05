# DidiYun Golang SDK

**UNOFFICIAL** Go Client for Didiyun

## Getting Started

### Installation

#### Requirement

This SDK requires Go 1.6 or higher versions.

#### Install from code

Use `go get` to retrieve thd SDK, and add into your `GOPATH`:

```sh
$ go get -u github.com/shonenada/didiyun-go
```

## Usage

`didiyun-go` require access token for invoking HTTP API, so you need to genearte on in [API Token](https://app.didiyun.com/#/api/authtoken) before developing with `didiyun-go`.

### List DC2

```go
package main

import (
	"fmt"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
	dc2 "github.com/shonenada/didiyun-go/dc2"
	. "github.com/shonenada/didiyun-go/schema"
)

func PrettyPrintDc2(data *[]Dc2Info) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tIP: %s\tEip: %s\tRegion: %s\n", i+1, e.Uuid, e.Name, e.Ip, e.Eip.Ip, e.Region.Name)
	}
}

func main() {
	accessToken := "<API_TOKEN>"
	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	builder := dc2.ListRequestBuilder{}
	builder.SetRegionId("gz")
	req := builder.Build()

	if r, e := client.Dc2().List(&req); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrintDc2(r)
	}

}
```

### List Ebs

```go
package main

import (
	"fmt"
	"os"

	didiyun "github.com/shonenada/didiyun-go"
	ebs "github.com/shonenada/didiyun-go/ebs"
)

func PrettyPrint(data *[]ebs.EbsInfo) {
	for i, e := range *data {
		fmt.Printf("[%d] - Uuid: %s\tName: %s\tAttr: %s\tDc2: %s\n", i+1, e.EbsUuid, e.Name, e.Attr, e.Dc2.Name)
	}
}

func main() {
	accessToken := os.Getenv("DIDIYUN_ACCESS_TOKEN")
	client := &didiyun.Client{
		AccessToken: accessToken,
	}
	request := ebs.ListRequest{
		RegionId: "gz",
	}

	if r, e := client.Ebs().List(&request); e != nil {
		fmt.Println(e)
	} else {
		PrettyPrint(r)
	}

}
```

For more examples please visit [examples/](examples).

## Contributing

TBD

## LICENSE

The Apache License (Version 2.0, January 2004).
