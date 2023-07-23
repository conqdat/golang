package main

import "fmt"
import "reflect"

func main() {
	s := []int{2, 2, 3}
	sliceString := []string{"1", "2", "3"}

	subSlice := s[:]      // get all items in slice s
	subSliceTwo := s[0:2] // from index = 1, get 2 items

	s = append(s, 4) // add item to slice

	d := append(s, subSlice...) // append slice to slice

	fmt.Println(s, reflect.TypeOf(s), sliceString, subSlice, subSliceTwo, d)
}
