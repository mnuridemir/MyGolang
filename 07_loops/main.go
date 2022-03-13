package main

import (
	"fmt"
	"time"
)

func main() {
	//loopOne()
	//loopExample()
	timer()

}

func loopOne() {
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
func loopExample() {

	mySlc := []int{1, 2, 3, 4, 5}

	for i, value := range mySlc {
		fmt.Printf("index : %d and value : %d \n", i, value)
	}

}

func timer() {
	c := time.After(5 * time.Second)

	for {
		b := false
		select {
		case <-c:
			b = true
		default:
			fmt.Println(time.Now().Second())
			time.Sleep(time.Second)
		}

		if b {
			break
		}
	}

}
