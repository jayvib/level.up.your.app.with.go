package v1

import "errors"

var (
	ErrorNotFound = errors.New("item not found")
	ErrorSessionExpired = errors.New("session expired")
	ErrorCredentials = errors.New("invalid credentials")
)
