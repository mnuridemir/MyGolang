package main

import "fmt"

type shape interface {
	area() float64
}

func printArea(shapes ...shape) {
	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}

type triangle struct {
	a, h float64
}

type square struct {
	a float64
}

type rectangle struct {
	a, b float64
}

func (t triangle) area() float64 {
	return (t.a * t.h) / 2
}
func (s square) area() float64 {
	return s.a * s.a
}
func (r rectangle) area() float64 {
	return r.a * r.b
}

func main() {

	t := triangle{3, 4}
	s := square{4}
	r := rectangle{4, 5}

	printArea(t, s, r)

}
