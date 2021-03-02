package array

import (
	"errors"
)

type LinearTable interface {
	Push(item int)
	Pop() (int, error)
	Shift() (int, error)
	UnShift(item int)
}

/*
 * 无序数组
 * 数组里面的元素为什么不能直接删除元素，因为数组是已经申请好的连续的内存空间，删除会导致内存重新分配
 */
type Array struct {
	items []int
}

var (
	ArrayEmpty    = errors.New("数组为空")
	ArrayOutIndex = errors.New("下标越界")
)

func NewArray(len int) *Array {
	return &Array{
		items: make([]int, 0, len),
	}
}

//删除指定位置的元素
func (ary *Array) RemoveByIndex(index int) error {
	if ary.empty() {
		return ArrayEmpty
	}

	if index < 0 || index > ary.len {
		return ArrayOutIndex
	}

	newAry := make([]int, ary.len-1, cap(ary.items))
	for k, val := range ary.items {
		if k > index {
			newAry[k-1] = val
		} else if k < index {
			newAry[k] = val
		}
	}

	ary.len -= 1
	return nil
}

//删除指定值的元素
func (ary *Array) Remove(item int) ([]int, error) {
	if ary.empty() {
		return nil, ArrayEmpty
	}

	newAry := make([]int, len(ary.items)-1, cap(ary.items))

	res := make([]int, 0)
	for k, val := range ary.items {
		if item == val {
			res = append(res, k)
		} else {
			newAry = append(newAry, val)
		}
	}

	return res, nil
}

//指定位置插入元素
func (ary *Array) InsertBeforeIndex(index, item int) error {
	if ary.empty() {
		return ArrayEmpty
	}

	//这里是否可以在原数组上通过修改位置来实现
	newAry := tmpCapacityExtend(ary)
	if index >= ary.Len() {
		newAry = append(newAry, ary.items...)
		newAry = append(newAry, item)
		return nil
	}

	newAry = append(newAry, ary.items[0:index]...)
	newAry = append(newAry, item)
	newAry = append(newAry, ary.items[index:]...)

	ary.items = newAry
	return nil
}

func tmpCapacityExtend(ary *Array) []int {
	var newAry []int
	if ary.fill() {
		newAry = make([]int, len(ary.items), len(ary.items)*2)
	} else {
		newAry = make([]int, len(ary.items), cap(ary.items))
	}
	return newAry
}

//push
func (ary *Array) Push(item int) {
	//这里只是手动玩玩哈，并无卵用
	if ary.fill() {
		ary.capacityExtend()
	}

	ary.items = append(ary.items, item)
}

//pop
func (ary *Array) Pop() (int, error) {
	var item int
	if !ary.empty() {
		item = ary.items[len(ary.items)-1]
		ary.items = ary.items[0 : len(ary.items)-1]
		return item, nil
	}

	return item, ArrayEmpty
}

//shift
func (ary *Array) Shift() (int, error) {
	var item int
	if !ary.empty() {
		item = ary.items[0]
		ary.items = ary.items[1:]
		return item, nil
	}

	return item, ArrayEmpty
}

//unshift
func (ary *Array) UnShift(item int) {

	if ary.empty() {
		ary.items = append(ary.items, item)
	} else {
		var newAry []int
		if ary.fill() {
			newAry = make([]int, len(ary.items), len(ary.items)*2)
		} else {
			newAry = make([]int, len(ary.items), cap(ary.items))
		}

		newAry = append(newAry, item)
		newAry = append(newAry, ary.items...)

		ary.items = newAry
	}
}

func (ary *Array) Len() int {
	return len(ary.items)
}

func (ary *Array) empty() bool {
	return len(ary.items) == 0
}

func (ary *Array) fill() bool {
	return cap(ary.items) == len(ary.items)
}

func (ary *Array) capacityExtend() {
	s := make([]int, len(ary.items), len(ary.items)*2)
	for k, v := range ary.items {
		s[k] = v
	}

	ary.items = s
}
