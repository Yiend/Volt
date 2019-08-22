package viewmodel

import "gopkg.in/go-playground/validator.v9"

type LoginViewModel struct {
	UserName         string     `json:"username" validate:"required"`
	PassWord     string     `json:"password" validate:"required"`
}

//自定义验证
func LoginViewModelValidation(sl validator.StructLevel) {
	//user := sl.Current().Interface().(LoginViewModel)

}