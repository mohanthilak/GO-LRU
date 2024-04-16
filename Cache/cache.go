package cache

import (
	"time"
)

type Cache struct {
	list             *LinkedList
	cacheMap         map[string]*Node
	Capacity         int
	EvictionDuration time.Duration
}

// Singleton Pattern
var instantiatedCache *Cache = nil

func InstantiateCache(Capacity int, evictionDurtion time.Duration) *Cache {
	// if instantiatedCache == nil {
	// 	instantiatedCache = &Cache{Capacity: Capacity, list: NewLinkedList(Capacity), cacheMap: make(map[string]*Node), EvictionDuration: evictionDurtion}
	// }

	// return instantiatedCache
	return &Cache{Capacity: Capacity, list: NewLinkedList(Capacity), cacheMap: make(map[string]*Node), EvictionDuration: evictionDurtion}
}

func (c *Cache) RemoveLRU() {
	delete(c.cacheMap, c.list.Tail.Prev.Key)
	c.list.remove(c.list.Tail.Prev)
}

func (c *Cache) Set(key string, value interface{}) {
	c.MakeEvictionCheck()
	node, ok := c.cacheMap[key]
	if ok {
		c.list.remove(node)
	} else {
		if c.list.size >= c.Capacity {
			c.RemoveLRU()
		}
	}
	c.list.add(key, value)
	c.cacheMap[key] = c.list.head.Next
}

func (c *Cache) Get(key string) interface{} {
	c.MakeEvictionCheck()
	node, ok := c.cacheMap[key]
	var value interface{} = -1
	if ok {
		value = node.Value
		c.list.swap(node)
		c.cacheMap[key] = c.list.head.Next
	}
	return value
}

func (c *Cache) PrintCache() {
	c.list.printList()
}

func (c *Cache) MakeEvictionCheck() {
	if c.list.size <= 0 {
		return
	}
	temp := c.list.Tail.Prev

	for temp != nil && c.list.size > 0 {
		if time.Since(temp.TimeAccessed) > c.EvictionDuration {
			temp = temp.Prev
			c.RemoveLRU()
		} else {
			break
		}
	}
}

func (c *Cache) ClearCache() {
	for k := range c.cacheMap {
		delete(c.cacheMap, k)
	}
	c.list = NewLinkedList(c.Capacity)
}
