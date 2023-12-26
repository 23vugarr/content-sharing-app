package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	User UserScheme
	db   *sql.DB
}

type UserScheme struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) PostUser() string {
	return "test"
}

func (user *User) GetUser() string {
	return "pass"
}

func (db *DB) FetchUserByUsername(username string) (UserScheme, error) {
	var user User

	err := db.DB.QueryRow("SELECT username, password FROM users WHERE username = $1", username).Scan(&user.User.Username, &user.User.Password)
	if err != nil {
		fmt.Println(err)
		return UserScheme{}, err
	}
	return user.User, nil
}
