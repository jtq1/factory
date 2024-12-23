package models

type ClientService interface {
	CreateClient(client Client) (int64, error)
	GetClients() ([]Client, error)
	//GetClientByID(id int64) (models.Client, error)
	//UpdateClient(client models.Client) error
	//DeleteClient(id int64) error
}

type Client struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (c *Client) Validate() bool {
	return !(c.Name == "" || c.Email == "" || c.Phone == "")
}
