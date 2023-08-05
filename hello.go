package main

import (
	"fmt"
)

func main() {
	var arr [3]int           // array with default value [0, 0, 0]
	arr2 := [3]int{1, 2, 34} // array with value
	arrBool := [1]bool{true}
	length := len(arr) // get length

	// Loop thought the array
	for i := 0; i < length; i++ {
		fmt.Println(arr2[i])
	}

	// another way to loop
	for index, value := range arr {
		fmt.Println(index, value)
	}

	fmt.Println(arr, arr2, arrBool, length)

	fmt.Println("********************")

	var myBooks [5]string
	myBooks[1] = "Book 1"
	myBooks[4] = "Book 5"

	var myNum []int8
	myNum = append(myNum, 0)
	myNum = append(myNum, 1)
	myNum = append(myNum, 2)
	myNum = append(myNum, 3)
	myNum = append(myNum, 4)

	mySlice := make([]int8, 8)
	mySlice[0] = 5
	mySlice[1] = 6
	mySlice[2] = 7
	mySlice[3] = 8
	mySlice[4] = 9

	fmt.Println(myBooks, myNum, cap(myNum), len(myNum), mySlice)

	for index, value := range myNum {
		fmt.Println(" i: ", index, " value: ", value)
	}

	fmt.Println("==========")
	name := "Cong Dat"

	for index, value := range name {
		fmt.Println(index, value)
	}

	hiDat := sayHi("Cong Dat")
	fmt.Println(hiDat)

	fmt.Println("--------------")

	myMap := make(map[string]string)

	myMap["name"] = "Dat"
	myMap["age"] = "20"
	myMap["add"] = "HCM"

	fmt.Println("my Map: ", myMap, len(myMap))

}
