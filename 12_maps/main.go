package main

import "fmt"

func main() {
	//maps pass by referance
	myMap := map[string]int{ // key sting value int
		"ali":   12,
		"veli":  14,
		"kamil": 45,
		"cemil": 77,
	}
	fmt.Println(myMap["alia"])

	yourMap := map[int]string{ // key int value string
		1: "ali",
		2: "veli",
		3: "nuri",
	}
	yourMap[2] = "kamile"
	yourMap[2] += "kamil"
	fmt.Println(yourMap)
	delete(yourMap, 2)
	yourMap[5] = "kamuran"
	fmt.Println(len(yourMap))
	fmt.Println(yourMap)
}
