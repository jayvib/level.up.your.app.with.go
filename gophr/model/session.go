package model

import (
	"net/http"
	"time"
)

const (
	SessionDuration   = 24 * 3 * time.Hour
	SessionCookieName = "gophr"
)

type Session struct {
	ID     string
	UserID string
	Expiry time.Time
}

func NewSession(w http.ResponseWriter) *Session {
	expiry := time.Now().Add(SessionDuration)

	session := &Session{
		ID: GenerateID(),
		Expiry: expiry,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    SessionCookieName,
		Value:   session.ID,
		Expires: expiry,
	})

	return session
}