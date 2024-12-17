package cache

import "time"

type CacheItem struct {
	Value interface{}
	Valid time.Time
}

type cache struct {
	items map[string]CacheItem
}

var Cache *cache = &cache{
	items: make(map[string]CacheItem),
}

func Get(key string) (interface{}, bool) {
	item, exists := Cache.items[key]
	if !exists {
		return nil, false
	}
	if item.Valid.Before(time.Now().UTC()) {
		Delete(key)
		return nil, false
	}

	return item.Value, true
}

func Set(key string, value interface{}, ttl time.Duration) {
	Cache.items[key] = CacheItem{
		Value: value,
		Valid: time.Now().Add(ttl),
	}
}

func Delete(key string) {
	delete(Cache.items, key)
}
