package domain

import "errors"

var (
	ErrMissingArgument    = errors.New("missing argument")
	ErrInvalidMessageType = errors.New("invalid message-type")
	ErrInvalidKey         = errors.New("invalid key")
	ErrNotFound           = errors.New("not found")
)
