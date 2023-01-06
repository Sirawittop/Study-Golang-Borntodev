package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type cercle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c cercle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c cercle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println("area : ", g.area())
	fmt.Println("Periter : ", g.perim())
}
func main() {
	r := rect{width: 5, height: 10}
	c := cercle{radius: 15}
	measure(r)
	measure(c)
}
