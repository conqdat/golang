package model

import (
	fake "github.com/brianvoe/gofakeit/v6"
	"time"
)

type Address struct {
	Line    string
	City    string
	Country string
}

type Person struct {
	Name        string
	Emails      []string
	DateOfBirth time.Time
	Addresses   []*fake.AddressInfo
}

type Book struct {
	Title  string
	Author string
	Pages  string
}
