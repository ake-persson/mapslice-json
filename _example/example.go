package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mickep76/mapslice-json"
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
