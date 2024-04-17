package models

type Login struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}
