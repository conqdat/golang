package main

import (
	"encoding/json"
	"fmt"
)

// Main has a default goroutine

func main() {
	var myMap = make(map[string]interface{})

	myMap["1"] = true
	myMap["2"] = 2
	myMap["3"] = "3"

	marshal, err := json.Marshal(myMap) // Marshal func is used to convert data to json
	if err != nil {
		return
	}
	marshalInString := string(marshal)

	var data = make(map[string]interface{})

	var payload = `{
		"age":18,
		"name": "Dat"
	}`

	// UnMarshal func is used to convert json to data
	err = json.Unmarshal([]byte(payload), &data)

	type People struct {
		name    string
		age     int
		address string
	}

	myBoy := People{age: 20, name: "Cong Dat", address: "HCM"}

	myBoyJson, _ := json.Marshal(myBoy)

	fmt.Println(string(myBoyJson))

	fmt.Println("=========")
	fmt.Println(data)
	fmt.Println("=======")
	fmt.Println(marshalInString)
	fmt.Println(myMap)
}
