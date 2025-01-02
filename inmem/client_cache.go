package inmem

import (
	models "appTalleres"
	"fmt"
	"sync"
)

var _ models.ClientService = (*clientCache)(nil)

type clientCache struct {
	mu    sync.Mutex
	cache map[int64]models.Client
}

func NewClientCache() *clientCache {
	return &clientCache{
		cache: make(map[int64]models.Client),
	}
}

func (c *clientCache) CreateClient(client models.Client) (int64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if client.ID == 0 {
		id := int64(len(c.cache) + 1)
		client.ID = id
	}
	c.cache[client.ID] = client
	return client.ID, nil
}

func (c *clientCache) GetClients() ([]models.Client, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var cliSlice []models.Client
	for _, value := range c.cache {
		cliSlice = append(cliSlice, value)
	}
	return cliSlice, nil
}

func (c *clientCache) Sync(objs interface{}) error {
	clients, ok := objs.([]models.Client)
	if !ok {
		return fmt.Errorf("invalid type: expected []models.Client")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache = make(map[int64]models.Client)
	for i := range clients {
		c.cache[clients[i].ID] = clients[i]
	}
	return nil
}
