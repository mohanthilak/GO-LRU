package main

import (
	"fmt"
	cache "go-lru/Cache"
)

func main() {

	cache := cache.IntantiateCache(3)
	cache.Set(1, 11)
	cache.Set(2, 22)
	cache.Set(1, 10)
	cache.PrintCache()
	fmt.Println("\n", cache.Get(1))
	cache.PrintCache()

	cache.Set(3, 33)
	cache.Set(4, 44)
	cache.PrintCache()

}
