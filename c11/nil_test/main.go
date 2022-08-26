package main

import "fmt"

func main() {

	var p *int = new(int)
	*p = 10

	fmt.Println(*p)

	//make, slcie, map
	////new 函数返回的是指针，make函数返回的是指定类型的实例
	var info map[string]string = make(map[string]string)
	info["fuck"] = "arjing"

	for k, v := range info {
		fmt.Printf("%s %s\n", k, v)
	}

	var info2 map[string]string
	if info2 == nil {
		fmt.Println("map default value is nil")
	}

	var slicie []string
	if slicie == nil {
		fmt.Println("slcie is nil")
	}

	var err error
	if err == nil {
		fmt.Println("err default is nil")
	}
	fmt.Println("endtest")

}
