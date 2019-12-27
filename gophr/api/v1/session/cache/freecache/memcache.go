package freecache

import (
	"encoding/json"
	"github.com/coocood/freecache"
	"gophr/model"
	"time"
)

func New()  *FreeCache {
	return &FreeCache{
		cache: freecache.NewCache(100*1024*1024),
	}
}

type FreeCache struct {
	cache *freecache.Cache
}

func (f *FreeCache) Get(id string) (*model.Session, error) {
	res, err := f.cache.Get([]byte(id))
	if err != nil {
		return nil, err
	}

	var sess model.Session

	err = json.Unmarshal(res, &sess)
	if err != nil {
		return nil, err
	}

	return &sess, nil
}

func (f *FreeCache) Set(id string, sess *model.Session, seconds time.Duration) error {

	payload, err := json.Marshal(sess)
	if err != nil {
		return err
	}

	err = f.cache.Set([]byte(sess.ID), payload, int(seconds))
	if err != nil {
		return err
	}

	return nil
}

func (f *FreeCache) Delete(id string) error {
	f.cache.Del([]byte(id))
	return nil
}