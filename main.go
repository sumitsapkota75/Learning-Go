package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rect struct {
	height float64
	width  float64
}

func (r *Rect) area() float64 {
	return r.height * r.width
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	c1 := &Circle{4.6}
	r1 := &Rect{2, 3}
	shapes := []shape{r1, c1}

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
