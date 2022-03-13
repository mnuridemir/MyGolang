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

	mySlc := make([]int, 0, 3)
	mySlc = append(mySlc, 1)
	mySlc = append(mySlc, 1)
	mySlc = append(mySlc, 1)
	fmt.Printf("slc len : %d cap: %d \n", len(mySlc), cap(mySlc))
	mySlc = append(mySlc, 1) // if cap is more than what assigned makes it double
	fmt.Printf("slc len : %d cap: %d ", len(mySlc), cap(mySlc))

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
