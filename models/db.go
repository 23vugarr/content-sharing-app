package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	Params string
	DB     *sql.DB
}

type DBparams struct {
	DBHost   string
	Port     int
	User     string
	Password string
	Dbname   string
}

func NewDB(params DBparams) *DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		params.DBHost, params.Port, params.User, params.Password, params.Dbname)

	return &DB{Params: psqlInfo, DB: nil}
}

func (db *DB) Connect() error {
	var err error

	db.DB, err = sql.Open("postgres", db.Params)
	if err != nil {
		return err
	}

	return nil
}
