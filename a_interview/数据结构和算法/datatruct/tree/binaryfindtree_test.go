package tree

import (
	"testing"
)

type People struct {
	Name string
	Age  int
	Addr string
}

var testTree *Tree

func TestTreePut(t *testing.T) {
	testTree = newTree()
	testTree.Put(30, People{
		Name: "李红卫",
		Age:  30,
		Addr: "宝鸡",
	})
	testTree.Put(13, People{
		Name: "张丽丽",
		Age:  13,
		Addr: "延安",
	})
	testTree.Put(45, People{
		Name: "张波",
		Age:  45,
		Addr: "上海",
	})
	testTree.Put(8, People{
		Name: "张八岁",
		Age:  8,
		Addr: "上海",
	})
	testTree.Put(21, People{
		Name: "李21岁",
		Age:  21,
		Addr: "上海",
	})
	testTree.Put(35, People{
		Name: "韩国平",
		Age:  35,
		Addr: "上海",
	})
	testTree.Put(24, People{
		Name: "张二十四",
		Age:  24,
		Addr: "上海",
	})

	testTree.Put(15, People{
		Name: "张十五",
		Age:  15,
		Addr: "上海",
	})

	testTree.Put(22, People{
		Name: "二十二",
		Age:  22,
		Addr: "上海",
	})

	testTree.Put(28, People{
		Name: "二十八",
		Age:  28,
		Addr: "上海",
	})

	t.Logf("根节点:%v\n", testTree.Root)
	t.Logf("|\n")
	t.Logf("|\n")
	t.Logf("------左子节点:%v\n", testTree.Root.Children[0])
	t.Logf("      |\n")
	t.Logf("      |\n")
	t.Logf("      ------左子节点:%v\n", testTree.Root.Children[0].Children[0])
	t.Logf("      |\n")
	t.Logf("      |\n")
	t.Logf("      ------右子节点:%v\n", testTree.Root.Children[0].Children[1])
	t.Logf("------右子节点:%v\n", testTree.Root.Children[1])
	t.Logf("      |\n")
	t.Logf("      |\n")
	t.Logf("      ------左子节点:%v\n", testTree.Root.Children[1].Children[0])
}

func TestTreeFind(t *testing.T) {
	item, found := testTree.Find(35)
	if found {
		t.Logf("%v\n", item)
	} else {
		t.Log("没找到\n")
	}
}

func TestTreeOrder(t *testing.T) {
	testTree.Order(1)
	testTree.Order(2)
	testTree.Order(3)
}

func _TestTreeRemove(t *testing.T) {
	//删除21
	testTree.Remove(21)
	testTree.Order(2)
}

func TestTreeGetHeight(t *testing.T) {
	t.Logf("树的高度:%d\n", testTree.GetHeight())
}
