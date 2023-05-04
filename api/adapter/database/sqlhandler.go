package database

type ISqlHandler interface {
	Query(string, ...interface{}) (IRows, error)
	Execute(string, ...interface{}) (IResult, error)
	Begin() (ITx, error)
}

// transaction
type ITx interface {
	Commit() error
	Execute(string, ...interface{}) (IResult, error)
	Rollback() error
	Query(string, ...interface{}) (IRows, error)
}

type IRows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

type IResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
