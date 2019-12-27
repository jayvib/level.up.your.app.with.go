package session

import (
	"gophr/model"
	"time"
)

type Cache interface {
	Set(id string, sess *model.Session, duration time.Duration) error
	Get(id string) (*model.Session, error)
	Delete(id string) error
}
