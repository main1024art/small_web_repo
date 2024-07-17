package controllers

import (
	"fmt"
	"pro01/bace/auth"
	"pro01/models"
	"pro01/utils"
)

type HomeController struct {
	auth.AuthorController
}

func (c *HomeController) Index() {
	c.TplName = "home/index.html"
}

func (c *HomeController) ChangerPassword() {
	if c.Ctx.Input.IsPost() {
		oldPassword := c.GetString("old_password")
		newPassword := c.GetString("new_password")
		confirmPassword := c.GetString("confirm_password")
		id := c.GetSession("user")
		userid, _ := id.(int)
		fmt.Println(userid)
		user := models.GetUserById(userid)
		if user.ValiPassword(oldPassword) && newPassword == confirmPassword {
			user.Password = utils.Md5Text(newPassword)
			models.ModifyUser(user)
			c.Data["Success"] = "Password changed successfully"
		} else {
			c.Data["Error"] = "New password and confirm password do not match"
			c.TplName = "changeword/change_password.html"
			return
		}
	}

	c.TplName = "changeword/change_password.html"

}
