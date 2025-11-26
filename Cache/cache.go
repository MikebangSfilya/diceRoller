package cache

import "sync"

type CacheRols struct {
	Name string
	Sum  int8
}

func New(name string, sum int8) CacheRols {
	return CacheRols{
		Name: name,
		Sum:  sum,
	}
}

type InitiativeCache struct {
	data []CacheRols
	mu   sync.RWMutex
}

func NewInitiativeCache() *InitiativeCache {
	return &InitiativeCache{
		data: make([]CacheRols, 0),
	}
}

func (c *InitiativeCache) Set(data []CacheRols) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = data
}

func (c *InitiativeCache) Get() []CacheRols {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data
}

func (c *InitiativeCache) Append(data []CacheRols) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = append(c.data, data...)
}
