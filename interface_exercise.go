package main

import "fmt"
import "math"

type Shape interface {
	getArea() float64
}

func printArea(shape Shape) {
	fmt.Printf("the area is %v cm\u00b2\n", shape.getArea())
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rec Rectangle) getArea() float64 {
	return rec.Width * rec.Height
}

type Circle struct {
	Radius float64
}

func (cir Circle) getArea() float64 {
	return math.Pi * math.Pow(cir.Radius, 2)
}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 7}

	printArea(r) // Rectangle Area: 50
	printArea(c) // Circle Area: 153.94 (approx)

}
