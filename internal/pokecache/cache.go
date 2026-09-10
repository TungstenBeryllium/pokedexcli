package pokecache

import (

	"sync"
	"time"

)


type Cache struct {

	Entry	map[string]cacheEntry
	Mu	sync.Mutex
}


type cacheEntry struct {

	createdAt	time.Time
	val	[]byte
}



func NewCache(interval time.Duration) *Cache {

	
	c := Cache{
		Entry: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return &c

}



func (c *Cache) Add(key string, val []byte) {
	
	c.Mu.Lock()
	defer c.Mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.Entry[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	
	c.Mu.Lock()
	defer c.Mu.Unlock()

	value, ok := c.Entry[key]
	if ok {
		return value.val, true
	} else {
		return nil, false
	} 
}

func (c *Cache) reapLoop(interval time.Duration) {
	

	ticker := time.NewTicker(interval)

	for range ticker.C {
	
		c.Mu.Lock()
       		

		for key, value := range c.Entry {

			if time.Since(value.createdAt) > interval {
				delete(c.Entry, key)
				
			}

		}
		c.Mu.Unlock()
	}

}
