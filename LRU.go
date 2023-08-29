package lru

import (
	"container/list"
)

type CachedItem struct {
	value string
	key   string
}

type Cache struct {
	capacity int
	storage  map[string]*list.Element
	queue    *list.List
}

func (c *Cache) Add(key, value string) bool {

	if v, ok := c.storage[key]; ok {
		c.queue.MoveToFront(v)
		v.Value.(*CachedItem).value = value
		return true
	}

	newItem := &CachedItem{
		value: value,
		key:   key,
	}
	c.storage[key] = c.queue.PushFront(newItem)
	if len(c.storage) > c.capacity {
		elem := c.queue.Remove(c.queue.Back()).(*CachedItem)
		delete(c.storage, elem.key)
	}
	return true
}

func (c *Cache) Get(key string) (value string, ok bool) {
	v, ok := c.storage[key]
	if ok {
		c.queue.MoveToFront(v)
		return v.Value.(*CachedItem).value, ok
	}
	return "", false
}

func (c *Cache) Remove(key string) (ok bool) {
	v, ok := c.storage[key]
	if ok {
		delete(c.storage, key)
		c.queue.Remove(v)
	}
	return ok
}

func NewLRUCahce(capacity int) LRUCache {
	return &Cache{capacity: capacity, storage: make(map[string]*list.Element), queue: list.New()}
}
