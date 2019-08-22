package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type HomeController struct {
}

func (this *HomeController) Get(c iris.Context) mvc.View {

	return mvc.View{
		Name:"ui/views/home/index.html",
		Layout:"ui/views/layout/ui_layout",
	}
}