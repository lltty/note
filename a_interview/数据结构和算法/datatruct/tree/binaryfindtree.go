package tree

import (
	"fmt"
	"math"
)

type treeItem interface{}

/* -------------二叉查找数----------------------
 * 满二叉树，完全二叉树
 * 二叉树的形态各式各样。对于同一组数组，会因为插入数据的先后不同，构造出不同的二叉树。
 * 不同形态的二叉树，查找、插入，删除的的时间复杂度是不一样的，最坏的情况是退化成链表，时间复杂度是O(n)
 * 理想的情况是完全二叉树或者满二叉树，时间复杂度是O(height)
 * 所以我们需要构建不管怎么删除、插入数据，在任何情况，都能保持任意节点左右子树都平衡的平衡二叉查找树
 * 平衡二叉查找树的高度近似logN,所以插入，删除，查找的时间复杂度也比较稳定，是O(logN)
 */

type Tree struct {
	Root *Node
	Size int
}

type Node struct {
	Key      int
	Val      treeItem
	Parent   *Node
	Children [2]*Node
}

func (t *Tree) CompareInt(m, n int) int {
	switch {
	case m > n:
		return 1
	case m < n:
		return -1
	default:
		return 0
	}
}

func newTree() *Tree {
	return &Tree{}
}

func (t *Tree) Put(key int, val interface{}) {
	t.put(key, val, &t.Root)
}

func (t *Tree) Find(key int) (*Node, bool) {
	return t.find(key, &t.Root)
}

func (t *Tree) Remove(key int) {
	node, found := t.Find(key)
	if !found {
		return
	}
	t.remove(key, &node)
}

func (t *Tree) Order(n int) {
	switch n {
	case 0:
		t.levelOrder()
	case 1:
		t.frontOrder(t.Root)
	case 2:
		t.middleOrder(t.Root)
	case 3:
		t.backendOrder(t.Root)
	}
	fmt.Println()
}

func (t *Tree) GetHeight() int {
	return int(t.getHeight(t.Root))
}

func (t *Tree) getHeight(root *Node) float64 {
	if root == nil {
		return 0
	}

	return math.Max(t.getHeight(root.Children[0]), t.getHeight(root.Children[1])) + 1
}

//按照层遍历
func (t *Tree) levelOrder() {

}

//前序遍历
func (t *Tree) frontOrder(root *Node) {

	if root == nil {
		return
	}

	fmt.Print(root.Key, ",")
	t.frontOrder(root.Children[0])
	t.frontOrder(root.Children[1])

}

//中序遍历
func (t *Tree) middleOrder(root *Node) {
	if root == nil {
		return
	}

	t.frontOrder(root.Children[0])
	fmt.Print(root.Key, ",")
	t.frontOrder(root.Children[1])
}

//后序遍历
func (t *Tree) backendOrder(root *Node) {
	if root == nil {
		return
	}

	t.frontOrder(root.Children[0])
	t.frontOrder(root.Children[1])
	fmt.Print(root.Key)
}

func (t *Tree) put(key int, val interface{}, parent **Node) {

	p := *parent

	if p == nil {
		node := &Node{
			Key: key,
			Val: val,
		}
		*parent = node
		t.Size += 1
		return
	}

	a := t.CompareInt(key, p.Key)
	if a == 0 {
		p.Val = val
	} else if a == -1 { //左
		if p.Children[0] == nil {
			p.Children[0] = &Node{
				Key:    key,
				Val:    val,
				Parent: p,
			}
		} else {
			t.put(key, val, &p.Children[0])
		}
	} else { //右
		if p.Children[1] == nil {
			p.Children[1] = &Node{
				Key:    key,
				Val:    val,
				Parent: p,
			}
		} else {
			t.put(key, val, &p.Children[1])
		}
	}
}

func (t *Tree) find(key int, parent **Node) (*Node, bool) {
	p := *parent
	if p == nil {
		return nil, false
	}

	a := t.CompareInt(key, p.Key)
	if a == 0 {
		return p, true
	} else if a == -1 { //左
		if p.Children[0] == nil {
			return nil, false
		} else {
			return t.find(key, &p.Children[0])
		}
	} else { //右
		if p.Children[1] == nil {
			return nil, false
		} else {
			return t.find(key, &p.Children[1])
		}
	}

}

/*
 * 删除节点的情况:
 * 1 如果要删除的节点没有子节点，我们只需要直接将父节点中。指向要删除节点的指针置为nil
 * 2 如果要删除的节点只有一个子节点，我们只需要更新父节点中，指向要删除节点的指针，让它指向要删除节点的子节点就可以了
 * 3 如果要删除的节点有两个子节点。我们需要找到这个节点的右子树的最小节点，把它替换到要删除的节点上。然后再删除掉那个最小节点。
 */

func (t *Tree) remove(key int, node **Node) {
	fmt.Printf("要删除的节点:%v\n", *node)
	defer func() {
		t.Size -= 1
	}()
	//要删除的节点有两个子节点
	pp := (*node).Parent
	p := *node
	if p.Children[0] != nil && p.Children[1] != nil {
		//查找右子节点中的最小节点
		minNp := *node                      //父节点
		minRightNode := (*node).Children[1] //右子节点
		//注意这里是在左边找,因为是找最小的
		for minRightNode.Children[0] != nil {
			minNp = minRightNode
			minRightNode = minRightNode.Children[0]
		}

		(*node).Val = minRightNode.Val //替换要删除节点的值为找到的值
		(*node).Key = minRightNode.Key
		minNp.Children[0] = nil
		return
	}

	//当前要删除的节点是否根节点
	if pp != nil {
		//判断当前要删除的节点是父节点的左子节点，还是右子节点
		replaceNodeIndex := 0
		if !LeftLeaf(p, pp) {
			replaceNodeIndex = 1
		}

		//删除节点是叶子节点或者只有一个子节点
		if p.Children[0] != nil {
			(*node).Parent.Children[replaceNodeIndex] = p.Children[0]
		} else if p.Children[1] != nil {
			(*node).Parent.Children[replaceNodeIndex] = p.Children[1]
		} else {
			(*node).Parent.Children[replaceNodeIndex] = nil
		}
	} else {
		if p.Children[0] != nil {
			t.Root = p.Children[0]
			(*node).Children[0].Parent = nil
		} else if p.Children[1] != nil {
			t.Root = p.Children[1]
			(*node).Children[1].Parent = nil
		} else {
			t.Root = nil
		}
	}
}

func LeftLeaf(p, pp *Node) bool {
	return pp.Children[0] != nil &&
		pp.Children[0].Key == p.Key
}

/*
 * 补充说明
 * 1 关于二叉树的删除操作，也可以使用将删除节点标记为删除，但是不真正删除节点的方式。这样比较浪费内存空间，但是删除操作变简单了
 * 2 如果要二叉树支持重复数据的存储，有两种方式
 *   A: 我们可以使用数组或者链表作为节点，存储相同键值的数据
 *   B: 每个节点仍然只存储一个数据。在查找插入时，如果遇到相同的节点，就将要插入的数据放到这个节点的右子树，也就是说把这个新插入的数据
 *       当做大于这个节点来处理
 *      当要查找数据的时候，遇到值相同的节点也不停止，而是继续在右子树中查找，直到遇到叶子节点为止
 *      删除操作也是一样，找出所有要删除的节点，然后按照前面将的删除操作的方法，依次删除。
 */
