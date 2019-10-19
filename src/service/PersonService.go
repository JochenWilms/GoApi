package service

import (
	"../entity"
	"../repository"
	"fmt"
)

func GetPerson(param map[string][]string) []entity.Person {
	fmt.Println(param)
	if val, ok := param["firstname"]; ok {
		fmt.Println(val)
		person := repository.GetPersonByName(val[0])
		fmt.Println(person)
		return person
	}
	persons := repository.GetAllPersons()
	fmt.Println(persons)
	return persons
}

func AddPerson(person entity.Person) entity.Person {
	repository.InsertNewPerson(person)
	return person
}
