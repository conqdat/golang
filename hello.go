package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

type Address struct {
	City    string
	Country string
}

func (add Address) String() string {
	return add.City + " - " + add.Country
}

func main() {
	// testInput()
	// testSwitchCase()

	// intArry := []int{1, 2, 3, 34, 3}

	// for index, value := range intArry {
	// 	fmt.Println("index: " + strconv.Itoa(index) + " Value: " + strconv.Itoa(value))
	// }

	// testDefer()
	// testChannel()

	// testMap()
}

func testChannel() {
	ch := make(chan int)

	// Start a goroutine to send a value to the channel
	go func() {
		ch <- 60
	}()

	// Receive the value from the channel
	value := <-ch

	fmt.Println(value)
}

func testMap() {
	my_map_1 := map[string]int{
		"one":    1,
		"two":    2,
		"threee": 3,
	}

	my_map_2 := make(map[string]int)

	my_map_2["1"] = 1
	my_map_2["2"] = 3
	my_map_2["3"] = 2

	fmt.Println(my_map_1, my_map_2)
}

func testInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input something: ")

	somthing, _ := reader.ReadString('\n')

	fmt.Print("Something is: " + somthing)
}

func testDefer() {
	a := [10]int{1, 2, 3, 4, 45, 5}
	for index, value := range a {
		defer fmt.Println(strconv.Itoa(index) + "-" + strconv.Itoa(value))
	}
}

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(3, 2)
	expected := 5
	if result != expected {
		t.Errorf("Test Failed!, expected: %d", expected)
	}
}

func testSwitchCase() {
	testBMIs := []float64{1, 2, 14, 16, 16.5, 18, 23, 27, 32, 40, 42}
	for _, bmi := range testBMIs {
		fmt.Println(ShowBMI(bmi))
	}
}

func ShowBMI(bmi float64) string {
	switch {
	case bmi < 16:
		return "Opps"
	case bmi < 16.9:
		return "Oppp 16.9"
	default:
		return "Some thing defaul"
	}
}
