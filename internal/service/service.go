package service

import (
	"context"
	"time"

	"github.com/rogaliiik/library/internal/model"
	"github.com/rogaliiik/library/internal/repository"
)

type Auth interface {
	CreateUser(ctx context.Context, user *model.User) (int, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, token string) (userId int, err error)
}

type Book interface {
	GetAll(ctx context.Context, userId int) ([]model.Book, error)
	GetById(ctx context.Context, bookId, userId int) (model.Book, error)
	Create(ctx context.Context, book *model.Book) (int, error)
	Update(ctx context.Context, bookId, userId int, bookUpdateInput *model.BookUpdateInput) error
	Delete(ctx context.Context, bookId, userId int) error
}

type Service struct {
	Auth
	Book
}

func NewServices(repo *repository.Repository, salt string, signingKey string, ttl time.Duration) *Service {
	return &Service{
		Auth: NewAuthService(repo, salt, signingKey, ttl),
		Book: NewBookService(repo),
	}
}
