package controllers

import (
	"github.com/revel/revel"
  "github.com/carlqt/revel_up/app/models"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	greeting := "Hello Carl!"
	users := models.All()
	return c.Render(greeting, users)
}

func (c User) New() revel.Result {
	return c.Render()
}

func (c User) Create(username string, email string, password string, passwordConfirmation string) revel.Result {
	user := &models.User{username, email, password, passwordConfirmation}
	user.Create()
	return c.Redirect(User.Index)
}