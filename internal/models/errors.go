package models

import (
	"errors"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")

	ErrInvalidCredentials = errors.New("modles: invalid credentials")

	ErrDuplicateEmail = errors.New("models:duplicate email")
)
