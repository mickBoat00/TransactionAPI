package models

type UserRequestParams struct {
	Email    string
	Password string
}

type UserResponseParams struct {
	Token string `json:"token"`
}
