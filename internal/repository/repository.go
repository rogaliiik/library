package repository

import (
	"context"
	"database/sql"

	"github.com/rogaliiik/library/internal/model"
)

type Auth interface {
	CreateUser(context.Context, *model.User) (int, error)
	GetUser(ctx context.Context, username, password string) (model.User, error)
}

type Book interface {
	GetAll(ctx context.Context, userId int) ([]model.Book, error)
	GetById(ctx context.Context, bookId, userId int) (model.Book, error)
	Create(ctx context.Context, book *model.Book) (int, error)
	Update(ctx context.Context, bookId, userId int, bookUpdateInput *model.BookUpdateInput) error
	Delete(ctx context.Context, bookId, userId int) error
}

type Repository struct {
	Auth
	Book
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
		Book: NewBookRepository(db),
	}
}
