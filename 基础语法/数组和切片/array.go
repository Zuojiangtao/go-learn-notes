package main

import "fmt"

func main() {
	// 数组
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [5]string{3: "tom", 4: "jack"} // 指定索引位置初始化
	var arr3 = new([10]int)
	fmt.Println(arr1, arr2, arr3)
}
