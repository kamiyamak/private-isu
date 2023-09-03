package main

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	mu    sync.Mutex // guards
	items map[int]int
}

func NewCache() *Cache {
	m := make(map[int]int)
	c := &Cache{
		items: m,
	}
	return c
}

func (c *Cache) Set(key int, value int) {
	c.mu.Lock()
	c.items[key] = value
	c.mu.Unlock()
}

func (c *Cache) Get(key int) int {
	c.mu.Lock()
	v, ok := c.items[key]
	c.mu.Unlock()

	if ok {
		return v
	}

	v = HeavyGet(key)

	c.Set(key, v)

	return v
}

func HeavyGet(key int) int {
	time.Sleep(time.Second)
	return key * 2
}

func main() {
	mCache := NewCache()
	log.Println(mCache.Get(12))
}
