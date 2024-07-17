package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"pro01/bace/auth"
	"pro01/models"
	"pro01/utils"
	"strconv"
)

type UserController struct {
	auth.AuthorController
}

func (c *UserController) Query() {
	users := models.Queryuser()
	c.Data["users"] = users
	c.TplName = "user/query.html"
}
func (c *UserController) DeleteUser() {
	id, _ := strconv.Atoi(c.GetString("id"))
	_ = models.DeleteUserById(id)
	c.Redirect(beego.URLFor("UserController.Query"), 302)
}

func (c *UserController) ModifyUsers() {
	id := c.GetSession("user")
	userid, _ := id.(int)
	fmt.Println(userid)
	user := models.GetUserById(userid)
	c.TplName = "modify/modify.html"
	if c.Ctx.Request.Method == "POST" {
		d := c.GetString("name")
		a := c.GetString("password")
		b := c.GetString("email")
		if a != "" && b != "" {
			user.Name = d
			user.Email = b
			user.Password = utils.Md5Text(a)
			models.ModifyUser(user)
			fmt.Println(user)
		}
	}
}
