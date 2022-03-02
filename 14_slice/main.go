package main

import "fmt"

func main() {

	var a int = 5
	fmt.Println(a)
	passByValue(&a)
	fmt.Println(a)

	slice := []int{1, 2}
	fmt.Println(slice)
	updateSlice(slice)
	fmt.Println(slice)
}

// if pass by ref then passed reference updated
func passByValue(p *int) {
	x := 61
	p = &x
}

// if pass by ref then passed reference updated
func updateSlice(a []int) {
	a = []int{3, 4}
}
