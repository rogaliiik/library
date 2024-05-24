package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/rogaliiik/library/internal/domain"
	"github.com/rogaliiik/library/internal/repository"
)

type customClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo       *repository.Repository
	salt       string
	signingKey string
	ttl        time.Duration
}

func (s *AuthService) CreateUser(ctx context.Context, user *domain.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.Auth.CreateUser(ctx, user)
}

func (s *AuthService) GenerateToken(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.Auth.GetUser(ctx, username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &customClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(s.signingKey))
}

func (s *AuthService) ParseToken(ctx context.Context, tokenString string) (userId int, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}

		return []byte(s.signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return 0, ErrUnexpectedTokenClaims
	}

	return claims.UserId, nil
}

func NewAuthService(repos *repository.Repository, salt string, signingKey string, ttl time.Duration) *AuthService {
	return &AuthService{
		repo:       repos,
		salt:       salt,
		signingKey: signingKey,
		ttl:        ttl,
	}
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.salt)))
}
