package main

import (
	"fmt"
	"unsafe"
)

type Course struct {
	 Name	string
	 Price	int
	 Url 	string
}

func (c *Course) setPrice(price int){
	c.Price = price
}

func main() {

	c := Course{
		Name: "dj",
		Price: 100,
		Url: "www.hlgnet.com",
	}

	fmt.Println(c.Name)

	var c7 *Course = new(Course)

	fmt.Println(c7.Name)

	fmt.Println(unsafe.Sizeof(c))

	//切片sizeof始终是24个字节，他是一个结构，由一个指针和两个整数构成
	s1 := []string{"django", "c", "java"}

	fmt.Println(unsafe.Sizeof(s1))


	var cc *Course = new (Course)//

	cc.Price = 100

	cc.setPrice(200)
	fmt.Println(cc.Price)

}
