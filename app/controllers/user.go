package controllers

import (
	"github.com/revel/revel"
  "github.com/carlqt/revel_up/app/models"
	"time"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	greeting := "Hello Carl!"
	users := models.All()
	revel.INFO.Println(users)
	return c.Render(greeting, users)
}

func (c User) New() revel.Result {
	now := time.Now()
	revel.ERROR.Println("Hello from down under")
	revel.INFO.Println("STATUS " + now.String())
	return c.Render()
}

func (c User) Create(username string, email string) revel.Result {
	user := &models.User{username, email}
	user.Create()
	return c.Redirect(User.Index)
}