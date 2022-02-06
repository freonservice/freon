package app

import "github.com/pkg/errors"

var (
	ErrNotFound           = errors.New("entity not found")
	ErrEmailIsUsed        = errors.New("email already used")
	ErrWrongPassword      = errors.New("wrong password")
	ErrUserNotActive      = errors.New("user not active")
	ErrUserIsBanned       = errors.New("user is banned")
	ErrPasswordNotCorrect = errors.New("password is not correct")
	ErrAutoTranslation    = errors.New("auto translation not supported")
)
