package main

import (
	"encoding/json"
	"fmt"
	"time"
)

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

type User struct {
	Name     string
	SurName  string
	Follower []User `json:"follower,omitempty"` // we can use omitempty to ignore zero values
	Likes    []Like
}

type Like struct {
	Date time.Time
}

func main() {

	//example1()
	example2()

}

func example1() {
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

func example2() {

	u := User{
		Name:    "Go",
		SurName: "Turkiye",
		Follower: []User{
			{
				Name:    "name",
				SurName: "1",
			},
			{
				Name:    "name",
				SurName: "2",
			},
		},
	}

	arr, _ := json.Marshal(u)
	fmt.Println(string(arr))

	fmt.Println(u)

}
