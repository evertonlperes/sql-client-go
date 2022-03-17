package sqlclient

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	goEnvironment = "GO_ENV"
	production    = "production"
)

var (
	isMocked bool
	dbClient SqlClient
)

type client struct {
	db *sql.DB
}

type SqlClient interface {
	Query(query string, args ...interface{}) (rows, error)
}

func StartMockServer() {
	isMocked = true
}

func StopMockServer() {
	isMocked = false
}

func isProduction() bool {
	return os.Getenv(goEnvironment) == "production"
}

func Open(driverName, dataSourceName string) (SqlClient, error) {
	if isMocked && !isProduction() {
		dbClient := &clientMock{}
		return dbClient, nil
	}
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
