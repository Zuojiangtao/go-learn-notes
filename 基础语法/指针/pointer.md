## 指针

我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。

Go 语言的取地址符是 **&**，放到一个变量前使用就会返回相应变量的内存地址。

以下实例演示了变量在内存中地址：
```go
package main

import "fmt"

func main() {
	var a = 100

	fmt.Println("变量a的地址：", &a)
}
```

得到： 变量a的地址： 0xc00001a098

### 指针定义

一个指针变量指向了一个值的内存地址。

类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：

```go
var ip *int /* 指向整型 */
var fq *float /* 指向浮点型 */
```

**指针的使用方法：**

* 定义指针变量；

* 为指针变量赋值；

* 访问指针变量中指向地址的值；

* 在指针类型前面加上\*号来获取指针所指向的内容。

在Go语言中，指针类型表示指向给定类型（称为指针的基础类型）的变量的所有指针的集合。 符号 \* 可以放在一个类型前，如 \*T，那么它将以类型T为基础，生成指针类型\*T。未初始化指针的值为nil。
```go
var a int = 100 // 声明实际变量

	fmt.Println("变量a的地址：", &a)

	var ip *int // 声明一个指针
	ip = &a

	fmt.Println("ip变量存储的指针地址：", ip)
	fmt.Println("*ip变量的值：", *ip)
```

**注意：不能得到一个数字或常量的地址**，下面的写法是错误的：

```go
const i = 5
ptr := &i // error: cannot take the address of i
ptr2 := &10 // error: cannot take the address of 10
```

> 虽然Go 语言和 C、C++ 这些语言一样，都有指针的概念，但是指针运算在语法上是不允许的。这样做的目的是保证内存安全。从这一点看，Go 语言的指针基本就是一种引用。

如果代码在运行中需要占用大量的内存，或很多变量，或者两者都有，这时使用指针会减少内存占用和提高运行效率。被指向的变量保存在内存中，直到没有任何指针指向它们。所以从它们被创建开始就具有相互独立的生命周期。

内存管理中的内存区域一般包括堆内存（heap）和栈内存（stack）， 栈内存主要用来存储当前调用栈用到的简单类型数据，如string，bool，int，float 等。这些类型基本上较少占用内存，容易回收，因此可以直接复制，进行垃圾回收时也比较容易做针对性的优化。 而复杂的复合类型占用的内存往往相对较大，存储在堆内存中，垃圾回收频率相对较低，代价也较大，因此传引用或指针可以避免进行成本较高的复制操作，并且节省内存，提高程序运行效率。

### 空指针

当一个指针被定义后没有分配到任何变量时，它的值为 nil。

nil 指针也称为空指针。

nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

一个指针变量通常缩写为 ptr。

```go
package main

import "fmt"

func main() {
   var  ptr *int

   fmt.Printf("ptr 的值为 : %x\n", ptr  )
}
```
以上实例输出结果为：

ptr 的值为 : 0

**空指针判断**：

if(ptr != nil)     /* ptr 不是空指针 */
if(ptr == nil)    /* ptr 是空指针 */
