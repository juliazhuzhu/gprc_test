package main


func main() {

	//指针，slice, map 接口的默认类型nil
	//如何开始的时候就分配内存
	var p * int = new(int)
	*p = 10
	//make 更加常用， make 用于slice , map
	//new指针返回的是指针, make函数返回的是指定类型的实例
	var info map[string]string = make(map[string]string)
	info["c"] = "fuck"




}


