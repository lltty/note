package tree

/*
 * 注意堆的下标是从1开始的，主要是为了便于计算
 * 这里构建的是大顶堆
 */

type Heap struct {
	List  []*HeapNode
	Index int //因为是完全二叉树，所以这里的下标也可以当size用
	Cap   int
}

type HeapNode struct {
	Key int
	Val treeItem
}

func NewHeap(n int) *Heap {
	return &Heap{
		List:  make([]*HeapNode, n+1, n+1),
		Index: 0,
		Cap:   n,
	}
}

func (h *Heap) Put(key int, val treeItem) bool {

	if h.Index == h.Cap {
		return false
	}

	//注意这里是先加的，所以第一个元素就是从下标1开始的
	h.Index++
	h.List[h.Index] = &HeapNode{
		Key: key,
		Val: val,
	}

	t := h.Index
	//这里需要注意弱类型和强类型语言的差异，如果是弱类型需要判断奇偶
	//调整堆,一直查找当前节点的父节点，如果需要则做值替换
	for t/2 > 0 && h.List[t].Key > h.List[t/2].Key {
		h.List[t], h.List[t/2] = h.List[t/2], h.List[t]
		t = t / 2
	}

	return true
}

func (h *Heap) Find(key int) (*HeapNode, bool) {
	if h.Index == 0 {
		return nil, false
	}
	//直接遍历数组
	for i := 1; i <= h.Index; i++ {
		if h.List[i].Key == key {
			return h.List[i], true
		}
	}

	return nil, false
}

//删除堆顶元素
func (h *Heap) RemoveTop(key int) (*HeapNode, bool) {
	if h.Index == 0 {
		return nil, false
	}

	//将最后一个元素和树顶元素替换，并且删除最后一个元素
	h.List[1] = h.List[h.Index]
	h.Index--

	i := 1
	for {
		maxPos := i
		if i*2 <= h.Index && h.List[i*2].Key > h.List[i].Key {
			maxPos = i * 2
		}
		if (i*2+1) <= h.Index && h.List[i*2+1].Key > h.List[i].Key {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}

		h.List[i], h.List[maxPos] = h.List[maxPos], h.List[i]

		//切记这一步
		i = maxPos

	}

	t := *(h.List[h.Index])
	h.List[h.Index] = nil

	return &t, true
}
