package driver

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxDBConnections = 10
const maxIdleDbConnection = 5
const maxDBLifeTime = 5 * time.Minute

func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxDBConnections)
	d.SetMaxIdleConns(maxIdleDbConnection)
	d.SetConnMaxLifetime(maxDBLifeTime)

	dbConn.SQL = d

	err = testDB(d)

	if err != nil{
		return nil, err
	}
	return dbConn, nil

}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}

	return nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}