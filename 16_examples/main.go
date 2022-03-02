package main

import "fmt"

type rectangle struct {
	a, b int
}

func main() {
	r1 := rectangle{2, 5}
	r1.display()
	fmt.Println(r1.area())  // this is method
	fmt.Println(circum(r1)) // this is func
}

func (r rectangle) display() {
	fmt.Println("kısa kenar ", r.a, "uzun kenar ", r.b)
}
func (r rectangle) area() int {
	return r.a * r.b
}

func circum(r rectangle) int {
	return (r.a + r.b) * 2
}
