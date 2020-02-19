[![GoDoc](https://godoc.org/github.com/mickep76/mapslice-json?status.svg)](https://godoc.org/github.com/mickep76/mapslice-json)
[![Go Report Card](https://goreportcard.com/badge/github.com/mickep76/mapslice-json)](https://goreportcard.com/report/github.com/mickep76/mapslice-json)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mickep76/mapslice-json/blob/master/LICENSE)

# mapslice-json

Go MapSlice for ordered marshal/ unmarshal of maps in JSON

# Example

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mickep76/mapslice-json"
)

func main() {
	ms := mapslice.MapSlice{
		mapslice.MapItem{Key: "abc\"", Value: 123},
		mapslice.MapItem{Key: "d ef", Value: 456},
		mapslice.MapItem{Key: "ghi", Value: 789},
	}

	b, err := json.Marshal(ms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	ms = mapslice.MapSlice{}
	if err := json.Unmarshal(b, &ms); err != nil {
		log.Fatal(err)
	}

	fmt.Println(ms)
}
```
