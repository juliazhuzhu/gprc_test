package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

type sub func(a, b int) int // sub 是一个函数类型

func subImpl(a, b int) int {

	return a - b

}

func filter(score []int, f func(int) bool) []int {

	reSlice := make([]int, 0)
	for _, v := range score {
		if f(v) {
			reSlice = append(reSlice, v)
		}
	}

	return reSlice
}

func main() {
	//函数作为变量被复制赋值
	myFunc := add
	fmt.Printf("%T \n", myFunc)

	fmt.Println(myFunc(1, 3))

	//匿名函数赋值
	myFunc2 := func(a, b int) int {
		return a - b
	}

	fmt.Println(myFunc2(3, 2))
	//匿名函数直接调用
	result := func(a, b int) int {
		return a * b
	}(1, 3)

	fmt.Printf("%T \n", result)

	var mySub1 sub = subImpl

	var mySub2 sub = func(a, b int) int {
		return a - b
	}

	mySub1(1, 2)
	mySub2(2, 1)

	//函数作为另外一个函数的参数
	score := []int{10, 20, 60, 70}

	standard := func(level int) bool {
		return level >= 60
	}

	fmt.Println(filter(score, standard))

}
