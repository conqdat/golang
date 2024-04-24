package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dsn := "user:pass@tcp(127.0.0.1:3307)/shopapp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // start := time.Now()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("db ", db)
	// genData()
	// elapsed := time.Since(start)
	// fmt.Printf("Execution time: %s\n", elapsed)
}
