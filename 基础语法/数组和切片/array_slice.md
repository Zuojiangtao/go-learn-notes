## 数组和切片

数组在大多数编程语言中均有，就不再重点去看了，简单过下。重点看切片相关内容。

### 数组
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列

语法格式: `var variable_name [SIZE] variable_type`

> Go中数组是相同类型且长度固定的序列，这点要注意。

```go
package main

import "fmt"

func main() {
	// 数组
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [5]string{3: "tom", 4: "jack"} // 指定索引位置初始化
	var arr3 = new([10]int) // 关键字创建
	fmt.Println(arr1, arr2, arr3)
}
```

### 切片
Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

**切片（slice）** 是对底层数组一个连续片段的引用，所以切片是一个引用类型。切片提供对该数组中编号的元素序列的访问。未初始化切片的值为nil。


**优点**

因为切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中切片比数组更常用。
#### 定义切片
可以声明一个未指定大小的数组来定义切片： `var identifier []type`

切片不需要指定长度。

也可以使用make()函数来创建切片： `var slice1 []type = make([]type, len)`

##### 切片初始化

```go
s := []int{1,2,3}
```    

##### len() 和 cap()函数
* len() 方法获取长度。 _len(slice)_
* 切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。 _cap(slice)_

#### 空切片(nil)
一个切片在未初始化之前默认为 nil，长度为 0
```go
package main

import "fmt"

func main() {
	var nilSlice []int
	if nilSlice == nil {
		fmt.Println("切片nilSlice是空的")
	}
}
```

#### 切片截取
可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]

切片截取和js数组的slice()方法很像，左开右闭。截取方式： `slice3[1:3]`

#### append() 和 copy()函数
append用于往切片追加元素，其底层实现会判断切片容量，如果容量不足，则触发扩容。append通常有两种写法：1）追加一个切片到另一个切片；2）追加元素到一个切片。
```go
package main

import "fmt"

func main()  {
    slice := make([]int, 0, 0)
	slice1 := append(slice, 3) // 单个元素
	slice2 := append(slice, 10, 20, 30) // 添加多个元素
	slice3 := append(slice1, slice...) // 一个切片
	slice4 := append(slice2[:2], slice3[1:2]...) // 切片片段
	
	fmt.Println(slice, slice1, slice2, slice3, slice4)
}
```

copy()方法用于拷贝
```go
package main

import "fmt"

func main()  {
    slice := make([]int, 10, 0)
	slice1 := append(slice, 1 , 2, 3, 4)
	
	fmt.Println(slice, slice1) // 将slice1 copy到slice
}
```

#### 切片重组
通过改变切片长度得到新切片的过程称之为切片重组 reslicing，做法如下：slice1 = slice1[0:end]，其中 end 是新的末尾索引（即长度）。
```go
package main
import "fmt"
func get() []byte {  
    raw := make([]byte, 10000)
    fmt.Println(len(raw), cap(raw), &raw[0]) // 显示: 10000 10000 数组首字节地址
    return raw[:3]  // 10000个字节实际只需要引用3个，其他空间浪费
}
func main() {  
    data := get()
    fmt.Println(len(data), cap(data), &data[0]) // 显示: 3 10000 数组首字节地址
}
```
当我们在一个切片基础上重新划分一个切片时，新的切片会继续引用原有切片的数组。如果你忘了这个行为的话，在你的应用分配大量临时的切片用于创建新的切片来引用原有数据的一小部分时，会导致难以预期的内存使用。
为了避免这个陷阱，我们需要从临时的切片中使用内置函数copy()，拷贝数据（而不是重新划分切片）到新切片。
