package main

import (
	"fmt"
	"sync"
)

// Main has a default goroutine

var wg = sync.WaitGroup{}

func main() {
	chanOne := make(chan int)
	chanTwo := make(chan int)

	go sumOne(chanOne)
	//go sumTwo(chanTwo)

	select {
	case result := <-chanTwo:
		fmt.Println("2: ", result)
	case result := <-chanOne:
		fmt.Println("1: ", result)
	}
}

func sumOne(chanOne chan int) {
	var result int
	for i := 0; i < 10; i++ {
		result += i
	}
	chanOne <- result
}

func sumTwo(chanTwo chan int) {
	var result int
	for i := 0; i < 10; i++ {
		result += i
	}
	chanTwo <- result
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
