package app

import cache "go-lru/Cache"

type App struct {
	cache *cache.Cache
}

func NewApp(cache *cache.Cache) *App {
	return &App{cache: cache}
}

func (a *App) GetValueFromKey(key int) int {
	return a.cache.Get(key)
}

func (a *App) SetKeyValuePair(key, value int) {
	a.cache.Set(key, value)
}
