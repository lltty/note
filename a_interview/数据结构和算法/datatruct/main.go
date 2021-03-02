package main

import (
	"fmt"

	"t/note/a_interview/数据结构和算法/datatruct/queue"
	"t/note/a_interview/数据结构和算法/datatruct/stack"
)

func isNumber(i string) bool {
	return i[0]-'0' >= 0 &&
		i[0]-'0' <= 9
}

var ExpPriority map[string]int = map[string]int{
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
	"%": 1,
}

func expPriority(i, j string) bool {
	return ExpPriority[i] <= ExpPriority[j]
}

func isLeft(i string) bool {
	return i == "(" ||
		i == "[" ||
		i == "{"
}

func isRight(i string) bool {
	return i == ")" ||
		i == "]" ||
		i == "}"
}

func isMatch(i, j string) bool {
	return (i == "(" && j == ")") ||
		(i == "[" && j == "]") ||
		(i == "{" && j == "}")
}

//测试括号匹配
func TestBracketsMath() {
	stack := stack.NewStack(10)
	expression := "{[()[()]]}"

	var res bool
	for _, v := range expression {

		if isLeft(string(v)) {
			stack.Push(string(v))
		} else {

			left, err := stack.Pop()

			if err != nil {
				res = false
				break
			}
			if !isMatch(left.(string), string(v)) {
				res = false
				break
			}
			res = true
		}
	}
	fmt.Printf("计算结果:%v", res)
}

func main() {
	q := queue.NewQueue(3)

	err := q.Push(0)
	err = q.Push(1)
	err = q.Push(2)
	err = q.Push(3)

	var qq interface{}
	qq, err = q.Pop()
	fmt.Printf("%v->err:%v\n", qq, err)
	qq, err = q.Pop()
	fmt.Printf("%v->err:%v\n", qq, err)
	qq, err = q.Pop()
	err = q.Push(10)
	fmt.Printf("%v->err:%v\n", qq, err)
	qq, err = q.Pop()
	fmt.Printf("%v->err:%v\n", qq, err)

}
