package userservice

import "github.com/google/uuid"

type UserDto struct {
	Uuid      uuid.UUID `json:"uuid"`
	Email     string    `json:"email"`
	Enabled   bool      `json:"enabled"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
}

type UserCreateRequestDto struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}
