package models

//Book represents a book object
type Employed struct {
	Id           string `json:"id"`
	NameEmployed string `json:"name_employed"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
}
