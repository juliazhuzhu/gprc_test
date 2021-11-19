package main

import (
	"fmt"
	"sort"
)
type Course struct {
	Name string
	Price int
	Url string
}

type Courses []Course


func (c Courses) Len() int {
	return len(c)
}

func (c Courses) Less(i, j int) bool {
	return c[i].Price < c[j].Price
}

func (c Courses) Swap(i,j int)  {
	c[i], c[j] = c[j], c[i]
}

func main() {

	courses := Courses{
		Course{"Python", 11, ""},
		Course{"Java", 15, ""},
		Course{"C++", 13, ""},
	}
	sort.Sort(courses)
	for _, v := range courses {
		fmt.Println(v)
	}
}