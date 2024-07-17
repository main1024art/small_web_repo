package controllers

import (
	"github.com/astaxie/beego"
	"pro01/errors"
	"pro01/form"
	"pro01/models"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	forms := &form.LoginForm{}
	errs := errors.New()
	if c.Ctx.Request.Method == "POST" {
		if err := c.ParseForm(forms); err == nil {
			user := models.GetUserByName(forms.Name)
			if user == nil {
				errs.Add("default", "用户不存在")
			} else if user.ValiPassword(forms.Password) {
				c.SetSession("user", user.Id)
				c.Redirect(beego.URLFor("HomeController.Index"), 302)
			} else {
				errs.Add("default", "密码错误")
			}
		}
	}
	c.Data["errors"] = errs
	c.Data["form"] = forms
	c.TplName = "auth/login.html"
}

func (c *AuthController) Logout() {
	c.DestroySession()
	c.Redirect(beego.URLFor("AuthController.Login"), 302)
}
