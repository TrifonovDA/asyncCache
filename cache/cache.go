package cache

import (
	"sync"
	"time"
)

const timeout = time.Microsecond * 2

type Cache struct {
	c map[string]string
	m *sync.RWMutex
}

func InitCache() *Cache {
	c := make(map[string]string)
	m := new(sync.RWMutex)
	return &Cache{
		c: c,
		m: m,
	}
}

func (c *Cache) Add(key, value string) {
	c.m.Lock()
	time.Sleep(timeout)
	c.c[key] = value
	c.m.Unlock()
}

func (c *Cache) Get(key string) (string, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	time.Sleep(timeout)
	value, ok := c.c[key]
	return value, ok
}

func (c *Cache) Delete(key string) *Cache {
	c.m.Lock()
	defer c.m.Unlock()
	time.Sleep(timeout)
	delete(c.c, key)
	return c
}
