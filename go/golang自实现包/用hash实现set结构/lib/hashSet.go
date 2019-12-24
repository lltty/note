package lib

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
)

type HashSet struct {
	sync.RWMutex
	items map[interface{}]bool
}

func newHashSet() *HashSet {
	return &HashSet{
		items: make(map[interface{}]bool),
	}
}

func (h *HashSet) Contains(i interface{}) bool {
	return h.items[i]
}

func (h *HashSet) Add(i interface{}) {
	h.Lock()
	defer h.Unlock()
	if !h.Contains(i) {
		h.items[i] = true
	}
}

func (h *HashSet) Remove(i interface{}) {
	h.Lock()
	defer h.Unlock()
	delete(h.items, i)
}

func (h *HashSet) Clear() {
	h.Lock()
	defer h.Unlock()
	h.items = make(map[interface{}]bool)
}

func (h *HashSet) Len() int {
	return len(h.items)
}

func (h *HashSet) Same(other *HashSet) bool {

	h.RLock()
	defer h.RUnlock()

	if other == nil {
		return false
	}

	if h.Len() != other.Len() {
		return false
	}

	for k := range other.items {
		if !h.Contains(k) {
			return false
		}
	}

	return true

}

func (h *HashSet) Elements() []interface{} {
	h.RLock()
	defer h.RUnlock()

	s := make([]interface{}, h.Len())
	index := 0
	for key := range h.items {
		s[index] = key
		index++
	}

	return s
}

func (h *HashSet) String() string {

	h.RLock()
	defer h.RUnlock()

	var buf bytes.Buffer
	buf.WriteString("{\n")

	index := 0
	for key := range h.items {
		str := fmt.Sprintf("%v", key)
		buf.WriteString(strings.Repeat(" ", 2) + str + "\n")
		index++
	}
	buf.WriteString("}\n")
	return buf.String()
}

func (h *HashSet) Union(other *HashSet) []interface{} {
	h.Lock()
	defer h.Unlock()

	union := make([]interface{}, 0, len(h.items))
	for k := range h.items {
		union = append(union, k)
	}

	for k := range other.items {
		if !h.Contains(k) {
			union = append(union, k)
		}
	}

	return union
}

func (h *HashSet) Intersect(other *HashSet) []interface{} {
	h.Lock()
	defer h.Unlock()

	filter := make([]interface{}, 0)
	for k := range other.items {
		if h.Contains(k) {
			filter = append(filter, k)
		}
	}
	return filter
}

func (h *HashSet) Diff(other *HashSet) []interface{} {
	h.Lock()
	defer h.Unlock()

	filter := make([]interface{}, 0)
	for k := range other.items {
		if !h.Contains(k) {
			filter = append(filter, k)
		}
	}
	return filter
}

func (h *HashSet) SymmetricDiff(other *HashSet) []interface{} {
	h.Lock()
	defer h.Unlock()

	filter := make([]interface{}, 0)

	for k := range other.items {
		if !h.Contains(k) {
			filter = append(filter, k)
		}
	}

	for k := range h.items {
		if !other.Contains(k) {
			filter = append(filter, k)
		}
	}

	return filter
}
