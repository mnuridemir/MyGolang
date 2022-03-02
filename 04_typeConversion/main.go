package main

import "fmt"

func main() {
	x := 10
	y := 10.0

	fmt.Printf("%v  %T\n", x, x)
	fmt.Printf("%v  %T\n", y, y)

	z := x + int(y) // y is int for only this line

	fmt.Printf("%v  %T\n", z, z)
	fmt.Println(z)

	var num1 uint16 = 256
	var num2 uint8

	num2 = uint8(num1)

	fmt.Println(num2)

	a := 10.9
	fmt.Println(5 + int(a))
}
