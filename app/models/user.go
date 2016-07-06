package models

import (
	"github.com/carlqt/revel_up/app"
	"github.com/revel/revel"
)

type User struct {
	Username string
	Email    string
	Password string
	PasswordConfirmation string
}

func (u *User) Validate(v *revel.Validation) {
	v.Required(u.Username)
	v.Required(u.Email)
	v.Required(u.Password)
	v.Required(u.PasswordConfirmation)
	v.Required(u.Password == u.PasswordConfirmation).Message("Password does not match")
}

func (u *User) Create() {
	stmnt, err := app.DB.Prepare("INSERT INTO users(username, email, password) VALUES(?, ?, ?)")

	if err == nil {
		stmnt.Exec(u.Username, u.Email, u.Password)
	} else {
		revel.ERROR.Println(err)
	}
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
