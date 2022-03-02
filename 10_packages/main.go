package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Sınav sonucunu giriniz")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	fmt.Println(value)
}
