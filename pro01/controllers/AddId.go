package controllers

import (
	"github.com/astaxie/beego"
	"pro01/errors"
	"pro01/models"
)

type AddController struct {
	beego.Controller
}

func (c *AddController) Adduser() {
	user := models.User{}
	errs := errors.New()
	if c.Ctx.Request.Method == "POST" {
		name := c.GetString("name")
		password := c.GetString("password")
		email := c.GetString("email")
		if name != "" && password != "" && email != "" {
			user.Name = name
			user.Email = email
			user.Password = password
			_ = user.GetFormUser(&user)
		} else {
			errs.Add("04", "不能为空")
		}
	}
	c.Data["errors"] = errs

	c.TplName = "add/adduser.html"
}
