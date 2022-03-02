package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := evenNum(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func evenNum(num int) (string, error) {
	if num%2 != 0 {
		return "", errors.New("hatalı giriş yaptınız")
	}
	return "tebrikler çift sayı girdiniz", nil
}
