package main

import "fmt"

func main() {
	var a int = 100 // 声明实际变量

	fmt.Println("变量a的地址：", &a)

	var ip *int // 声明一个指针
	ip = &a

	fmt.Println("ip变量存储的指针地址：", ip)
	fmt.Println("*ip变量的值：", *ip)

	//const b = 12 // 常量不可以获取指针地址
	//t := &b
	//
	//fmt.Println(t)

	var ptr *int

	fmt.Println("空指针展示：", ptr)
}
