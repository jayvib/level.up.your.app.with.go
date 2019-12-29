package session

import (
	"context"
	"github.com/jayvib/golog"
	"gophr/api/v1"
	"gophr/api/v1/user"
	"gophr/model"
	"net/http"
	"time"
)

func GetRequestSession(c Cache, r *http.Request) (*Session, error) {
	golog.Debugf("%#v\n", r.Header)
	// get cookie from the request
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return nil, err
	}
	golog.Debugf("cookie: %#v\n", cookie)

	// check if expired, then return an error if expired
	sessId := cookie.Value
	val, err := c.Get(sessId)
	golog.Debugf("get session: %s %#v\n", sessId, val)
	if err != nil {
		if err == v1.ErrorNotFound {
			return nil, v1.ErrorSessionExpired
		}
	}

	return val, nil
}

func GetSession(c Cache, w http.ResponseWriter, r *http.Request) *Session {
	session, err := GetRequestSession(c, r)
	if err != nil {
		session = NewSession(w)
	}
	return session
}

func GetUserFromSession(userRepo user.Service, c Cache, r *http.Request) *model.User {
		sess, err := GetRequestSession(c, r)
		if err != nil {
			golog.Error(err)
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
	DurationInSecond = Duration*60*60
	CookieName = "Gophr"
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
		Path: "/",
	})

	return session
}
