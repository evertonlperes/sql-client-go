package sqlclient

import "errors"

func AddMock(mock Mock) {
	client := dbClient.(*clientMock)
	if client.mocks == nil {
		client.mocks = make(map[string]Mock, 0)
	}
	client.mocks[mock.Query] = mock
}

type clientMock struct {
	mocks map[string]Mock
}

type Mock struct {
	Query   string
	Args    []interface{}
	Error   error
	Columns []string
	Rows    [][]interface{}
}

func (c *clientMock) Query(query string, args ...interface{}) (rows, error) {
	mock, exists := c.mocks[query]
	if !exists {
		return nil, errors.New("no mock found")
	}

	if mock.Error != nil {
		return nil, mock.Error
	}

	return nil, nil
}
