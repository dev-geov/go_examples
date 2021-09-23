package main

import (
	"fmt"
	"math"
	"reflect"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * 2 * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("Geometry: ", g)
	fmt.Println("Type:", reflect.TypeOf(g))
	fmt.Println("Area:", g.area())
	fmt.Println("Perim:", g.perim())
}

func main() {
	r := rect{12, 22}
	c := circle{30}

	measure(r)
	measure(c)
}
