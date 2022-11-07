## 接口

和其它语言一样，Go也有接口的概念。

很多面向对象的语言都有相似的接口概念，但Go语言中接口类型的独特之处在于它是满足隐式实现的。也就是说，我们没有必要对于给定的具体类型定义所有满足的接口类型；简单地拥有一些必需的方法就足够了。这种设计可以让你创建一个新的接口类型满足已经存在的具体类型却不会去改变这些类型的定义；当我们使用的类型来自于不受我们控制的包时这种设计尤其有用。

本节是接口的入门，细节和高阶使用在后续章节中学习。

### 接口类型

Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例。


```go
/** 定义接口 */
type interface_name interface {
	methods_name1 [return type]
	methods_name2 [return type]
	methods_name3 [return type]
	...
}
/** 定义结构体 */
type struct_name struct {
    /* variables */
}
/* 实现接口方法 */
func (struct_name_variable struct_name) method_name1() [return_type] {
    /* 方法实现 */
}
...
func (struct_name_variable struct_name) method_namen() [return_type] {
    /* 方法实现*/
}
```

### 空接口

Go语言将接口分为两种：带方法的接口，一般比较复杂，用iface表示；不带方法的接口也就是空接口，一般当我们不知道变量类型时，会声明变量类型为空接口（interface{}），其余类型可以转化为空接口类型。将某一类型变量转化为空接口时，依然需要维护原始变量类型，以及数据，Go语言用eface表示空接口变量，定义如下：

```go
type eface struct {
	_type *type
	data unsafe.Pointer
}
```
