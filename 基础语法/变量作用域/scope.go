package main

import "fmt"

const (
	v int = 98
)

func main() {
	// 局部变量
	var name = "jack"
	fmt.Println("局部变量:", name)

	// 全局变量
	fmt.Println("全局变量:", v)

	// 形式变量
	var a int = 23
	var b = 34
	var c = 0
	c = sum(a, b)
	fmt.Println("形式变量:", c)
}

func sum(a int, b int) int {
	return a + b
}
