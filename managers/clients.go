package managers

import (
	models "appTalleres"
	"fmt"
)

type masterClient struct {
	MySQLClientDB   models.ClientService
	CacheClient     models.ClientService
	CacheController models.Cache
	eventManager    models.EventService
	dbEnabled       bool
	inMemEnabled    bool
}

func NewManagerClient(db, cache models.ClientService, cacheCont models.Cache, events models.EventService, dbEnabled, inMem bool) *masterClient {
	return &masterClient{
		MySQLClientDB:   db,
		CacheClient:     cache,
		CacheController: cacheCont,
		eventManager:    events,
		dbEnabled:       dbEnabled,
		inMemEnabled:    inMem,
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
		clients, err := m.CacheClient.GetClients()
		if err != nil {
			return nil, err
		}
		if len(clients) != 0 {
			return m.CacheClient.GetClients()
		}
	}

	if m.dbEnabled {
		return m.MySQLClientDB.GetClients()
	}

	return nil, nil
}

func (m *masterClient) SyncCache() error {
	if m.inMemEnabled && m.dbEnabled {
		clients, err := m.MySQLClientDB.GetClients()
		if err != nil {
			return fmt.Errorf("error getting clients: %v", err)
		}

		err = m.CacheController.Sync(clients)
		if err != nil {
			return fmt.Errorf("error syncing cache: %v", err)
		}

		m.eventManager.Send(models.ClientCacheRefreshedSubject)
	}

	return nil
}
