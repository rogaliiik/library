package service

import (
	"context"
	"time"

	"github.com/rogaliiik/library/internal/domain"
	"github.com/rogaliiik/library/internal/repository"
)

type Auth interface {
	CreateUser(ctx context.Context, user *domain.User) (int, error)
	GenerateToken(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, token string) (userId int, err error)
}

type Book interface {
	GetAll(ctx context.Context, userId int) ([]domain.Book, error)
	GetById(ctx context.Context, bookId, userId int) (domain.Book, error)
	Create(ctx context.Context, book *domain.Book) (int, error)
	Update(ctx context.Context, bookId, userId int, bookUpdateInput *domain.BookUpdateInput) error
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
