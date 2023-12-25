package models

import (
	"database/sql"
)

type User struct {
	User UserScheme
	db   *sql.DB
}

type UserScheme struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (user *User) PostUser() string {
	return "test"
}
