package cerr

import "errors"

var ErrNoTokenFound = errors.New("no token found")
var ErrInvalidToken = errors.New("invalid token")
var ErrAlreadyRegistered = errors.New("already registered")
var ErrAlreadyAuthorized = errors.New("already authorized")
var ErrClaims = errors.New("error claims")
var ErrPasswordsNotEqual = errors.New("passwords not equal")
var ErrWrongPassword = errors.New("wrong password")
var ErrUserNotFound = errors.New("user not found")
