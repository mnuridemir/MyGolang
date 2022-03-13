package main

import "fmt"

func main() {
	// x := 5
	// y := 10

	// x, y = y, x // this is valid

	// fmt.Println(x, y)

	x := 5
	s := string(rune(65))
	if true {
		x = 10
		x++
		fmt.Println(x) // ascii code
	}
	fmt.Println(x)
	fmt.Printf("%T %v\n", s, s)

	str := "lorem ipsum dolor amet"

	str_1 := str[:5] // first five char

	str_2 := str[len(str)-4:] // last four char

	str_3 := fmt.Sprintf("string -> %s -- first five char -> %s -- last four char -> %s \n", str, str_1, str_2)

	fmt.Println(str_3)

}
