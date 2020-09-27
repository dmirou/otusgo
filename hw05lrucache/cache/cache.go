package cache

import (
	"fmt"

	"github.com/dmirou/otusgo/hw05lrucache/list"
)

type Item struct {
	Key   string
	Value interface{}
}

type Cache struct {
	queue *list.List
	cache map[string]*list.Item
	size  int
}

// NewCache creates new empty cache with the specified size.
func NewCache(size int) (*Cache, error) {
	if size < 1 {
		return nil, fmt.Errorf("cache size must be greater or equal than one")
	}

	return &Cache{
		queue: list.NewList(),
		cache: make(map[string]*list.Item),
		size:  size,
	}, nil
}

// Set saves item with key and value into cache.
// It returns true if item with key was in cache
// before setting, else false.
// If cache already reached its size and we try
// to add new item into cache, the oldest element
// will be removed to get a memory.
func (c *Cache) Set(key string, value interface{}) bool {
	if el, ok := c.cache[key]; ok {
		item := el.Value().(*Item)
		item.Value = value

		c.queue.MoveToFront(el)

		return true
	}

	c.cache[key] = c.queue.PushFront(&Item{Key: key, Value: value})

	if c.size < c.queue.Len() {
		last := c.queue.Back()
		c.queue.Remove(last)
		delete(c.cache, last.Value().(*Item).Key)
	}

	return false
}

// Get returns cache value and true if item with
// key found in cache, else nil and false.
func (c *Cache) Get(key string) (interface{}, bool) {
	if el, ok := c.cache[key]; ok {
		c.queue.MoveToFront(el)

		return el.Value().(*Item).Value, true
	}

	return nil, false
}

// Clear clears cache.
func (c *Cache) Clear() {
	c.queue = list.NewList()
	c.cache = make(map[string]*list.Item)
}
