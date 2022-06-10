package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Info struct {
	Name string
	Age int	`json:"age,omitempty"`
	Gender string `json:"-"`//-表示不需要参与序列化

}


type CustomInfo struct {
	Name string	`custom:"name, max_length=17, min_length=5"`
	Age int	`custom:"age, min=18, max=70"`
	Gender string `custom:"gender, required"`
}



func main()  {
	//结构体标签
	//比如我们在解析/生成json文件的时候，常用到的encoding/json包
	//例如: omitempty标签在序列化的时候忽略0值或者空值
	//

	info := Info {
		Name: "boddy",
		Gender: "male",
	}
	re, _ := json.Marshal(info)
	fmt.Println(string(re))

	//通过反射包来识别这些tag
	custom_info := CustomInfo {
		Name: "boddy",
		Gender: "male",
	}
	t := reflect.TypeOf(custom_info)
	fmt.Println(t.Name())
	fmt.Println(t.Kind())

	for i:=0 ; i<t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("custom")
		fmt.Printf("%d. %v (%v), tag: '%v' \n", i+1, field.Name, field.Type.Name(), tag)
	}
}
