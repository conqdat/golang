package main

import "sync"

// Main has a default goroutine

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1) // Wait one go routine
	go sum(9, 1)
	wg.Wait()
}

func sum(a, b int) int {
	return a + b
}
