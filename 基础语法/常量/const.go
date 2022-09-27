package main

import "fmt"

func main() {
	// 常量定义
	const name = "tom"
	const age int = 12
	const name2, num, bool = "jack", 23, false

	fmt.Println(name, age, name2, num, bool)

	// iota
	const (
		a = iota     // 0
		b            // 1
		c            // 2
		d = "hahaha" // 独立值，iota += 1
		e            // "hahaha", iota += 1
		f = 100      // iota += 1
		g            // 100
		h = iota     // 7， 恢复计数
		i            // 8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
