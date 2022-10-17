package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	var a = 12
	var b = 24
	var res int
	res = max(a, b)

	fmt.Println("最大值是: ", res)

	x, y := swap("这本书是：", "《三体》")

	fmt.Println(x, y)

	var md5func = MD5("12345")

	fmt.Println("md5:", md5func)

	times := getTimeInt()

	fmt.Println("获取时间戳:", times)
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return x, y
}

// MD5
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 获取时间戳
func getTimeInt() int64 {
	return time.Now().Unix()
}
