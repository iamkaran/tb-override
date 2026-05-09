// Package core contains core objects like errors
package core

import (
	"errors"
)

var (
	// I/O
	ErrNoRootPrivilages = errors.New("root privilages are required")
	ErrAlreadyExists    = errors.New("directory or file already exists")
	ErrInvalidTheme     = errors.New("theme does not exist")
)
