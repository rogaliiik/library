package domain

import (
	"time"
)

type User struct {
	Id        int       `json:"id" validate:"-"`
	Username  string    `json:"username" validate:"required,alphanum"`
	Password  string    `json:"password" validate:"required,alphanum,min=8,max=20"`
	Email     string    `json:"email" validate:"omitempty,email" example:"user@gmail.com"`
	CreatedAt time.Time `json:"-" validate:"-"`
}

func (u *User) Validate() error {
	return validate.Struct(u)
}
