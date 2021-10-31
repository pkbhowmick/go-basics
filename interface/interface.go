// A simple example of interface in Golang:

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	GetArea() float64
}

type Rectangle struct {
	Length float64
	Height float64
}

func NewRectangle(l, h float64) *Rectangle {
	return &Rectangle{
		Length: l,
		Height: h,
	}
}

func (r *Rectangle) GetArea() float64 {
	return r.Height * r.Length
}

type Circle struct {
	Radius float64
}

func NewCircle(r float64) *Circle {
	return &Circle{
		Radius: r,
	}
}

func (c *Circle) GetArea() float64 {
	return c.Radius * c.Radius * math.Pi
}

var _ Shape = &Rectangle{}
var _ Shape = &Circle{}

func main() {
	rec := NewRectangle(10, 6)
	crl := NewCircle(2.5)

	shapes := []Shape{rec, crl}

	for _, o := range shapes {
		fmt.Printf("Interface value: %v and type: %T\n", o, o)
		fmt.Println(o.GetArea())
	}

}

// Output:
// Interface value: &{10 6} and type: *main.Rectangle
// 60
// Interface value: &{2.5} and type: *main.Circle
// 19.634954084936208
