package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go printY()
	//fmt.Println("--------------------")
	wg.Wait()
	printX()

	//time.Sleep((time.Second))

}

func printX() {
	for i := 0; i < 15; i++ {
		fmt.Print("X")
	}
}

func printY() {
	for i := 0; i < 100; i++ {
		fmt.Print("Y")
	}
	wg.Done()
}
