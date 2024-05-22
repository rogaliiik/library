package v1

import "errors"

var (
	ErrInvalidParameter = errors.New("invalid url parameter")

	ErrEmptyAuthHeader   = errors.New("empty auth header")
	ErrInvalidAuthHeader = errors.New("invalid auth header")
	ErrInvalidIdType     = errors.New("invalid user id type")
)
