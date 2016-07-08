package models

import (
	"github.com/carlqt/revel_up/app"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
  "database/sql"
	"strconv"
)

type User struct {
	ID 			 string
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
		result, _ := stmnt.Exec(u.Username, u.Email, hashedPassword)
		u.ID = getInsertId(result)
	} else {
		revel.ERROR.Println(err)
	}
}

func (u *User) Authenticate() bool{
	var hashedPassword []byte

	rows, err := app.DB.Query("SELECT id, password FROM users WHERE username = ?", u.Username)
	defer rows.Close()

	if err != nil {
		revel.ERROR.Println(err)
	}

	for rows.Next() {
		err := rows.Scan(&u.ID, &hashedPassword)

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

func getInsertId(r sql.Result) string {
	id, _ := r.LastInsertId()
	s := strconv.FormatInt(id, 10)

	return s
}