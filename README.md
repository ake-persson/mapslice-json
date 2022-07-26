[![GoDoc](https://godoc.org/github.com/ake-persson/mapslice-json?status.svg)](https://godoc.org/github.com/ake-persson/mapslice-json)
[![Go Report Card](https://goreportcard.com/badge/github.com/ake-persson/mapslice-json)](https://goreportcard.com/report/github.com/ake-persson/mapslice-json)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/ake-persson/mapslice-json/blob/master/LICENSE)
[![Coverage Status](https://coveralls.io/repos/github/ake-persson/mapslice-json/badge.svg?branch=master)](https://coveralls.io/github/ake-persson/mapslice-json?branch=master)
[![Build Status](https://travis-ci.org/ake-persson/mapslice-json.svg?branch=master)](https://travis-ci.org/ake-persson/mapslice-json)

# mapslice-json

Go MapSlice for ordered marshal/ unmarshal of maps in JSON

# Example

https://go.dev/play/p/wzWP5vufCvL

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ake-persson/mapslice-json"
)

func main() {
	ms := mapslice.MapSlice{
		mapslice.MapItem{Key: "abc", Value: 123},
		mapslice.MapItem{Key: "def", Value: 456},
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
