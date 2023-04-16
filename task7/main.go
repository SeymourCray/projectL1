package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	//встраиваем sync.RWMutex
	sync.RWMutex
	items map[int]string
}

// создаем мапу
func NewCache() *Cache {
	return &Cache{items: make(map[int]string)}
}

func (c *Cache) Insert(key int, value string) {
	//блокируем дпугие Lock и RLock
	c.Lock()

	defer c.Unlock()

	c.items[key] = value
}

func (c *Cache) Get(key int) string {
	//блокируем дпугие Lock
	c.RLock()

	defer c.RUnlock()

	return c.items[key]
}

func main() {
	cache := NewCache()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			cache.Insert(i, fmt.Sprintf("example %d", i))
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			fmt.Println(cache.Get(i))
		}(i)
	}

	wg.Wait()
}
