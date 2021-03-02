package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var unorderedAry []int = []int{10, 4, 10, 5, 11, 9, 10, 6, 3, 2, 9, 1, 9, 9, 10, 7, 40, 10}

func bubbleSort(a []int) {
	l := len(a)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if a[j] > a[j+1] {
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
			}
		}
	}
}

func selectSort(a []int) {
	l := len(a)
	//4, 5, 6, 3, 2, 1
	for i := 0; i < l-1; i++ {
		max := 0
		var j int
		for j = 1; j < l-i; j++ {
			if a[j] > a[max] {
				max = j
			}
		}
		if max != j-1 {
			tmp := a[max]
			a[max] = a[j-1]
			a[j-1] = tmp
		}

	}
}

func insertSort(a []int) {
	l := len(a)
	for i := 1; i < l; i++ {
		tmp := a[i]
		for j := i - 1; j >= 0; j-- {
			if tmp < a[j] {
				a[j+1] = a[j]
				a[j] = tmp
			} else {
				break
			}
		}
	}
}

func merge(b, c []int) []int {
	i := 0
	bl := len(b)
	j := 0
	cl := len(c)

	a := make([]int, 0, len(b)+len(c))
	for i < bl && j < cl {
		if b[i] <= c[j] {
			a = append(a, b[i])
			i++
		} else {
			a = append(a, c[j])
			j++
		}
	}

	if i < bl {
		for ; i < bl; i++ {
			a = append(a, b[i])
		}
	}

	if j < cl {
		for ; j < cl; j++ {
			a = append(a, c[j])
		}
	}

	return a
}

//归并排序
func mergeSort(a []int) []int {
	start := 0
	end := len(a) - 1

	if start >= end {
		return a
	}

	middle := start + (end-start)>>1
	b := mergeSort(a[start : middle+1])
	c := mergeSort(a[middle+1:])

	t := merge(b, c)

	return t
}

type People struct {
	Name  string
	Age   int
	Phone string
}

//对每个桶里的人群按照age进行插入排序
func sortPeopleByAge(ps []People) {
	l := len(ps)
	for i := 1; i < l; i++ {
		tmp := ps[i]
		for j := i - 1; j >= 0; j-- {
			if tmp.Age < ps[j].Age {
				ps[j+1] = ps[j]
				ps[j] = tmp
			} else {
				break
			}
		}
	}
}

func sortPeopleByPhone(ps []People) {
	l := len(ps)
	for i := 1; i < l; i++ {
		tmp := ps[i]
		iPhone, _ := strconv.Atoi(tmp.Phone)
		for j := i - 1; j >= 0; j-- {
			jPhone, _ := strconv.Atoi(ps[j].Phone)
			if iPhone < jPhone {
				ps[j+1] = ps[j]
				ps[j] = tmp
			} else {
				break
			}
		}
	}
}

//桶排序
func bucketSort(ps []People) {
	//扫描数据
	min := 10
	max := 0
	for _, v := range ps {
		if v.Age < min {
			min = v.Age
		}

		if v.Age > max {
			max = v.Age
		}
	}

	//划分数据到10个桶里
	cap := (max-min)/10 + 1
	b0 := make([]People, 0, cap)
	b1 := make([]People, 0, cap)
	b2 := make([]People, 0, cap)
	b3 := make([]People, 0, cap)
	b4 := make([]People, 0, cap)
	b5 := make([]People, 0, cap)
	b6 := make([]People, 0, cap)
	b7 := make([]People, 0, cap)
	b8 := make([]People, 0, cap)
	b9 := make([]People, 0, cap)

	for _, v := range ps {
		age := v.Age
		switch {
		case age >= min && age <= (min+cap):
			b0 = append(b0, v)
			break
		case age > (min+cap) && age <= (min+2*cap):
			b1 = append(b1, v)
			break
		case age > (min+2*cap) && age <= (min+3*cap):
			b2 = append(b2, v)
			break
		case age > (min+3*cap) && age <= (min+4*cap):
			b3 = append(b3, v)
			break
		case age > (min+4*cap) && age <= (min+5*cap):
			b4 = append(b4, v)
			break
		case age > (min+5*cap) && age <= (min+6*cap):
			b5 = append(b5, v)
			break
		case age > (min+6*cap) && age <= (min+7*cap):
			b6 = append(b6, v)
			break
		case age > (min+7*cap) && age <= (min+8*cap):
			b7 = append(b7, v)
			break
		case age > (min+8*cap) && age <= (min+9*cap):
			b8 = append(b8, v)
			break
		case age >= (min+9*cap) && age <= max:
			b9 = append(b9, v)
			break
		}
	}

	sortPeopleByAge(b0)
	sortPeopleByAge(b1)
	sortPeopleByAge(b2)
	sortPeopleByAge(b3)
	sortPeopleByAge(b4)
	sortPeopleByAge(b5)
	sortPeopleByAge(b6)
	sortPeopleByAge(b7)
	sortPeopleByAge(b8)
	sortPeopleByAge(b9)
	fmt.Printf("b0:%v->%v\n", b0[0], b0[len(b0)-1])
	fmt.Printf("b0:%v->%v\n", b1[0], b1[len(b1)-1])
	fmt.Printf("b0:%v->%v\n", b2[0], b2[len(b2)-1])
	fmt.Printf("b0:%v->%v\n", b3[0], b3[len(b3)-1])
	fmt.Printf("b0:%v->%v\n", b4[0], b4[len(b4)-1])
	fmt.Printf("b0:%v->%v\n", b5[0], b5[len(b5)-1])
	fmt.Printf("b0:%v->%v\n", b6[0], b6[len(b6)-1])
	fmt.Printf("b0:%v->%v\n", b7[0], b7[len(b7)-1])
	fmt.Printf("b0:%v->%v\n", b8[0], b8[len(b8)-1])
	fmt.Printf("b0:%v->%v\n", b9[0], b9[len(b9)-1])
}

