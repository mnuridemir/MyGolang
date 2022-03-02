package main

import "fmt"

func main() {
	// cities := [...]string{"ist", "anka", "maras", "antep"} //fixed size array
	// fmt.Println(cities[1])
	// fmt.Println(len(cities))
	// var myArr [4]int
	// var myArr2 [4]int
	// myArr[0] = 3
	// myArr2[0] = 3

	// if myArr == myArr2 {
	// 	fmt.Println("equal")
	// }

	myArr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mySqrArr := mySqr(myArr)
	fmt.Println(mySqrArr)

	underArray := myArr[2:5] //takes 2 parameters 1 start index 2 end index(not included)
	fmt.Println(underArray)
	fmt.Println(len(underArray))

	var mySlice []int
	fmt.Println(mySlice)

}

func mySqr(arr [10]int) [10]int { // pass by value
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}
