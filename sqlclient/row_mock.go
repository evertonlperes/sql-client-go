package sqlclient

type rowMock struct{}

type sqlRowsMock struct {
	rows rowMock
}

type rows interface {
	HasNext() bool
	Close() error
	Scan(destinations ...interface{}) error
}

//TODO: Finish mock implementation
