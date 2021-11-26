package main

import "fmt"

type Teacher struct {
	Name string `json:"age,omitempty"`
}

func (t *Teacher) teachInfo() {
	fmt.Println("hello teacher.")
}

type HighSchoolTeacher struct {
	Teacher
	Level int
}

func (ht *HighSchoolTeacher) highLevel() int {
	fmt.Printf("hello teacher %d.\n", ht.Level)
	return 1
}

func main() {

	ht := HighSchoolTeacher{}
	ht.Name = "li"
	ht.Level = 3
	ht.teachInfo()
	ht.highLevel()
}
