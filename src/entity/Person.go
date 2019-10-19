package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID        *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	FirstName string              `json:"firstName"`
	LastName  string              `json:"lastName"`
	Email     string              `json:"email"`
	Age       int                 `json:"age"`
	Address   Address             `json:"address"`
}

func (person *Person) GoCrazy() {
	person.Age = 25
	person.Address = Address{"Bornem", "K.V.Doorslearlaan", 45, 2880}
}
