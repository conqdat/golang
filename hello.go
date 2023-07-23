package main

import "fmt"

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
}
