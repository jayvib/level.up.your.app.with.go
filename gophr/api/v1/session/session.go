package session

import (
	"context"
	"gophr/api/v1"
	"gophr/api/v1/user"
	"gophr/model"
	"net/http"
	"time"
)

func GetRequestSession(c Cache, r *http.Request) (*Session, error) {
	// get cookie from the request
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return nil, err
	}

	// check if expired, then return an error if expired
	sessId := cookie.Value
	val, err := c.Get(sessId)
	if err != nil {
		if err == v1.ErrorNotFound {
			return nil, v1.ErrorSessionExpired
		}
	}
	return val, nil
}

func GetUserFromSession(userRepo user.Repository, c Cache, r *http.Request) *model.User {
		sess, err := GetRequestSession(c, r)
		if err != nil {
			return nil
		}

		usr, err := userRepo.GetByID(context.Background(), sess.UserID)
		if err != nil {
			return nil
		}

		return usr
}

const (
	Duration   = 8 * time.Hour
	CookieName = "gophr"
)

type Session struct {
	ID     string
	UserID string
	Expiry time.Time
}

func NewSession(w http.ResponseWriter) *Session {
	expiry := time.Now().Add(Duration)

	session := &Session{
		ID: model.GenerateID(),
		Expiry: expiry,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    CookieName,
		Value:   session.ID,
		Expires: expiry,
	})

	return session
}
