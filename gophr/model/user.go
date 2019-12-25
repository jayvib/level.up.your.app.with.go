package model

import (
	"github.com/rs/xid"
)

type User struct {
	Username string
	Email string
	Password string
}

func GenerateID() string {
	guid := xid.New()
	return guid.String()
}