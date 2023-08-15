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

type Student struct {
	name string
	age  int8
}

type StudentRepo struct {
	students []Student
}

func (studentRepo *StudentRepo) add(student Student) []Student {
	studentRepo.students = append(studentRepo.students, student)
	return studentRepo.students
}

func main() {
	a := rect{width: 10, height: 10}

	fmt.Println(a.area())
	fmt.Println(a.per())
	fmt.Println("===============")
	congdat := Student{name: "Cong Dat", age: 20}

	c := StudentRepo{}
	c.add(congdat)
	c.add(congdat)

	fmt.Println(c)

}

func isTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}
