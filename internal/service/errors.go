package service

import "errors"

var (
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrUnexpectedTokenClaims   = errors.New("token claims of wrong type")
)
