package main

import "fmt"

type Teacher struct {
	Name string
	Age	int
	Title string
}
func(t Teacher) teacherInfo() {
	fmt.Printf("name: %s", t.Name)
}

//匿名嵌套
type Course struct {
	Teacher
	Name	string
	Price	int
	Url 	string
}

func (c Course) courseInfo() {

}

func main()  {


}