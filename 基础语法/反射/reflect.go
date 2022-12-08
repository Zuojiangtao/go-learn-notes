package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func Sprint(x interface{}) string {
	type stringer interface {
		String() string
	}
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	// ...similar cases for int16, uint32, and so on...
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map, pointer, slice, struct
		return "???"
	}
}

func main() {
	name := "tom"
	nums := 123131
	boo := true
	fmt.Println(Sprint(name), Sprint(nums), Sprint(boo))

	// 通过reflect.TypeOf()获取变量类型
	fmt.Println(reflect.TypeOf(name), reflect.TypeOf(nums), reflect.TypeOf(boo))

	// 通过reflect.ValueOf()获取变量值
	fmt.Println(reflect.ValueOf(name), reflect.ValueOf(nums), reflect.ValueOf(boo))
}
