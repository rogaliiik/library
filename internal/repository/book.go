package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/rogaliiik/library/internal/domain"
)

type BookRepository struct {
	db *sql.DB
}

func (r *BookRepository) GetAll(ctx context.Context, userId int) ([]domain.Book, error) {
	var books []domain.Book
	rows, err := r.db.Query("SELECT id, user_id, name, content, author, created_at FROM books WHERE user_id = $1",
		userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book domain.Book
		err = rows.Scan(&book.Id, &book.UserId, &book.Name, &book.Content, &book.Author, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, rows.Err()
	}

	if len(books) == 0 {
		return nil, sql.ErrNoRows
	}

	return books, nil
}

func (r *BookRepository) GetById(ctx context.Context, bookId, userId int) (domain.Book, error) {
	var book domain.Book
	err := r.db.QueryRow("SELECT id, user_id, name, content, author, created_at FROM books WHERE user_id = $1 AND id = $2",
		userId, bookId).Scan(&book.Id, &book.UserId, &book.Name, &book.Content, &book.Author, &book.CreatedAt)

	return book, err
}

func (r *BookRepository) Create(ctx context.Context, book *domain.Book) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO books (user_id, name, content, author) VALUES ($1, $2, $3, $4) RETURNING id",
		&book.UserId, &book.Name, &book.Content, &book.Author).Scan(&id)

	return id, err
}

func (r *BookRepository) Update(ctx context.Context, bookId, userId int, input *domain.BookUpdateInput) error {
	var (
		query   = "UPDATE books SET "
		changes []string
	)

	if input.Name != "" {
		changes = append(changes, fmt.Sprintf("name = '%s'", input.Name))
	}

	if input.Content != "" {
		changes = append(changes, fmt.Sprintf("content = '%s'", input.Content))
	}

	if input.Author != "" {
		changes = append(changes, fmt.Sprintf("author = '%s'", input.Author))
	}

	if len(changes) == 0 {
		return nil
	}

	query += strings.Join(changes, ",")
	query += " WHERE id = $1 AND user_id = $2"
	_, err := r.db.Exec(query, bookId, userId)

	return err
}

func (r *BookRepository) Delete(ctx context.Context, bookId, userId int) error {
	_, err := r.db.Exec("DELETE FROM books where id = $1 AND user_id = $2", bookId, userId)

	return err
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}
