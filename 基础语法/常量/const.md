## 常量

常量是一个简单值的标识符，在程序运行时，不会被修改的量。

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量的定义格式：`const identifier [type] = value`

### 常量定义

Go的常量定义可以限定常量类型，但不是必需的。编译器可以根据变量的值来推断其类型。

- 显式类型定义： `const name string = "tom""`
- 隐式类型定义： `const age = 12`

常量还可以用做枚举：
```go
const (
    unkown = 0
    female = 1
    male = 2
)
```

### iota

iota，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

**1. iota 可以被用作枚举值：**
```go
const (
    a = iota
    b = iota
    c = iota
)
```
第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
```go
const (
    a = iota
    b
    c
)
```

**2. 根据定义推断**

可以简单理解为在一个const块中，每换一行定义个常量，iota 都会自动+1。
```go
package main

import "fmt"
const (
    i=1<<iota
    j=3<<iota
    k
    l
)

func main() {
    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}
```

得到：
```
i= 1
j= 6
k= 12
l= 24
```
iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。

简单表述:

* i=1：左移 0 位，不变仍为 1。
* j=3：左移 1 位，变为二进制 110，即 6。
* k=3：左移 2 位，变为二进制 1100，即 12。
* l=3：左移 3 位，变为二进制 11000，即 24。

> 注：<<n==*(2^n)。

iota的使用可以理解为默认数学值得加法，也可以根据其他运算来定义后续推导的规则。
