package domain

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrBadRequest   = errors.New("bad request")
	ErrUnknownError = errors.New("unknown error")
)
