package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Main has a default goroutine

func main() {
	fmt.Println("Hello Guy !!!")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/api", apiPage)
	err := http.ListenAndServe(":333", nil)
	if err != nil {
		return
	}
}

func apiPage(writer http.ResponseWriter, request *http.Request) {
	var data = map[string]interface{}{
		"name": "cong Dat",
	}
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		return
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Home page !!!")
	if err != nil {
		return
	}
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "About page !!!")
	if err != nil {
		return
	}
}
