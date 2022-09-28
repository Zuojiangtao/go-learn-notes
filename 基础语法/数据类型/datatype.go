package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	// 变量声明
	const name string = "tom" // const 声明常量
	fmt.Println(name)

	const age int16 = 31
	fmt.Println(age)

	const name1, name2 string = "Tom", "Jack"
	fmt.Println(name1, name2)

	var count int16 = 1 // var 声明变量
	count = 3           // 更改变量值
	fmt.Println(count)

	// 字符串
	var str1 = "hello "
	str2 := "go"
	fmt.Println(str1 + str2)

	// 数组，容量固定
	var arr = [5]int{1, 2, 3, 4, 5}
	arr[1] = 100
	fmt.Println(len(arr), arr)

	// 切片，容量可扩充
	var slice = []int{1, 2, 3}
	slice[1] = 100
	slice = append(slice, 4, 5, 6)
	fmt.Println(len(slice), cap(slice), slice) // len返回切片长度，cap返回切片容量

	// map
	var score = map[string]int{
		"age":   32,
		"money": 100,
	}
	score["newVal"] = 100    // 新变量赋值
	s, ok := score["newVal"] // map访问，s对应value值，ok表示key是否存在（不存在返回空值）
	delete(score, "lang")    // 删除元素
	fmt.Println(s, ok, score)
}
