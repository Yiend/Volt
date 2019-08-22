package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"volt/data/entitys"
)

const currentUserKey = "User"

type BaseController struct {
	//上下文
	Ctx iris.Context
	//会话
	Session *sessions.Session

}

func (this *BaseController)isLoggedIn()bool  {
	i := this.Session.Get(currentUserKey)
	user,ok:= i.(*entitys.User)
	if ok && user.ID > 0 {
		return true
	}
	return false
}

func (this *BaseController) logout() {
	this.Session.Destroy()
}
