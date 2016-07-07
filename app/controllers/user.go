package controllers

import (
	"github.com/carlqt/revel_up/app/models"
	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	if c.Session["session_id"] == "" {
		c.Flash.Error("You are not logged in")
		return c.Redirect(Sessions.New)
	}

	greeting := "Hello Carl!"
	users := models.All()
	return c.Render(greeting, users)
}

func (c User) New() revel.Result {
	return c.Render()
}

func (c User) Create(username string, email string, password string, passwordConfirmation string) revel.Result {
	user := &models.User{Username: username, Email: email, Password: password, PasswordConfirmation: passwordConfirmation}
	user.Create()
	return c.Redirect(User.Index)
}

func checkUser(c *revel.Controller) revel.Result {
	revel.INFO.Println(c.Session.Id())
	return nil
}

func init() {
	revel.InterceptFunc(checkUser, revel.BEFORE, &User{})
}
