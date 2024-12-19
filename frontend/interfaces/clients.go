package interfaces

import "appTalleres/backend/models"

type ClientMaster interface {
	CreateClient(client models.Client) (int64, error)
	GetClients() ([]models.Client, error)
	//GetClientByID(id int64) (models.Client, error)
	//UpdateClient(client models.Client) error
	//DeleteClient(id int64) error
}
