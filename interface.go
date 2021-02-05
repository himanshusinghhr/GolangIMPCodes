package main

import (
	"fmt"
	"math"
)

// var i=0;
// type

type rectangle struct {
	Length float32
	Width  float32
}

type circle struct {
	Radius float32
}

func (r rectangle) area() float32 {
	return r.Length * r.Width
}
func (c circle) area() float32 {
	return math.Pi * math.Pi * float32(c.Radius)
}
func (r rectangle) perimeter() float32 {
	return 2 * (r.Length + r.Width)
}
func (c circle) perimeter() float32 {
	return 2 * math.Pi * float32(c.Radius)
}

type geometry interface {
	area() float32
	perimeter() float32
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	rect := rectangle{3, 3}
	fmt.Println(rect.area())
	circ := circle{3}
	fmt.Println((circ.area()))
	rectper := rectangle{3, 3}
	fmt.Println(rectper.perimeter())
	circper := circle{3}
	fmt.Println((circper.perimeter()))

	measure(rect)

}
