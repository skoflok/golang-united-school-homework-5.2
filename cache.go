package cache

import (
	"time"
)

type Item struct {
	value    string
	deadline time.Time
	infinite bool
}

func NewItem(value string, deadline time.Time, infinite bool) Item {
	return Item{value: value, deadline: deadline, infinite: infinite}
}

type Cache struct {
	store map[string]Item
}

func NewCache() Cache {
	return Cache{make(map[string]Item)}
}

func (c *Cache) Get(key string) (string, bool) {
	item, ok := c.store[key]
	if ok != true {
		return "", false
	}

	if item.infinite == true {
		return item.value, true
	}

	if item.deadline.Before(time.Now()) == true {
		delete(c.store, key)
		return "", false
	}
	return item.value, true
}

func (c *Cache) Put(key, value string) {
	newItem := NewItem(value, time.Now(), true)
	c.store[key] = newItem
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0)
	for k, _ := range c.store {
		keys = append(keys, k)
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	newItem := NewItem(value, deadline, false)
	c.store[key] = newItem
}
