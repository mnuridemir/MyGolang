package main

import "fmt"

func main() {

	var x, y int
	y = 1

	for x = 2; x < 50; x++ {
		for y = 2; y < (x / y); y++ {
			if x%y == 0 {
				break
			}
		}
		if y > (x / y) {
			fmt.Println(x)
		}
	}

}
