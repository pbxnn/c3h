package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/VictoriaMetrics/fastcache"
)

type Cache interface {
	Get(key string, result interface{}) error
	Set(key string, value interface{}) error
}

type cache struct {
	data   *fastcache.Cache
	expire int64
}

type cacheValue struct {
	Detail   interface{}
	Deadline int64
}

func NewCache(size, expire int64) Cache {
	return &cache{
		data:   fastcache.New((int)(size)),
		expire: expire,
	}
}

func (c *cache) Get(key string, result interface{}) error {
	k := []byte(key)
	var v []byte
	c.data.Get(k, v)
	cv := &cacheValue{
		Detail: result,
	}
	err := json.Unmarshal(v, cv)
	if err != nil {
		return err
	}

	if time.Now().Unix() > cv.Deadline {
		result = nil
		return fmt.Errorf("cache expired. key:%s, deadline:%s", key, time.Unix(cv.Deadline, 0).Format("2006-01-02 15:04:05"))
	}

	return nil
}

func (c *cache) Set(key string, value interface{}) error {
	k := []byte(key)
	cv := cacheValue{
		Detail:   value,
		Deadline: time.Now().Unix() + c.expire,
	}
	v, err := json.Marshal(cv)
	if err != nil {
		return err
	}
	c.data.Set(k, v)
	return nil
}
