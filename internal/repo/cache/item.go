package cache

import (
	"context"
	"fmt"

	"github.com/Enthreeka/L0/internal/entity"
	"github.com/Enthreeka/L0/internal/repo"
)

type cacheItem struct {
	cache map[string][]entity.Item
}

func NewItemCache(cache map[string][]entity.Item) repo.Item {
	return &cacheItem{
		cache: cache,
	}
}

func (c *cacheItem) Create(ctx context.Context, id string, item entity.Item) error {

	c.cache[id] = append(c.cache[id], item)

	return nil
}

func (c *cacheItem) DeleteByID(ctx context.Context, id string) error {
	delete(c.cache, id)

	return nil
}

func (c *cacheItem) GetByID(ctx context.Context, id string) (*[]entity.Item, error) {

	data, ok := c.cache[id]
	if !ok {
		return nil, fmt.Errorf("%s", "Order number for item invalible")
	}

	return &data, nil
}

func (*cacheItem) GetAll(ctx context.Context) (*[]entity.Item, error) {
	panic("unimplemented")
}
