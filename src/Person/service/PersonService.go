package service

import (
	"desyco.com/Person/entity"
	"desyco.com/Person/repository"
	"fmt"
)

func FindPersonByMail(email string) entity.Person {
	person := repository.FindPersonByMail(email)
	return person
}

func GetPerson(param map[string][]string) entity.Person {
	fmt.Println(param)
	if val, ok := param["firstname"]; ok {
		fmt.Println(val)
		person := repository.GetPersonByName(val[0])
		fmt.Println(person)
		return person
	}
	persons := repository.GetAllPersons()
	fmt.Println(persons)
	return persons[0]
}

func AddPerson(person entity.Person) entity.Person {
	repository.InsertNewPerson(person)
	return person
}

func UpdatePassword(personId string, password string) bool {
	hashedPW := entity.HashPassword(password)
	fmt.Println(hashedPW)
	return repository.UpdatePassword(personId, hashedPW)

}
