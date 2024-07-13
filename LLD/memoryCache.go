package main

import (
	"container/list"
	"fmt"
	"sync"
)

// CacheItem holds the value and the reference to the list element
type CacheItem struct {
	value interface{}
	node  *list.Element
}

// LRUCache struct
type LRUCache struct {
	capacity int
	cache    map[string]*CacheItem
	lruList  *list.List
	mu       sync.Mutex
}

// NewLRUCache creates a new LRU cache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*CacheItem),
		lruList:  list.New(),
	}
}

// Get retrieves an item from the cache
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, found := c.cache[key]; found {
		c.lruList.MoveToFront(item.node)
		return item.value, true
	}
	return nil, false
}

// Put adds an item to the cache
func (c *LRUCache) Put(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, found := c.cache[key]; found {
		// Update the value and move the node to the front
		item.value = value
		c.lruList.MoveToFront(item.node)
	} else {
		// Add new item
		if len(c.cache) >= c.capacity {
			// Remove the oldest item
			oldest := c.lruList.Back()
			if oldest != nil {
				c.lruList.Remove(oldest)
				delete(c.cache, oldest.Value.(string))
			}
		}
		// Insert the new item
		node := c.lruList.PushFront(key)
		c.cache[key] = &CacheItem{value, node}
	}
}

// Delete removes an item from the cache
func (c *LRUCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if item, found := c.cache[key]; found {
		c.lruList.Remove(item.node)
		delete(c.cache, key)
	}
}

// Display the cache for debugging purposes
func (c *LRUCache) Display() {
	for e := c.lruList.Front(); e != nil; e = e.Next() {
		key := e.Value.(string)
		fmt.Printf("%s: %v\n", key, c.cache[key].value)
	}
}

func main() {
	cache := NewLRUCache(3)

	cache.Put("A", 1)
	cache.Put("B", 2)
	cache.Put("C", 3)
	cache.Display()
	fmt.Println()

	cache.Get("A") // Access A, A becomes the most recently used
	cache.Put("D", 4) // Evicts B
	cache.Display()
	fmt.Println()

	cache.Delete("A") // Remove A
	cache.Display()
	fmt.Println()
}
