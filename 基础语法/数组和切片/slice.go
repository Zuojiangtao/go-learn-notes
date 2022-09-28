package main

import "fmt"

var arr = [5]int{1, 2, 3, 4, 5}

func main() {
	// 定义及初始化
	var slice1 = []int{1, 2, 3}    // 直接初始化切片
	var slice2 = make([]string, 5) // 通过make函数创建
	slice3 := arr[:]               // 数组的引用
	slice4 := arr[1:3]             // 截取数组创建
	slice5 := arr[1]               // 从数组第一个开始
	slice6 := arr[:4]              // 一直到数组最后一个元素

	fmt.Println(slice1, slice2, slice3, slice4, slice5, slice6)
	fmt.Println("slice2, len=", len(slice2), "slice1, cap=", cap(slice1))

	// nil切片
	var nilSlice []int

	printSlice(nilSlice)

	if nilSlice == nil {
		fmt.Println("切片nilSlice是空的")
	}

	// 切片截取
	slice7 := slice3[1:3] // 左开右闭
	fmt.Println("切片slice3截取1-3:", slice7)

	// 扩容
	slice := make([]int, 3, 0)
	slice = append(slice, 10, 20, 30)

	fmt.Println("扩容:", slice)

	// 拷贝
	nums := make([]int, len(slice), (cap(slice))*2) // 创建nums是之前slice切片的2倍容量
	copy(nums, slice)

	fmt.Println("拷贝:", nums)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%V\n", len(x), cap(x), x)
}
