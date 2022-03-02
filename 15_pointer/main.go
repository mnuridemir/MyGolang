package main

import "fmt"

func main() {
	name := "mehmet"
	x := &name
	fmt.Println(&name)    //address of name
	fmt.Println(x)        //address of name
	fmt.Printf("%T\n", x) // string pointer => *string
	fmt.Println(*(x))

	name2 := &name
	*name2 = "nuri" // both changes (pass by referance)
	fmt.Println(name, *name2)

	number := 4
	fmt.Println(number)
	makeDouble(number)  // pass by value
	fmt.Println(number) // not changed

}

func makeDouble(num int) {
	fmt.Println(num * 2)
}
