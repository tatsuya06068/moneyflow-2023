package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	database "github.com/tatsuya06068/moneyflow-2023/internal/adapter/gateway"
	"github.com/tatsuya06068/moneyflow-2023/internal/constants"
	"github.com/tatsuya06068/moneyflow-2023/internal/entity"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(baseInfo entity.BaseDbInfo) database.ISqlHandler {
	jst, err := time.LoadLocation(constants.LocationAisia)
	if err != nil {
		panic(err)
	}

	// DB接続情報
	config := mysql.Config{
		DBName:    baseInfo.HostName,
		User:      baseInfo.User,
		Passwd:    baseInfo.Password,
		Addr:      "db",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	conn, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

/********************
    sqlHandler methods
**************************/

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.IRows, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRows), err
	}
	row := new(SqlRows)
	row.Rows = rows
	return rows, nil
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.IResult, error) {
	res := SqlResult{}
	stmt, err := handler.Conn.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	exe, err := stmt.Exec(args...)
	if err != nil {
		return res, err
	}
	res.Result = exe
	return res, nil
}

// begin transaction
func (handler *SqlHandler) Begin() (database.ITx, error) {
	res := SqlTransaction{}
	transaction, err := handler.Conn.Begin()
	if err != nil {
		return res, err
	}
	res.Tx = transaction
	return res, err
}

type SqlRows struct {
	Rows *sql.Rows
}

type SqlResult struct {
	Result sql.Result
}

type SqlTransaction struct {
	Tx *sql.Tx
}

/***************************
         transaction
***************************/

func (t SqlTransaction) Commit() error {
	return t.Tx.Commit()
}

func (t SqlTransaction) Rollback() error {
	return t.Tx.Rollback()
}

func (t SqlTransaction) Execute(statement string, args ...interface{}) (database.IResult, error) {
	res := SqlResult{}
	stmt, err := t.Tx.Prepare(statement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()
	exe, err := stmt.Exec(args...)
	if err != nil {
		return res, err
	}
	res.Result = exe
	return res, nil
}

func (t SqlTransaction) Query(statement string, args ...interface{}) (database.IRows, error) {
	rows, err := t.Tx.Query(statement, args...)
	if err != nil {
		return new(SqlRows), err
	}
	row := new(SqlRows)
	row.Rows = rows
	return rows, nil
}

/********************
    Rows methods
**************************/

func (r SqlRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRows) Next() bool {
	return r.Rows.Next()
}

func (r SqlRows) Close() error {
	return r.Rows.Close()
}

/********************
    Results methods
**************************/

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}
