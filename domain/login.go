package domain

type Login struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

func NewLogin(
	id string,
	email string,
	password string,
	salt string,
) *Login {
	return &Login{
		Id:       id,
		Email:    email,
		Password: password,
		Salt:     salt,
	}
}
