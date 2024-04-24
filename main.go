package main

import (
	"fmt"
	"os"
	"time"
)

func genData() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // Defer closing the file until the function exits

	for i := 0; i < 1000000; i++ {
		fmt.Println("inserting : ", i)
		line := fmt.Sprintf("insert into categories (id, name) values (%d, 'Cassie %d');\n", i, i)
		_, err = file.WriteString(line)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	fmt.Println("Data has been written to the file.")
}

func main() {
	start := time.Now()
	genData()
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}