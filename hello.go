package main

import "fmt"

func main() {
	var i interface{} = "Hello"

	b, _ := i.(float64)

	fmt.Println(i, b)

}
