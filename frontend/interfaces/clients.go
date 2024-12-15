package interfaces

import "appTalleres/backend/models"

type ClientDB interface {
	CreateClient(client models.Client) (int64, error)
	GetClientByID(id int64) (models.Client, error)
	UpdateClient(client models.Client) error
	DeleteClient(id int64) error
}
