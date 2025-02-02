package models

import "errors"

var (
	// ErrInvalidArgument for invalid arguments
	ErrInvalidArgument = errors.New("invalid argument")
	// ErrBadRequest for http bad requests
	ErrBadRequest = errors.New("bad request")
)
