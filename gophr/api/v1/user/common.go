package user

import "errors"

var (
	ErrorNotFound = errors.New("item not found")
	ErrorUsernameExists = errors.New("username already exists")
	ErrorEmailExists = errors.New("email already exists")
)
