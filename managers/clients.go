package managers

import (
	"appTalleres"
	"fmt"
)

type masterClient struct {
	MySQLClientDB models.ClientService
	CacheClient   models.ClientService
	dbEnabled     bool
	inMemEnabled  bool
}

func NewManagerClient(db, cache models.ClientService, dbEnabled, inMem bool) *masterClient {
	return &masterClient{
		MySQLClientDB: db,
		CacheClient:   cache,
		dbEnabled:     dbEnabled,
		inMemEnabled:  inMem,
	}
}

func (m *masterClient) CreateClient(client models.Client) (int64, error) {
	if !client.Validate() {
		return 0, fmt.Errorf("el cliente introducido no tiene los campos validos")
	}

	if m.dbEnabled {
		idDB, err := m.MySQLClientDB.CreateClient(client)
		if err != nil {
			return 0, err
		}
		client.ID = idDB
	}

	if m.inMemEnabled {
		idCache, err := m.CacheClient.CreateClient(client)
		if err != nil {
			return 0, err
		}
		client.ID = idCache
	}

	return client.ID, nil
}

func (m *masterClient) GetClients() ([]models.Client, error) {
	if m.inMemEnabled {
		return m.CacheClient.GetClients()
	}

	if m.dbEnabled {
		return m.MySQLClientDB.GetClients()
	}

	return nil, nil
}
