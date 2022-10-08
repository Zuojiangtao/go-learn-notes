package main

import (
	"fmt"
	"time"
)

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

	// 修改key对应的值
	countryCapitalMap["England"] = "伦敦"
	fmt.Println("England首都修正为: ", countryCapitalMap["England"])

	// 删除键值对
	theCountry, ok := countryCapitalMap["England"]
	fmt.Println(theCountry, ok)
	delete(countryCapitalMap, "England")
	theCountry, ok = countryCapitalMap["England"]
	fmt.Println(theCountry, ok)

	testSync()
}

func testSync() {
	var m = make(map[string]int, 0)
	//创建10个协程
	for i := 0; i <= 10; i++ {
		go func() {
			//协程内，循环操作map
			for j := 0; j <= 100; j++ {
				m[fmt.Sprintf("test_%v", j)] = j
			}

		}()
	}
	//主协程休眠3秒，否则主协程结束了，子协程没有机会执行
	time.Sleep(time.Second * 3)
	fmt.Println(m)
}
