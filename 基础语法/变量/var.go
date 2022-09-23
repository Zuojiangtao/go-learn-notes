package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var name string = "tom"
	var age = 23
	status := "success"
	var sex, level, time = "man", 1, "2022-1-1"

	fmt.Println(name, age, status, sex, level, time)

	//var x = nil // 不指定类型会报错

	//fmt.Println(x)
}
