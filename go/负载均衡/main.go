package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

//经数学验证的随机算法
func shuffle(indexs []int) {
	for i := len(indexs); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		indexs[lastIdx], indexs[idx] = indexs[idx], indexs[lastIdx]
	}
}

//上述算法已经在系统包实现了
func shuffleSys(len int) []int {
	b := rand.Perm(len)
	return b
}

func main() {
	var cnt1 = map[int]int{}
	for i := 0; i < 1000000; i++ {
		/*var s1 = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle(s1)*/
		s1 := shuffleSys(7)
		cnt1[s1[1]]++
	}
	fmt.Println(cnt1)
}
