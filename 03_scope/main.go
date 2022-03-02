package main

import "fmt"

var packVar = "Package Scope" //cant use := shorthand decleration

func main() {

	var funcVar = "Func Scope"
	fmt.Println(funcVar)
	fmt.Println(packVar)

	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}

	myFunc()

	var name = "kamil"
	name, surname := "mehmet", "demir" // assigment and declaration
	fmt.Println(name)
	fmt.Println(surname)
}

func myFunc() {
	fmt.Println(packVar)
}
