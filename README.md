# DidiYun Golang SDK

**Unofficial** Go Client for Didiyun

## Getting Started

### Installation

#### Requirement

This SDK requires Go 1.6 or higher versions.

#### Install from code

Use `go get` to retrieve thd SDK, and add into your `GOPATH`:

```sh
$ go get -u github.com/shonenada/didiyun-go
```

## Example

### List your DC2

Before developing with SDK, you need to create a API Token in [API Token](https://app.didiyun.com/#/api/authtoken).

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

View [examples/](examples) for more detail.

## Contributing

TBD

## LICENSE

The Apache License (Version 2.0, January 2004).
