package models

import (
	"github.com/carlqt/revel_up/app"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Email    string
	Password string
	PasswordConfirmation string
}

func (u *User) Validate(v *revel.Validation) {
	v.Required(u.Username).Message("Username is required")
	v.Required(u.Email).Message("Email is required")
	v.Required(u.Password).Message("Password is required")
	v.Required(u.PasswordConfirmation).Message("Password Confirmation is required")
	v.Required(u.Password == u.PasswordConfirmation).Message("Password does not match")
}

func (u *User) Create() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	stmnt, err := app.DB.Prepare("INSERT INTO users(username, email, password) VALUES(?, ?, ?)")

	if err == nil {
		stmnt.Exec(u.Username, u.Email, hashedPassword)
	} else {
		revel.ERROR.Println(err)
	}
}

func (u *User) Authenticate() bool{
	var hashedPassword []byte

	rows, err := app.DB.Query("SELECT password FROM users WHERE username = ?", u.Username)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&hashedPassword)

		if err != nil {
			revel.ERROR.Println(err)	
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(u.Password))

	if err != nil {
		return false
	} else {
		return true
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
