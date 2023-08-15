package main

import "fmt"

type rect struct {
	width, height int8
}

func (r rect) area() int {
	return int(r.width * r.height)
}

func (r rect) per() int {
	return int(2*r.width + 2*r.height)
}

func main() {
	a := rect{width: 10, height: 10}

	fmt.Println(a.area())
	fmt.Println(a.per())
}

func isTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}
