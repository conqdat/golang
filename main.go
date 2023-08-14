package main

import "fmt"

func main() {
	fmt.Println(isTriangle(3, 4, 5))
	fmt.Println(isTriangle(1, 1, 13))
}

func isTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}
