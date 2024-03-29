package cache

import (
	"errors"
	"sync"
)

type Cache struct {
	sync.Mutex
	Data map[string]string
}

func (c *Cache) Put(uid string, order string) {
	c.Lock()
	defer c.Unlock()
	c.Data[uid] = order
}

func (c *Cache) Get(uid string) (string, error) {
	c.Lock()
	defer c.Unlock()

	var order string
	if order, ok := c.Data[uid]; ok {
		return order, nil
	}
	return order, errors.New("error")
}

func (c *Cache) Del(uid string) {
	delete(c.Data, uid)
}
