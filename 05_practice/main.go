package main

import "fmt"

func main() {
	// x := 5
	// y := 10

	// x, y = y, x

	// fmt.Println(x, y)

	x := 5
	s := string(65)
	if true {
		x = 10
		x++
		fmt.Println(x) // ascii code
	}
	fmt.Println(x)
	fmt.Printf("%T %v\n", s, s)
}
