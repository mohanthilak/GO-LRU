package cache

type cache struct {
	list     *LinkedList
	cacheMap map[int]*Node
	Capacity int
}

// Singleton Pattern
var intantiatedCache *cache = nil

func IntantiateCache(Capacity int) *cache {
	if intantiatedCache == nil {
		intantiatedCache = &cache{Capacity: Capacity, list: NewLinkedList(Capacity), cacheMap: make(map[int]*Node)}
	}

	return intantiatedCache
}

func (c *cache) Set(key int, value int) {
	node, ok := c.cacheMap[key]
	if ok {
		c.list.remove(node)
	} else {
		if c.list.size >= c.Capacity {
			c.list.remove(c.list.Tail.Prev)
		}
	}
	c.list.add(key, value)
	c.cacheMap[key] = c.list.head.Next
}

func (c *cache) Get(key int) int {
	node, ok := c.cacheMap[key]
	var value int = -1
	if ok {
		value = node.Value
		c.list.swap(node)
		c.cacheMap[key] = c.list.head.Next
	}
	return value
}

func (c *cache) PrintCache() {
	c.list.printList()
}