//计数排序
func countSort(ps []People) {
	//扫描数据
	min := 1
	max := 0
	for _, v := range ps {
		if v.Age < min {
			min = v.Age
		}

		if v.Age > max {
			max = v.Age
		}
	}

	if max > 10 {
		return
	}

	//划分数据到10个桶里
	cap := (max-min)/10 + 1
	b0 := make([]People, 0, cap)
	b1 := make([]People, 0, cap)
	b2 := make([]People, 0, cap)
	b3 := make([]People, 0, cap)
	b4 := make([]People, 0, cap)
	b5 := make([]People, 0, cap)
	b6 := make([]People, 0, cap)
	b7 := make([]People, 0, cap)
	b8 := make([]People, 0, cap)
	b9 := make([]People, 0, cap)
	b10 := make([]People, 0, cap)

	for _, v := range ps {
		switch v.Age {
		case 0:
			b0 = append(b0, v)
			break
		case 1:
			b1 = append(b1, v)
			break
		case 2:
			b2 = append(b2, v)
			break
		case 3:
			b3 = append(b3, v)
			break
		case 4:
			b4 = append(b4, v)
			break
		case 5:
			b5 = append(b5, v)
			break
		case 6:
			b6 = append(b6, v)
			break
		case 7:
			b7 = append(b7, v)
			break
		case 8:
			b8 = append(b8, v)
			break
		case 9:
			b9 = append(b9, v)
			break
		case 10:
			b10 = append(b10, v)
			break
		}
	}

	fmt.Printf("b0:%v\n", b0)
	fmt.Printf("b1:%v\n", b1)
	fmt.Printf("b2:%v\n", b2)
	fmt.Printf("b3:%v\n", b3)
	fmt.Printf("b4:%v\n", b4)
	fmt.Printf("b5:%v\n", b5)
	fmt.Printf("b6:%v\n", b6)
	fmt.Printf("b7:%v\n", b7)
	fmt.Printf("b8:%v\n", b8)
	fmt.Printf("b9:%v\n", b9)
	fmt.Printf("b10:%v\n", b10)

}

//基数排序
func radixSort(ps []People) {
	mod := 10
	dev := 1
	digit := 10
	l := len(ps)
	for i := 0; i < digit; i++ {
		//这里用一个map的简单结构表示10个桶
		counter := [10][]People{}
		for j := 0; j < l; j++ {
			//电话号码转换为数字
			t, _ := strconv.Atoi(ps[j].Phone)
			bucket := t % mod / dev
			counter[bucket] = append(counter[bucket], ps[j])
		}

		mod *= 10
		dev *= 10

		pos := 0
		for _, v := range counter {
			for _, vv := range v {
				ps[pos] = vv
				pos++
			}
		}
	}

	printlnPeople(ps)

}

func phoneSort(ps []People) {
	l := len(ps)
	b0p := make([]People, 0, l)
	b1p := make([]People, 0, l)
	b2p := make([]People, 0, l)
	b3p := make([]People, 0, l)
	b4p := make([]People, 0, l)
	b5p := make([]People, 0, l)
	b6p := make([]People, 0, l)
	b7p := make([]People, 0, l)
	b8p := make([]People, 0, l)
	b9p := make([]People, 0, l)
	//按照手机号码第二位分成10个桶
	for i := 0; i < len(ps); i++ {
		b0 := ps[i].Phone[1] - '0'
		switch b0 {
		case 0:
			b0p = append(b0p, ps[i])
			break
		case 1:
			b1p = append(b1p, ps[i])
			break
		case 2:
			b2p = append(b2p, ps[i])
			break
		case 3:
			b3p = append(b3p, ps[i])
			break
		case 4:
			b4p = append(b4p, ps[i])
			break
		case 5:
			b5p = append(b5p, ps[i])
			break
		case 6:
			b6p = append(b6p, ps[i])
			break
		case 7:
			b7p = append(b7p, ps[i])
			break
		case 8:
			b8p = append(b8p, ps[i])
			break
		case 9:
			b9p = append(b9p, ps[i])
			break
		}
	}

	//对每个桶的数据按插入排序
	sortPeopleByPhone(b0p)
	sortPeopleByPhone(b1p)
	sortPeopleByPhone(b2p)
	sortPeopleByPhone(b3p)
	sortPeopleByPhone(b4p)
	sortPeopleByPhone(b5p)
	sortPeopleByPhone(b6p)
	sortPeopleByPhone(b7p)
	sortPeopleByPhone(b8p)
	sortPeopleByPhone(b9p)

	printPeople(b0p)
	printPeople(b1p)
	printPeople(b2p)
	printPeople(b3p)
	printPeople(b4p)
	printPeople(b5p)
	printPeople(b6p)
	printPeople(b7p)
	printPeople(b8p)
	printPeople(b9p)

}

