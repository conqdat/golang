package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i int8 = -1 // signed int
	var u uint8 = 1 // unsigned int => no negative number
	var f float32 = 2.3
	isHappy := true
	var s string = "TruE" // false

	str, _ := strconv.ParseBool(s)

	fmt.Print(isHappy, i, u, f, str)
}
