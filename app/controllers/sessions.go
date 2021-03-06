package controllers

import (
  "net/url"
  "github.com/revel/revel"
  "github.com/carlqt/revel_up/app/models"
)

type Sessions struct {
  *revel.Controller
}

func (c Sessions) New() revel.Result {

  return c.Render()
}

func (c Sessions) Create(username string, password string) revel.Result {
  u := models.User{Username: username, Password: password}

  if u.Authenticate() {
    c.Flash.Success("Login successful")
    return c.Redirect(User.Index)
  } else {
    c.Flash.Error("User or Password is incorrect")
    return c.RenderTemplate("Sessions/New.html")
  }
}

func (c Sessions) Register() revel.Result {
  u := newUser(c.Params.Form)

  u.Validate(c.Validation)

  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
    return c.Redirect(Sessions.New)
  } else {
    u.Create()
    c.Flash.Success("User created")
    return c.Redirect(User.Index)
  }
}

func (c Sessions) Destroy() revel.Result {
  return c.Redirect(Sessions.New)
}

func newUser(params url.Values) *models.User {
  return &models.User{
    Username: params.Get("username"),
    Email: params.Get("email"),
    Password: params.Get("password"),
    PasswordConfirmation: params.Get("passwordConfirmation"),
  }
}