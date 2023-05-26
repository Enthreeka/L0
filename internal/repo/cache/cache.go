package cache

import (
	"sync"

	"github.com/Enthreeka/L0/internal/entity"
)

type OrderCache struct {
	cache map[string][]entity.Order
	mu    sync.RWMutex
}

func (d *OrderCache) GetByID(id string) data []entity.Order {
	d.mu.RLock()
	//data, _ := d.cache[id]
	//var data []entity.Order
	d.cache[id] = data
	d.mu.RUnlock()

	return data
}

func (d *OrderCache) SearchByID(id string) bool {

	_, ok := d.cache[id]
	if !ok {
		return false
	}
	return true
}
