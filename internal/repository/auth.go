package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/rogaliiik/library/internal/domain"
)

type AuthRepository struct {
	db *sql.DB
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *domain.User) (int, error) {
	id := 0
	err := r.db.QueryRow("SELECT id FROM users WHERE username = $1", user.Username).Scan(&id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	if id != 0 {
		return 0, ErrAlreadyExists
	}

	if err = r.db.QueryRow("INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id",
		user.Username, user.Password, user.Email).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetUser(ctx context.Context, username, password string) (domain.User, error) {
	var user domain.User
	err := r.db.QueryRow("SELECT id, username, password, email, created_at FROM users WHERE username = $1 AND password = $2",
		username, password).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt)

	return user, err
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}
