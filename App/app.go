package app

import cache "go-lru/Cache"

type App struct {
	cache *cache.Cache
}

func NewApp(cache *cache.Cache) *App {
	return &App{cache: cache}
}

func (a *App) GetValueFromKey(key string) interface{} {
	return a.cache.Get(key)
}

func (a *App) SetKeyValuePair(key string, value interface{}) {
	a.cache.Set(key, value)
}
