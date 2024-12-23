package mock

import (
	models "appTalleres"
	"github.com/stretchr/testify/mock"
)

type clientMock struct {
	mock.Mock
}

func NewClientMock() *clientMock {
	return &clientMock{}
}

func (c *clientMock) CreateClient(client models.Client) (int64, error) {
	args := c.Called(client)
	return args.Get(0).(int64), args.Error(1)
}

func (c *clientMock) GetClients() ([]models.Client, error) {
	args := c.Called()
	return args.Get(0).([]models.Client), args.Error(1)
}
