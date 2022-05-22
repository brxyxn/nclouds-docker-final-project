package data

import (
	"database/sql"

	u "github.com/brxyxn/go_gpr_nclouds/backend/utils"
)

type User struct {
	UserID   uint64 `json:"userId"`
	Username string `json:"userName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Counter struct {
	Value int `json:"counter"`
}

func CreateUser(db *sql.DB, v *User, c *Counter) error {
	err := db.QueryRow(
		"INSERT INTO users(username, email, password) VALUES($1, $2, $3) RETURNING user_id;",
		&v.Username, &v.Email, &v.Password,
	).Scan(
		&v.UserID,
	)
	if err != nil {
		return err
	}

	return CountUsers(db, c)
}

func CountUsers(db *sql.DB, c *Counter) error {
	u.Log.Debug("Handling CountUsers")
	return db.QueryRow(
		"SELECT COUNT(*) FROM users;",
	).Scan(&c.Value)
}
