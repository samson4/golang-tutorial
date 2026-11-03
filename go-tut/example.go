package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Reactangle struct {
	width  float64
	height float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {

	fmt.Println("hello world")

}

// circle1 := Circle{Radius: 5}
// rectangle1 := Reactangle{width: 4, height: 6}
// ci
