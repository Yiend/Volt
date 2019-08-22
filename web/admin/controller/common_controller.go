package controller

import (
	"github.com/kataras/iris"
	"io"
	"os"
)

type CommonController struct {
   BaseController
}

func (this *CommonController) PostUpload() interface{}{
	file, info, err := this.Ctx.FormFile("upload")
	if err != nil {
		return iris.Map{"uploaded":"0","error":iris.Map{"message":"请选择文件"}}
	}
	defer file.Close()
	fname := info.Filename
	out, err := os.OpenFile("./assets/uploads/"+fname,os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return iris.Map{"uploaded":"0","error":iris.Map{"message":err.Error()}}
	}
	defer out.Close()
	io.Copy(out, file)

	return iris.Map{"uploaded":"1","fileName":fname,"url":"/static/uploads/"+fname}
}
