package entity

type Person struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	Age       int     `json:"age"`
	Address   Address `json:"address"`
}

func (person *Person) GoCrazy() {
	person.Age = 25
	person.Address = Address{"Bornem", "K.V.Doorslearlaan", 45, 2880}
}
