package entity

type Address struct {
	City    string `json:"city"`
	Street  string `json:"street"`
	Number  int    `json:"number"`
	ZipCode int    `json:"zipCode"`
}
