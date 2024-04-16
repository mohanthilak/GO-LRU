package cache

type Cache struct {
	list     *LinkedList
	cacheMap map[int]*Node
	Capacity int
}

// Singleton Pattern
var intantiatedCache *Cache = nil

func IntantiateCache(Capacity int) *Cache {
	if intantiatedCache == nil {
		intantiatedCache = &Cache{Capacity: Capacity, list: NewLinkedList(Capacity), cacheMap: make(map[int]*Node)}
	}

	return intantiatedCache
}

func (c *Cache) Set(key int, value int) {
	node, ok := c.cacheMap[key]
	if ok {
		c.list.remove(node)
	} else {
		if c.list.size >= c.Capacity {
			delete(c.cacheMap, c.list.Tail.Prev.Key)
			c.list.remove(c.list.Tail.Prev)
		}
	}
	c.list.add(key, value)
	c.cacheMap[key] = c.list.head.Next
}

func (c *Cache) Get(key int) int {
	node, ok := c.cacheMap[key]
	var value int = -1
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
