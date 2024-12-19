package managers

import (
	"appTalleres/backend/models"
	"fmt"
)

type ClientDB interface {
	CreateClient(client models.Client) (int64, error)
	GetClients() ([]models.Client, error)
	GetClientByID(id int64) (models.Client, error)
	UpdateClient(client models.Client) error
	DeleteClient(id int64) error
}

type MasterClient struct {
	clientDB ClientDB
}

func NewManagerClient(db ClientDB) *MasterClient {
	return &MasterClient{
		clientDB: db,
	}
}

func (m *MasterClient) CreateClient(client models.Client) (int64, error) {
	if !client.Validate() {
		return 0, fmt.Errorf("el cliente introducido no tiene los campos validos")
	}

	return m.clientDB.CreateClient(client)
}

func (m *MasterClient) GetClients() ([]models.Client, error) {
	return m.clientDB.GetClients()
}
