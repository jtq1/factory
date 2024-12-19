package models

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
