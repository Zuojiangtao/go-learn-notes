package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// function
func distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// method
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {
	var p Point
	var q Point

	p.x = 12
	p.y = 23
	q.x = 5
	q.y = 6

	r := distance(p, q) // function
	t := p.Distance(q)  // method

	fmt.Println("使用函数得到结果：", r)
	fmt.Println("使用方法得到结果：", t)

	_distance := p.Distance // 方法值

	fmt.Println("方法值：", _distance(q))
}
