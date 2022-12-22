package main

import "fmt"

func main() {

	ret := twoValueSum[int](100, 200)
	fmt.Println(ret)

	ret1 := twoValueSum[string]("hello ", "world")
	fmt.Println(ret1)
}

func twoValueSum[T int | float64 | string](a T, b T) T {
	return a + b
}
