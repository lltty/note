package main

import (
	"fmt"
)

/*
 * 函数的传参是空接口类型时，传入空接口的切片时需要注意展开的问题
 */
func t1Run(a ...interface{}) {
	fmt.Println(a)
}

func test1() {
	t1Run(10, 20, 30) //[10, 20, 30]
	ts := []interface{}{10, 20, 30}
	t1Run(ts)    //[[10, 20, 30]] 这里整个数组被当成了第一个参数
	t1Run(ts...) //[10, 20, 30]
}

/*
 * 数组和切片虽然都是值传递，但是切片底层是对数组的引用，所以可以直接修改值
 * 但是数组直接修改值是不起作用的
 */
func test2() {
	t1 := []int{1, 2, 3}
	t2 := [...]int{1, 2, 3}

	func(s []int) {
		s[0] = 999
	}(t1)

	func(s [3]int) {
		s[0] = 999
	}(t2)

	fmt.Println(t1) //999 2 3
	fmt.Println(t2) // 1 2 3

}

//recover捕获的是祖父级调用时的异常，直接调用无效

func main() {

}
