package cache

import (
	"encoding/json"

	"github.com/VictoriaMetrics/fastcache"
)

type Cache interface {
	Get(key string, result interface{}) error
	Set(key string, value interface{}) error
}

type cache struct {
	data *fastcache.Cache
}

func NewCache(size int) Cache {
	return &cache{
		data: fastcache.New(size),
	}
}

func (c *cache) Get(key string, result interface{}) error {
	k := []byte(key)
	var v []byte
	c.data.Get(k, v)
	err := json.Unmarshal(v, result)
	return err
}

func (c *cache) Set(key string, value interface{}) error {
	k := []byte(key)
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	c.data.Set(k, v)
	return nil
}
