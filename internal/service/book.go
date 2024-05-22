package service

import (
	"context"

	"github.com/rogaliiik/library/internal/domain"
	"github.com/rogaliiik/library/internal/repository"
)

type BookService struct {
	repo *repository.Repository
}

func (s *BookService) GetAll(ctx context.Context, userId int) ([]domain.Book, error) {
	return s.repo.Book.GetAll(ctx, userId)
}

func (s *BookService) GetById(ctx context.Context, bookId, userId int) (domain.Book, error) {
	return s.repo.Book.GetById(ctx, bookId, userId)
}

func (s *BookService) Create(ctx context.Context, book *domain.Book) (int, error) {
	return s.repo.Book.Create(ctx, book)
}

func (s *BookService) Update(ctx context.Context, bookId, userId int, bookUpdateInput *domain.BookUpdateInput) error {
	return s.repo.Book.Update(ctx, bookId, userId, bookUpdateInput)
}

func (s *BookService) Delete(ctx context.Context, bookId, userId int) error {
	return s.repo.Book.Delete(ctx, bookId, userId)
}

func NewBookService(repos *repository.Repository) *BookService {
	return &BookService{repo: repos}
}
