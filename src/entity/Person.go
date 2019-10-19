package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Person struct {
	ID        *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	FirstName string              `json:"firstName"`
	LastName  string              `json:"lastName"`
	Password  string              `json:"-"`
	Email     string              `json:"email"`
	Age       int                 `json:"age"`
	Address   Address             `json:"address"`
}

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hashedPassword)
}

func (person *Person) VerifyPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(password)); err != nil {
		return false
	}
	return true
}
