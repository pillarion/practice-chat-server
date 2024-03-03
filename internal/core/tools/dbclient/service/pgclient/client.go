package service

import (
	db "github.com/pillarion/practice-chat-server/internal/core/tools/dbclient/port/pgclient"
)

type pgClient struct {
	masterDBC db.DB
}

// New initializes a new user repository using the provided database configuration.
func New(dbc db.DB) (db.Client, error) {
	return &pgClient{
		masterDBC: dbc,
	}, nil
}

// DB returns the database client
func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

// Close closes the database client
func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
