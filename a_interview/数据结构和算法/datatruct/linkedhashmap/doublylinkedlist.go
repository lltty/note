package doublylinkedlist

import "t/note/a_interview/数据结构和算法/datatruct/gods/utils"

/*
 * 保存键值的双向链表
 */

type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	key   interface{}
	value interface{}
	prev  *element
	next  *element
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values)%2 == 0 {
		list.Append(values...)
	}
	return list
}

func (list *List) checkDouble(values ...interface{}) bool {
	l := len(values)
	return l%2 != 0
}

//键值对保存
func (list *List) Append(values ...interface{}) {

	if !list.checkDouble(values) {
		return
	}

	for i, l := 0, len(values); i < l; i += 2 {
		newElement := &element{key: values[i], value: values[i+1], prev: list.last}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

func (list *List) Prepend(values ...interface{}) {

	if !list.checkDouble(values) {
		return
	}

	l := len(values)
	for v := l - 1; v >= 0; v -= 2 {
		newElement := &element{key: values[v-1], value: values[v], next: list.first}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.first.prev = newElement
			list.first = newElement
		}
		list.size++
	}
}

func (list *List) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var element *element
	// determine traversal direction, last to first or first to last
	if list.size-index < index {
		element = list.last
		for e := list.size - 1; e != index; e, element = e-1, element.prev {
		}
	} else {
		element = list.first
		for e := 0; e != index; e, element = e+1, element.next {
		}
	}

	if element == list.first {
		list.first = element.next
	}
	if element == list.last {
		list.last = element.prev
	}
	if element.prev != nil {
		element.prev.next = element.next
	}
	if element.next != nil {
		element.next.prev = element.prev
	}

	element = nil

	list.size--
}

func (list *List) ReplaceByKey(key interface{}, val interface{}) {
	for element := list.first; element != nil; element = element.next {
		if element.key == key {
			element.value = val
		}
	}
}

func (list *List) ReplaceByIndex(index int, val interface{}) {
	if !list.withinRange(index) {
		return
	}

	for i, element := 0, list.first; element != nil; i, element = i+1, element.next {
		if i == index {
			element.value = val
		}
	}
}

func (list *List) Find(key interface{}) (interface{}, int, bool) {

	if list.size == 0 {
		return nil, -1, false
	}

	for index, element := 0, list.first; element != nil; index, element = index+1, element.next {
		if element.key == key {
			return element.value, index, true
		}
	}

	return nil, -1, false

}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

func (list *List) Map() ([]interface{}, map[interface{}]interface{}) {
	m := make(map[interface{}]interface{}, list.Size())
	keys := make([]interface{}, 0, list.Size())
	for element := list.first; element != nil; element = element.next {
		m[element.key] = element.value
		keys = append(keys, element.key)
	}

	return keys, m
}

func (list *List) Sort(comparator utils.Comparator) {

	if list.size < 2 {
		return
	}

	keys, m := list.Map()
	utils.Sort(keys, comparator)

	list.Clear()

	values := make([]interface{}, 0, list.size*2)
	for _, key := range keys {
		values = append(values, key, m[key])
	}

	list.Append(values...)

}

func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
