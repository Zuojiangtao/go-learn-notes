## 流程控制

条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。

### if语句

* if 语句由布尔表达式后紧跟一个或多个语句组成。

```go
package main

import "fmt"

func main() {
	// if条件语句
	var a int = 10
	if a < 20 {
		fmt.Println("a 小于 20")
	}
	fmt.Println("a的值为:", a)
}
```

### switch语句

* switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。

* switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。

* switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。

```go
package main

import "fmt"

func main() {
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
}
```

### for循环语句

最简单的基于计数器的迭代，基本形式为：

```go
for  初始化语句; 条件语句; 修饰语句 {}
```

```go
package main

import "fmt"

func main() {
	// for循环语句
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("偶数", i)
		}
	}
}
```

### select语句

* select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。

* select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// select语句
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	case <-time.After(time.Second * 3): //超时退出
		fmt.Println("request time out")
	}
}
```

### for-range 结构

* for-range 结构是 Go 语言特有的一种迭代结构，它在许多情况下都非常有用。它可以迭代任何一个集合，包括数组（array）和字典（map），同时可以获得每次迭代所对应的索引和值。一般形式为：

```go
for ix, val := range coll { }
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// for-range结构
	data := []string{"one", "two", "three"}
	for _, v := range data {
		go func(in string) {
			fmt.Println(in)
		}(v)
	}
	time.Sleep(3 * time.Second)
}
```

> 注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断。

### 关键字

**break**

一个 break 的作用范围为该语句出现后的最内部的结构，它可以被用于任何形式的 for 循环（计数器、条件判断等）。
但在 switch 或 select 语句中，break 语句的作用结果是跳过整个代码块，执行后续的代码。

**continue**

关键字 continue 忽略剩余的循环体而直接进入下一次循环的过程，但不是无条件执行下一次循环，执行之前依旧需要满足循环的判断条件。
关键字 continue 只能被用于 for 循环中。

**label**

for、switch 或 select 语句都可以配合标签（label）形式的标识符使用，即某一行第一个以冒号（:）结尾的单词（Gofmt 会将后续代码自动移至下一行）
（标签的名称是大小写敏感的，为了提升可读性，一般建议使用全部大写字母）
continue 语句指向 LABEL1，当执行到该语句的时候，就会跳转到 LABEL1 标签的位置。

使用标签和 Goto 语句是不被鼓励的：它们会很快导致非常糟糕的程序设计，而且总有更加可读的替代方案来实现相同的需求。