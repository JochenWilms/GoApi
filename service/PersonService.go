package service

import (
	"../entity"
	"encoding/json"
	"fmt"
)

func GetPerson() string {
	person := entity.Person{FirstName: "Jochen", LastName: "Wilms", Email: "jochen@mariekerke.be"}
	person.GoCrazy()
	json, err := json.Marshal(person)
	fmt.Println(person)

	if err != nil {
		fmt.Errorf("error", err)
		return "error"
	}
	return string(json)
}
