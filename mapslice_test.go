package mapslice

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	ms := MapSlice{
		MapItem{Key: "abc", Value: 123},
		MapItem{Key: "def", Value: 456},
		MapItem{Key: "ghi", Value: 789},
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
	ms := MapSlice{
		MapItem{Key: "abc", Value: make(chan int)},
	}

	e := "json: error calling MarshalJSON for type mapslice.MapSlice: json: unsupported type: chan int"
	if _, err := json.Marshal(ms); err != nil && e != err.Error() {
		t.Errorf("expected: %s\ngot: %v", e, err)
	}
}

func TestUnmarshal(t *testing.T) {
	ms := MapSlice{}
	if err := json.Unmarshal([]byte("{\"abc\":123,\"def\":456,\"ghi\":789}"), &ms); err != nil {
		t.Fatal(err)
	}

	e := "[{abc 123} {def 456} {ghi 789}]"
	r := fmt.Sprintf("%v", ms)

	if r != e {
		t.Errorf("expected: %s\ngot: %s", e, r)
	}
}

func TestConcurrentUnmarshal(t *testing.T) {
	tests := []struct {
		name     string
		data     json.RawMessage
		expected string
	}{
		{
			name:     "data1",
			data:     json.RawMessage(`{"s1":"v1", "s2":"v2", "s3":"v3"}`),
			expected: "[{s1 v1} {s2 v2} {s3 v3}]",
		},
		{
			name:     "data2",
			data:     json.RawMessage(`{"s1":"v1", "s2":"v2", "s3":"v3"}`),
			expected: "[{s1 v1} {s2 v2} {s3 v3}]",
		},
		{
			name:     "data3",
			data:     json.RawMessage(`{"s1":"v1", "s2":"v2", "s3":"v3"}`),
			expected: "[{s1 v1} {s2 v2} {s3 v3}]",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ms := MapSlice{}
			if err := json.Unmarshal(tt.data, &ms); err != nil {
				t.Fatal(err)
			}
			actual := fmt.Sprintf("%v", ms)
			if tt.expected != actual {
				t.Errorf("expected: %s\ngot: %s", tt.expected, actual)
				return
			}
		})
	}
}
