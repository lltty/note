package doublylinkedlist

import (
	"fmt"
	"math"
	"strconv"
)

type HashTable struct {
	table []*List
	size  int
}

func HashTableNew(size int) *HashTable {
	return &HashTable{
		table: make([]*List, size, size),
		size:  size}
}

func ToString(value interface{}) string {
	switch value := value.(type) {
	case string:
		return value
	case int8, int16, int32, int64:
		return strconv.FormatInt(value.(int64), 10)
	case uint8, uint16, uint32, uint64:
		return strconv.FormatUint(value.(uint64), 10)
	case float32, float64:
		return strconv.FormatFloat(value.(float64), 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		return fmt.Sprintf("%+v", value)
	}
}

//取余法
func (ht *HashTable) hash(key interface{}) int {
	str := ToString(key)
	l, hash := len(str), 5381
	for i := 0; i < l; i++ {
		hash += int(str[i])
	}

	return hash % ht.size
}

//乘余取整法
func (ht *HashTable) hash1(key interface{}) int {

	str := ToString(key)
	l, hash := len(str), 5381
	for i := 0; i < l; i++ {
		hash += int(str[i])
	}

	_, div := math.Modf(float64(hash) * 0.765)
	fl, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", div), 10)
	intPart, _ := math.Modf(float64(ht.size) * fl)
	return int(intPart)
}

//限制了数据容量只能是10,100,1000,10000...
func (ht *HashTable) getHashBit(n int) int {

	if n <= ht.size {
		return n
	} else {
		if n <= ht.size*10 {
			return n / 10
		} else {
			return n / 100
		}
	}
}

//平方取中法
func (ht *HashTable) hash2(key interface{}) int {
	//字符串的Ascii值相加
	str := ToString(key)
	l, hash := len(str), 5381
	for i := 0; i < l; i++ {
		hash += int(str[i])
	}

	//计算出的值求平方
	hash *= hash

	//取hash值的中间几位
	return ht.getHashBit(hash)
}

/*
 * 数字分析法(只适用于特殊场景)
 * 关键字是数字组成，且可以估算数据的值分布，eg:
 * hash表长度100
 * K1=61317602 K2=61326875 K3=62739628 K4=61343634
 * K5=62706815 K6=62774638 K7=61381262 K8=61394220
 * 观察关键字可以发现1，2，3，6位取值比较集中，4，5，7，8位可选取其中的两位作为hash值
 */
func (ht *HashTable) hash3(n int) int {

	b0 := n / 1 % 10
	b1 := n / 10 % 10
	_ = n / 1000 % 10
	_ = n / 10000 % 10

	return b0 + b1*10
}

func (ht *HashTable) Put(key interface{}, value interface{}) {
	hash := ht.hash(key)
	val := ht.table[hash]
	if val == nil {
		ht.table[hash] = New(key, value)
	} else {
		//查找该键值是否存在,如果存在则覆盖，不存在则添加
		list := ht.table[hash]
		_, index, found := list.Find(key)
		if !found {
			ht.table[hash].Append(key, value)
		} else {
			list.Remove(index)
		}
	}
	ht.size += 1
}

func (ht *HashTable) Get(key interface{}) (interface{}, int, bool) {
	hash := ht.hash(key)
	val := ht.table[hash]
	if val == nil {
		return nil, -1, false
	} else {
		list := ht.table[hash]
		return list.Find(key)
	}
}

func (ht *HashTable) Remove(key interface{}) {
	hash := ht.hash(key)
	val := ht.table[hash]
	if val != nil {
		//查找该键值是否存在
		list := ht.table[hash]
		item, _, found := list.Find(key)
		if found {
			if item.(element).prev == nil {
				ht.table[hash].first = item.(element).next
			} else {
				item.(element).prev = item.(element).next
			}
			item = nil
		}
	}
	ht.size += 1
}
