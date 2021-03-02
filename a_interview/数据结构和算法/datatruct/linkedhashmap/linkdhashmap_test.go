package doublylinkedlist

import (
	"testing"
)

type People struct {
	Name  string
	Age   int
	Addr  string
	Hobby []string
}

var table *HashTable

func TestHashTablePut(t *testing.T) {
	table = HashTableNew(10)
	table.Put("李红卫", People{
		Name: "李红卫",
		Age:  30,
		Addr: "宝鸡",
	})
	table.Put("李红卫", People{
		Name: "李红卫",
		Age:  50,
		Addr: "上海",
	})
	table.Put("赵DE", People{
		Name: "赵DE",
		Age:  30,
		Addr: "宝鸡",
	})
	table.Put("赵CF", People{
		Name: "赵CF",
		Age:  30,
		Addr: "宝鸡",
	})
	table.Put("赵BG", People{
		Name: "赵BG",
		Age:  30,
		Addr: "宝鸡",
	})
	table.Put("赵AH", People{
		Name: "赵AH",
		Age:  30,
		Addr: "宝鸡",
	})

	t.Logf("hashTable:%v\n", table)

}

func TestHashTableGet(t *testing.T) {
	people, index, found := table.Get("赵BG")
	t.Logf("赵BG found:%v\n", found)
	t.Logf("赵BG index:%v\n", index)
	t.Logf("赵BG:%v\n", people)

	people1, index1, found1 := table.Get("李红卫")
	t.Logf("李红卫 found:%v\n", found1)
	t.Logf("李红卫 index:%v\n", index1)
	t.Logf("李红卫:%v\n", people1)
	t.Logf("李红卫:%v\n", table.table[0])
}
