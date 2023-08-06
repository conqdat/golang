package main

import (
	"fmt"
	"sync"
)

// Main has a default goroutine

var wg = sync.WaitGroup{}

func main() {

	sumChan := make(chan int)

	wg.Add(2) // Wait one go routine
	go loggerTwo(sumChan)
	go loggerOne(sumChan)

	fmt.Println(<-sumChan)
	fmt.Println(<-sumChan)

	//wg.Wait()
}

func loggerOne(sumChan chan int) {
	var result int
	for i := 0; i <= 10; i++ {
		fmt.Println("i: ", i)
		result += i
	}
	sumChan <- result
	wg.Done()
}

func loggerTwo(sumChan chan int) {
	var result int
	for i := 0; i <= 10; i++ {
		fmt.Println("i_2: ", i)
		result += i
	}
	sumChan <- result
	wg.Done()
}
