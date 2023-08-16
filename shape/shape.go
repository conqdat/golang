package shape

import "fmt"

type Shape interface {
	Perimeter() float64 // function prototype =>
	Area() float64
	String() string
}

func DemoShape() {
	rec := React{Height: 10, Width: 20}

	var shape Shape
	shape = rec

	fmt.Println(shape)
}
