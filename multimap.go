package ge

import (
	"fmt"
	"strconv"
	"strings"
)

// MapStringToIntSlice desc
type MapStringToIntSlice interface {
	Add(key string, value int) MapStringToIntSlice
	ReprItems() []string
}

type mapStringToIntSlice struct {
	cache map[string][]int
}

func (m *mapStringToIntSlice) Add(key string, value int) MapStringToIntSlice {
	v, ok := m.cache[key]
	if !ok {
		v = make([]int, 1)
	}
	v = append(v, value)
	m.cache[key] = v
	return m
}

func (m *mapStringToIntSlice) ReprItems() []string {
	var items []string
	for k, v := range m.cache {
		var vs []string
		for _, vi := range v {
			vs = append(vs, strconv.Itoa(vi))
		}
		items = append(items, fmt.Sprintf("%s[%s]", k, strings.Join(vs, ",")))
	}
	return items
}

// MakeMapStringToIntSlice desc
func MakeMapStringToIntSlice() MapStringToIntSlice {
	return &mapStringToIntSlice{
		cache: make(map[string][]int),
	}
}
