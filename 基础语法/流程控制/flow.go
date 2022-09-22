package main

import (
	"fmt"
	"time"
)

func main() {
	// if条件语句
	var a int = 10
	if a < 20 {
		fmt.Println("a 小于 20")
	}
	fmt.Println("a的值为:", a)

	// switch语句
	expr := "tom"
	switch expr {
	case "tom":
		fmt.Println("tom")
	case "jack":
		fmt.Println("jack")
	default:
		fmt.Println("name")
	}

	// for循环语句
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("偶数", i)
		}
	}

	// select语句
	//var c1, c2, c3 chan int
	//var i1, i2 int
	//select {
	//case i1 = <-c1:
	//	fmt.Printf("received ", i1, " from c1\n")
	//case c2 <- i2:
	//	fmt.Printf("sent ", i2, " to c2\n")
	//case i3, ok := (<-c3):
	//	if ok {
	//		fmt.Printf("received ", i3, " from c3\n")
	//	} else {
	//		fmt.Printf("c3 is closed\n")
	//	}
	//case <-time.After(time.Second * 3): //超时退出
	//	fmt.Println("request time out")
	//}

	// for-range结构
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func(in string) {
			fmt.Println(in)
		}(v)
	}
	time.Sleep(3 * time.Second)
}
