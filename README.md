# worldota-sdk-go

Golang SDK for [ETG APIv3](https://docs.emergingtravel.com/).

## Installation
```shell
go get github.com/molodsom/worldota-sdk-go
```

## Quickstart
To start using ETG APIv3 you need a key, which you received after registration.
A key is a combination of an id and uuid. These are passed into each request as a Basic Auth header after
initialization.

Then you can use almost all available methods.
Say you want to check an overview of the available methods (which is `api/b2b/v3/overview` endpoint), you do:

```go
package main

import "github.com/molodsom/worldota-sdk-go"

func main() {
	c := worldota.Client("0000", "b85ae9bc-30d2-4607-8a7f-53c673a001dd")
	c.Overview() // models.Overview
}
```

Another example is downloading hotels dump with `api/b2b/v3/hotel/info/dump` endpoint:

```go
package main

import (
	"fmt"
	"github.com/molodsom/worldota-sdk-go"
	"github.com/molodsom/worldota-sdk-go/models"
)

func main() {
	c := worldota.Client("0000", "b85ae9bc-30d2-4607-8a7f-53c673a001dd")
	r := c.HotelInfoDump(models.LanguageRequest{Language: "pt"}) // models.HotelDump
	fmt.Println(*r.Data.URL)
}
```

You can also save and unpack the file

```go
path, err := c.HotelInfoIncrementalDump(models.LanguageRequest{Language: "de"}).Decompress("/path/to/dir")
```