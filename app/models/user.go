package models

import (
  "github.com/carlqt/revel_up/app"
  _ "github.com/mattn/go-sqlite3"
  "database/sql"
  "github.com/revel/revel"
)

type User struct {
  Username string
  Email string
}

func (u *User) Create() {
  stmnt, _ := app.DB.Prepare("INSERT INTO users(username, email) VALUES(?, ?)")
  stmnt.Exec(u.Username, u.Email)
}

func (u *User) All() *sql.Rows{
  stmnt, _ := app.DB.Query("Select * from users")
  return stmnt
}

func All() []map[string]interface{}{
  var username, email string

  users := make([]map[string]interface{}, 2)

  i := 0
  rows, _ := app.DB.Query("Select username, email from users")

  for rows.Next() {
    rows.Scan(&username, &email)
    users[i] = map[string]interface{}{"username": username, "email": email}
    i++
    revel.INFO.Println(users[0]["email"])
  } 

  return users
}