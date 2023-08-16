package shape

import "fmt"

type React struct {
	Width  float64
	Height float64
}

func (rec React) Area() float64 {
	return rec.Width * rec.Height
}

func (rec React) Perimeter() float64 {
	return 2 * (rec.Width + rec.Height)
}

func (rec React) String() string {
	return fmt.Sprintf("Reactangle W=%2.f, h=%2.f", rec.Width, rec.Height)
}
