package string

import (
	"testing"
)

func TestFindBf(t *testing.T) {
	s1 := "我是你n23们吗手动删掉所"
	s2 := "比手动"
	index, found := findBf(s1, s2)
	t.Logf("find res:%+v, %+v", index, found)
}

func TestFindRk(t *testing.T) {
	s1 := "abcASSDFSCqwqsdsqdef"
	s2 := "ASSD"
	index, found := findRkPerf(s1, s2)
	t.Logf("find res:%+v, %+v", index, found)
}
