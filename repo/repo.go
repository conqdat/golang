package repo

import (
	"GoLang/model"
	fake "github.com/brianvoe/gofakeit/v6"
)

var (
	people []*model.Person
)

func init() {
	people = []*model.Person{}
	for i := 0; i < 30; i++ {
		people = append(people, &model.Person{
			Name:   fake.Name(),
			Emails: []string{fake.Email(), fake.Email()},
		})
	}
}

func ListPeople() []*model.Person {
	return people
}
