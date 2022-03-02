package main

import "fmt"

type employee struct {
	name      string
	age       int
	isMarried bool
	kids      []string
}

type menager struct {
	employee
	hasDegree bool
}

func main() {

	e1 := employee{
		name:      "mehmet",
		age:       25,
		isMarried: false,
		kids: []string{
			"ali",
			"veli",
		},
	}
	fmt.Println(e1)
	e2 := e1
	e2.name = "kemal"
	fmt.Println(e2)
	fmt.Println(e1)

	m1 := menager{
		employee: employee{ // we can use another struct inside of struct
			name:      "ayse",
			age:       22,
			isMarried: false,
			kids: []string{
				"salih",
				"kamuran",
			},
		},
		hasDegree: true,
	}

	fmt.Println(m1)
	fmt.Println(m1.employee.name)

}
