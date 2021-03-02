package string

import (
	"fmt"
	"math"
)

//字符串匹配算法
//主串n，模式串m

/*
 * BF算法(单模式匹配，一个串跟另一个串是否匹配)
 * 在主串中查找下标0,1,2,n-m且长度是m的n-m+1个字符串
 * 算法复杂度O(n*m)
 */
func findBf(res, tar string) ([]int, bool) {

	resource := []rune(res)
	target := []rune(tar)

	n := len(resource)
	m := len(target)

	if n < m {
		return nil, false
	}

	if n == m {
		return []int{0, m}, res == tar
	}

	for i := 0; i < n-m; i++ {
		found := true
		j := 0
		for ; j < m; j++ {
			if resource[i+j] != target[j] {
				found = false
				break
			}
		}

		if found {
			return []int{i, i + j - 1}, true
		}

	}

	return nil, false
}

//字符对应字典表
var characterDict map[rune]int = map[rune]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
	'i': 8,
	'j': 9,
	'k': 10,
	'l': 11,
	'm': 12,
	'n': 13,
	'o': 14,
	'p': 15,
	'q': 16,
	'r': 17,
	's': 18,
	't': 19,
	'u': 20,
	'v': 21,
	'w': 22,
	'x': 23,
	'y': 24,
	'z': 25,
	'A': 26,
	'B': 27,
	'C': 28,
	'D': 29,
	'E': 30,
	'F': 31,
	'G': 32,
	'H': 33,
	'I': 34,
	'J': 35,
	'K': 36,
	'L': 37,
	'M': 38,
	'N': 39,
	'O': 40,
	'P': 41,
	'Q': 42,
	'R': 43,
	'S': 44,
	'T': 45,
	'U': 46,
	'V': 47,
	'W': 48,
	'X': 49,
	'Y': 50,
	'Z': 51,
}

//字符对应字典表,以素数为字典
var characterDict1 map[rune]int = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 5,
	'e': 7,
	'f': 11,
	'g': 13,
	'h': 17,
	'i': 19,
	'j': 23,
	'k': 29,
	'l': 31,
	'm': 37,
	'n': 41,
	'o': 43,
	'p': 47,
	'q': 53,
	'r': 57,
	's': 59,
	't': 61,
	'u': 63,
	'v': 67,
	'w': 71,
	'x': 73,
	'y': 79,
	'z': 83,
	'A': 85, //太累了，从这里开始乱写
	'B': 89,
	'C': 90,
	'D': 93,
	'E': 95,
	'F': 97,
	'G': 100,
	'H': 101,
	'I': 103,
	'J': 104,
	'K': 105,
	'L': 107,
	'M': 110,
	'N': 113,
	'O': 115,
	'P': 116,
	'Q': 118,
	'R': 120,
	'S': 128,
	'T': 130,
	'U': 140,
	'V': 150,
	'W': 151,
	'X': 153,
	'Y': 157,
	'Z': 159,
}

//不会hash冲突
func charHash1(resource []rune) int {
	l := len(resource)
	hash := 0
	for i := 0; i < l; i++ {
		hash += characterDict[resource[i]] * int(math.Pow(float64(52), float64(i)))
	}
	return hash
}

//会hash冲突
func charHash2(resource []rune) int {
	l := len(resource)
	hash := 0
	for i := 0; i < l; i++ {
		hash += characterDict1[resource[i]]
	}
	return hash
}

//RK算法(单模式匹配，一个串跟另一个串是否匹配)
func findRk(res, tar string) ([]int, bool) {

	resource := []rune(res)
	target := []rune(tar)

	n := len(resource)
	m := len(target)

	//计算n-m+1个字段的hash
	childrenL := n - m
	childrenCharacters := make([]int, 0, childrenL)
	//第一个子串
	firstChild := charHash1(resource[0:m])
	childrenCharacters = append(childrenCharacters, firstChild)

	//10 - 6 = 4
	//0
	// 1
	//2
	//3 3 + 6 - 1

	for i := 1; i < n-m; i++ {
		//这里计算子串的hash值我了减少遍历,采用的是用上一个hash值推到的方式
		hash := (childrenCharacters[i-1]-int(math.Pow(float64(52), float64(m-1)))-'a')*52 +
			(characterDict[resource[i+m-1]] - 'a')
		childrenCharacters = append(childrenCharacters, hash)
	}

	targetHash := charHash1(target)

	//将子串和模式串比对
	for i := 0; i < childrenL; i++ {
		if childrenCharacters[i] == targetHash {
			return []int{i, i + m - 1}, true
		}
	}

	fmt.Printf("%v\n", targetHash)
	fmt.Printf("%v\n", childrenCharacters)

	return nil, false
}

//优化的RK算法(单模式匹配，一个串跟另一个串是否匹配),允许冲突
func findRkPerf(res, tar string) ([]int, bool) {

	resource := []rune(res)
	target := []rune(tar)

	n := len(resource)
	m := len(target)

	//计算n-m+1个字段的hash
	targetHash := charHash1(target)
	for i := 0; i < n-m; i++ {
		//不保存了直接比较
		hash := charHash2(resource[i:m])
		if hash == targetHash {
			//因为hash冲突，所以还需要值比较
			found := true
			j := 0
			for ; j < m; j++ {
				if resource[i+j] != target[j] {
					found = false
					break
				}
			}

			if found {
				return []int{i, i + j - 1}, true
			}
		}
	}

	return nil, false
}

//BM算法(单模式匹配，一个串中同时查找多个串)

//RK算法(单模式匹配，一个串中同时查找多个串)
