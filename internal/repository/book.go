package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rogaliiik/library/internal/model"
)

type BookRepository struct {
	db *sql.DB
}

func (r *BookRepository) GetAll(ctx context.Context, userId int) ([]model.Book, error) {
	var books []model.Book
	rows, err := r.db.Query("SELECT id, user_id, name, content, author, created_at FROM books WHERE user_id = $1",
		userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.Id, &book.UserId, &book.Name, &book.Content, &book.Author, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, rows.Err()
}

func (r *BookRepository) GetById(ctx context.Context, bookId, userId int) (model.Book, error) {
	var book model.Book
	err := r.db.QueryRow("SELECT id, user_id, name, content, author, created_at FROM books WHERE user_id = $1 AND id = $2",
		userId, bookId).Scan(&book.Id, &book.UserId, &book.Name, &book.Content, &book.Author, &book.CreatedAt)

	return book, err
}

func (r *BookRepository) Create(ctx context.Context, book *model.Book) (int, error) {
	var id int
	rows, err := r.db.Query("INSERT INTO books (user_id, name, content, author) VALUES ($1, $2, $3, $4) RETURNING id",
		&book.UserId, &book.Name, &book.Content, &book.Author)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	err = rows.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, rows.Err()
}

func (r *BookRepository) Update(ctx context.Context, bookId, userId int, input *model.BookUpdateInput) error {
	var query = "UPDATE books SET "

	if input.Name != "" {
		query += fmt.Sprintf("name = %s ", input.Name)
	}

	if input.Content != "" {
		query += fmt.Sprintf("content = %s ", input.Content)
	}

	if input.Author != "" {
		query += fmt.Sprintf("author = %s ", input.Author)
	}

	query += "WHERE id = $1 AND user_id = $2"
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
