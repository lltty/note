package lib

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	hs := newHashSet()
	hs.Add(1)
	hs.Add(2)
	hs.Add("爸爸")
	hs.Add([...]int{1, 2, 3})
	es := hs.String()
	println(es)
}

func TestUnion(t *testing.T) {
	h := newHashSet()
	h.Add(1)
	h.Add(2)
	h.Add(3)
	h.Add("01")

	h1 := newHashSet()
	h1.Add(1)
	h1.Add(2)
	h1.Add(4)
	h1.Add(5)
	h1.Add(6)

	fmt.Printf("union:%v\n", h.Union(h1))
	fmt.Printf("Intersect:%v\n", h.Intersect(h1))
	fmt.Printf("diff:%v\n", h.Diff(h1))
	fmt.Printf("diff:%v\n", h.SymmetricDiff(h1))
}
