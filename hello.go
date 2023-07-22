package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // point to i
	fmt.Println(*p, i)
	i = 43
	fmt.Println(*p) // read i through the pointer
	fmt.Println(*p, i)
	*p = 21 // set i through the pointer
	//fmt.Println(i)  // see the new value of i
	fmt.Println(*p, i)
	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer
	//fmt.Println(j) // see the new value of j

}
