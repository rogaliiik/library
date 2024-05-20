package model

import "time"

type Book struct {
	Id        int       `json:"id" validate:"-"`
	UserId    int       `json:"userId" validate:"-"`
	Name      string    `json:"name" validate:"min=1,max=6,alphanum"`
	Content   string    `json:"content" validate:"-"`
	Author    string    `json:"author" validate:"required"`
	CreatedAt time.Time `json:"-" validate:"-"`
}

type BookUpdateInput struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (U *Book) Validate(book *Book) error {
	return validate.Struct(book)
}
