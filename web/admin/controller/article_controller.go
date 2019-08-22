package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gopkg.in/go-playground/validator.v9"
	"time"
	"volt/data/entitys"
	"volt/data/model"
	"volt/services"
	"volt/web/admin/viewmodel"
)

type ArticleController struct {
	BaseController
	ArticleService *services.ArticleService
}

const article_viewpath  = "admin/views/article/"

func (this *ArticleController)Get() mvc.View{
	return mvc.View{
		Name:article_viewpath + "index.html",
	}
}

func (this *ArticleController) GetPage()interface{} {
	var articleSearch model.ArticleSearch
	if err := this.Ctx.ReadQuery(&articleSearch); err != nil {
		return viewmodel.NewDataTablesResult(0,[]entitys.Article{}).ToJson()
	}

	data,total,err :=this.ArticleService.GetListByPage(articleSearch)
	if err!=nil {
		return viewmodel.NewDataTablesResult(0,[]entitys.Article{}).ToJson()
	}
   return viewmodel.NewDataTablesResult(total,data).ToJson()
}


func (this *ArticleController)GetCreate()mvc.View  {
	tags := this.ArticleService.GetTags()
	return mvc.View{
		Name:article_viewpath + "create.html",
		Data:iris.Map{"tags":tags},
	}
}

func (this *ArticleController)PostCreate()interface{}  {
	//实例验证器
	validate := validator.New()

	var articleViewModel viewmodel.ArticleViewModel
	if err := this.Ctx.ReadJSON(&articleViewModel); err != nil {
		return iris.Map{"status":false,"msg":"请检查输入"}
	}

	err := validate.Struct(articleViewModel)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return iris.Map{"status":false,"msg":err.Error()}
		}
	}
	tagIds := make([]int,0,3)
	var tagStr = ""
	for _,v := range articleViewModel.Tags{
		tagIds = append(tagIds, v.ID)
		tagStr += (v.Name + ",")
	}

	article := new(entitys.Article)
	article.Status = 1
	article.Author = articleViewModel.Author
	article.Color = articleViewModel.TitleColor
	article.Content = articleViewModel.Content
	article.PublishTime = time.Now()
	article.Title = articleViewModel.Title
	article.Tags = tagStr[0:len(tagStr)-1]

   b := this.ArticleService.AddArticle(tagIds,article)
	if b {
		return iris.Map{"status":true,"msg":"创建成功"}
	}
	return iris.Map{"status":false,"msg":"创建失败"}
}
