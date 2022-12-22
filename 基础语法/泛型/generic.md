## 泛型

Golang在1.18版本支持了泛型，写过java/c++等语言的可能对泛型有一定的了解。那么泛型到底是什么呢？他有什么作用呢

### 为什么需要泛型
为什么需要泛型呢？Golang是强类型语言，任何变量或者函数参数，都需要定义明确的参数类型。假设我们需要实现这么一个函数，输入两个参数，函数返回其相加的值，输入参数可以是两个整型int，浮点数float，还有可能是字符串等等，这时候通常怎么办？定义多个函数实现吗？如下面程序所示：

```go
//定义多个函数实现
func twoIntValueSum(a, b int) int {
    return a + b
}

func twoFloatValueSum(a, b float32) float32 {
    return a + b
}

func twoStrValueSum(a, b string) string {
    return a + b
}

//定义一个函数，类型是interface{}
```

这会导致大量重复代码，而且调用方好需要根据参数类型决定先调用哪种方法。那怎么办？定义一个函数，知识参数是interface{}:

```go
func twoValueSum(a, b interface{}) (interface{}, error) {
    if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
        return nil, errors.New("two value type different")
    }
    
    switch reflect.TypeOf(a).Kind() {
    case reflect.Int:
        return reflect.ValueOf(a).Int() + reflect.ValueOf(b).Int(), nil
    case reflect.Float64:
        return reflect.ValueOf(a).Float() + reflect.ValueOf(b).Float(), nil
	case reflect.String:
        return reflect.ValueOf(a).String() + " " + reflect.ValueOf(b).String(), nil
	default:
        return nil, errors.New("unknow value type")
    }
}
```

使用反射实现的话，依赖反射性能较低，二来可以看到输入参数和返回值都是interface{}，使用方还需要多执行一步返回值类型转换，而且反射相对而言还是比较复杂的。

### 泛型初体验
泛型相当于定义了一个函数模板，真正调用函数的时候，再确定参数以及返回值等具体类型，基于泛型实现上述功能如下：

```go
func generic[T int | string | bool] (a T, b T) T { /* ... */}
```

泛型类型或者泛型函数定义的语法格式可以描述为[Identifier TypeConstraint]，上述程序中的T就是标识符（Identifier），int等就是TypeConstraint（类型限制，也就是说twoValueSum函数的输入参数类型只能是这几种，不能是其他的），注意在调用具体函数时，需要声明真正的类型。

介绍下泛型类型常见定义方式:

```go
// 定义切片类型，元素类型可以是int，float64或者string
type Slice[T int|float64|string ] []T
// 实例化变量arr1，注意声明了切片元素类型为int
var arr1 Slice[int] = []int{1, 2, 3} 
// 实例化变量arr2，注意声明了切片元素类型为string
var arr2 Slice[string] = []string{"Hello", "World"} 

// 自定义map类型，key只能是string，value可以是int、float32或者float64
type DefineMap[KEY string, VALUE int | float32 | float64] map[KEY]VALUE  
var m DefineMap[string, int] = map[string]int{
    "zhangsan": 89,
    "lisi":  80,
}
```
