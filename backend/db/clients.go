package db

import (
	"appTalleres/backend/models"
	"database/sql"
	"fmt"
)

type clientDB struct {
	db *sql.DB
}

func NewClientDB(db *sql.DB) *clientDB {
	return &clientDB{
		db: db,
	}
}

func (c *clientDB) CreateClient(client models.Client) (int64, error) {
	query := "INSERT INTO clients (name, email, phone, address) VALUES (?, ?, ?, ?)"
	result, err := c.db.Exec(query, client.Name, client.Email, client.Phone, client.Address)
	if err != nil {
		return 0, fmt.Errorf("failed to insert client: %v", err)
	}
	return result.LastInsertId()
}

func (c *clientDB) GetClientByID(id int64) (models.Client, error) {
	query := "SELECT id, name, email, phone, address FROM clients WHERE id = ?"
	var client models.Client
	row := c.db.QueryRow(query, id)
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.Address); err != nil {
		if err == sql.ErrNoRows {
			return client, fmt.Errorf("client with id %d not found", id)
		}
		return client, fmt.Errorf("failed to get client: %v", err)
	}
	return client, nil
}

func (c *clientDB) UpdateClient(client models.Client) error {
	query := "UPDATE clients SET name = ?, email = ?, phone = ?, address = ? WHERE id = ?"
	_, err := c.db.Exec(query, client.Name, client.Email, client.Phone, client.Address, client.ID)
	if err != nil {
		return fmt.Errorf("failed to update client: %v", err)
	}
	return nil
}

func (c *clientDB) DeleteClient(id int64) error {
	query := "DELETE FROM clients WHERE id = ?"
	_, err := c.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete client: %v", err)
	}
	return nil
}
