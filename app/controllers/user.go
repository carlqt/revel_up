package controllers

import (
	"github.com/revel/revel"
	"time"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	greeting := "Hello baby!"
	return c.Render(greeting)
}

func (c User) New() revel.Result {
	now := time.Now()
	revel.INFO.Println("STATUS " + now.String())
	return c.Render()
}

func (c User) Create(username string, email string) revel.Result {
	revel.INFO.Println(email)
	return c.Redirect(User.Index)
}
