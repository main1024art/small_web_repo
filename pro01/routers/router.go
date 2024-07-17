package routers

import (
	"github.com/astaxie/beego"
	"pro01/controllers"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.AddController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.Router("/", &controllers.AuthController{}, "*:Login")
}
