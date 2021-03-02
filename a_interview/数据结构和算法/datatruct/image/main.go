package main

import (
	"fmt"
)

/*
 * 图: 无向图，有向图，带权图
 *  入度和出度
 * 稀疏图：顶点很多，但是每个顶点的边不多
 */

//使用邻接矩阵存储图
//浪费存储空间
func image1() {
	//image := [][4]int{}
	for i := 0; i < 4; i++ {
		t := []int{}
		for j := 0; j < i; j++ {
			t = append(t, 0)
		}
		t = append(t, 0)
		for x := i + 1; x < 4; x++ {
			t = append(t, 0)
		}
		fmt.Println(t)
	}
}

//使用邻接表存储图

func main() {
	test()
}
