package sqlclient

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type client struct {
	db *sql.DB
}

type SqlClient interface {
	Query(query string, args ...interface{}) (rows, error)
}

func Open(driverName, dataSourceName string) (SqlClient, error) {
	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}

	database, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	client := &client{
		db: database,
	}

	return client, nil
}

func (c *client) Query(query string, args ...interface{}) (rows, error) {
	returnedRows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	results := sqlRows{
		rows: returnedRows,
	}
	return &results, nil
}
