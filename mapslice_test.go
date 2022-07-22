package mapslice

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	ms := MapSlice[int]{
		MapItem[int]{Key: "abc", Value: 123},
		MapItem[int]{Key: "def", Value: 456},
		MapItem[int]{Key: "ghi", Value: 789},
	}

	b, err := json.Marshal(ms)
	if err != nil {
		t.Fatal(err)
	}

	e := "{\"abc\":123,\"def\":456,\"ghi\":789}"
	r := string(b)

	if r != e {
		t.Errorf("expected: %s\ngot: %s", e, r)
	}
}

func TestMarshalError(t *testing.T) {
	ms := MapSlice[chan int]{
		MapItem[chan int]{Key: "abc", Value: make(chan int)},
	}

	e := "json: error calling MarshalJSON for type mapslice.MapSlice[chan int]: json: unsupported type: chan int"
	if _, err := json.Marshal(ms); err != nil && e != err.Error() {
		t.Errorf("expected: %s\ngot: %v", e, err)
	}
}

func TestUnmarshal(t *testing.T) {
	ms := MapSlice[int]{}
	if err := json.Unmarshal([]byte("{\"abc\":123,\"def\":456,\"ghi\":789}"), &ms); err != nil {
		t.Fatal(err)
	}

	e := "[{abc 123} {def 456} {ghi 789}]"
	r := fmt.Sprintf("%v", ms)

	if r != e {
		t.Errorf("expected: %s\ngot: %s", e, r)
	}
}
