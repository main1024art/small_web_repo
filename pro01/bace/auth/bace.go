package auth

import (
	"github.com/astaxie/beego"
	"net/http"
	"pro01/bace/bace"
)

type AuthorController struct {
	bace.BaceConteoller
}

func (ac *AuthorController) Prepare() {
	user := ac.GetSession("user")
	if user == nil {
		ac.Redirect(beego.URLFor("LoginController.Login"), http.StatusFound)
	}
}
