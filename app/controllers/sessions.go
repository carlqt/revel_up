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

func (c Sessions) Create() revel.Result {
  return c.Redirect(User.Index)
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