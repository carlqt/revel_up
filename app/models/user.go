package models

import (
	_ "database/sql"
	"github.com/carlqt/revel_up/app"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
)

type User struct {
	Username string
	Email    string
}

func (u *User) Create() {
	stmnt, _ := app.DB.Prepare("INSERT INTO users(username, email) VALUES(?, ?)")
	stmnt.Exec(u.Username, u.Email)
}

func All() []map[string]interface{} {
	var username, email string

	rows, _ := app.DB.Query("Select username, email from users")

	users := make([]map[string]interface{}, 0)

	for rows.Next() {
		rows.Scan(&username, &email)
		//users[i] = map[string]interface{}{"username": username, "email": email}
		users = append(users, map[string]interface{}{"username": username, "email": email})

		revel.INFO.Println(users[0]["email"])
	}

	return users
}
