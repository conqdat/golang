package main

import "fmt"

func main() {
	// Anonymous functions
	goMod := func(a, b int) int {
		return a % b
	}

	fmt.Println(goMod(10, 34))
}

func sum(a, b int) int {
	return a + b
}

// function return multiple value
func logger(a, b int) (int, int) {
	return a, b
}
