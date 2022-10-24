## 方法

在 Go 语言中，结构体就像是类的一种简化形式，那么面向对象程序员可能会问：类的方法在哪里呢？在 Go 语言中有一个概念，它和方法有着同样的名字，并且大体上意思相近。

Go 语言中方法和函数在形式上很像，它是作用在接收器（receiver）上的一个函数，接收器是某种类型的变量。因此方法是一种特殊类型的函数，方法只是比函数多了一个接收器（receiver），当然在接口中定义的函数我们也称为方法（因为最终还是要通过绑定到类型来实现）。

正是因为有了接收器，方法才可以作用于接收器的类型（变量）上，类似于面向对象中类的方法可以作用于类属性上。

定义方法的一般格式如下：
```go
func (recv receiver_type) methodName(param_list) (return_value_type) {
	// ...
}
```
在方法名之前，func 关键字之后的括号中指定接收器 receiver。

### 接收器
1. 接收器类型除了不能是指针类型或接口类型外，可以是其他任何类型，不仅仅是结构体类型，也可以是函数类型，还可以是 int、bool、string 等等为基础的自定义类型。
2. 接收器不能是一个接口类型，因为接口是一个抽象定义，但是方法却是具体实现；如果这样做会引发一个编译错误：invalid receiver type…。
3. 接收器不能是一个指针类型，但是它可以是任何其他允许类型的指针。
4. 接收器不能是一个指针类型，但可以是类型的指针。

### 方法值与方法表达式

和 JS使用函数类似，Go的方法也有使用方法赋值和方法表达式的2种方式。

方法值调用：将方法以“值”的形式传入或者赋值，然后使用。
```go
type Rocket Struct { /*...*/ }
func (r *Rocket) Launch() {
    /* ... */
}
r := new(Rocket)
time.AfterFunc(10 * time.Second, r.launch)
```

方法表达式：使用方法时，与调用一个普通函数相比，要用选择器语法指定方法的接收器。
```go
f1 := T.mv
f2(t, 5)

f2 := (T).mv
f2(t, 6)
```

这个函数值的第一个参数必须是一个接收器。

> 和JS种函数的使用方式一样，可以选择将一个函数计算赋值给一个变量，也可以将一个函数表达式赋值给变量。

当你根据一个变量来决定调用同一个类型的哪个函数时，方法表达式就显得很有用了。你可以根据选择来调用接收器各不相同的方法。

```go
type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
    var op func(p, q Point) Point
    if add {
        op = Point.Add
    } else {
        op = Point.Sub
    }
    for i := range path {
        // Call either path[i].Add(offset) or path[i].Sub(offset).
        path[i] = op(path[i], offset)
    }
}
```
