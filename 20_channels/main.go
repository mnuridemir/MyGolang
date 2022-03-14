package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle")
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func merhaba(chan1 chan string) {
	chan1 <- "merhabaa"
}

func main() {
	//myFunc1()
	myFunc2()

}

func myFunc1() {
	c1 := circle{5}
	c1.display()
	area := c1.area()
	fmt.Printf("area = %.2f\n", area)
}

func myFunc2() {
	myChannel := make(chan string)

	go merhaba(myChannel)

	fmt.Println(<-myChannel)
}
