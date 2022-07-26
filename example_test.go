package mapslice_test

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

func ExampleMapSlice_MarshalJSON() {
	ms := mapslice.MapSlice{
		mapslice.MapItem{Key: "abc", Value: 123},
		mapslice.MapItem{Key: "def", Value: 456},
		mapslice.MapItem{Key: "ghi", Value: 789},
	}

	b, err := json.Marshal(ms)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", b)

	// Output:
	// {"abc":123,"def":456,"ghi":789}
}

func ExampleMapSlice_UnmarshalJSON() {
	var ms = mapslice.MapSlice{}

	if err := json.Unmarshal([]byte(`{"abc":123,"def":456,"ghi":789}`), &ms); err != nil {
		panic(err)
	}

	fmt.Printf("%s", ms)

	// Output:
	// [{abc 123} {def 456} {ghi 789}]
}
