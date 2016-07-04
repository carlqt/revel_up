package controllers

import (
  "github.com/revel/revel"
  _ "github.com/carlqt/revel_up/app/models"
)

type Sessions struct {
  *revel.Controller
}

func (c Sessions) New() revel.Result {
  return c.Render()
}
