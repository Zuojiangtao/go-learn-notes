## 结构体

Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
结构体是由一系列具有*相同类型或不同类型*的数据构成的数据集合。

与C++,Java有所不同，Go没有class的概念，只有结构体struct，其可以拥有属性和方法，我们可以利用这一特性来实现面向对象编程。

### 定义结构体

结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

```go
type Student struct {
	member definition
	member definition
	...
}
```

#### 匿名字段
对于匿名字段，**必须将匿名字段指定为类型名称或指向非接口类型名称 *T的指针，并且T可能不是指针类型。**
```go
struct {
	T1          // 字段名 T1
	*T2       // 字段名 T2
	P.T3      // 字段名 T3
	*P.T4     // f字段名T4
    x, y int    // 字段名 x 和 y
}
```

#### 使用new函数创建结构体
和其他语言一样，Go也可以使用关键字new来实例化一个结构体，并分配内存：
```go
type Car struct {
	name string
	price int
}

func main()  {
    my_car := new(Car)  // new函数声明并初始化一个实例
	
	var your_car Car    // 这样声明的变量是类型T
	
	// 在这两种方式中，t 通常被称做类型 T 的一个实例（instance）或对象（object）。
}
```

#### 初始化结构体
```go
intr := Car{"大众", 200000} // 按照字段顺序初始化
intr := Car{name: "长城", price: 190000} // 显示化声明
intr := Car{name: "奥迪"} // 也可以只声明部分字段
```

#### 命名冲突
当结构体的字段一样时，外层名字覆盖内层名称；当同级冲突时，会报错，不过可以采用逐级选择使用的方式来避免这个错误。
```go
type A struct {
	name string
}
type B struct {
	name string
}

type C struct {
	A
	B
}

var c C
```
上面如果直接使用c.name时会报错的，但是可以c.A.name的形式就ok了。

### 访问结构体成员

如果要访问结构体成员，需要使用点号 . 操作符，格式为：

```go
package main

import "fmt"

type Book struct {
	title string
	author string
	book_id int
}

func main()  {
    var book1 Books
	
	book1.title = "Go学习笔记"
	book1.author = "zuojt"
	book1.book_id = 12938139439
	
	fmt.Println(book1.title)
}
```

### 结构体特性

#### 内存布局
Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。

#### 递归结构体
递归结构体类型可以通过引用自身指针来定义。这在定义链表或二叉树的节点时特别有用，此时节点包含指向临近节点的链接。例如：

```go
type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element
	// The list to which this element belongs.
	list *List
	// The value stored with this element.
	Value interface{}
}
```

#### 可见性
通过参考应用可见性规则，如果结构体名不能导出，可使用 new 函数使用工厂方法的方法达到同样的目的。例如：

```go
type bitmap struct {
	Size int
	data []byte
}
func NewBitmap(size int) *bitmap {
	div, mod := size/8, size%8
	if mod > 0 {
		div++
	}
	return &bitmap{size, make([]byte, div)}
}
```

#### 标签
Go的字段还有一个标签(tag),用于编辑文档和内容，但是只有reflect包能获取它。
reflect包可以在运行时反射得到类型、属性和方法。如变量是结构体类型，可以通过 Field() 方法来索引结构体的字段，然后就可以得到Tag 属性。

```go
package main

import (
    "fmt"
    "reflect"
)
type Student struct {
	name string "学生名字"
	age  int    "学生年龄"
}

func main() {
	stu := Student{"tom", 13}

	fmt.Println(reflect.TypeOf(stu).Field(0).Tag)
}
```

### 结构体作为函数参数

你可以像其他数据类型一样将结构体类型作为参数传递给函数。

```go
...

func printBook(book Book)  {
    fmt.Println(book1.title)
}
```

### 结构体指针

- 你可以定义指向结构体的指针类似于其他指针变量，格式如下：

```go
var struct_pointer *Book
```

- 以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：

```go
struct_pointer = &Book
```

### 结构体的继承

和其他语言的类一样，Go的结构体也可以继承属性和方法。

```go
package main

import "fmt"

type Human struct {
	name string
	age int
}

type Man struct {
	Human
	sex int
	score int
}

func main()  {
    var student Man
	student.score = 90
	student.age = 26
	
	fmt.Println(student.score, student.age)
}
```
