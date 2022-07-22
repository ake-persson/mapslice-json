package mapslice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

// MapItem representation of one map item.
type MapItem[V any] struct {
	Key   string
	Value V
	index uint64
}

// MapSlice of map items.
type MapSlice[V any] []MapItem[V]

func (ms MapSlice[V]) Len() int           { return len(ms) }
func (ms MapSlice[V]) Less(i, j int) bool { return ms[i].index < ms[j].index }
func (ms MapSlice[V]) Swap(i, j int)      { ms[i], ms[j] = ms[j], ms[i] }

var indexCounter uint64

func nextIndex() uint64 {
	indexCounter++
	return indexCounter
}

// MapItem as a string.
func (mi MapItem[V]) String() string {
	return fmt.Sprintf("{%v %v}", mi.Key, mi.Value)
}

// MarshalJSON for map slice.
func (ms MapSlice[V]) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})
	for i, mi := range ms {
		b, err := json.Marshal(&mi.Value)
		if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprint(mi.Key)))
		buf.Write(b)
		if i < len(ms)-1 {
			buf.Write([]byte{','})
		}
	}
	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

// UnmarshalJSON for map slice.
func (ms *MapSlice[V]) UnmarshalJSON(b []byte) error {
	m := map[string]MapItem[V]{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	for k, v := range m {
		*ms = append(*ms, MapItem[V]{Key: k, Value: v.Value, index: v.index})
	}
	sort.Sort(*ms)
	return nil
}

// UnmarshalJSON for map item.
func (mi *MapItem[V]) UnmarshalJSON(b []byte) error {
	var v V
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	mi.Value = v
	mi.index = nextIndex()
	return nil
}
