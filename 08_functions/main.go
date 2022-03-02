package main

import (
	"fmt"
	"time"
)

func main() {
	sum := mySum(3, 5)
	fmt.Println(sum)

	var x int = 4
	var today time.Time = time.Now()

	fmt.Println(x, today)

}

// moduler programming
// readable code

func mySum(x, y int) int {
	return x + y
}