func bucketSortTest() {
	len := 150
	ps := make([]People, len, len)
	for i := 0; i < len; i++ {
		ps[i] = People{
			Age:  rand.Intn(120),
			Name: fmt.Sprintf("toby_%d", i),
		}
	}
	bucketSort(ps)
}

func countSortTest() {
	len := 30
	max := 10
	ps := make([]People, len, len)
	for i := 0; i < len; i++ {
		ps[i] = People{
			Age:  rand.Intn(max),
			Name: fmt.Sprintf("toby_%d", i),
		}
	}
	countSort(ps)
}

func createPhone() string {
	phone := 10000000000 +
		rand.Intn(10)*1000000000 +
		rand.Intn(10)*100000000 +
		rand.Intn(10)*10000000 +
		rand.Intn(10)*1000000 +
		rand.Intn(10)*100000 +
		rand.Intn(10)*10000 +
		rand.Intn(10)*1000 +
		rand.Intn(10)*100 +
		rand.Intn(10)*10 +
		rand.Intn(10)
	return strconv.Itoa(phone)

}

func printPeople(obj []People) {
	j := len(obj)
	for i := 0; i < j; i++ {
		fmt.Printf("%v", obj[i])
	}
	fmt.Println()
}

func printlnPeople(obj []People) {
	j := len(obj)
	for i := 0; i < j; i++ {
		fmt.Printf("%v\n", obj[i])
	}
}

func phoneSortTest() {
	len := 32
	max := 150
	ps := make([]People, len, len)
	for i := 0; i < len; i++ {
		ps[i] = People{
			Age:   rand.Intn(max),
			Name:  fmt.Sprintf("toby_%d", i),
			Phone: createPhone(),
		}
	}
	//phoneSort(ps)
	radixSort(ps)
}

func bubbleSortTest() {
	t := unorderedAry[:]
	bubbleSort(t)
	fmt.Printf("排序之后的元素:%v", t)
}

func selectSortTest() {
	t := unorderedAry[:]
	selectSort(t)
	fmt.Printf("排序之后的元素:%v", t)
}

func insertSortTest() {
	t := unorderedAry[:]
	insertSort(t)
	fmt.Printf("排序之后的元素:%v", t)
}

func mergeSortTest() {
	t := unorderedAry[:]
	t = mergeSort(t)
	fmt.Printf("排序之后的元素:%v", t)
}

/*
 * ty代表查找的方式：0 查找改值是否存在;1 查找给定值第一次出现的位置;2 查找给定值最后一次出现的位置 3 第一个大于给定值的位置；
 * 4最后一个小于给定值的位置
 */
func bSearchTest(n, ty int) int {
	t := unorderedAry[:]
	t = mergeSort(t)
	fmt.Printf("要排序的数据:%v\n", t)

	l := len(t)
	start := 0
	end := l - 1
	index := -1
	for start <= end {
		middle := start + (end-start)>>1
		if t[middle] == n {
			index = middle
			switch ty {
			case 0:
				break
			case 1:
				//往前找
				index--
				fmt.Println("fuck")
				for ; index >= start; index-- {
					if t[index] != n {
						goto Loop
					}
				}
			Loop:
				index++
				break
			case 2:
				//往后找
				index++
				for ; index <= end; index++ {
					if t[index] != n {
						goto Loo1
					}
				}
			Loo1:
				index--
				break
			case 3:
				//往后找
				index++
				hasFind := false
				for ; index <= end; index++ {
					if t[index] != n {
						hasFind = true
						goto Loo2
					}
				}
			Loo2:
				if !hasFind {
					index = -1
				}
				break
			case 4:
				//往后找
				index++
				hasFind := false
				for ; index <= end; index++ {
					if t[index] != n {
						hasFind = true
						newN := t[index]
						//往后找
						for ; index <= l-1; index++ {
							if t[index] != newN {
								index--
								goto Loo3
							}
						}
						goto Loo3
					}
				}
			Loo3:
				if !hasFind {
					index = -1
				}
				break
			}

			return index
		} else if t[middle] < n {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return index
}

func main() {
	fmt.Println(bSearchTest(9, 4))
}
