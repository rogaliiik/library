package repository

import (
	"context"
	"database/sql"

	"github.com/rogaliiik/library/internal/model"
)

type AuthRepository struct {
	db *sql.DB
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *model.User) (int, error) {
	id := 0
	err := r.db.QueryRow("SELECT id FROM users WHERE username = $1", user.Username).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if id != 0 {
		return 0, ErrAlreadyExists
	}

	if err = r.db.QueryRow("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)",
		user.Username, user.Password, user.Email).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetUser(ctx context.Context, username, password string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, username, password, email, created_at FROM users WHERE username = $1 AND password = $2",
		username, password).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt)

	return user, err
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}
