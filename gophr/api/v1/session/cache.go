package session

import (
	"time"
)

type Cache interface {
	Set(id string, sess *Session, duration time.Duration) error
	Get(id string) (*Session, error)
	Delete(id string) error
}
