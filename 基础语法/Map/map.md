## Map

Map 是一种无序的键值对的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

### 定义Map
可以使用内建函数 make 也可以使用 map 关键字来定义 Map:

```go
var map_var map[key_type]value_type

map_var2 := make(map[key_type]value_type)
```

> 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对。

也可以直接在定义的时候初始化。
```go
ages := map[string]int{
    "alice":   31,
    "charlie": 34,
}
```

### map方法

#### 可直接用覆盖的方式修改map值。
```go
score := map[string]int{
	"age": 20
}

score["age"] = 32
```

#### 使用delete方法删除map键值对。
```go
score := map[string]int{
	"alice": 23
	"charlie": 34
}

delete(score["alice"]) // remove element ages["alice"]
```

#### 迭代/遍历
```go
func main() {
    var countryCapitalMap map[string]string     // 关键字创建一个map
    countryCapitalMap = make(map[string]string) // make 初始化map
    
    // k-v存储
    countryCapitalMap["France"] = "巴黎"
    countryCapitalMap["Italy"] = "罗马"
    countryCapitalMap["Japan"] = "东京"
    countryCapitalMap["China"] = "北京"
    countryCapitalMap["England"] = "库斯科"
    
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是: ", countryCapitalMap[country])
    }
}
```

> map遍历顺序随机。
> 如果要按顺序遍历key/value对，我们必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序。下面是常见的处理方式：

```go
import "sort"

var names []string
for name := range ages {
    names = append(names, name)
}
sort.Strings(names)
for _, name := range names {
    fmt.Printf("%s\t%d\n", name, ages[name])
}

```

#### 指针取址问题
```go
_ = &ages["bob"] // compile error: cannot take address of map element
```
禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。

#### 零值
map类型的零值是nil，也就是没有引用任何哈希表。
```go
var times map[string]int

fmt.Println(times == nil) // "true"
fmt.Println(len(times) == 0) // "true"
```

map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。但是向一个nil值的map存入元素将导致一个panic异常：

```go
ages["carol"] = 21 // panic: assignment to entry in nil map
```

#### 判断是否存在key对应的值
```go
var age, ok := times["name"]
if !ok {/* // void */}
```
在这种场景下，map的下标语法将产生两个值；第二个是一个布尔值，用于报告元素是否真的存在。布尔变量一般命名为ok，特别适合马上用于if条件判断部分。

#### map之间比较
和slice一样，map之间也不能进行相等比较；唯一的例外是和nil进行比较。要判断两个map是否包含相同的key和value，我们必须通过一个循环实现：
```go
func equal(x, y map[string]int) bool {
    if len(x) != len(y) {
        return false
    }
    for k, xv := range x {
        if yv, ok := y[k]; !ok || yv != xv {
            return false
        }
    }
    return true
}
```

### 容量
和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。但是你也可以选择标明 map 的初始容量 capacity，就像这样：`make(map[keytype]valuetype，cap)`。

当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。

### 并发问题
我们都知道Go语言具有天然的并发优势，go func可以很方便的创建一个并发协程。那么，假如有一个全局的map，在多个协程中并发操作map呢？可以这么操作吗？我们写个小程序测试一下：
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    var m = make(map[string]int, 0)
    //创建10个协程
    for i := 0; i <= 10; i ++ {
        go func() {
            //协程内，循环操作map
            for j := 0; j <= 100; j ++ {
                m[fmt.Sprintf("test_%v", j)] = j
            }

        }()
    }
    //主协程休眠3秒，否则主协程结束了，子协程没有机会执行
    time.Sleep(time.Second * 3)
    fmt.Println(m)
}
```
运行之前可以猜一下执行结果。程序panic了，异常信息为："fatal error: concurrent map writes"。很明确，并发写map导致的panic，也就是说，我们不能在多个协程并发执行map写操作，这一点要特别注意，初次写Go语言很容易犯这个错。

那假如确实想在多个协程并发操作map怎么办？这就需要采取其他方案了，比如加锁，比如使用并发安全map（sync.Map），这些将在后面章节并发编程详细介绍。
