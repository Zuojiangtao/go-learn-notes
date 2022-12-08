## 反射

Go语言提供了一种机制，能够在运行时更新变量和检查它们的值、调用它们的方法和它们支持的内在操作，而不需要在编译时就知道这些变量的具体类型。这种机制被称为反射。

反射使得Go语言具备一些动态特性，比如不知道参数类型怎么办？当然你可以定义多个函数，分别传递不同参数；你也可以定义一个函数就行，参数类型为interface{}，函数内通过反射操作变量。

> 就是语言层面的switch...case...语句，用于不确定类型值的'兜底'处理。

### 为何需要反射
在处理一类不满足定义接口类型的值或者未知数据类型的时候，需要检查未知类型的表示方式。因此需要`反射`来对未知类型响应的处理。

### reflect.Type
reflect.Type，表示Go语言类型，这是一个接口，定义了很多方法，可以帮助我们获取到该类型拥有的属性、方法，占用内存大小等等。

```go
t := reflect.TypeOf(3)  // a reflect.Type
fmt.Println(t.String()) // "int"
fmt.Println(t)          // "int"
```

### reflect.Value
reflect.Value，表示Go语言中的值，其不仅包含变量的值，还包含该变量的类型信息。

```go
v := reflect.ValueOf(3) // a reflect.Value
fmt.Println(v)          // "3"
fmt.Printf("%v\n", v)   // "3"
fmt.Println(v.String()) // NOTE: "<int Value>"
```
