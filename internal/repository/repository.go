package repository

import (
	"context"
	"database/sql"

	"github.com/rogaliiik/library/internal/domain"
)

type Auth interface {
	CreateUser(context.Context, *domain.User) (int, error)
	GetUser(ctx context.Context, username, password string) (domain.User, error)
}

type Book interface {
	GetAll(ctx context.Context, userId int) ([]domain.Book, error)
	GetById(ctx context.Context, bookId, userId int) (domain.Book, error)
	Create(ctx context.Context, book *domain.Book) (int, error)
	Update(ctx context.Context, bookId, userId int, bookUpdateInput *domain.BookUpdateInput) error
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
