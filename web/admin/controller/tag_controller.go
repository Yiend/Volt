package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
	"volt/data/entitys"
	"volt/data/model"
	"volt/services"
	"volt/web/admin/viewmodel"
)

type TagController struct {
	BaseController
	TagService *services.TagService
}
const tag_viewpath  = "admin/views/tag/"

func (this *TagController) Get() mvc.View {
	return mvc.View{
		Name:tag_viewpath + "index.html",
	}
}


func (this *TagController) GetPage()interface{} {
	var tagSearch model.TagSearch
	if err := this.Ctx.ReadQuery(&tagSearch); err != nil {
		return viewmodel.NewDataTablesResult(0,[]entitys.Tag{}).ToJson()
	}

	data,total,err :=this.TagService.GetListByPage(tagSearch)
	if err!=nil {
		return viewmodel.NewDataTablesResult(0,[]entitys.Tag{}).ToJson()
	}
	return viewmodel.NewDataTablesResult(total,data).ToJson()
}


func (this *TagController)GetCreate()mvc.View  {
	return mvc.View{
		Name:tag_viewpath + "create.html",
		Layout:iris.NoLayout,
	}
}

func (this *TagController)PostCreate() interface{} {
	tagName := this.Ctx.FormValue("tagName")
	if tagName =="" {
		return iris.Map{"status":false,"msg":"请输入标签名称"}
	}
   b := this.TagService.AddTag(tagName)
	if b {
		return iris.Map{"status":true,"msg":""}
	}

	return iris.Map{"status":false,"msg":"添加失败"}
}

func (this *TagController) GetEdit()mvc.View  {
	tagId,_ := this.Ctx.URLParamInt64("id")
	if tagId == 0 {
		return mvc.View{
			Name:tag_viewpath + "edit.html",
			Data:iris.Map{"tagId":0,"tagName":""},
			Layout:iris.NoLayout,
		}
	}

	model,err := this.TagService.GetByID(tagId)
	if err!=nil {
		return mvc.View{
			Name:tag_viewpath + "edit.html",
			Data:iris.Map{"tagId":0,"tagName":""},
			Layout:iris.NoLayout,
		}
	}

	return mvc.View{
		Name:tag_viewpath + "edit.html",
		Data:iris.Map{"tagId":model.ID,"tagName":model.Name},
		Layout:iris.NoLayout,
	}
}

func (this *TagController)PostEdit()interface{}  {
	tagId := this.Ctx.FormValue("tagId")
	tagName := this.Ctx.FormValue("tagName")
	v,e:= strconv.ParseInt(tagId,10,64)
	if e != nil || v <= 0{
		return iris.Map{"status":false,"msg":"记录不存在，请刷新页面"}
	}
	if tagName =="" {
		return iris.Map{"status":false,"msg":"请输入标签名称"}
	}
	b := this.TagService.UpdateTag(v,tagName)
	if b {
		return iris.Map{"status":true,"msg":""}
	}

	return iris.Map{"status":false,"msg":"更新失败"}
}