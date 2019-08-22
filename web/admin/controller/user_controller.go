package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gopkg.in/go-playground/validator.v9"
	"volt/services"
	"volt/web/admin/viewmodel"

	_ "gopkg.in/go-playground/validator.v9"
)

type UserController struct {
	BaseController
	//服务
	UserService *services.UserService
}

func (this *UserController) GetLogin() mvc.View {
	if this.isLoggedIn() {
		this.logout()
	}
	return mvc.View{
		Name:"admin/views/user/login.html",
		Layout:iris.NoLayout,
	}
}

func (this *UserController) PostLogin() mvc.Result{
	returnurl := this.Ctx.URLParam("returnurl")
	//实例验证器
	validate := validator.New()
	//注入自定义验证器
	validate.RegisterStructValidation(viewmodel.LoginViewModelValidation, viewmodel.LoginViewModel{})

	var loginViewModel viewmodel.LoginViewModel
	if err := this.Ctx.ReadForm(&loginViewModel); err != nil {
		return mvc.View{
			Name:"admin/views/user/login.html",
			Layout:iris.NoLayout,
			Data:iris.Map{"errormsg": "您输入有误"},
		}
	}

	err := validate.Struct(loginViewModel)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return mvc.View{
				Name:"admin/views/user/login.html",
				Layout:iris.NoLayout,
				Data:iris.Map{"errormsg": err.Error()},
			}
		}
	}

	u, b,s := this.UserService.Login(loginViewModel.UserName, loginViewModel.PassWord)
	if !b {
		return mvc.View{
			Name:"admin/views/user/login.html",
			Layout:iris.NoLayout,
			Data:iris.Map{"errormsg": s},
		}
	}

	this.Session.Set(currentUserKey,u)
	if returnurl != "" {
		return mvc.Response	{
			Path: returnurl,
		}
	}else {
		return mvc.Response	{
			Path: "/admin",
		}
	}

}

func (this *UserController) Logout() {
	this.Session.Destroy()
	this.Ctx.Redirect("/admin/login")
}