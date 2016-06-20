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

func All() (users map[string]string){
  var (
    uname string
    email string
    )

  rows, _ := app.DB.Query("Select username, email from users where id = 1")
  for rows.Next() {
    rows.Scan(&uname, &email)
    revel.INFO.Println(email)
  } 

  return users
}