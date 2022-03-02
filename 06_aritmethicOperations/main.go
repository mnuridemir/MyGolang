package main

import "fmt"

func main() {
	x, y := 10, 5

	fmt.Printf("%T , %v\n", x, x)
	fmt.Printf("%T , %v\n", y, y)
	fmt.Printf("%T , %v\n", (x / y), (x / y))

	z := (5 % 2)

	fmt.Printf("%T , %v\n", z, z)

	w := 10
	w++ // omly postfix is validate
	fmt.Println(w)
	w--
	fmt.Println(w)

	const c = 3 // const ---> compile time    var ---> run time
	fmt.Printf("%T, %v", c, c)

}
