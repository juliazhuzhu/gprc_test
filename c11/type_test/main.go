package main

import "fmt"

type Course struct {
	name string
	price int
}

type Callable interface {

}

type handle func(str string)

func main()  {

	//1.给一个类型定义别名,增加代码可读性
	type myByte = byte

	var b myByte

	fmt.Printf("%T \n" , b)

	//2.基于一个已有类型定义一个新的类型
	type myInt int
	var i myInt
	fmt.Printf("%T \n" , i)

	//3.定义结构体
	//4.定义接口
	//5.定义函数别名
}
