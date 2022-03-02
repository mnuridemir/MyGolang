package main

import "fmt"

func main() {

	var firstName, lastName string = "mnd", "demir" // multi assigment
	var age float32 = 25.2
	isMarried := false // short hand declaration

	var empty string

	fmt.Println("selaam", firstName, lastName, "your age is :", age, isMarried)

	//fmt.Printf("%T", age)

	fmt.Println(empty)
}
