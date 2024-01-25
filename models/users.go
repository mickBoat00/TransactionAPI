package models

type UserRequestParams struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8"`
}

type UserResponseParams struct {
	Token string `json:"token"`
}

type ErrorJsonParams struct {
	Error string `json:"error"`
}
