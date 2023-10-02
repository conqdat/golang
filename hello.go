package main

import (
	"bufio"
	"fmt"
	"os"
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

}

func testInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input something: ")

	somthing, _ := reader.ReadString('\n')

	fmt.Print("Something is: " + somthing)
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
