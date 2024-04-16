package main

import (
	app "go-lru/App"
	cache "go-lru/Cache"
	server "go-lru/Server"
)

func main() {

	cache := cache.IntantiateCache(3)

	App := app.NewApp(cache)

	HttpServer := server.NewHTTPServer(8080, App)
	HttpServer.StartServer()

	// cache.Set(1, 11)
	// cache.Set(2, 22)
	// cache.Set(1, 10)
	// cache.PrintCache()
	// fmt.Println("\n", cache.Get(1))
	// cache.PrintCache()

	// cache.Set(3, 33)
	// cache.Set(4, 44)
	// cache.PrintCache()

}
