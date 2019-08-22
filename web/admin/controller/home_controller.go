package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get(c iris.Context) mvc.Result {
	//if !this.isLoggedIn() {
	//	return mvc.Response{
	//		Path:"/admin/login",
	//	}
	//}
	return mvc.View{
		Name:"admin/views/home/index.html",
	}
}