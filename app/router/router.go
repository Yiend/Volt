package router

import (
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/mvc"
	"volt/app/middleware"
	"volt/app/startup"
	"volt/data/repository"
	"volt/services"
	adminCtr "volt/web/admin/controller"
	webCtr "volt/web/ui/controller"
)
var (
	userRepository *repository.UserRepository
	articleRepository *repository.ArticleRepository
	tagRepository *repository.TagRepository
	tagArticleRepository *repository.TagArticleRepository
)
var(
	userService *services.UserService
	articleService *services.ArticleService
	tagService *services.TagService
)


func Configure(s *startup.Startup) {
	initRepository(s.Xorm)

	//前台
	mvc.Configure(s.Party("/"), func(app *mvc.Application) {
		app.Register(articleService)
		app.Handle(new(webCtr.HomeController))
	})

	//后台
	admin := mvc.New(s.Party("/admin"))
	{
		hero.Register(s.Sessions.Start)
		//设置 这个路由组下面的路由 全部走这个中间件
		admin.Router.Use(hero.Handler(middleware.Authentication))

		admin.Configure(func(app *mvc.Application) {
			app.Register(articleRepository)
			app.Handle(new(adminCtr.HomeController))
		})

		//article
		admin.Party("/article").Configure(func(app *mvc.Application) {
			app.Register(articleService)
			app.Handle(new(adminCtr.ArticleController))
		})

		//tag
		admin.Party("/tag").Configure(func(app *mvc.Application) {
			app.Register(tagService)
			app.Handle(new(adminCtr.TagController))
		})

         //common
		admin.Party("/common").Configure(func(app *mvc.Application) {
			app.Handle(new(adminCtr.CommonController))
		})
	}

	//登录
	tourist := mvc.New(s.Party("/admin"))
	tourist.Configure(func(app *mvc.Application) {
		app.Register(userService)
		app.Handle(new(adminCtr.UserController))
	})
}

func initRepository(db *xorm.Engine)  {
	userRepository = repository.NewUserRepository(db)
	articleRepository = repository.NewArticleRepository(db)
	tagRepository = repository.NewTagRepository(db)
	tagArticleRepository = repository.NewTagArticleRepository(db)

	userService = services.NewUserService(userRepository)
	articleService = services.NewArticleService(articleRepository,tagRepository,tagArticleRepository)
	tagService = services.NewTagService(tagRepository)
}



