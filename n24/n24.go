package main

import (
	"fmt"
	"math"
)

// структурка
type Point struct {
	x int
	y int
}

// констуктор структурки
func newPoint(k int, n int) *Point {
	p := &Point{
		x: k,
		y: n,
	}
	return p
}

func main() {
	var p1 Point
	var p2 Point

	//создаем точки, с помощью конструктора
	p1 = *newPoint(10, 10)
	p2 = *newPoint(5, 5)

	//находим дистанцию между ними
	fmt.Println(distance(p1, p2))
}

func distance(p1 Point, p2 Point) float64 {
	x := p1.x - p2.x
	y := p1.y - p2.y

	return math.Sqrt(float64(x*x + y*y))
}
