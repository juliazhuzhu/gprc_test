package main

import "fmt"

type Programmer interface {
	Debug() string
}

type JavaDev struct {

}

func (j JavaDev) Debug() string{
	return "'javadev'"
}

func printInfo(x interface{})  {

	v, ok := x.(int)
	if  ok {
		fmt.Println(v)
	}
}

func print(x interface{}){

	switch v:=x.(type) {
	case string:
		fmt.Printf("%s string \n",v)

	case int:
		fmt.Printf("%d integer \n",v)
	}
}

func main()  {

	var p Programmer
	p = &JavaDev{}

	fmt.Println(p.Debug())
	printInfo(99)
	printInfo("xiaoye")

	print(2)
}
