package main

import "fmt"

// Main has a default goroutine

func main() {
	// init a map
	var myMap = make(map[int8]string)

	// assign key value to map
	myMap[1] = "One"
	myMap[2] = "Two"

	// check the length of map
	var myLength = len(myMap)

	// delete key / value in map
	delete(myMap, 1)

	// Check key exist
	_, isExist := myMap[1]

	fmt.Println(isExist)
	fmt.Println("=====")
	fmt.Println(myMap, myLength)
}
