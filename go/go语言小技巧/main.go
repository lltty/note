package main

import (
	"fmt"
	"sync/atomic"
)

//三元运算符
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}

	return falseVal
}

//map键是否存在不一定用ok判断
func mapOk() {
	s := make(map[interface{}]bool)
	s["a"] = false
	if s["b"] {
		println("Yes")
	} else {
		println("Yes")
	}

	//冗余方式
	if _, ok := s["b"]; ok {
		println("Yes")
	} else {
		println("Yes")
	}
}

//对uint类型的数值原子的减操纵
func aotmicSubUint() {
	var s uint32
	s = 100
	var gg uint32
	gg = 20
	atomic.AddUint32(&s, ^uint32(gg-1))
	fmt.Println(s) //80
}

func main() {
	var a int
	b := 10
	max := If(a == 0, a, b).(int)
	fmt.Println(max)
}
