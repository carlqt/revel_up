package models

import (
	"github.com/carlqt/revel_up/app"
)

type User struct {
	Username string
	Email    string
}

func (u *User) Create() {
	stmnt, _ := app.DB.Prepare("INSERT INTO users(username, email) VALUES(?, ?)")
	stmnt.Exec(u.Username, u.Email)
}

func All() []User {
	var user User

	rows, _ := app.DB.Query("Select username, email from users")

	users := make([]User, 0)

	for rows.Next() {
		rows.Scan(&user.Username, &user.Email)
		users = append(users, user)
	}

	return users
}
