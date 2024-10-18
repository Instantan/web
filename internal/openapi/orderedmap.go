package openapi

import (
	"bytes"
	"encoding/json"
)

type OrderedMap[K comparable, V any] struct {
	unique map[K]int
	keys   []K
	values []V
}

func (m *OrderedMap[K, V]) Get(k K) (V, bool) {
	index, ok := m.unique[k]
	if !ok {
		return [1]V{}[0], ok
	}
	return m.values[index], true
}

func (m *OrderedMap[K, V]) Set(k K, v V) *OrderedMap[K, V] {
	if m.unique == nil {
		m.unique = map[K]int{}
		m.keys = []K{}
		m.values = []V{}
	}
	index, ok := m.unique[k]
	if !ok {
		m.unique[k] = len(m.keys)
		m.keys = append(m.keys, k)
		m.values = append(m.values, v)
	} else {
		m.keys[index] = k
		m.values[index] = v
	}
	return m
}

func (m *OrderedMap[K, V]) Iterate() func(func(key K, value V) bool) {
	return func(yield func(status K, value V) bool) {
		for i := 0; i < len(m.keys); i++ {
			if !yield(m.keys[i], m.values[i]) {
				return
			}
		}
	}
}

func (m *OrderedMap[K, V]) Len() int {
	return len(m.unique)
}

func (m OrderedMap[K, V]) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("{")

	i := -1
	for k, v := range m.Iterate() {
		i++
		if i != 0 {
			buf.WriteString(",")
		}
		key, err := json.Marshal(k)
		if err != nil {
			return nil, err
		}
		buf.Write(key)
		buf.WriteString(":")
		val, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}

	buf.WriteString("}")
	return buf.Bytes(), nil
}
