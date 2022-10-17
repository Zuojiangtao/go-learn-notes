## 函数

函数是基本的代码块，用于执行一个任务。

Go 语言最少有个 main() 函数。

你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。

函数声明告诉了编译器函数的名称，返回类型，和参数。

Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

本章节学习简单的函数入门，后续会有专门的章节来学习函数。

### 函数定义

Go 语言函数定义格式如下:
```go
func function_name([params list]) [return_types] {
	函数体
}
```

* 函数定义解析：

  * func：函数由 func 开始声明
  * function_name：函数名称，参数列表和返回值类型构成了函数签名。
  * parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
  * return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
  * 函数体：函数定义的代码集合。

### 函数返回多个值
Go 函数可以返回多个值，例如:
```go
func ([params list]) ([return_types list]) {
	函数体
}
```

### 函数参数

函数如果使用参数，该变量可称为函数的形参。

形参就像定义在函数体内的局部变量。

调用函数，可以通过两种方式来传递参数：

1. 值传递: 值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
2. 引用传递: 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

> 默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

### 函数使用
1. 一个函数作为另一个函数的参数
```go
package main

import (
	"fmt"
	"math"
)

func main()  {
  getSquareRoot := func(x float64) float64 {
	  return math.Sqrt(x)
  }
  fmt.Println(getSquareRoot(9))
}
```
2. 闭包

```go
package main

import "fmt"

func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i  
   }
}

func main(){
   /* nextNumber 为一个函数，函数 i 为 0 */
   nextNumber := getSequence()  

   /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   
   /* 创建新的函数 nextNumber1，并查看结果 */
   nextNumber1 := getSequence()  
   fmt.Println(nextNumber1())
   fmt.Println(nextNumber1())
}
```
