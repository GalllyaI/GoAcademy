package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (s Square) NewSquare(side float64) Square {

	return Square{side}
}

func (sq *Square) Area() float64 {
	return sq.side * sq.side
}

type Circle struct {
	radius float64
}

func (c *Circle) NewCircle(radius float64) Circle {

	return Circle{radius}
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	sideA float64
	sideB float64
}

func (r *Rectangle) NewRectangle(sideA, sideB float64) Rectangle {

	return Rectangle{sideA, sideB}
}

func (r *Rectangle) Area() float64 {
	return r.sideA * r.sideB
}

type Shape interface {
	Area() float64
}

type Shapes []Shape

func (s Shapes) LargestArea() float64 {

	var maxArea float64 = 0

	for _, shape := range s {
		temp := shape.Area()
		if temp > maxArea {
			maxArea = temp
		}
	}
	return maxArea
}

func main() {

	var sq Square
	sq = sq.NewSquare(3.2)
	sq2 := sq.NewSquare(6)

	var circle Circle
	circle = circle.NewCircle(2)

	sh := Shapes{&sq, &sq2, &circle}

	var r Rectangle
	r = r.NewRectangle(5, 5)

	sh = append(sh, &r)

	result := sh.LargestArea()

	fmt.Println(result)
}
